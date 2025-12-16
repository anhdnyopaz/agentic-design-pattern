# Hướng dẫn Thực hành: Reflection Pattern với ADK-Go

## Mục tiêu

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu mẫu Reflection (Phản chiếu / Tự phê bình) trong thiết kế Agentic
- Triển khai mô hình Producer-Critic với ADK-Go
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

## Phần 2: Mô hình Producer-Critic

### 2.1 Tại sao tách Producer và Critic?

| Đặc điểm | Self-Reflection | Producer-Critic |
|----------|-----------------|-----------------|
| Khách quan | Thấp (bias) | Cao |
| Chuyên biệt | Chung chung | Mỗi agent tối ưu cho vai trò |
| Hiệu quả | Trung bình | Cao hơn |
| Debugging | Khó | Dễ theo dõi |

### 2.2 Vai trò của từng Agent

**Producer Agent:**
- Tập trung hoàn toàn vào việc **tạo nội dung**
- Nhận task ban đầu và feedback từ Critic
- Tạo version mới dựa trên phê bình

**Critic Agent:**
- Chuyên **đánh giá và phê bình**
- Có tiêu chí rõ ràng để đánh giá
- Cung cấp feedback có cấu trúc
- Xác định khi nào output đạt yêu cầu

---

## Phần 3: Bài tập - Code Review Agent

### Mô tả bài tập

Xây dựng một **Code Review System** với:

1. **Code Producer** - Viết code theo yêu cầu
2. **Code Critic** - Review code và đưa ra feedback
3. **Orchestrator** - Điều phối vòng lặp reflection

Hệ thống sẽ tự động cải thiện code qua nhiều vòng review.

### Cấu trúc file

```
cmd/reflection/
├── main.go
```

---

## Phần 4: Hướng dẫn từng bước

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
    "google.golang.org/adk/cmd/launcher"
    "google.golang.org/adk/cmd/launcher/full"
    "google.golang.org/adk/model"
    "google.golang.org/adk/model/gemini"
    "google.golang.org/adk/tool"
    "google.golang.org/adk/tool/agenttool"
    "google.golang.org/genai"
)
```

### Bước 2: Tạo Code Producer Agent

```go
func createCodeProducer(ctx context.Context, m model.LLM) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "code_producer",
        Model:       m,
        Description: "Chuyên gia viết code. Tạo code ban đầu hoặc cải thiện code dựa trên feedback.",
        Instruction: `Bạn là một lập trình viên chuyên nghiệp.

NHIỆM VỤ:
Khi nhận yêu cầu viết code hoặc feedback để cải thiện:

1. NẾU LÀ YÊU CẦU MỚI:
   - Phân tích yêu cầu kỹ lưỡng
   - Viết code sạch, có comment
   - Xử lý edge cases
   - Tuân thủ best practices

2. NẾU CÓ FEEDBACK:
   - Đọc kỹ từng điểm feedback
   - Sửa tất cả các vấn đề được chỉ ra
   - Giải thích những thay đổi đã làm

ĐỊNH DẠNG TRẢ VỀ:
` + "```" + `go
// Code ở đây
` + "```" + `

GIẢI THÍCH:
[Giải thích ngắn gọn về code/thay đổi]`,
    })
}
```

### Bước 3: Tạo Code Critic Agent

```go
func createCodeCritic(ctx context.Context, m model.LLM) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "code_critic",
        Model:       m,
        Description: "Chuyên gia review code. Đánh giá chất lượng code và đưa ra feedback chi tiết.",
        Instruction: `Bạn là một Senior Software Engineer với 15 năm kinh nghiệm.
Vai trò của bạn là thực hiện code review tỉ mỉ.

TIÊU CHÍ ĐÁNH GIÁ:
1. **Tính đúng đắn (Correctness):** Code có hoạt động đúng không?
2. **Xử lý lỗi (Error Handling):** Có xử lý các trường hợp lỗi không?
3. **Hiệu suất (Performance):** Có vấn đề về hiệu suất không?
4. **Khả năng đọc (Readability):** Code có dễ đọc, dễ hiểu không?
5. **Best Practices:** Code có tuân thủ best practices không?
6. **Edge Cases:** Có xử lý các trường hợp đặc biệt không?

QUY TRÌNH ĐÁNH GIÁ:
1. Đọc và hiểu code được cung cấp
2. Kiểm tra từng tiêu chí trên
3. Liệt kê các vấn đề cần sửa
4. Nếu code đã hoàn hảo, trả lời "CODE_APPROVED"

ĐỊNH DẠNG PHẢN HỒI:

**NẾU CÓ VẤN ĐỀ:**
📋 CODE REVIEW REPORT
━━━━━━━━━━━━━━━━━━━━━

🔴 VẤN ĐỀ NGHIÊM TRỌNG:
• [Vấn đề 1 - giải thích và cách sửa]
• [Vấn đề 2 - giải thích và cách sửa]

🟡 CẦN CẢI THIỆN:
• [Điểm cần cải thiện 1]
• [Điểm cần cải thiện 2]

🟢 ĐIỂM TỐT:
• [Điểm tốt của code]

📊 ĐÁNH GIÁ TỔNG QUAN: [X]/10
━━━━━━━━━━━━━━━━━━━━━

**NẾU CODE ĐÃ HOÀN HẢO:**
✅ CODE_APPROVED

Lý do: [Giải thích ngắn gọn tại sao code đạt yêu cầu]`,
    })
}
```

### Bước 4: Tạo Reflection Orchestrator

```go
func createReflectionOrchestrator(ctx context.Context, m model.LLM, producer, critic agent.Agent) (agent.Agent, error) {
    // Wrap Producer và Critic thành tools
    producerTool := agenttool.New(producer, nil)
    criticTool := agenttool.New(critic, nil)

    return llmagent.New(llmagent.Config{
        Name:        "reflection_orchestrator",
        Model:       m,
        Description: "Điều phối viên Code Review - Thực hiện Reflection Loop",
        Instruction: `Bạn là điều phối viên hệ thống Code Review tự động.

**QUY TRÌNH REFLECTION (BẮT BUỘC TUÂN THỦ):**

Khi người dùng yêu cầu viết code, thực hiện VÒNG LẶP sau:

BƯỚC 1: TẠO CODE
- Gọi code_producer với yêu cầu của người dùng
- Lưu code được tạo

BƯỚC 2: REVIEW CODE
- Gọi code_critic để review code vừa tạo
- Đọc kỹ feedback

BƯỚC 3: KIỂM TRA KẾT QUẢ
- NẾU critic trả về "CODE_APPROVED" → Chuyển sang BƯỚC 4
- NẾU critic có feedback → Quay lại BƯỚC 1 với feedback

BƯỚC 4: TRÌNH BÀY KẾT QUẢ
Hiển thị cho người dùng:

═══════════════════════════════════════════════════════
🔄 QUÁ TRÌNH REFLECTION HOÀN TẤT
═══════════════════════════════════════════════════════

📊 THỐNG KÊ:
• Số vòng lặp: [X]
• Trạng thái: ✅ Đã được phê duyệt

📝 CODE CUỐI CÙNG:
[Code đã được approve]

📋 LỊCH SỬ CẢI TIẾN:
• Vòng 1: [Tóm tắt thay đổi]
• Vòng 2: [Tóm tắt thay đổi]
...

═══════════════════════════════════════════════════════

**GIỚI HẠN:**
- Tối đa 3 vòng lặp
- Nếu sau 3 vòng vẫn chưa approve, trả về code tốt nhất với ghi chú

**KHI NGƯỜI DÙNG CHÀO HỎI:**
Giới thiệu bản thân:
"Xin chào! Tôi là Code Review System với khả năng Reflection.
Tôi sẽ viết code cho bạn và tự động review, cải thiện qua nhiều vòng.

Ví dụ yêu cầu:
• 'Viết hàm tính fibonacci trong Go'
• 'Tạo REST API handler cho user registration'
• 'Implement binary search tree'"`,
        Tools: []tool.Tool{producerTool, criticTool},
    })
}
```

### Bước 5: Main function

```go
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

    // Tạo Producer và Critic agents
    producer, err := createCodeProducer(ctx, geminiModel)
    if err != nil {
        log.Fatal(err)
    }

    critic, err := createCodeCritic(ctx, geminiModel)
    if err != nil {
        log.Fatal(err)
    }

    // Tạo Reflection Orchestrator
    orchestrator, err := createReflectionOrchestrator(ctx, geminiModel, producer, critic)
    if err != nil {
        log.Fatal(err)
    }

    config := &launcher.Config{
        AgentLoader: agent.NewSingleLoader(orchestrator),
    }

    lch := full.NewLauncher()
    fmt.Println("=== Code Review System - Reflection Demo ===")
    fmt.Println("Code sẽ được tự động review và cải thiện qua nhiều vòng")
    fmt.Println("Khởi động server...")

    err = lch.Execute(ctx, config, os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## Phần 5: Các biến thể Reflection

### 5.1 Multi-Critic Reflection

Sử dụng nhiều Critic với các chuyên môn khác nhau:

```go
// Security Critic - Tập trung vào bảo mật
func createSecurityCritic(ctx context.Context, m model.LLM) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "security_critic",
        Model:       m,
        Description: "Chuyên gia bảo mật - Review code về các lỗ hổng bảo mật",
        Instruction: `Bạn là Security Expert. Đánh giá code về:
- SQL Injection
- XSS vulnerabilities
- Input validation
- Authentication/Authorization issues
- Sensitive data exposure
- OWASP Top 10

Trả về "SECURITY_APPROVED" nếu không có vấn đề bảo mật.`,
    })
}

// Performance Critic - Tập trung vào hiệu suất
func createPerformanceCritic(ctx context.Context, m model.LLM) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "performance_critic",
        Model:       m,
        Description: "Chuyên gia hiệu suất - Review code về performance",
        Instruction: `Bạn là Performance Expert. Đánh giá code về:
- Time complexity (Big O)
- Space complexity
- Memory leaks
- Unnecessary allocations
- Database query optimization
- Caching opportunities

Trả về "PERFORMANCE_APPROVED" nếu hiệu suất tốt.`,
    })
}
```

### 5.2 Hierarchical Reflection

Reflection nhiều cấp độ:

```
                    ┌──────────────────┐
                    │ FINAL REVIEWER   │
                    │ (Quality Gate)   │
                    └────────┬─────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
        ┌──────────┐  ┌──────────┐  ┌──────────┐
        │ Security │  │ Quality  │  │ Perf     │
        │ Review   │  │ Review   │  │ Review   │
        └──────────┘  └──────────┘  └──────────┘
              │              │              │
              └──────────────┼──────────────┘
                             ▼
                    ┌──────────────────┐
                    │    PRODUCER      │
                    └──────────────────┘
```

### 5.3 Adaptive Reflection

Điều chỉnh số vòng lặp dựa trên độ phức tạp:

```go
type ReflectionConfig struct {
    MaxIterations   int
    QualityThreshold float64
    AdaptiveMode    bool
}

func determineIterations(taskComplexity string) int {
    switch taskComplexity {
    case "simple":
        return 1
    case "medium":
        return 2
    case "complex":
        return 3
    default:
        return 2
    }
}
```

---

## Phần 6: Xử lý Stopping Conditions

### 6.1 Các điều kiện dừng phổ biến

```go
type StoppingCondition int

const (
    QualityMet       StoppingCondition = iota // Chất lượng đạt yêu cầu
    MaxIterations                              // Đạt số vòng tối đa
    NoImprovement                              // Không có cải thiện
    UserInterrupt                              // Người dùng yêu cầu dừng
    Timeout                                    // Hết thời gian
)

func shouldStop(iteration int, config ReflectionConfig, improvement float64) (bool, StoppingCondition) {
    if improvement >= config.QualityThreshold {
        return true, QualityMet
    }
    if iteration >= config.MaxIterations {
        return true, MaxIterations
    }
    if improvement <= 0 {
        return true, NoImprovement
    }
    return false, -1
}
```

### 6.2 Quality Scoring

```go
type QualityScore struct {
    Correctness   float64 `json:"correctness"`    // 0-10
    ErrorHandling float64 `json:"error_handling"` // 0-10
    Performance   float64 `json:"performance"`    // 0-10
    Readability   float64 `json:"readability"`    // 0-10
    BestPractices float64 `json:"best_practices"` // 0-10
}

func (q QualityScore) Overall() float64 {
    return (q.Correctness*0.3 + q.ErrorHandling*0.2 +
            q.Performance*0.2 + q.Readability*0.15 +
            q.BestPractices*0.15)
}

func (q QualityScore) MeetsThreshold(threshold float64) bool {
    return q.Overall() >= threshold
}
```

---

## Phần 7: Bài tập mở rộng

### Bài tập 1: Content Writer với Reflection

Xây dựng hệ thống viết blog với:
- **Writer Agent**: Viết nội dung
- **Editor Agent**: Chỉnh sửa grammar, flow
- **SEO Critic**: Tối ưu SEO
- **Fact Checker**: Kiểm tra tính chính xác

### Bài tập 2: Test-Driven Reflection

Tích hợp test execution vào vòng lặp:
```
Producer → Code → Run Tests → Critic → Feedback → Producer
```

### Bài tập 3: Human-in-the-Loop Reflection

Thêm khả năng nhận feedback từ người dùng:
- Sau mỗi vòng, hỏi người dùng có feedback thêm không
- Kết hợp feedback của Critic và người dùng

### Bài tập 4: Reflection với Memory

Lưu lại các patterns lỗi phổ biến:
- Nếu Producer lặp lại lỗi cũ, Critic nhắc lại mạnh hơn
- Học từ các lần review trước

---

## Phần 8: Best Practices

### 8.1 Khi nào KHÔNG nên dùng Reflection

- Task đơn giản, một bước
- Cần response nhanh (real-time)
- Chi phí API là concern chính
- Output không cần chất lượng cao

### 8.2 Thiết kế Critic hiệu quả

1. **Tiêu chí rõ ràng:** Liệt kê cụ thể những gì cần đánh giá
2. **Actionable feedback:** Feedback phải có thể thực hiện được
3. **Structured output:** Format feedback nhất quán
4. **Clear approval signal:** Định nghĩa rõ khi nào đạt yêu cầu

### 8.3 Tránh Infinite Loops

```go
// Anti-pattern: Không có giới hạn
for !approved {
    // Có thể chạy mãi
}

// Best practice: Luôn có giới hạn
for i := 0; i < maxIterations && !approved; i++ {
    // An toàn
}
```

### 8.4 Logging và Debugging

```go
type ReflectionLog struct {
    Iteration    int
    ProducerOutput string
    CriticFeedback string
    Timestamp    time.Time
    Duration     time.Duration
}

func logIteration(log ReflectionLog) {
    fmt.Printf("[Iteration %d] Duration: %v\n", log.Iteration, log.Duration)
    fmt.Printf("Producer: %s\n", truncate(log.ProducerOutput, 100))
    fmt.Printf("Critic: %s\n", truncate(log.CriticFeedback, 100))
}
```

---

## Phần 9: Code Mẫu Đầy Đủ (Solution)

Dưới đây là code hoàn chỉnh cho file `cmd/reflection/main.go`:

```go
// Package main demonstrates Reflection pattern using Google ADK-Go
//
// Reflection cho phép agent tự đánh giá và cải thiện output qua nhiều vòng lặp.
// Sử dụng mô hình Producer-Critic:
// - Producer: Tạo code
// - Critic: Review và đưa ra feedback
// - Orchestrator: Điều phối vòng lặp cho đến khi đạt chất lượng
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
	"google.golang.org/genai"
)

// ============================================================================
// PRODUCER AGENT - Tạo code
// ============================================================================

func createCodeProducer(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "code_producer",
		Model:       m,
		Description: "Chuyên gia viết code. Tạo code ban đầu hoặc cải thiện code dựa trên feedback từ reviewer.",
		Instruction: `Bạn là một Senior Software Engineer chuyên viết code Go.

NHIỆM VỤ:
Khi nhận yêu cầu viết code hoặc feedback để cải thiện:

1. NẾU LÀ YÊU CẦU MỚI:
   - Phân tích yêu cầu kỹ lưỡng
   - Viết code sạch, có comment giải thích
   - Xử lý tất cả edge cases
   - Tuân thủ Go best practices và idioms
   - Thêm error handling đầy đủ

2. NẾU CÓ FEEDBACK TỪ REVIEWER:
   - Đọc kỹ TỪNG điểm feedback
   - Sửa TẤT CẢ các vấn đề được chỉ ra
   - Không bỏ sót bất kỳ feedback nào
   - Giải thích những thay đổi đã làm

ĐỊNH DẠNG TRẢ VỀ:

📝 CODE:
` + "```go" + `
// Code của bạn ở đây
// Phải có comment giải thích logic phức tạp
` + "```" + `

📌 GIẢI THÍCH:
[Giải thích ngắn gọn về code hoặc các thay đổi đã thực hiện]

⚠️ LƯU Ý:
[Các điểm cần lưu ý khi sử dụng code này]`,
	})
}

// ============================================================================
// CRITIC AGENT - Review code
// ============================================================================

func createCodeCritic(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "code_critic",
		Model:       m,
		Description: "Senior Code Reviewer. Đánh giá chất lượng code và đưa ra feedback chi tiết, có thể hành động được.",
		Instruction: `Bạn là một Principal Engineer với 20 năm kinh nghiệm.
Vai trò của bạn là thực hiện code review tỉ mỉ và KHÁCH QUAN.

TIÊU CHÍ ĐÁNH GIÁ (theo thứ tự ưu tiên):

1. 🔴 CORRECTNESS (Tính đúng đắn) - QUAN TRỌNG NHẤT
   - Code có hoạt động đúng với mọi input không?
   - Logic có chính xác không?
   - Có bug tiềm ẩn không?

2. 🟠 ERROR HANDLING (Xử lý lỗi)
   - Có xử lý tất cả các lỗi có thể xảy ra không?
   - Error messages có rõ ràng không?
   - Có return error thay vì panic không?

3. 🟡 EDGE CASES (Trường hợp đặc biệt)
   - Empty input, nil values
   - Boundary conditions (0, negative, max values)
   - Concurrent access (nếu applicable)

4. 🟢 CODE QUALITY (Chất lượng code)
   - Naming conventions (Go idioms)
   - Code organization
   - Comments và documentation
   - DRY principle

5. 🔵 PERFORMANCE (Hiệu suất)
   - Time complexity
   - Space complexity
   - Unnecessary allocations

QUY TRÌNH ĐÁNH GIÁ:
1. Đọc và hiểu TOÀN BỘ code
2. Kiểm tra TỪNG tiêu chí ở trên
3. Liệt kê CỤ THỂ các vấn đề
4. Đưa ra cách sửa CHI TIẾT

QUAN TRỌNG:
- Nếu code ĐÃ ĐẠT TẤT CẢ tiêu chí: Trả lời CHÍNH XÁC "✅ CODE_APPROVED"
- Nếu CÒN vấn đề: Liệt kê CHI TIẾT để developer sửa được

ĐỊNH DẠNG PHẢN HỒI:

═══════════════════════════════════════════════════════
📋 CODE REVIEW REPORT
═══════════════════════════════════════════════════════

🔴 VẤN ĐỀ NGHIÊM TRỌNG (phải sửa):
1. [Vấn đề]: [Mô tả]
   → Cách sửa: [Hướng dẫn cụ thể]

🟡 CẦN CẢI THIỆN (nên sửa):
1. [Điểm cải thiện]: [Mô tả]
   → Gợi ý: [Hướng dẫn]

🟢 ĐIỂM TỐT:
• [Những gì code đã làm tốt]

📊 ĐIỂM ĐÁNH GIÁ: [X]/10
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

HOẶC NẾU CODE HOÀN HẢO:

═══════════════════════════════════════════════════════
✅ CODE_APPROVED

Lý do phê duyệt:
• [Điểm tốt 1]
• [Điểm tốt 2]
• [Điểm tốt 3]

📊 ĐIỂM ĐÁNH GIÁ: 10/10
═══════════════════════════════════════════════════════`,
	})
}

// ============================================================================
// REFLECTION ORCHESTRATOR - Điều phối vòng lặp
// ============================================================================

func createReflectionOrchestrator(ctx context.Context, m model.LLM, producer, critic agent.Agent) (agent.Agent, error) {
	// Wrap Producer và Critic thành tools
	producerTool := agenttool.New(producer, nil)
	criticTool := agenttool.New(critic, nil)

	return llmagent.New(llmagent.Config{
		Name:        "reflection_orchestrator",
		Model:       m,
		Description: "Điều phối viên Code Review System - Thực hiện Reflection Loop để cải thiện code",
		Instruction: `Bạn là điều phối viên hệ thống Code Review tự động với khả năng REFLECTION.

**QUY TRÌNH REFLECTION (TUÂN THỦ NGHIÊM NGẶT):**

Khi người dùng yêu cầu viết code, thực hiện VÒNG LẶP sau:

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
VÒNG LẶP 1:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
BƯỚC 1.1: Gọi code_producer với yêu cầu của người dùng
BƯỚC 1.2: Gọi code_critic để review code vừa tạo
BƯỚC 1.3: Kiểm tra kết quả:
          - Nếu critic trả về "CODE_APPROVED" → Kết thúc
          - Nếu có feedback → Tiếp tục VÒNG LẶP 2

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
VÒNG LẶP 2 (nếu cần):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
BƯỚC 2.1: Gọi code_producer với feedback từ vòng 1
BƯỚC 2.2: Gọi code_critic để review code cải tiến
BƯỚC 2.3: Kiểm tra kết quả:
          - Nếu critic trả về "CODE_APPROVED" → Kết thúc
          - Nếu có feedback → Tiếp tục VÒNG LẶP 3

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
VÒNG LẶP 3 (cuối cùng):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
BƯỚC 3.1: Gọi code_producer với feedback từ vòng 2
BƯỚC 3.2: Gọi code_critic lần cuối
BƯỚC 3.3: Dừng lại dù kết quả thế nào

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
KẾT THÚC - TRÌNH BÀY KẾT QUẢ:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

╔═══════════════════════════════════════════════════════════════╗
║           🔄 REFLECTION PROCESS COMPLETED                     ║
╠═══════════════════════════════════════════════════════════════╣
║                                                               ║
║ 📊 THỐNG KÊ:                                                  ║
║ • Số vòng lặp: [X]                                           ║
║ • Trạng thái: [✅ Approved / ⚠️ Best effort sau 3 vòng]      ║
║                                                               ║
║ 📝 CODE CUỐI CÙNG:                                           ║
║ [Code đã được approve hoặc version tốt nhất]                  ║
║                                                               ║
║ 📋 LỊCH SỬ CẢI TIẾN:                                         ║
║ • Vòng 1: [Tóm tắt feedback và thay đổi]                     ║
║ • Vòng 2: [Tóm tắt feedback và thay đổi]                     ║
║ • Vòng 3: [Kết quả cuối]                                     ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝

**KHI NGƯỜI DÙNG CHÀO HỎI HOẶC HỎI THÔNG TIN:**
Giới thiệu hệ thống:

"Xin chào! 👋 Tôi là Code Review System với khả năng Reflection.

🔄 Cách hoạt động:
1. Bạn yêu cầu → Tôi viết code
2. Code được tự động review
3. Nếu có vấn đề → Tự động sửa và review lại
4. Lặp lại đến khi code hoàn hảo (tối đa 3 vòng)

📝 Ví dụ yêu cầu:
• 'Viết hàm tính fibonacci trong Go'
• 'Tạo function validate email'
• 'Implement stack data structure'
• 'Viết HTTP handler cho user registration'

Bạn muốn tôi viết code gì?"

**LƯU Ý QUAN TRỌNG:**
- Mỗi vòng phải gọi CẢ producer VÀ critic
- Không bỏ qua bất kỳ vòng nào khi còn feedback
- Tối đa 3 vòng - sau đó trả về kết quả tốt nhất`,
		Tools: []tool.Tool{producerTool, criticTool},
	})
}

// ============================================================================
// MAIN
// ============================================================================

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

	// Tạo Producer Agent
	producer, err := createCodeProducer(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	// Tạo Critic Agent
	critic, err := createCodeCritic(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	// Tạo Reflection Orchestrator
	orchestrator, err := createReflectionOrchestrator(ctx, geminiModel, producer, critic)
	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(orchestrator),
	}

	lch := full.NewLauncher()
	fmt.Println("=== Code Review System - Reflection Pattern Demo ===")
	fmt.Println("Code sẽ được tự động review và cải thiện qua nhiều vòng lặp")
	fmt.Println("Producer → Code → Critic → Feedback → Producer → ...")
	fmt.Println("Khởi động server...")

	err = lch.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
```

---

## Phần 10: Chạy thử và kiểm tra

### 10.1 Chuẩn bị

```bash
# Set API key
export GOOGLE_API_KEY="your-api-key"

# Chạy chương trình
go run cmd/reflection/main.go
```

### 10.2 Test cases

**Test 1: Greeting**
```
User: Xin chào
Expected: Agent giới thiệu hệ thống Reflection
```

**Test 2: Simple Function**
```
User: Viết hàm tính giai thừa trong Go
Expected:
- Vòng 1: Code ban đầu + Review
- Vòng 2 (nếu cần): Code cải tiến + Review
- Kết quả: Code được approve với đầy đủ error handling
```

**Test 3: Complex Task**
```
User: Viết function validate email với regex trong Go
Expected:
- Nhiều vòng review
- Xử lý edge cases
- Code cuối cùng robust
```

---

## Tài liệu tham khảo

1. [ADK-Go Multi-Agent Documentation](https://google.github.io/adk-docs/agents/multi-agents/)
2. [Training Language Models to Self-Correct via Reinforcement Learning](https://arxiv.org/abs/2409.12917)
3. [Chapter 4: Reflection - Agentic Design Patterns](../doc_vi/04_Chapter_4_Reflection.md)
