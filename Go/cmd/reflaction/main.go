package main

import (
	"fmt"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/loopagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// State keys
const (
	stateCode      = "current_code"
	stateCritique  = "critique"
	approvedPhrase = "CODE_APPROVED"
)

type ExitLoopArgs struct {
}
type ExitLoopResult struct{}

func ExitLoop(ctx tool.Context, input ExitLoopArgs) (ExitLoopResult, error) {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Printf("[exitLoop] Code đã được phê duyệt!\n")
	fmt.Printf("[exitLoop] Triggered by agent: %s\n", ctx.AgentName())
	fmt.Println("═══════════════════════════════════════════════════════")
	ctx.Actions().Escalate = true
	return ExitLoopResult{}, nil

}

// ============================================================================
// INITIAL CODE PRODUCER
// ============================================================================

func createInitialProducer(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "InitialCodeProducer",
		Model:       m,
		Description: "Viết code ban đầu dựa trên yêu cầu của user.",
		Instruction: `Bạn là một Senior Software Engineer chuyên viết code Go.

NHIỆM VỤ:
Viết code theo yêu cầu của người dùng.

YÊU CẦU:
- Viết code sạch, có comment giải thích
- Xử lý các edge cases cơ bản
- Tuân thủ Go best practices
- Thêm error handling

ĐỊNH DẠNG:
Output CHỈ code Go trong block, không giải thích thêm.`,
		OutputKey: stateCode,
	})
}

// ============================================================================
// INITIAL CODE PRODUCER
// ============================================================================

func createInitialProducer(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "InitialCodeProducer",
		Model:       m,
		Description: "Viết code ban đầu dựa trên yêu cầu của user.",
		Instruction: `Bạn là một Senior Software Engineer chuyên viết code Go.

NHIỆM VỤ:
Viết code theo yêu cầu của người dùng.

YÊU CẦU:
- Viết code sạch, có comment giải thích
- Xử lý các edge cases cơ bản
- Tuân thủ Go best practices
- Thêm error handling

ĐỊNH DẠNG:
Output CHỈ code Go trong block, không giải thích thêm.`,
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
		Description: "Review code và đưa ra feedback hoặc approve.",
		Instruction: fmt.Sprintf(`Bạn là một Principal Engineer với 20 năm kinh nghiệm.
Vai trò của bạn là review code một cách tỉ mỉ và KHÁCH QUAN.

**CODE CẦN REVIEW:**
"""
{%s}
"""

**TIÊU CHÍ ĐÁNH GIÁ:**
1. Tính đúng đắn - Code có hoạt động đúng không?
2. Error handling - Có xử lý lỗi đầy đủ không?
3. Edge cases - Có xử lý các trường hợp đặc biệt không?
4. Code quality - Code có sạch, dễ đọc không?

**HÀNH ĐỘNG:**

NẾU code có 1-2 điểm cần cải thiện:
→ Liệt kê cụ thể các điểm cần sửa.
→ Output CHỈ feedback text.

NẾU code đã hoàn hảo, đạt TẤT CẢ tiêu chí:
→ Trả lời CHÍNH XÁC: %s
→ KHÔNG thêm bất kỳ text nào khác.`, stateCode, approvedPhrase),
		OutputKey: stateCritique,
	})
}

func createCodeRefiner(m model.LLM, exitLoopTool tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "CodeRefinder",
		Model:       m,
		Description: "Cải thiện code hoặc gọi exitLoop nếu code được approve.",
		Instruction: fmt.Sprintf(`Bạn là một Software Engineer cải thiện code dựa trên feedback.
**CODE HIỆN TẠI:**
"""
{%s}
"""
**FEEDBACK TỪ REVIEWER:**
{%s}
**HÀNH ĐỘNG:**
1. Đọc kỹ feedback từ reviewer.
2. NẾU feedback CHÍNH XÁC là "%s":
   → Gọi function 'exitLoop' NGAY LẬP TỨC.
   → KHÔNG output bất kỳ text nào.
3. NẾU feedback chứa gợi ý cải thiện:
   → Áp dụng TẤT CẢ các gợi ý.
   → Output CHỈ code đã cải thiện.
→ KHÔNG giải thích, KHÔNG gọi exitLoop.`,
			stateCode,
			stateCritique,
			approvedPhrase),
		Tools:     []tool.Tool{exitLoopTool},
		OutputKey: stateCode,
	})
}

func createReflectionPipline(m model.LLM) (agent.Agent, error) {
	initialProducer, err := createInitialProducer(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create initial producer: %w", err)
	}
	exitLoopTool, err := functiontool.New(
		functiontool.Config{
			Name:        "exitLoop",
			Description: fmt.Sprintf("Gọi function này KHI VÀ CHỈ KHI %s CHÍNH XÁC là %s. Khi gọi, KHÔNG output text.", stateCritique, approvedPhrase),
		}, ExitLoop)

	if err != nil {
		return nil, fmt.Errorf("failed to create exitLoop function: %w", err)
	}
	critic, err := createCodeCritic(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create code critic: %w", err)
	}
	refiner, err := createCodeRefiner(m, exitLoopTool)
	if err != nil {
		return nil, fmt.Errorf("failed to create refiner: %w", err)
	}
	refinementLoop, err := loopagent.New(loopagent.Config{
		MaxIterations: 3,
		AgentConfig: agent.Config{
			Name: "RefinementLoop",
			Description: "Vòng lặp: Critic review → Refiner cải thiện hoặc exit",
			SubAgents: []agent.Agent{critic, refiner},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create loop agent: %w", err)
	}
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name: "Sequential",
		}
	})
}
