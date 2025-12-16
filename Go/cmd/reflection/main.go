// Package main demonstrates Reflection pattern using LoopAgent in Google ADK-Go
//
// Reflection cho phep agent tu danh gia va cai thien output qua nhieu vong lap.
// Su dung LoopAgent + functiontool de dieu khien vong lap:
// - InitialProducer: Tao code ban dau
// - Critic: Review va dua ra feedback hoac approve
// - Refiner: Cai thien code hoac goi exitLoop neu approved
// - LoopAgent: Lap cho den khi exitLoop() duoc goi hoac dat max iterations
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/loopagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// State keys
const (
	stateCode      = "current_code"
	stateCritique  = "critique"
	approvedPhrase = "CODE_APPROVED"
)

// ============================================================================
// EXIT LOOP TOOL
// ============================================================================

// ExitLoopArgs dinh nghia arguments cho exitLoop tool (rong)
type ExitLoopArgs struct{}

// ExitLoopResults dinh nghia ket qua tra ve (rong)
type ExitLoopResults struct{}

// ExitLoop la function tool signal cho LoopAgent dung
// Khi duoc goi, set ctx.Actions().Escalate = true
func ExitLoop(ctx tool.Context, input ExitLoopArgs) (ExitLoopResults, error) {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Printf("[exitLoop] Code da duoc phe duyet!\n")
	fmt.Printf("[exitLoop] Triggered by agent: %s\n", ctx.AgentName())
	fmt.Println("═══════════════════════════════════════════════════════")
	ctx.Actions().Escalate = true
	return ExitLoopResults{}, nil
}

// ============================================================================
// INITIAL CODE PRODUCER
// ============================================================================

func createInitialProducer(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "InitialCodeProducer",
		Model:       m,
		Description: "Viet code ban dau dua tren yeu cau cua user.",
		Instruction: `Ban la mot Senior Software Engineer chuyen viet code Go.

NHIEM VU:
Viet code theo yeu cau cua nguoi dung.

YEU CAU:
- Viet code sach, co comment giai thich
- Xu ly cac edge cases co ban
- Tuan thu Go best practices
- Them error handling

DINH DANG:
Output CHI code Go trong block, khong giai thich them.`,
		OutputKey: stateCode,
	})
}

// ============================================================================
// CODE CRITIC
// ============================================================================

func createCodeCritic(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "CodeCritic",
		Model:       m,
		Description: "Review code va dua ra feedback hoac approve.",
		Instruction: fmt.Sprintf(`Ban la mot Principal Engineer voi 20 nam kinh nghiem.
Vai tro cua ban la review code mot cach ti mi va KHACH QUAN.

**CODE CAN REVIEW:**
"""
{%s}
"""

**TIEU CHI DANH GIA:**
1. Tinh dung dan - Code co hoat dong dung khong?
2. Error handling - Co xu ly loi day du khong?
3. Edge cases - Co xu ly cac truong hop dac biet khong?
4. Code quality - Code co sach, de doc khong?

**HANH DONG:**

NEU code co 1-2 diem can cai thien:
-> Liet ke cu the cac diem can sua.
-> Output CHI feedback text.

NEU code da hoan hao, dat TAT CA tieu chi:
-> Tra loi CHINH XAC: %s
-> KHONG them bat ky text nao khac.`, stateCode, approvedPhrase),
		OutputKey: stateCritique,
	})
}

// ============================================================================
// CODE REFINER (voi exitLoop tool)
// ============================================================================

func createCodeRefiner(m model.LLM, exitLoopTool tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "CodeRefiner",
		Model:       m,
		Description: "Cai thien code hoac goi exitLoop neu code duoc approve.",
		Instruction: fmt.Sprintf(`Ban la mot Software Engineer cai thien code dua tren feedback.

**CODE HIEN TAI:**
"""
{%s}
"""

**FEEDBACK TU REVIEWER:**
{%s}

**HANH DONG:**

1. Doc ky feedback tu reviewer.

2. NEU feedback CHINH XAC la "%s":
   -> Goi function 'exitLoop' NGAY LAP TUC.
   -> KHONG output bat ky text nao.

3. NEU feedback chua goi y cai thien:
   -> Ap dung TAT CA cac goi y.
   -> Output CHI code da cai thien.
   -> KHONG giai thich, KHONG goi exitLoop.`, stateCode, stateCritique, approvedPhrase),
		Tools:     []tool.Tool{exitLoopTool},
		OutputKey: stateCode,
	})
}

// ============================================================================
// REFLECTION PIPELINE
// ============================================================================

func createReflectionPipeline(m model.LLM) (agent.Agent, error) {
	// 1. Tao Initial Producer (chay 1 lan dau)
	initialProducer, err := createInitialProducer(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create initial producer: %w", err)
	}

	// 2. Tao exitLoop tool
	exitLoopTool, err := functiontool.New(
		functiontool.Config{
			Name:        "exitLoop",
			Description: "Goi function nay KHI VA CHI KHI critique CHINH XAC la CODE_APPROVED. Khi goi, KHONG output text.",
		},
		ExitLoop,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create exitLoop tool: %w", err)
	}

	// 3. Tao Critic Agent (trong loop)
	critic, err := createCodeCritic(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create critic: %w", err)
	}

	// 4. Tao Refiner Agent voi exitLoop tool (trong loop)
	refiner, err := createCodeRefiner(m, exitLoopTool)
	if err != nil {
		return nil, fmt.Errorf("failed to create refiner: %w", err)
	}

	// 5. Tao Refinement Loop
	refinementLoop, err := loopagent.New(loopagent.Config{
		MaxIterations: 3,
		AgentConfig: agent.Config{
			Name:        "RefinementLoop",
			Description: "Vong lap: Critic review -> Refiner cai thien hoac exit",
			SubAgents:   []agent.Agent{critic, refiner},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create loop agent: %w", err)
	}

	// 6. Tao Pipeline tong the
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "CodeReviewPipeline",
			Description: "Pipeline: InitialProducer -> RefinementLoop",
			SubAgents:   []agent.Agent{initialProducer, refinementLoop},
		},
	})
}

// ============================================================================
// MAIN
// ============================================================================

func printBanner() {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  Code Review System - Reflection Pattern with LoopAgent")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("")
	fmt.Println("  Pipeline:")
	fmt.Println("  ┌─────────────────────────────────────────────────┐")
	fmt.Println("  │  1. InitialProducer → {current_code}            │")
	fmt.Println("  │                 ↓                               │")
	fmt.Println("  │  2. LoopAgent (max=3)                           │")
	fmt.Println("  │     ┌─────────────────────────────────────┐     │")
	fmt.Println("  │     │ Critic → {critique}                 │     │")
	fmt.Println("  │     │            ↓                        │     │")
	fmt.Println("  │     │ Refiner → exitLoop() or refine      │     │")
	fmt.Println("  │     └─────────────────────────────────────┘     │")
	fmt.Println("  └─────────────────────────────────────────────────┘")
	fmt.Println("")
	fmt.Println("Khoi dong server...")
	fmt.Println("Vi du: 'Viet ham tinh fibonacci trong Go'")
}

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("Vui long set GOOGLE_API_KEY environment variable")
	}

	geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Khong the tao model: %v", err)
	}

	// Tao Pipeline
	pipeline, err := createReflectionPipeline(geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(pipeline),
	}

	lch := full.NewLauncher()
	printBanner()

	err = lch.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
