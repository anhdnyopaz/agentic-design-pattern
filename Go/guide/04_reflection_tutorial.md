# Hướng dẫn Thực hành: Reflection Pattern với ADK-Go

## Mục tiêu

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu mẫu Reflection (Phản chiếu / Tự phê bình) trong thiết kế Agentic
- Triển khai mô hình Producer-Critic với ADK-Go
- **Sử dụng LoopAgent và functiontool để điều khiển vòng lặp một cách deterministic**
- Xây dựng vòng lặp phản hồi để cải thiện chất lượng output
- Áp dụng Reflection vào các bài toán thực tế

---

## Phần 1: Giới thiệu Reflection Pattern

### 1.1 Reflection là gì?

**Reflection (Phản chiếu)** là kỹ thuật cho phép agent tự đánh giá công việc của mình và sử dụng đánh giá đó để cải thiện output. Đây là một hình thức **tự sửa lỗi (self-correction)**.

### 1.2 Quy trình Reflection

```
┌─────────────────────────────────────────────────────────────┐
│                    REFLECTION LOOP                          │
│                                                             │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐              │
│  │ PRODUCER │───▶│  OUTPUT  │───▶│  CRITIC  │              │
│  │  Agent   │    │ (Draft)  │    │  Agent   │              │
│  └──────────┘    └──────────┘    └──────────┘              │
│       ▲                               │                     │
│       │         ┌──────────┐          │                     │
│       └─────────│ CRITIQUE │◀─────────┘                     │
│                 │(Feedback)│                                │
│                 └──────────┘                                │
│                      │                                      │
│              ┌───────▼───────┐                              │
│              │  Đạt yêu cầu? │                              │
│              └───────┬───────┘                              │
│                 Yes  │  No                                  │
│                  ▼   └──────────▶ Lặp lại                   │
│              OUTPUT                                         │
│              CUỐI CÙNG                                      │
└─────────────────────────────────────────────────────────────┘
```

### 1.3 Các bước trong Reflection

1. **Thực thi (Execution):** Producer tạo output ban đầu
2. **Đánh giá (Evaluation):** Critic phân tích và phê bình output
3. **Tinh chỉnh (Refinement):** Producer cải thiện dựa trên phê bình
4. **Lặp lại (Iteration):** Tiếp tục cho đến khi đạt yêu cầu

### 1.4 Khi nào sử dụng Reflection?

- **Chất lượng quan trọng hơn tốc độ** (viết nội dung, tạo code)
- Cần **độ chính xác cao** (báo cáo, phân tích)
- Output cần **tuân thủ tiêu chuẩn** phức tạp
- **Sáng tạo nội dung** cần trau chuốt nhiều lần

---

## Phần 2: Tại sao dùng LoopAgent?

### 2.1 Vấn đề với LLM-based Loop Control

Cách tiếp cận cũ sử dụng một LLM orchestrator để điều khiển vòng lặp qua prompt:

```go
// ❌ Cách cũ - LLM-based (không khuyến khích)
Instruction: `Thực hiện VÒNG LẶP:
BƯỚC 1: Gọi producer...
BƯỚC 2: Gọi critic...
BƯỚC 3: Nếu approved → dừng, nếu không → lặp lại...`
```

**Vấn đề:**
- ❌ **Không deterministic**: LLM có thể không tuân thủ đúng số vòng
- ❌ **Tốn token**: LLM phải reasoning về loop logic mỗi vòng
- ❌ **Khó debug**: Loop state ẩn trong LLM reasoning
- ❌ **Unreliable**: LLM có thể "quên" lặp hoặc lặp vô hạn

### 2.2 Giải pháp: LoopAgent + exitLoop Tool

**LoopAgent** là workflow agent của ADK, kết hợp với **functiontool** để điều khiển vòng lặp:

```go
// ✅ Cách mới - LoopAgent + exitLoop tool (khuyến khích)
// 1. Tạo exitLoop tool
exitLoopTool, _ := functiontool.New(
    functiontool.Config{Name: "exitLoop", Description: "..."},
    ExitLoop,  // Function sets ctx.Actions().Escalate = true
)

// 2. Agent có tool này sẽ gọi khi cần dừng
refinerAgent, _ := llmagent.New(llmagent.Config{
    Tools: []tool.Tool{exitLoopTool},
    // ...
})

// 3. LoopAgent bọc ngoài
loopagent.New(loopagent.Config{
    MaxIterations: 3,
    AgentConfig: agent.Config{
        SubAgents: []agent.Agent{criticAgent, refinerAgent},
    },
})
```

**Ưu điểm:**
- ✅ **Deterministic**: Loop dừng khi tool được gọi hoặc đạt max
- ✅ **Tiết kiệm token**: Không tốn token cho loop reasoning
- ✅ **Dễ debug**: Có thể log khi exitLoop được gọi
- ✅ **ADK-idiomatic**: Sử dụng đúng API của ADK-Go

### 2.3 So sánh chi tiết

| Tiêu chí | LLM-based | LoopAgent + exitLoop |
|----------|-----------|----------------------|
| Max iterations | Phụ thuộc prompt | `MaxIterations` parameter |
| Stopping condition | LLM parse text | `exitLoop` tool call |
| Token cost | Cao (reasoning mỗi vòng) | Thấp |
| Reliability | ~70-80% | ~99%+ |
| Implementation | Complex prompt | Simple function |

---

## Phần 3: Kiến trúc với LoopAgent

### 3.1 Kiến trúc tổng quan

```
┌─────────────────────────────────────────────────────────────────┐
│                     SequentialAgent (Pipeline)                  │
│                                                                 │
│  ┌────────────────┐                                            │
│  │ Initial        │  Tạo code ban đầu                          │
│  │ Producer       │  → {current_code}                          │
│  └───────┬────────┘                                            │
│          ↓                                                      │
│  ┌────────────────────────────────────────────────────────┐    │
│  │              LoopAgent (max_iterations=3)              │    │
│  │                                                        │    │
│  │  ┌─────────────┐     ┌─────────────────────────────┐  │    │
│  │  │   Critic    │────▶│   Refiner                   │  │    │
│  │  │   Agent     │     │   Agent                     │  │    │
│  │  │             │     │                             │  │    │
│  │  │ Review code │     │ IF critique="CODE_APPROVED":│  │    │
│  │  │ ↓           │     │   → call exitLoop()         │  │    │
│  │  │ {critique}  │     │ ELSE:                       │  │    │
│  │  │             │     │   → refine code             │  │    │
│  │  └─────────────┘     │   → {current_code}          │  │    │
│  │                      │                             │  │    │
│  │                      │ Tools: [exitLoop]           │  │    │
│  │                      └─────────────────────────────┘  │    │
│  │                                                        │    │
│  │  exitLoop() → ctx.Actions().Escalate = true → EXIT    │    │
│  └────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
```

### 3.2 Sequence Diagram - Luồng xử lý ReflectionPipeline

```mermaid
sequenceDiagram
    autonumber
    participant User
    participant Pipeline as SequentialAgent<br/>(CodeReviewPipeline)
    participant Producer as InitialCodeProducer
    participant State as State Store
    participant Loop as LoopAgent<br/>(RefinementLoop)
    participant Critic as CodeCritic
    participant Refiner as CodeRefiner
    participant ExitTool as exitLoop Tool

    %% Phase 1: Initial Code Generation
    User->>Pipeline: "Viết hàm fibonacci"
    activate Pipeline

    Pipeline->>Producer: Execute
    activate Producer
    Producer->>Producer: Generate initial code
    Producer->>State: Save {current_code}
    Producer-->>Pipeline: Done
    deactivate Producer

    %% Phase 2: Refinement Loop
    Pipeline->>Loop: Execute (max_iterations=3)
    activate Loop

    %% === Iteration 1 ===
    rect rgb(255, 245, 238)
        Note over Loop: Iteration 1

        Loop->>Critic: Execute
        activate Critic
        Critic->>State: Read {current_code}
        State-->>Critic: code v1
        Critic->>Critic: Review code
        Critic->>State: Save {critique} = "Thiếu xử lý n < 0"
        Critic-->>Loop: Feedback
        deactivate Critic

        Loop->>Refiner: Execute
        activate Refiner
        Refiner->>State: Read {current_code}, {critique}
        State-->>Refiner: code v1, feedback
        Refiner->>Refiner: critique != "CODE_APPROVED"<br/>→ Refine code
        Refiner->>State: Save {current_code} = code v2
        Refiner-->>Loop: Refined code
        deactivate Refiner
    end

    %% === Iteration 2 ===
    rect rgb(240, 255, 240)
        Note over Loop: Iteration 2

        Loop->>Critic: Execute
        activate Critic
        Critic->>State: Read {current_code}
        State-->>Critic: code v2
        Critic->>Critic: Review code
        Critic->>State: Save {critique} = "Cần thêm comment"
        Critic-->>Loop: Feedback
        deactivate Critic

        Loop->>Refiner: Execute
        activate Refiner
        Refiner->>State: Read {current_code}, {critique}
        State-->>Refiner: code v2, feedback
        Refiner->>Refiner: critique != "CODE_APPROVED"<br/>→ Refine code
        Refiner->>State: Save {current_code} = code v3
        Refiner-->>Loop: Refined code
        deactivate Refiner
    end

    %% === Iteration 3 (Final - Approved) ===
    rect rgb(240, 248, 255)
        Note over Loop: Iteration 3 (Final)

        Loop->>Critic: Execute
        activate Critic
        Critic->>State: Read {current_code}
        State-->>Critic: code v3
        Critic->>Critic: Review code<br/>✓ All criteria passed!
        Critic->>State: Save {critique} = "CODE_APPROVED"
        Critic-->>Loop: Approved
        deactivate Critic

        Loop->>Refiner: Execute
        activate Refiner
        Refiner->>State: Read {critique}
        State-->>Refiner: "CODE_APPROVED"
        Refiner->>Refiner: critique == "CODE_APPROVED"<br/>→ Call exitLoop
        Refiner->>ExitTool: exitLoop()
        activate ExitTool
        ExitTool->>ExitTool: ctx.Actions().Escalate = true
        ExitTool-->>Refiner: Done
        deactivate ExitTool
        Note over Loop,Refiner: Escalate=true → Loop exits immediately
        Refiner-->>Loop: Exit signal
        deactivate Refiner
    end

    Loop-->>Pipeline: Final code (v3)
    deactivate Loop

    Pipeline->>State: Read {current_code}
    State-->>Pipeline: Final refined code
    Pipeline-->>User: Return final code
    deactivate Pipeline
```

**Giải thích luồng:**

1. **User** gửi yêu cầu viết code
2. **Pipeline** (SequentialAgent) điều phối luồng xử lý
3. **InitialProducer** tạo code ban đầu (v1) → lưu vào State
4. **RefinementLoop** (LoopAgent) bắt đầu với max 3 iterations:
   - **Critic** đọc code từ State, review, output feedback hoặc "CODE_APPROVED"
   - **Refiner** đọc feedback:
     - Nếu có feedback → cải thiện code → tiếp tục loop
     - Nếu "CODE_APPROVED" → gọi `exitLoop()` → loop dừng
5. **Pipeline** trả về code cuối cùng cho User

### 3.3 Vai trò của từng thành phần

**Initial Producer Agent:**
- Chạy **một lần** ở đầu pipeline
- Tạo code ban đầu dựa trên yêu cầu của user
- Output lưu vào `{current_code}`

**Critic Agent (trong loop):**
- **Đánh giá code** hiện tại
- Nếu có vấn đề: Đưa ra feedback cụ thể
- Nếu code hoàn hảo: Trả về **chính xác** chuỗi `CODE_APPROVED`
- Output lưu vào `{critique}`

**Refiner Agent (trong loop):**
- Đọc `{critique}` từ Critic
- **Nếu critique = "CODE_APPROVED"**: Gọi `exitLoop()` tool
- **Nếu có feedback**: Cải thiện code và lưu vào `{current_code}`
- Có tool: `exitLoop`

**exitLoop Tool:**
- Function tool đơn giản
- Set `ctx.Actions().Escalate = true`
- Khi được gọi, LoopAgent sẽ dừng ngay lập tức

---

## Phần 4: Bài tập - Code Review Agent với LoopAgent

### Mô tả bài tập

Xây dựng một **Code Review System** với:

1. **Initial Code Producer** - Viết code ban đầu
2. **Code Critic** - Review code và đưa ra feedback hoặc approve
3. **Code Refiner** - Cải thiện code hoặc gọi exitLoop nếu approved
4. **LoopAgent** - Điều phối vòng lặp refinement

### Cấu trúc file

```
cmd/reflection/
├── main.go
```

---

## Phần 5: Hướng dẫn từng bước

### Bước 1: Import packages

```go
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
    stateCode     = "current_code"
    stateCritique = "critique"
    approvedPhrase = "CODE_APPROVED"
)
```

### Bước 2: Tạo exitLoop Tool

Đây là cách **đúng** để signal cho LoopAgent dừng:

```go
// ExitLoopArgs định nghĩa arguments (rỗng)
type ExitLoopArgs struct{}

// ExitLoopResults định nghĩa kết quả trả về (rỗng)
type ExitLoopResults struct{}

// ExitLoop là function tool signal cho loop dừng
// Khi được gọi, set Escalate = true để LoopAgent biết cần exit
func ExitLoop(ctx tool.Context, input ExitLoopArgs) (ExitLoopResults, error) {
    fmt.Printf("[exitLoop] Triggered by agent: %s\n", ctx.AgentName())
    ctx.Actions().Escalate = true
    return ExitLoopResults{}, nil
}
```

**Giải thích:**
- `tool.Context` cung cấp `Actions()` method
- `ctx.Actions().Escalate = true` signal cho LoopAgent dừng
- Function trả về empty struct vì không cần output

### Bước 3: Tạo Initial Code Producer Agent

```go
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

ĐỊNH DẠNG TRẢ VỀ:
Chỉ output code Go, không thêm giải thích.

` + "```go" + `
// Code ở đây
` + "```",
        OutputKey: stateCode,
    })
}
```

### Bước 4: Tạo Code Critic Agent

```go
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
- NẾU code có vấn đề cần sửa:
  Liệt kê 1-2 điểm cần cải thiện cụ thể. Output CHỈ feedback.

- NẾU code đã đạt tất cả tiêu chí:
  Trả lời CHÍNH XÁC chuỗi: %s
  Không thêm bất kỳ text nào khác.`, stateCode, approvedPhrase),
        OutputKey: stateCritique,
    })
}
```

### Bước 5: Tạo Code Refiner Agent với exitLoop Tool

```go
func createCodeRefiner(m model.LLM, exitLoopTool tool.Tool) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "CodeRefiner",
        Model:       m,
        Description: "Cải thiện code hoặc gọi exitLoop nếu code được approve.",
        Instruction: fmt.Sprintf(`Bạn là một Creative Software Engineer cải thiện code dựa trên feedback.

**CODE HIỆN TẠI:**
"""
{%s}
"""

**FEEDBACK TỪ REVIEWER:**
{%s}

**HÀNH ĐỘNG:**

1. Đọc kỹ feedback từ reviewer.

2. NẾU feedback CHÍNH XÁC là "%s":
   → Gọi function 'exitLoop' ngay lập tức.
   → KHÔNG output bất kỳ text nào.

3. NẾU feedback chứa gợi ý cải thiện:
   → Áp dụng các gợi ý để cải thiện code.
   → Output CHỈ code đã cải thiện, không giải thích.`, stateCode, stateCritique, approvedPhrase),
        Tools:     []tool.Tool{exitLoopTool},
        OutputKey: stateCode,
    })
}
```

### Bước 6: Tạo Reflection Loop

```go
func createReflectionPipeline(m model.LLM) (agent.Agent, error) {
    // 1. Tạo Initial Producer (chạy 1 lần đầu)
    initialProducer, err := createInitialProducer(m)
    if err != nil {
        return nil, fmt.Errorf("failed to create initial producer: %w", err)
    }

    // 2. Tạo exitLoop tool
    exitLoopTool, err := functiontool.New(
        functiontool.Config{
            Name:        "exitLoop",
            Description: "Gọi function này KHI VÀ CHỈ KHI critique là CODE_APPROVED. Đừng output text nếu gọi function này.",
        },
        ExitLoop,
    )
    if err != nil {
        return nil, fmt.Errorf("failed to create exitLoop tool: %w", err)
    }

    // 3. Tạo Critic Agent (trong loop)
    critic, err := createCodeCritic(m)
    if err != nil {
        return nil, fmt.Errorf("failed to create critic: %w", err)
    }

    // 4. Tạo Refiner Agent với exitLoop tool (trong loop)
    refiner, err := createCodeRefiner(m, exitLoopTool)
    if err != nil {
        return nil, fmt.Errorf("failed to create refiner: %w", err)
    }

    // 5. Tạo Refinement Loop
    // Loop chạy: Critic → Refiner → Critic → Refiner → ...
    // Dừng khi: Refiner gọi exitLoop() HOẶC đạt MaxIterations
    refinementLoop, err := loopagent.New(loopagent.Config{
        MaxIterations: 3,
        AgentConfig: agent.Config{
            Name:        "RefinementLoop",
            Description: "Vòng lặp refinement: critique → refine → repeat",
            SubAgents:   []agent.Agent{critic, refiner},
        },
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create loop agent: %w", err)
    }

    // 6. Tạo Pipeline tổng thể
    // Pipeline: InitialProducer → RefinementLoop
    return sequentialagent.New(sequentialagent.Config{
        AgentConfig: agent.Config{
            Name:        "CodeReviewPipeline",
            Description: "Pipeline: tạo code → refinement loop",
            SubAgents:   []agent.Agent{initialProducer, refinementLoop},
        },
    })
}
```

### Bước 7: Main function

```go
func printBanner() {
    fmt.Println("═══════════════════════════════════════════════════════")
    fmt.Println("  Code Review System - Reflection Pattern với LoopAgent")
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
    fmt.Println("Khởi động server...")
}

func main() {
    ctx := context.Background()

    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        log.Fatal("Vui lòng set GOOGLE_API_KEY environment variable")
    }

    geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
        APIKey: apiKey,
    })
    if err != nil {
        log.Fatalf("Không thể tạo model: %v", err)
    }

    // Tạo Pipeline
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
```

---

## Phần 6: Cơ chế exitLoop chi tiết

### 6.1 tool.Context là gì?

`tool.Context` là context được truyền vào function tool, cung cấp:

```go
type Context interface {
    context.Context
    AgentName() string      // Tên agent đang gọi tool
    Actions() *EventActions // Actions để điều khiển flow
    // ... các methods khác
}
```

### 6.2 EventActions.Escalate

```go
type EventActions struct {
    StateDelta        map[string]any
    ArtifactDelta     map[string]int64
    SkipSummarization bool
    TransferToAgent   string
    Escalate          bool  // ← Khi = true, LoopAgent dừng
}
```

### 6.3 Luồng xử lý

```
Iteration 1:
  Critic → "Thiếu error handling cho n < 0"
  Refiner → đọc feedback, cải thiện code → {current_code}

Iteration 2:
  Critic → "Cần thêm comment cho logic phức tạp"
  Refiner → đọc feedback, thêm comment → {current_code}

Iteration 3:
  Critic → "CODE_APPROVED"
  Refiner → đọc "CODE_APPROVED" → gọi exitLoop()
           → ctx.Actions().Escalate = true
           → LoopAgent dừng ngay lập tức
```

---

## Phần 7: State Management

### 7.1 OutputKey và State

Mỗi agent có thể lưu output vào state thông qua `OutputKey`:

```go
llmagent.New(llmagent.Config{
    OutputKey: "current_code",  // Output được lưu vào state["current_code"]
    // ...
})
```

### 7.2 Đọc State trong Instruction

Sử dụng `{key}` syntax để đọc từ state:

```go
Instruction: `
**CODE CẦN REVIEW:**
{current_code}

**FEEDBACK:**
{critique}
`
```

### 7.3 State flow trong pipeline

```
User Input: "Viết hàm fibonacci"
    ↓
InitialProducer:
    → state["current_code"] = "func fibonacci(n int) int {...}"
    ↓
[Loop Iteration 1]
Critic:
    ← đọc state["current_code"]
    → state["critique"] = "Thiếu xử lý n < 0"
    ↓
Refiner:
    ← đọc state["current_code"], state["critique"]
    → state["current_code"] = "func fibonacci(n int) (int, error) {...}"
    ↓
[Loop Iteration 2]
Critic:
    ← đọc state["current_code"]
    → state["critique"] = "CODE_APPROVED"
    ↓
Refiner:
    ← đọc state["critique"] = "CODE_APPROVED"
    → gọi exitLoop() → STOP
```

---

## Phần 8: Các biến thể nâng cao

### 8.1 Multi-Critic với ParallelAgent

```go
// Nhiều critics chạy song song
parallelCritics, _ := parallelagent.New(parallelagent.Config{
    AgentConfig: agent.Config{
        Name:      "ParallelCritics",
        SubAgents: []agent.Agent{
            securityCritic,    // → state["security_critique"]
            performanceCritic, // → state["performance_critique"]
            styleCritic,       // → state["style_critique"]
        },
    },
})

// Aggregator tổng hợp feedback
aggregator, _ := llmagent.New(llmagent.Config{
    Instruction: `Tổng hợp feedback từ:
- Security: {security_critique}
- Performance: {performance_critique}
- Style: {style_critique}

Nếu TẤT CẢ đều approve → output "CODE_APPROVED"
Nếu có feedback → tổng hợp thành danh sách`,
    OutputKey: "critique",
})
```

### 8.2 Conditional Exit với Quality Score

```go
// ExitIfHighQuality chỉ exit nếu score >= threshold
func ExitIfHighQuality(ctx tool.Context, input QualityCheckArgs) (QualityCheckResults, error) {
    if input.Score >= 9.0 {
        fmt.Printf("[exitLoop] Quality score %.1f >= 9.0, exiting\n", input.Score)
        ctx.Actions().Escalate = true
        return QualityCheckResults{ShouldExit: true}, nil
    }
    return QualityCheckResults{ShouldExit: false}, nil
}

type QualityCheckArgs struct {
    Score float64 `json:"score" description:"Quality score from 0-10"`
}

type QualityCheckResults struct {
    ShouldExit bool `json:"should_exit"`
}
```

### 8.3 Exit với Reason Logging

```go
type ExitWithReasonArgs struct {
    Reason string `json:"reason" description:"Lý do exit loop"`
}

func ExitWithReason(ctx tool.Context, input ExitWithReasonArgs) (ExitLoopResults, error) {
    log.Printf("[exitLoop] Agent %s exiting. Reason: %s", ctx.AgentName(), input.Reason)
    ctx.Actions().Escalate = true
    return ExitLoopResults{}, nil
}
```

---

## Phần 9: Best Practices

### 9.1 Thiết kế Critic hiệu quả

```go
// ✅ Tốt: Approval phrase rõ ràng, không ambiguous
Instruction: `
NẾU code hoàn hảo:
Trả lời CHÍNH XÁC: CODE_APPROVED
KHÔNG thêm bất kỳ text nào khác.

NẾU có vấn đề:
Liệt kê cụ thể 1-2 điểm cần sửa.`

// ❌ Xấu: Approval phrase có thể bị nhầm
Instruction: `
Nếu tốt thì nói "code này tốt rồi" hoặc "approved"...`
```

### 9.2 Thiết kế Refiner với exitLoop

```go
// ✅ Tốt: Điều kiện gọi exitLoop rõ ràng
Instruction: fmt.Sprintf(`
NẾU feedback CHÍNH XÁC là "%s":
→ Gọi function 'exitLoop' NGAY LẬP TỨC
→ KHÔNG output bất kỳ text nào

NẾU feedback chứa gợi ý:
→ Cải thiện code
→ Output code đã cải thiện`, approvedPhrase)

// ❌ Xấu: Điều kiện mơ hồ
Instruction: `Nếu code tốt thì có thể gọi exitLoop...`
```

### 9.3 Logging cho debugging

```go
func ExitLoop(ctx tool.Context, input ExitLoopArgs) (ExitLoopResults, error) {
    // Log chi tiết để debug
    log.Printf("═══════════════════════════════════════")
    log.Printf("[exitLoop] TRIGGERED")
    log.Printf("[exitLoop] Agent: %s", ctx.AgentName())
    log.Printf("[exitLoop] Setting Escalate = true")
    log.Printf("═══════════════════════════════════════")

    ctx.Actions().Escalate = true
    return ExitLoopResults{}, nil
}
```

### 9.4 Xử lý edge cases

```go
// Đảm bảo MaxIterations hợp lý
loopagent.New(loopagent.Config{
    MaxIterations: 5, // Không quá nhỏ (có thể chưa đủ), không quá lớn (tốn token)
    // ...
})

// Trong Refiner: xử lý case critique rỗng
Instruction: `
NẾU feedback rỗng hoặc không rõ ràng:
→ Giữ nguyên code hiện tại
→ KHÔNG gọi exitLoop (để loop tiếp tục)`
```

---

## Phần 10: Code Mẫu Đầy Đủ (Solution)

Dưới đây là code hoàn chỉnh cho file `cmd/reflection/main.go`:

```go
// Package main demonstrates Reflection pattern using LoopAgent in Google ADK-Go
//
// Reflection cho phép agent tự đánh giá và cải thiện output qua nhiều vòng lặp.
// Sử dụng LoopAgent + functiontool để điều khiển vòng lặp:
// - InitialProducer: Tạo code ban đầu
// - Critic: Review và đưa ra feedback hoặc approve
// - Refiner: Cải thiện code hoặc gọi exitLoop nếu approved
// - LoopAgent: Lặp cho đến khi exitLoop() được gọi hoặc đạt max iterations
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

// ExitLoopArgs định nghĩa arguments cho exitLoop tool (rỗng)
type ExitLoopArgs struct{}

// ExitLoopResults định nghĩa kết quả trả về (rỗng)
type ExitLoopResults struct{}

// ExitLoop là function tool signal cho LoopAgent dừng
// Khi được gọi, set ctx.Actions().Escalate = true
func ExitLoop(ctx tool.Context, input ExitLoopArgs) (ExitLoopResults, error) {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Printf("[exitLoop] Code đã được phê duyệt!\n")
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

// ============================================================================
// CODE REFINER (với exitLoop tool)
// ============================================================================

func createCodeRefiner(m model.LLM, exitLoopTool tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "CodeRefiner",
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
   → KHÔNG giải thích, KHÔNG gọi exitLoop.`, stateCode, stateCritique, approvedPhrase),
		Tools:     []tool.Tool{exitLoopTool},
		OutputKey: stateCode,
	})
}

// ============================================================================
// REFLECTION PIPELINE
// ============================================================================

func createReflectionPipeline(m model.LLM) (agent.Agent, error) {
	// 1. Tạo Initial Producer (chạy 1 lần đầu)
	initialProducer, err := createInitialProducer(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create initial producer: %w", err)
	}

	// 2. Tạo exitLoop tool
	exitLoopTool, err := functiontool.New(
		functiontool.Config{
			Name:        "exitLoop",
			Description: "Gọi function này KHI VÀ CHỈ KHI critique CHÍNH XÁC là CODE_APPROVED. Khi gọi, KHÔNG output text.",
		},
		ExitLoop,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create exitLoop tool: %w", err)
	}

	// 3. Tạo Critic Agent (trong loop)
	critic, err := createCodeCritic(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create critic: %w", err)
	}

	// 4. Tạo Refiner Agent với exitLoop tool (trong loop)
	refiner, err := createCodeRefiner(m, exitLoopTool)
	if err != nil {
		return nil, fmt.Errorf("failed to create refiner: %w", err)
	}

	// 5. Tạo Refinement Loop
	refinementLoop, err := loopagent.New(loopagent.Config{
		MaxIterations: 3,
		AgentConfig: agent.Config{
			Name:        "RefinementLoop",
			Description: "Vòng lặp: Critic review → Refiner cải thiện hoặc exit",
			SubAgents:   []agent.Agent{critic, refiner},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create loop agent: %w", err)
	}

	// 6. Tạo Pipeline tổng thể
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "CodeReviewPipeline",
			Description: "Pipeline: InitialProducer → RefinementLoop",
			SubAgents:   []agent.Agent{initialProducer, refinementLoop},
		},
	})
}

// ============================================================================
// MAIN
// ============================================================================

func printBanner() {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  Code Review System - Reflection Pattern với LoopAgent")
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
	fmt.Println("Khởi động server...")
	fmt.Println("Ví dụ: 'Viết hàm tính fibonacci trong Go'")
}

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("Vui lòng set GOOGLE_API_KEY environment variable")
	}

	geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Không thể tạo model: %v", err)
	}

	// Tạo Pipeline
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
```

---

## Phần 11: Chạy thử và kiểm tra

### 11.1 Chuẩn bị

```bash
# Set API key
export GOOGLE_API_KEY="your-api-key"

# Chạy chương trình
go run cmd/reflection/main.go
```

### 11.2 Test cases

**Test 1: Simple Function**
```
User: Viết hàm tính giai thừa trong Go
Expected:
- InitialProducer tạo code v1
- Loop 1: Critic review → feedback → Refiner cải thiện
- Loop 2: Critic review → CODE_APPROVED → Refiner gọi exitLoop()
- Output: Code hoàn chỉnh
```

**Test 2: Complex Function**
```
User: Viết function validate email với regex
Expected:
- Nhiều vòng review hơn do logic phức tạp
- exitLoop được gọi khi code đạt chuẩn
```

### 11.3 Debugging

Xem logs để theo dõi:
```bash
go run cmd/reflection/main.go 2>&1 | grep "exitLoop"
```

Output mẫu:
```
═══════════════════════════════════════════════════════
[exitLoop] Code đã được phê duyệt!
[exitLoop] Triggered by agent: CodeRefiner
═══════════════════════════════════════════════════════
```

---

## Tài liệu tham khảo

1. [ADK-Go Multi-Agent Documentation](https://google.github.io/adk-docs/agents/multi-agents/)
2. [LoopAgent Documentation](https://google.github.io/adk-docs/agents/multi-agents/#loopagent)
3. [Function Tools Documentation](https://google.github.io/adk-docs/tools/)
4. [Training Language Models to Self-Correct via Reinforcement Learning](https://arxiv.org/abs/2409.12917)
5. [Chapter 4: Reflection - Agentic Design Patterns](../../doc_vi/04_Chapter_4_Reflection.md)
