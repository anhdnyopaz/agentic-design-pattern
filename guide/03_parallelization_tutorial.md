# Hướng dẫn Thực hành: Parallelization Pattern với ADK-Go

## Mục tiêu

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu mẫu Parallelization trong thiết kế Agentic
- Phân biệt giữa xử lý tuần tự và xử lý song song
- Triển khai Parallelization với ADK-Go sử dụng goroutines
- Xây dựng hệ thống phân tích đa nguồn với xử lý đồng thời

---

## Phần 1: Giới thiệu Parallelization Pattern

### 1.1 Parallelization là gì?

**Parallelization (Song song hóa)** là kỹ thuật thực thi nhiều tác vụ độc lập **cùng lúc** thay vì tuần tự. Điều này giúp giảm đáng kể thời gian thực thi tổng thể.

### 1.2 So sánh Tuần tự vs Song song

**Cách tiếp cận tuần tự:**
```
Bắt đầu → Tác vụ A (2s) → Tác vụ B (2s) → Tác vụ C (2s) → Tổng hợp → Kết thúc
Tổng thời gian: ~6 giây
```

**Cách tiếp cận song song:**
```
           ┌→ Tác vụ A (2s) ─┐
Bắt đầu ───┼→ Tác vụ B (2s) ─┼→ Tổng hợp → Kết thúc
           └→ Tác vụ C (2s) ─┘
Tổng thời gian: ~2 giây (+ thời gian tổng hợp)
```

### 1.3 Khi nào sử dụng Parallelization?

- Các tác vụ **không phụ thuộc** vào output của nhau
- Cần thu thập thông tin từ **nhiều nguồn** độc lập
- Muốn **giảm độ trễ** khi gọi nhiều API/dịch vụ bên ngoài
- Cần xử lý **nhiều phân đoạn dữ liệu** khác nhau

---

## Phần 2: Kiến trúc Parallelization trong ADK-Go

### 2.1 Mô hình Orchestrator + Parallel Workers

Trong ADK-Go, chúng ta có thể triển khai Parallelization theo hai cách:

**Cách 1: Go-level Parallelization (Goroutines)**
```
                    ┌→ [Goroutine 1: Agent A] ─┐
[User Input] → [Go] ┼→ [Goroutine 2: Agent B] ─┼→ [Aggregate] → [Response]
                    └→ [Goroutine 3: Agent C] ─┘
```

**Cách 2: LLM-orchestrated Parallelization**
```
                              ┌→ [Sub-Agent A (Tool)] ─┐
[User Input] → [Orchestrator] ┼→ [Sub-Agent B (Tool)] ─┼→ [Synthesize] → [Response]
                              └→ [Sub-Agent C (Tool)] ─┘
```

### 2.2 Ưu nhược điểm của từng cách

| Tiêu chí | Go-level | LLM-orchestrated |
|----------|----------|------------------|
| Tốc độ | Nhanh hơn | Chậm hơn (nhiều LLM calls) |
| Độ phức tạp | Phức tạp hơn | Đơn giản hơn |
| Linh hoạt | Cố định | LLM quyết định |
| Chi phí | Thấp hơn | Cao hơn |

### 2.3 Kiến trúc với Real-time Search (Decentralized)

Để các Analyst có khả năng tìm kiếm thông tin **thực** từ internet, chúng ta sử dụng **Google Search Tool** với kiến trúc **Decentralized Search**:

```
┌─────────────────────────────────────────────────────────────────┐
│                      ORCHESTRATOR                               │
│              (Điều phối các Analysts)                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │   Financial     │  │    Market       │  │     Risk        │ │
│  │    Analyst      │  │    Analyst      │  │    Analyst      │ │
│  │       │         │  │       │         │  │       │         │ │
│  │       ▼         │  │       ▼         │  │       ▼         │ │
│  │ ┌───────────┐   │  │ ┌───────────┐   │  │ ┌───────────┐   │ │
│  │ │  Search   │   │  │ │  Search   │   │  │ │  Search   │   │ │
│  │ │  Agent    │   │  │ │  Agent    │   │  │ │  Agent    │   │ │
│  │ │(as Tool)  │   │  │ │(as Tool)  │   │  │ │(as Tool)  │   │ │
│  │ └───────────┘   │  │ └───────────┘   │  │ └───────────┘   │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘ │
│                                                                 │
│     Mỗi Analyst TỰ TÌM KIẾM thông tin riêng họ cần             │
│     → Chạy SONG SONG thực sự với DỮ LIỆU THẬT                  │
└─────────────────────────────────────────────────────────────────┘
```

**Lợi ích của Decentralized Search:**
- Mỗi analyst search đúng thông tin họ cần (financial news, market data, risk reports...)
- Chạy song song thực sự - không cần đợi search xong mới phân tích
- Tự chủ hơn, giống như team nghiên cứu thật

**⚠️ Lưu ý quan trọng về Google Search Tool:**
- `geminitool.GoogleSearch{}` là built-in tool của ADK-Go
- **KHÔNG THỂ** kết hợp trực tiếp với các function tools khác trong cùng 1 agent
- Giải pháp: Tạo Search Agent riêng, wrap thành tool cho các Analysts

---

## Phần 3: Bài tập - Xây dựng Research Agent với Real Search

### Mô tả bài tập

Xây dựng một **Research Agent** phân tích công ty từ nhiều góc độ song song với **khả năng tìm kiếm thông tin thật** từ internet:

1. **Search Agent** - Agent chuyên tìm kiếm với Google Search (shared)
2. **Financial Analyst** - Phân tích tài chính (có thể search)
3. **Market Analyst** - Phân tích thị trường (có thể search)
4. **Risk Analyst** - Đánh giá rủi ro (có thể search)

Mỗi analyst có thể **tự tìm kiếm** thông tin họ cần và chạy **song song**.

### Cấu trúc file

```
cmd/parallelization/
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
    "google.golang.org/adk/tool/geminitool"  // ← Built-in tools (Google Search)
    "google.golang.org/genai"
)
```

**Lưu ý quan trọng:**
- `geminitool` - Chứa các built-in tools của Gemini như **Google Search**
- `agenttool` - Để wrap Agent thành Tool

### Bước 2: Tạo Search Agent (Shared)

Đây là agent **dùng chung** cho tất cả analysts, có khả năng tìm kiếm Google:

```go
// ============================================================================
// SEARCH AGENT - Agent chuyên tìm kiếm với Google Search
// ============================================================================

func createSearchAgent(ctx context.Context, m model.LLM) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "web_search",
        Model:       m,
        Description: "Tìm kiếm thông tin từ internet. Cung cấp query và nhận kết quả tìm kiếm mới nhất.",
        Instruction: `Bạn là chuyên gia tìm kiếm thông tin.

NHIỆM VỤ:
Khi nhận yêu cầu tìm kiếm, hãy:
1. Tìm kiếm thông tin liên quan
2. Trả về kết quả ngắn gọn, có nguồn
3. Ưu tiên thông tin mới nhất

ĐỊNH DẠNG TRẢ VỀ:
🔍 KẾT QUẢ TÌM KIẾM:
[Thông tin tìm được, có ghi nguồn nếu có]`,
        Tools: []tool.Tool{
            geminitool.GoogleSearch{},  // ← Built-in Google Search
        },
    })
}
```

**⚠️ Quan trọng:**
- `geminitool.GoogleSearch{}` chỉ hoạt động với **Gemini 2.x** trở lên
- Agent này **chỉ có** Google Search tool, không có tools khác

### Bước 3: Tạo các Specialist Agents (với Search capability)

Mỗi analyst nhận `searchAgent` làm tool để có thể tự tìm kiếm thông tin:

**Financial Analyst Agent:**

```go
// ============================================================================
// SPECIALIST AGENTS - Có khả năng tự tìm kiếm thông tin
// ============================================================================

func createFinancialAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
    // Wrap Search Agent thành tool cho analyst này
    searchTool := agenttool.New(searchAgent, nil)

    return llmagent.New(llmagent.Config{
        Name:        "financial_analyst",
        Model:       m,
        Description: "Chuyên gia phân tích tài chính. Có khả năng tìm kiếm thông tin tài chính mới nhất.",
        Instruction: `Bạn là chuyên gia phân tích tài chính với 20 năm kinh nghiệm.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin tài chính MỚI NHẤT về công ty
2. Phân tích dữ liệu tìm được
3. Đưa ra đánh giá

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] financial report 2024"
- "[Tên công ty] revenue profit quarterly"
- "[Tên công ty] stock price analysis"
- "[Tên công ty] debt ratio"

**NHIỆM VỤ PHÂN TÍCH:**
- Doanh thu và tăng trưởng
- Biên lợi nhuận
- Dòng tiền và thanh khoản
- Cấu trúc nợ

**ĐỊNH DẠNG TRẢ VỀ:**
📊 PHÂN TÍCH TÀI CHÍNH
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Điểm mạnh: [liệt kê]
• Điểm yếu: [liệt kê]
• Đánh giá: [1-10]/10
• Khuyến nghị: [Mua/Giữ/Bán]`,
        Tools: []tool.Tool{searchTool},  // ← Search Agent as tool
    })
}
```

**Market Analyst Agent:**

```go
func createMarketAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
    searchTool := agenttool.New(searchAgent, nil)

    return llmagent.New(llmagent.Config{
        Name:        "market_analyst",
        Model:       m,
        Description: "Chuyên gia phân tích thị trường. Có khả năng tìm kiếm thông tin thị trường mới nhất.",
        Instruction: `Bạn là chuyên gia phân tích thị trường với kiến thức sâu rộng.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin thị trường MỚI NHẤT
2. Phân tích vị thế cạnh tranh
3. Đánh giá xu hướng ngành

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] market share 2024"
- "[Tên công ty] competitors analysis"
- "[Industry] market trends"
- "[Tên công ty] expansion news"

**NHIỆM VỤ PHÂN TÍCH:**
- Thị phần và vị thế
- Đối thủ cạnh tranh
- Xu hướng ngành
- Cơ hội mở rộng

**ĐỊNH DẠNG TRẢ VỀ:**
🏆 PHÂN TÍCH THỊ TRƯỜNG
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Vị thế: [Dẫn đầu/Top 3/Trung bình/Theo sau]
• Đối thủ chính: [danh sách]
• Xu hướng ngành: [Tăng/Ổn định/Giảm]
• Tiềm năng: [Cao/Trung bình/Thấp]`,
        Tools: []tool.Tool{searchTool},
    })
}
```

**Risk Analyst Agent:**

```go
func createRiskAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
    searchTool := agenttool.New(searchAgent, nil)

    return llmagent.New(llmagent.Config{
        Name:        "risk_analyst",
        Model:       m,
        Description: "Chuyên gia đánh giá rủi ro. Có khả năng tìm kiếm tin tức và thông tin rủi ro.",
        Instruction: `Bạn là chuyên gia quản lý rủi ro với kinh nghiệm đánh giá doanh nghiệp.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin về rủi ro và tin tức tiêu cực
2. Đánh giá các loại rủi ro
3. Đề xuất biện pháp giảm thiểu

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] risks challenges"
- "[Tên công ty] lawsuit legal issues"
- "[Tên công ty] controversy scandal"
- "[Industry] regulatory risks"

**NHIỆM VỤ PHÂN TÍCH:**
- Rủi ro hoạt động
- Rủi ro pháp lý
- Rủi ro thị trường
- Rủi ro danh tiếng

**ĐỊNH DẠNG TRẢ VỀ:**
⚠️ ĐÁNH GIÁ RỦI RO
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Rủi ro CAO: [liệt kê nếu có]
• Rủi ro TRUNG BÌNH: [liệt kê]
• Rủi ro THẤP: [liệt kê]
• Điểm rủi ro: [1-10]/10
• Biện pháp giảm thiểu: [khuyến nghị]`,
        Tools: []tool.Tool{searchTool},
    })
}
```

### Bước 4: Tạo Orchestrator Agent

Orchestrator điều phối các Analysts - mỗi analyst sẽ tự search khi cần:

```go
// ============================================================================
// ORCHESTRATOR AGENT - Điều phối phân tích song song
// ============================================================================

func createResearchOrchestrator(ctx context.Context, m model.LLM, analysts []agent.Agent) (agent.Agent, error) {
    // Wrap các analyst agents thành tools
    var analystTools []tool.Tool
    for _, analyst := range analysts {
        analystTools = append(analystTools, agenttool.New(analyst, nil))
    }

    return llmagent.New(llmagent.Config{
        Name:        "research_orchestrator",
        Model:       m,
        Description: "Trưởng nhóm nghiên cứu đầu tư - Điều phối phân tích song song với dữ liệu thực",
        Instruction: `Bạn là trưởng nhóm nghiên cứu đầu tư chuyên nghiệp.

**ĐỘI NGŨ CỦA BẠN:**
Mỗi analyst có khả năng TỰ TÌM KIẾM thông tin từ internet:
- financial_analyst: Phân tích tài chính (tự search financial data)
- market_analyst: Phân tích thị trường (tự search market data)
- risk_analyst: Đánh giá rủi ro (tự search risk news)

**QUY TRÌNH KHI NHẬN YÊU CẦU PHÂN TÍCH:**

BƯỚC 1: PHÂN TÍCH SONG SONG
Gọi TẤT CẢ analysts CÙNG LÚC với tên công ty.
Mỗi analyst sẽ:
1. Tự tìm kiếm thông tin họ cần từ internet
2. Phân tích dữ liệu tìm được
3. Trả về kết quả

QUAN TRỌNG: Gọi cả 3 agent trong CÙNG MỘT LƯỢT (parallel execution)

BƯỚC 2: TỔNG HỢP BÁO CÁO
Sau khi nhận đủ kết quả (có dữ liệu thực từ search), tổng hợp:

═══════════════════════════════════════════════════════
📋 BÁO CÁO NGHIÊN CỨU: [TÊN CÔNG TY]
(Dựa trên dữ liệu thực từ internet)
═══════════════════════════════════════════════════════

📌 TÓM TẮT ĐIỀU HÀNH
[3-4 câu tóm tắt điểm quan trọng nhất - dựa trên data thực]

📊 PHÂN TÍCH TÀI CHÍNH
[Kết quả từ Financial Analyst - có nguồn]

🏆 PHÂN TÍCH THỊ TRƯỜNG
[Kết quả từ Market Analyst - có nguồn]

⚠️ ĐÁNH GIÁ RỦI RO
[Kết quả từ Risk Analyst - có nguồn]

💡 KHUYẾN NGHỊ ĐẦU TƯ
• Đánh giá tổng hợp: [điểm/10]
• Khuyến nghị: [MUA/GIỮ/BÁN]
• Lý do: [dựa trên data thực]

═══════════════════════════════════════════════════════

**KHI NGƯỜI DÙNG CHÀO HỎI:**
"Xin chào! Tôi là Research Orchestrator với khả năng phân tích THỰC.
Đội ngũ của tôi có thể tìm kiếm thông tin MỚI NHẤT từ internet:
• Financial Analyst - Tìm & phân tích dữ liệu tài chính
• Market Analyst - Tìm & phân tích dữ liệu thị trường
• Risk Analyst - Tìm & đánh giá rủi ro

Hãy cho tôi tên công ty (ví dụ: 'Phân tích Tesla' hoặc 'Nghiên cứu VinGroup')"`,
        Tools: analystTools,
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

    // Sử dụng Gemini 2.x cho Google Search support
    geminiModel, err := gemini.NewModel(ctx, "gemini-2.0-flash", &genai.ClientConfig{
        APIKey: apiKey,
    })
    if err != nil {
        log.Fatalf("Không thể tạo model: %v", err)
    }

    // 1. Tạo Search Agent (shared) - có Google Search
    searchAgent, err := createSearchAgent(ctx, geminiModel)
    if err != nil {
        log.Fatal(err)
    }

    // 2. Tạo các Analyst Agents - mỗi agent có Search Agent làm tool
    financialAnalyst, err := createFinancialAnalyst(ctx, geminiModel, searchAgent)
    if err != nil {
        log.Fatal(err)
    }

    marketAnalyst, err := createMarketAnalyst(ctx, geminiModel, searchAgent)
    if err != nil {
        log.Fatal(err)
    }

    riskAnalyst, err := createRiskAnalyst(ctx, geminiModel, searchAgent)
    if err != nil {
        log.Fatal(err)
    }

    // 3. Tạo Orchestrator - điều phối các Analysts
    orchestrator, err := createResearchOrchestrator(ctx, geminiModel, []agent.Agent{
        financialAnalyst,
        marketAnalyst,
        riskAnalyst,
    })
    if err != nil {
        log.Fatal(err)
    }

    config := &launcher.Config{
        AgentLoader: agent.NewSingleLoader(orchestrator),
    }

    lch := full.NewLauncher()
    fmt.Println("=== Company Research Agent - Parallelization with Real Search ===")
    fmt.Println("Mỗi Analyst có khả năng tự tìm kiếm thông tin từ Google")
    fmt.Println("Khởi động server...")

    err = lch.Execute(ctx, config, os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
}
```

**Luồng hoạt động:**
```
User: "Phân tích Tesla"
         │
         ▼
    Orchestrator
         │
    ┌────┼────┐
    ▼    ▼    ▼
Financial  Market   Risk
Analyst    Analyst  Analyst
    │        │        │
    ▼        ▼        ▼
 Search   Search   Search
 Agent    Agent    Agent
    │        │        │
    ▼        ▼        ▼
 Google   Google   Google
 Search   Search   Search
    │        │        │
    └────┬───┴────────┘
         ▼
   Tổng hợp báo cáo
   (với data thực)
```

---

## Phần 5: Parallelization với Goroutines (Nâng cao)

Nếu bạn muốn kiểm soát việc thực thi song song ở mức Go (không phụ thuộc vào LLM), đây là cách tiếp cận:

### 5.1 Parallel Executor Pattern

```go
// ParallelExecutor thực thi nhiều agents song song
type ParallelExecutor struct {
    agents []agent.Agent
}

// AnalysisTask đại diện cho một tác vụ phân tích
type AnalysisTask struct {
    Agent  agent.Agent
    Input  string
    Result string
    Error  error
    Duration time.Duration
}

// ExecuteParallel chạy tất cả agents song song và thu thập kết quả
func (pe *ParallelExecutor) ExecuteParallel(ctx context.Context, input string) []AnalysisTask {
    var wg sync.WaitGroup
    results := make([]AnalysisTask, len(pe.agents))

    for i, ag := range pe.agents {
        wg.Add(1)
        go func(index int, a agent.Agent) {
            defer wg.Done()

            start := time.Now()

            // Thực thi agent
            // Lưu ý: Đây là pseudo-code, cần adapter phù hợp với ADK-Go API
            result, err := executeAgent(ctx, a, input)

            results[index] = AnalysisTask{
                Agent:    a,
                Input:    input,
                Result:   result,
                Error:    err,
                Duration: time.Since(start),
            }
        }(i, ag)
    }

    wg.Wait()
    return results
}
```

### 5.2 Sử dụng Channels cho kết quả streaming

```go
// StreamingParallelExecutor với channels
func StreamingParallelExecutor(ctx context.Context, agents []agent.Agent, input string) <-chan AnalysisTask {
    results := make(chan AnalysisTask, len(agents))

    go func() {
        var wg sync.WaitGroup

        for _, ag := range agents {
            wg.Add(1)
            go func(a agent.Agent) {
                defer wg.Done()

                start := time.Now()
                result, err := executeAgent(ctx, a, input)

                results <- AnalysisTask{
                    Agent:    a,
                    Input:    input,
                    Result:   result,
                    Error:    err,
                    Duration: time.Since(start),
                }
            }(ag)
        }

        wg.Wait()
        close(results)
    }()

    return results
}

// Sử dụng:
// for result := range StreamingParallelExecutor(ctx, agents, "Apple Inc") {
//     fmt.Printf("Received result from %s\n", result.Agent.Name())
// }
```

---

## Phần 6: Xử lý lỗi trong Parallelization

### 6.1 Fail-fast vs Fail-safe

**Fail-fast:** Dừng tất cả nếu một tác vụ lỗi
```go
func FailFastParallel(ctx context.Context, agents []agent.Agent, input string) ([]string, error) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    results := make([]string, len(agents))
    errChan := make(chan error, len(agents))

    var wg sync.WaitGroup
    for i, ag := range agents {
        wg.Add(1)
        go func(index int, a agent.Agent) {
            defer wg.Done()

            select {
            case <-ctx.Done():
                return
            default:
                result, err := executeAgent(ctx, a, input)
                if err != nil {
                    errChan <- err
                    cancel() // Hủy tất cả goroutines khác
                    return
                }
                results[index] = result
            }
        }(i, ag)
    }

    wg.Wait()
    close(errChan)

    if err := <-errChan; err != nil {
        return nil, err
    }
    return results, nil
}
```

**Fail-safe:** Tiếp tục với các tác vụ còn lại
```go
func FailSafeParallel(ctx context.Context, agents []agent.Agent, input string) []AnalysisTask {
    var wg sync.WaitGroup
    results := make([]AnalysisTask, len(agents))

    for i, ag := range agents {
        wg.Add(1)
        go func(index int, a agent.Agent) {
            defer wg.Done()

            result, err := executeAgent(ctx, a, input)
            results[index] = AnalysisTask{
                Agent:  a,
                Result: result,
                Error:  err, // Lưu lỗi nhưng không dừng
            }
        }(i, ag)
    }

    wg.Wait()
    return results // Trả về tất cả, kể cả những task lỗi
}
```

### 6.2 Timeout handling

```go
func ParallelWithTimeout(ctx context.Context, agents []agent.Agent, input string, timeout time.Duration) []AnalysisTask {
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    results := make([]AnalysisTask, len(agents))
    var wg sync.WaitGroup

    for i, ag := range agents {
        wg.Add(1)
        go func(index int, a agent.Agent) {
            defer wg.Done()

            start := time.Now()
            result, err := executeAgent(ctx, a, input)

            if ctx.Err() == context.DeadlineExceeded {
                results[index] = AnalysisTask{
                    Agent: a,
                    Error: fmt.Errorf("timeout after %v", timeout),
                }
                return
            }

            results[index] = AnalysisTask{
                Agent:    a,
                Result:   result,
                Error:    err,
                Duration: time.Since(start),
            }
        }(i, ag)
    }

    wg.Wait()
    return results
}
```

---

## Phần 7: Bài tập mở rộng

### Bài tập 1: Thêm Performance Metrics

Mở rộng code để theo dõi:
- Thời gian thực thi của từng analyst
- Tổng thời gian (so với thời gian nếu chạy tuần tự)
- Speedup factor

### Bài tập 2: Weighted Aggregation

Thêm trọng số cho mỗi analyst:
- Financial: 30%
- Market: 25%
- Tech: 25%
- Risk: 20%

### Bài tập 3: Conditional Parallelization

Chỉ chạy song song các analyst phù hợp với ngành:
- Tech company → ưu tiên Tech Analyst
- Financial company → ưu tiên Financial Analyst

### Bài tập 4: Rate Limiting

Thêm rate limiting để không quá tải API:
- Maximum 3 concurrent requests
- Retry với exponential backoff

---

## Phần 8: Best Practices

### 8.1 Khi nào KHÔNG nên Parallelization

- Các tác vụ **phụ thuộc** vào nhau (dùng Prompt Chaining thay thế)
- Tác vụ đơn giản, overhead parallelization > lợi ích
- Resource constraints (memory, API rate limits)

### 8.2 Design Guidelines

1. **Xác định dependencies:** Vẽ dependency graph trước khi implement
2. **Graceful degradation:** Hệ thống vẫn hoạt động nếu một số task fail
3. **Timeout sensible:** Set timeout hợp lý cho từng task
4. **Logging:** Log đầy đủ để debug parallel execution

### 8.3 Testing Parallel Code

```go
func TestParallelExecution(t *testing.T) {
    // Sử dụng mock agents để test
    // Verify tất cả agents được gọi
    // Verify kết quả được aggregate đúng
    // Test timeout scenarios
    // Test error handling
}
```

---

## Phần 9: Code Mẫu Đầy Đủ (Solution)

Dưới đây là code hoàn chỉnh cho file `cmd/parallelization/main.go` với **Real Search capability**:

```go
// Package main demonstrates Parallelization pattern using Google ADK-Go
//
// Parallelization với Real Search:
// - Search Agent: Có Google Search built-in tool (shared)
// - Mỗi Analyst Agent có Search Agent làm tool → tự tìm kiếm thông tin
// - Orchestrator điều phối các Analysts chạy song song
// - Kết quả phân tích dựa trên DỮ LIỆU THỰC từ internet
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
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

// ============================================================================
// SEARCH AGENT - Agent chuyên tìm kiếm với Google Search (Shared)
// ============================================================================

func createSearchAgent(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "web_search",
		Model:       m,
		Description: "Tìm kiếm thông tin từ internet. Cung cấp query và nhận kết quả tìm kiếm mới nhất.",
		Instruction: `Bạn là chuyên gia tìm kiếm thông tin.

NHIỆM VỤ:
Khi nhận yêu cầu tìm kiếm, hãy:
1. Tìm kiếm thông tin liên quan và MỚI NHẤT
2. Trả về kết quả ngắn gọn, có nguồn
3. Ưu tiên thông tin từ nguồn uy tín

ĐỊNH DẠNG TRẢ VỀ:
🔍 KẾT QUẢ TÌM KIẾM:
[Thông tin tìm được - có ghi nguồn]`,
		Tools: []tool.Tool{
			geminitool.GoogleSearch{}, // Built-in Google Search
		},
	})
}

// ============================================================================
// SPECIALIST AGENTS - Có khả năng tự tìm kiếm thông tin
// ============================================================================

func createFinancialAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
	searchTool := agenttool.New(searchAgent, nil)

	return llmagent.New(llmagent.Config{
		Name:        "financial_analyst",
		Model:       m,
		Description: "Chuyên gia phân tích tài chính. Có khả năng tìm kiếm thông tin tài chính mới nhất.",
		Instruction: `Bạn là chuyên gia phân tích tài chính với 20 năm kinh nghiệm.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin tài chính MỚI NHẤT về công ty
2. Phân tích dữ liệu tìm được
3. Đưa ra đánh giá dựa trên data thực

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] financial report 2024"
- "[Tên công ty] revenue profit quarterly"
- "[Tên công ty] stock price analysis"

**NHIỆM VỤ PHÂN TÍCH:**
- Doanh thu và tăng trưởng
- Biên lợi nhuận
- Dòng tiền và thanh khoản
- Cấu trúc nợ

**ĐỊNH DẠNG TRẢ VỀ:**
📊 PHÂN TÍCH TÀI CHÍNH
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Điểm mạnh: [liệt kê]
• Điểm yếu: [liệt kê]
• Đánh giá: [1-10]/10
• Khuyến nghị: [Mua/Giữ/Bán]`,
		Tools: []tool.Tool{searchTool},
	})
}

func createMarketAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
	searchTool := agenttool.New(searchAgent, nil)

	return llmagent.New(llmagent.Config{
		Name:        "market_analyst",
		Model:       m,
		Description: "Chuyên gia phân tích thị trường. Có khả năng tìm kiếm thông tin thị trường mới nhất.",
		Instruction: `Bạn là chuyên gia phân tích thị trường với kiến thức sâu rộng.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin thị trường MỚI NHẤT
2. Phân tích vị thế cạnh tranh
3. Đánh giá xu hướng ngành

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] market share 2024"
- "[Tên công ty] competitors analysis"
- "[Industry] market trends"

**NHIỆM VỤ PHÂN TÍCH:**
- Thị phần và vị thế
- Đối thủ cạnh tranh
- Xu hướng ngành
- Cơ hội mở rộng

**ĐỊNH DẠNG TRẢ VỀ:**
🏆 PHÂN TÍCH THỊ TRƯỜNG
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Vị thế: [Dẫn đầu/Top 3/Trung bình/Theo sau]
• Đối thủ chính: [danh sách]
• Xu hướng ngành: [Tăng/Ổn định/Giảm]
• Tiềm năng: [Cao/Trung bình/Thấp]`,
		Tools: []tool.Tool{searchTool},
	})
}

func createRiskAnalyst(ctx context.Context, m model.LLM, searchAgent agent.Agent) (agent.Agent, error) {
	searchTool := agenttool.New(searchAgent, nil)

	return llmagent.New(llmagent.Config{
		Name:        "risk_analyst",
		Model:       m,
		Description: "Chuyên gia đánh giá rủi ro. Có khả năng tìm kiếm tin tức và thông tin rủi ro.",
		Instruction: `Bạn là chuyên gia quản lý rủi ro với kinh nghiệm đánh giá doanh nghiệp.

**CÁCH LÀM VIỆC:**
1. Sử dụng tool web_search để tìm thông tin về rủi ro và tin tức
2. Đánh giá các loại rủi ro
3. Đề xuất biện pháp giảm thiểu

**GỢI Ý TÌM KIẾM:**
- "[Tên công ty] risks challenges 2024"
- "[Tên công ty] lawsuit legal issues"
- "[Tên công ty] controversy news"

**NHIỆM VỤ PHÂN TÍCH:**
- Rủi ro hoạt động
- Rủi ro pháp lý
- Rủi ro thị trường
- Rủi ro danh tiếng

**ĐỊNH DẠNG TRẢ VỀ:**
⚠️ ĐÁNH GIÁ RỦI RO
━━━━━━━━━━━━━━━━━━━━━
📈 Dữ liệu tìm được: [tóm tắt từ search]
• Rủi ro CAO: [liệt kê nếu có]
• Rủi ro TRUNG BÌNH: [liệt kê]
• Rủi ro THẤP: [liệt kê]
• Điểm rủi ro: [1-10]/10
• Biện pháp giảm thiểu: [khuyến nghị]`,
		Tools: []tool.Tool{searchTool},
	})
}

// ============================================================================
// ORCHESTRATOR AGENT - Điều phối phân tích song song
// ============================================================================

func createResearchOrchestrator(ctx context.Context, m model.LLM, analysts []agent.Agent) (agent.Agent, error) {
	var analystTools []tool.Tool
	for _, analyst := range analysts {
		analystTools = append(analystTools, agenttool.New(analyst, nil))
	}

	return llmagent.New(llmagent.Config{
		Name:        "research_orchestrator",
		Model:       m,
		Description: "Trưởng nhóm nghiên cứu đầu tư - Điều phối phân tích song song với dữ liệu thực",
		Instruction: `Bạn là trưởng nhóm nghiên cứu đầu tư chuyên nghiệp.

**ĐỘI NGŨ CỦA BẠN:**
Mỗi analyst có khả năng TỰ TÌM KIẾM thông tin từ internet:
- financial_analyst: Phân tích tài chính (tự search financial data)
- market_analyst: Phân tích thị trường (tự search market data)
- risk_analyst: Đánh giá rủi ro (tự search risk news)

**QUY TRÌNH KHI NHẬN YÊU CẦU PHÂN TÍCH:**

BƯỚC 1: PHÂN TÍCH SONG SONG
Gọi TẤT CẢ analysts CÙNG LÚC với tên công ty.
Mỗi analyst sẽ tự tìm kiếm và phân tích.

QUAN TRỌNG: Gọi cả 3 agent trong CÙNG MỘT LƯỢT (parallel execution)

BƯỚC 2: TỔNG HỢP BÁO CÁO
Sau khi nhận đủ kết quả (có dữ liệu thực từ search), tổng hợp:

═══════════════════════════════════════════════════════
📋 BÁO CÁO NGHIÊN CỨU: [TÊN CÔNG TY]
(Dựa trên dữ liệu thực từ internet)
═══════════════════════════════════════════════════════

📌 TÓM TẮT ĐIỀU HÀNH
[3-4 câu tóm tắt - dựa trên data thực]

📊 PHÂN TÍCH TÀI CHÍNH
[Kết quả từ Financial Analyst]

🏆 PHÂN TÍCH THỊ TRƯỜNG
[Kết quả từ Market Analyst]

⚠️ ĐÁNH GIÁ RỦI RO
[Kết quả từ Risk Analyst]

💡 KHUYẾN NGHỊ ĐẦU TƯ
• Đánh giá: [điểm/10]
• Khuyến nghị: [MUA/GIỮ/BÁN]
• Lý do: [dựa trên data thực]

═══════════════════════════════════════════════════════

**KHI NGƯỜI DÙNG CHÀO HỎI:**
"Xin chào! Tôi là Research Orchestrator với khả năng phân tích THỰC.
Đội ngũ có thể tìm kiếm thông tin MỚI NHẤT từ internet:
• Financial Analyst - Tìm & phân tích dữ liệu tài chính
• Market Analyst - Tìm & phân tích dữ liệu thị trường
• Risk Analyst - Tìm & đánh giá rủi ro

Cho tôi tên công ty (ví dụ: 'Phân tích Tesla' hoặc 'Nghiên cứu VinGroup')"`,
		Tools: analystTools,
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

	// Sử dụng Gemini 2.x cho Google Search support
	geminiModel, err := gemini.NewModel(ctx, "gemini-2.0-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Không thể tạo model: %v", err)
	}

	// 1. Tạo Search Agent (shared) - có Google Search
	searchAgent, err := createSearchAgent(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Tạo các Analyst Agents - mỗi agent có Search Agent làm tool
	financialAnalyst, err := createFinancialAnalyst(ctx, geminiModel, searchAgent)
	if err != nil {
		log.Fatal(err)
	}

	marketAnalyst, err := createMarketAnalyst(ctx, geminiModel, searchAgent)
	if err != nil {
		log.Fatal(err)
	}

	riskAnalyst, err := createRiskAnalyst(ctx, geminiModel, searchAgent)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Tạo Orchestrator - điều phối các Analysts
	orchestrator, err := createResearchOrchestrator(ctx, geminiModel, []agent.Agent{
		financialAnalyst,
		marketAnalyst,
		riskAnalyst,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(orchestrator),
	}

	lch := full.NewLauncher()
	fmt.Println("=== Company Research Agent - Parallelization with Real Search ===")
	fmt.Println("Mỗi Analyst có khả năng tự tìm kiếm thông tin từ Google")
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
go run cmd/parallelization/main.go
```

### 10.2 Test cases

**Test 1: Greeting**
```
User: Xin chào
Expected: Agent giới thiệu bản thân và đội ngũ
```

**Test 2: Company Analysis**
```
User: Phân tích công ty Apple Inc
Expected:
- 4 analyst được gọi (có thể song song tùy LLM)
- Báo cáo tổng hợp đầy đủ 4 phần
- Khuyến nghị đầu tư cuối cùng
```

**Test 3: Vietnamese Company**
```
User: Nghiên cứu VinGroup
Expected: Phân tích phù hợp với context Việt Nam
```

---

## Tài liệu tham khảo

1. [ADK-Go Multi-Agent Documentation](https://google.github.io/adk-docs/agents/multi-agents/)
2. [Go Concurrency Patterns](https://go.dev/blog/pipelines)
3. [Chapter 3: Parallelization - Agentic Design Patterns](../doc_vi/03_Chapter_3_Parallelization.md)
