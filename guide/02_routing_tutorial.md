# Hướng Dẫn Thực Hành: Routing Pattern với ADK-Go

## Mục Tiêu Bài Học

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu rõ mẫu thiết kế Routing (Định tuyến)
- Phân biệt được các cơ chế routing khác nhau
- Xây dựng được hệ thống multi-agent với routing tự động
- Tự tay implement một Customer Support Bot với routing

---

## Phần 1: Tổng Quan Về Routing

### 1.1 Routing là gì?

**Routing (Định tuyến)** là kỹ thuật cho phép tác nhân AI **đưa ra quyết định động** về việc chuyển yêu cầu đến đâu dựa trên nội dung/ý định của input.

```
                    ┌─────────────────┐
                    │  Agent Đặt vé   │
                    └────────▲────────┘
                             │
[User Input] ──► [Router] ───┼────► [Agent Thông tin]
                             │
                    ┌────────▼────────┐
                    │ Agent Hỗ trợ KT │
                    └─────────────────┘
```

### 1.2 So sánh Prompt Chaining vs Routing

| Tiêu chí | Prompt Chaining | Routing |
|----------|-----------------|---------|
| Luồng xử lý | Tuần tự, cố định | Động, có điều kiện |
| Quyết định | Không có | Dựa trên ý định/nội dung |
| Phù hợp với | Tác vụ có bước rõ ràng | Tác vụ cần phân loại |
| Ví dụ | ETL pipeline | Customer support bot |

### 1.3 Khi nào sử dụng Routing?

- Bot hỗ trợ khách hàng (phân loại: bán hàng / kỹ thuật / khiếu nại)
- Trợ lý ảo đa năng (đặt vé / tra cứu / tư vấn)
- Hệ thống xử lý tài liệu (hóa đơn / hợp đồng / báo cáo)
- API Gateway thông minh

---

## Phần 2: Các Cơ Chế Routing

### 2.1 LLM-based Routing (Phổ biến nhất)

Sử dụng LLM để phân loại ý định và quyết định route.

```
Prompt: "Phân tích yêu cầu và trả về MỘT trong các danh mục:
- 'booking': đặt vé, đặt phòng
- 'info': hỏi thông tin
- 'support': hỗ trợ kỹ thuật
- 'unclear': không rõ ràng"

Input: "Tôi muốn đặt vé máy bay đi Đà Nẵng"
Output: "booking"
```

**Ưu điểm:** Linh hoạt, hiểu ngữ cảnh
**Nhược điểm:** Tốn token, có độ trễ

### 2.2 Rule-based Routing

Sử dụng từ khóa, regex, hoặc logic if-else.

```go
func routeByKeyword(input string) string {
    input = strings.ToLower(input)

    if strings.Contains(input, "đặt vé") || strings.Contains(input, "book") {
        return "booking"
    }
    if strings.Contains(input, "giá") || strings.Contains(input, "bao nhiêu") {
        return "info"
    }
    return "unclear"
}
```

**Ưu điểm:** Nhanh, xác định (deterministic)
**Nhược điểm:** Cứng nhắc, khó xử lý ngữ cảnh phức tạp

### 2.3 Embedding-based Routing (Semantic Router)

So sánh vector embedding của input với các route mẫu.

```
Route "booking" embeddings: ["đặt vé", "book phòng", "reservation"]
Route "info" embeddings: ["thông tin", "giá cả", "địa chỉ"]

Input embedding ──► Cosine Similarity ──► Route có similarity cao nhất
```

**Ưu điểm:** Hiểu ngữ nghĩa, không cần gọi LLM
**Nhược điểm:** Cần chuẩn bị embeddings trước

### 2.4 ML Model-based Routing

Huấn luyện classifier riêng cho việc phân loại.

**Ưu điểm:** Chính xác cao nếu có dữ liệu tốt
**Nhược điểm:** Cần dataset, effort huấn luyện

---

## Phần 3: Routing trong ADK-Go

### 3.1 Cơ chế Auto-Flow

ADK-Go sử dụng **Auto-Flow** - LLM tự động chọn sub-agent phù hợp dựa trên:
- `Description` của mỗi sub-agent
- `Instruction` của parent agent
- Nội dung tin nhắn user

```go
// Parent agent với sub_agents sẽ tự động routing
coordinator := Agent{
    Name: "Coordinator",
    Instruction: "Phân tích yêu cầu và chuyển đến agent phù hợp",
    SubAgents: []Agent{bookingAgent, infoAgent, supportAgent},
}
```

### 3.2 Cấu trúc Multi-Agent với Routing

```
┌─────────────────────────────────────────┐
│           Coordinator Agent             │
│  (Instruction: phân tích & điều phối)   │
└─────────────────┬───────────────────────┘
                  │ SubAgents (Auto-Flow)
        ┌─────────┼─────────┐
        ▼         ▼         ▼
   ┌─────────┐ ┌─────────┐ ┌─────────┐
   │ Booking │ │  Info   │ │ Support │
   │  Agent  │ │  Agent  │ │  Agent  │
   │(đặt vé) │ │(tra cứu)│ │(hỗ trợ) │
   └─────────┘ └─────────┘ └─────────┘
```

---

## Phần 4: Bài Tập Thực Hành

### Bài 1: Tạo cấu trúc project

**Yêu cầu:** Tạo cấu trúc thư mục cho Customer Support Bot

```
cmd/routing/
├── main.go
└── agents/
    ├── coordinator.go    # Agent điều phối
    ├── booking.go        # Agent đặt vé/phòng
    ├── info.go           # Agent tra cứu thông tin
    └── support.go        # Agent hỗ trợ kỹ thuật
```

**Bài tập:** Tạo các file trống với package declaration

```go
// agents/coordinator.go
package agents

// TODO: Implement coordinator agent
```

---

### Bài 2: Định nghĩa Sub-Agent chuyên biệt

**Yêu cầu:** Tạo 3 sub-agents với mô tả rõ ràng

**Kiến thức cần biết:**

```go
import "google.golang.org/adk/agent/llmagent"

agent, err := llmagent.New(llmagent.Config{
    Name:        "agent_name",
    Model:       model,
    Description: "Mô tả QUAN TRỌNG - LLM dùng để routing",
    Instruction: "Hướng dẫn chi tiết cho agent",
    Tools:       []tool.Tool{...},
})
```

**Bài tập:** Hoàn thành 3 agents

```go
// agents/booking.go
package agents

// BookingAgent - Xử lý đặt vé máy bay và khách sạn
//
// Description cần rõ ràng để Coordinator biết khi nào route đến
// Ví dụ: "Xử lý tất cả yêu cầu đặt vé máy bay, tàu, và đặt phòng khách sạn"

func CreateBookingAgent(ctx context.Context, model model.Model) (agent.Agent, error) {
    // TODO: Tạo booking tool
    // TODO: Tạo agent với:
    //   - Name: "booking_agent"
    //   - Description: mô tả rõ nhiệm vụ đặt vé/phòng
    //   - Instruction: hướng dẫn cách xử lý yêu cầu đặt chỗ
    //   - Tools: booking tool

    return nil, nil // Thay bằng code thực
}
```

```go
// agents/info.go
package agents

// InfoAgent - Tra cứu thông tin chung
//
// Description: "Cung cấp thông tin về giá cả, lịch trình, địa điểm,
// và trả lời các câu hỏi kiến thức chung"

func CreateInfoAgent(ctx context.Context, model model.Model) (agent.Agent, error) {
    // TODO: Implement
    return nil, nil
}
```

```go
// agents/support.go
package agents

// SupportAgent - Hỗ trợ kỹ thuật
//
// Description: "Xử lý các vấn đề kỹ thuật, khiếu nại,
// và yêu cầu hỗ trợ từ khách hàng"

func CreateSupportAgent(ctx context.Context, model model.Model) (agent.Agent, error) {
    // TODO: Implement
    return nil, nil
}
```

**Gợi ý Description tốt:**
- Cụ thể về loại yêu cầu xử lý
- Liệt kê các từ khóa/tình huống
- Không trùng lặp với agent khác

---

### Bài 3: Tạo Coordinator Agent với Auto-Flow

**Đây là phần quan trọng nhất!**

**Kiến thức cần biết:**

```go
// Coordinator KHÔNG cần tools riêng
// Chỉ cần instruction và sub_agents

coordinator, err := llmagent.New(llmagent.Config{
    Name:        "coordinator",
    Model:       model,
    Description: "Điều phối viên chính",
    Instruction: `Bạn là điều phối viên. Nhiệm vụ:
1. Phân tích yêu cầu của người dùng
2. Xác định ý định (đặt chỗ / thông tin / hỗ trợ)
3. Chuyển đến agent chuyên biệt phù hợp

Quy tắc:
- Đặt vé, đặt phòng → booking_agent
- Hỏi thông tin, giá cả → info_agent
- Khiếu nại, lỗi kỹ thuật → support_agent`,
    SubAgents: []agent.Agent{bookingAgent, infoAgent, supportAgent},
})
```

**Bài tập:** Hoàn thành file coordinator.go

```go
// agents/coordinator.go
package agents

import (
    // TODO: Thêm imports cần thiết
)

// CreateCoordinator tạo agent điều phối với routing tự động
func CreateCoordinator(
    ctx context.Context,
    model model.Model,
    subAgents []agent.Agent,
) (agent.Agent, error) {

    // TODO: Viết instruction cho coordinator
    // Instruction cần:
    // 1. Giải thích vai trò điều phối
    // 2. Liệt kê các loại yêu cầu và agent tương ứng
    // 3. Hướng dẫn xử lý trường hợp không rõ ràng

    instruction := `
    // Viết instruction ở đây
    `

    // TODO: Tạo coordinator agent với SubAgents
    return llmagent.New(llmagent.Config{
        // Hoàn thành config
    })
}
```

**Template Instruction cho Coordinator:**

```
Bạn là [VAI TRÒ] của hệ thống [TÊN HỆ THỐNG].

NHIỆM VỤ CHÍNH:
- Phân tích yêu cầu người dùng
- Xác định ý định và chuyển đến agent phù hợp

QUY TẮC ROUTING:
1. [Loại yêu cầu 1] → [agent_name_1]
   Ví dụ: [ví dụ cụ thể]

2. [Loại yêu cầu 2] → [agent_name_2]
   Ví dụ: [ví dụ cụ thể]

3. [Loại yêu cầu 3] → [agent_name_3]
   Ví dụ: [ví dụ cụ thể]

XỬ LÝ TRƯỜNG HỢP KHÔNG RÕ RÀNG:
- Nếu không xác định được ý định, hỏi lại người dùng
- Nếu yêu cầu thuộc nhiều loại, ưu tiên [loại nào]
```

---

### Bài 4: Hoàn thành Main và Test

**Yêu cầu:** Kết nối tất cả và chạy thử

```go
// cmd/routing/main.go
package main

import (
    // TODO: Imports
)

func main() {
    ctx := context.Background()

    // 1. Lấy API key
    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        log.Fatal("GOOGLE_API_KEY is required")
    }

    // 2. Tạo model
    // TODO: Tạo Gemini model

    // 3. Tạo sub-agents
    // TODO: Gọi CreateBookingAgent, CreateInfoAgent, CreateSupportAgent

    // 4. Tạo coordinator với sub-agents
    // TODO: Gọi CreateCoordinator

    // 5. Cấu hình và chạy
    config := &adk.Config{
        AgentLoader: services.NewSingleAgentLoader(coordinator),
    }

    launcher := full.NewLauncher()
    err := launcher.Execute(ctx, config, os.Args[1:])
    if err != nil {
        log.Fatalf("Failed: %v", err)
    }
}
```

---

### Bài 5: Test Cases

**Chạy ứng dụng:**
```bash
go run cmd/routing/main.go web
```

**Test với các input sau và ghi nhận kết quả:**

| # | Input | Route mong đợi |
|---|-------|----------------|
| 1 | "Đặt cho tôi vé máy bay đi Đà Nẵng ngày 25/12" | booking_agent |
| 2 | "Giá phòng khách sạn 5 sao ở Hà Nội là bao nhiêu?" | info_agent |
| 3 | "Tôi không đăng nhập được vào tài khoản" | support_agent |
| 4 | "Book 2 vé tàu đi Sapa" | booking_agent |
| 5 | "Thời tiết Đà Lạt thế nào?" | info_agent |
| 6 | "Tôi muốn khiếu nại về chuyến bay bị delay" | support_agent |
| 7 | "Hello" | ??? (quan sát) |

**Bài tập:**
- Ghi lại kết quả routing thực tế
- So sánh với kết quả mong đợi
- Điều chỉnh instruction/description nếu routing sai

---

## Phần 5: Các Import Cần Thiết

```go
import (
    "context"
    "fmt"
    "log"
    "os"

    "google.golang.org/adk/agent"
    "google.golang.org/adk/agent/llmagent"
    "google.golang.org/adk/cmd/launcher/adk"
    "google.golang.org/adk/cmd/launcher/full"
    "google.golang.org/adk/model"
    "google.golang.org/adk/model/gemini"
    "google.golang.org/adk/server/restapi/services"
    "google.golang.org/adk/tool"
    "google.golang.org/adk/tool/functiontool"
    "google.golang.org/genai"
)
```

---

## Phần 6: Bài Tập Nâng Cao

### Nâng cao 1: Thêm Fallback Agent

Tạo agent xử lý các yêu cầu không rõ ràng:

```go
// Fallback agent hỏi lại người dùng để làm rõ ý định
fallbackAgent := Agent{
    Name: "clarification_agent",
    Description: "Xử lý các yêu cầu không rõ ràng, hỏi lại người dùng",
    Instruction: "Khi nhận yêu cầu không rõ, hãy hỏi lại một cách lịch sự...",
}
```

### Nâng cao 2: Hybrid Routing

Kết hợp Rule-based và LLM-based:

```go
// Bước 1: Kiểm tra từ khóa (nhanh)
route := checkKeywords(input)
if route != "unknown" {
    return route
}

// Bước 2: Nếu không match, dùng LLM (chậm hơn nhưng chính xác)
return llmClassify(input)
```

### Nâng cao 3: Routing với Context/State

Routing dựa trên cả lịch sử hội thoại:

```go
// Nếu user đã hỏi về booking trước đó,
// câu hỏi tiếp theo có thể liên quan đến booking
// dù không có từ khóa rõ ràng

// User: "Đặt vé đi Đà Nẵng"  → booking
// User: "Còn ngày nào khác không?" → vẫn booking (theo context)
```

### Nâng cao 4: Logging và Monitoring

Thêm logging để theo dõi routing decisions:

```go
type RoutingLog struct {
    Timestamp   time.Time
    UserInput   string
    DetectedIntent string
    RoutedTo    string
    Confidence  float64
}

// Log mỗi quyết định routing để phân tích sau
```

### Nâng cao 5: A/B Testing Routes

Thử nghiệm các instruction khác nhau:

```go
// Version A: Instruction ngắn gọn
// Version B: Instruction chi tiết với ví dụ

// So sánh độ chính xác routing giữa 2 versions
```

---

## Phần 7: Troubleshooting

### Vấn đề 1: Routing sai agent

**Nguyên nhân có thể:**
- Description của agents trùng lặp hoặc mơ hồ
- Instruction của coordinator không rõ ràng

**Giải pháp:**
- Viết description cụ thể, không overlap
- Thêm ví dụ vào instruction
- Thêm negative examples ("KHÔNG xử lý: ...")

### Vấn đề 2: Coordinator không chuyển đến sub-agent

**Nguyên nhân có thể:**
- SubAgents chưa được add đúng cách
- Model không đủ mạnh để hiểu instruction

**Giải pháp:**
- Kiểm tra SubAgents array
- Thử model mạnh hơn (gemini-2.5-flash → gemini-2.5-pro)

### Vấn đề 3: Response chậm

**Nguyên nhân:**
- Mỗi routing decision cần gọi LLM

**Giải pháp:**
- Implement rule-based pre-filter
- Cache routing decisions cho queries tương tự

---

## Phần 8: Best Practices

### 8.1 Viết Description Tốt

```go
// ❌ Không tốt - quá chung chung
Description: "Xử lý các yêu cầu của khách hàng"

// ✅ Tốt - cụ thể và rõ ràng
Description: "Xử lý yêu cầu đặt vé máy bay, tàu hỏa, và đặt phòng khách sạn. " +
    "Bao gồm: tìm chuyến, chọn ghế, thanh toán, hủy/đổi vé."
```

### 8.2 Instruction Coordinator

```go
// ❌ Không tốt
Instruction: "Chuyển yêu cầu đến agent phù hợp"

// ✅ Tốt
Instruction: `Bạn là điều phối viên của Travel Assistant.

PHÂN LOẠI YÊU CẦU:
1. ĐẶT CHỖ (→ booking_agent):
   - Đặt vé máy bay, tàu, xe
   - Đặt phòng khách sạn
   - Từ khóa: "đặt", "book", "reservation"

2. THÔNG TIN (→ info_agent):
   - Hỏi giá, lịch trình
   - Thông tin điểm đến
   - Từ khóa: "giá", "bao nhiêu", "khi nào"

3. HỖ TRỢ (→ support_agent):
   - Khiếu nại, báo lỗi
   - Vấn đề tài khoản
   - Từ khóa: "lỗi", "không được", "khiếu nại"

LƯU Ý:
- Nếu không chắc chắn, hỏi lại người dùng
- Ưu tiên trải nghiệm người dùng`
```

---

## Phần 9: Tài Liệu Tham Khảo

1. **ADK-Go GitHub**: https://github.com/google/adk-go
2. **ADK Documentation**: https://google.github.io/adk-docs/
3. **Semantic Router Paper**: Semantic Router concept
4. **LangChain Routing**: https://python.langchain.com/docs/how_to/routing/

---

## Checklist Hoàn Thành

- [ ] Hiểu sự khác biệt giữa Prompt Chaining và Routing
- [ ] Biết 4 cơ chế routing (LLM, Rule, Embedding, ML)
- [ ] Tạo được sub-agents với description rõ ràng
- [ ] Viết được instruction cho Coordinator
- [ ] Kết nối sub-agents vào Coordinator
- [ ] Chạy và test routing với nhiều inputs
- [ ] Điều chỉnh instruction để cải thiện độ chính xác
- [ ] Hoàn thành ít nhất 1 bài tập nâng cao

---

## So Sánh với Chapter 1

| Aspect | Ch1: Prompt Chaining | Ch2: Routing |
|--------|---------------------|--------------|
| Pattern | Sequential | Conditional |
| ADK Component | `sequentialagent` | `SubAgents` + Auto-Flow |
| Flow | A → B → C | Router → {A or B or C} |
| Use case | Pipeline xử lý | Phân loại & điều phối |

**Tip:** Có thể kết hợp cả hai!
```
Router → Booking Agent → [Chain: Search → Select → Pay]
       → Info Agent → [Chain: Query → Format → Response]
```

---

## Phần 10: Code Mẫu Đầy Đủ (Solution)

> **Lưu ý:** Hãy tự làm bài tập trước khi xem code mẫu!

```go
// cmd/routing/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher/adk"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/server/restapi/services"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// ============================================================================
// TOOLS CHO CÁC SUB-AGENTS
// ============================================================================

// createBookingTool - Tool xử lý đặt vé/phòng
func createBookingTool() (tool.Tool, error) {
	type Input struct {
		BookingType string `json:"booking_type" description:"Loại đặt chỗ: flight, train, hotel"`
		Destination string `json:"destination" description:"Điểm đến"`
		Date        string `json:"date" description:"Ngày đi/đến"`
		Passengers  int    `json:"passengers" description:"Số lượng hành khách/phòng"`
	}
	type Output struct {
		Status      string `json:"status"`
		BookingID   string `json:"booking_id"`
		Message     string `json:"message"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		// Mô phỏng xử lý đặt chỗ
		bookingID := fmt.Sprintf("BK-%s-%d", input.BookingType[:3], 12345)
		return Output{
			Status:    "confirmed",
			BookingID: bookingID,
			Message:   fmt.Sprintf("Đã đặt %s đến %s ngày %s cho %d người. Mã đặt chỗ: %s",
				input.BookingType, input.Destination, input.Date, input.Passengers, bookingID),
		}
	}

	return functiontool.New(functiontool.Config{
		Name:        "process_booking",
		Description: "Xử lý yêu cầu đặt vé máy bay, tàu, hoặc phòng khách sạn",
	}, handler)
}

// createInfoTool - Tool tra cứu thông tin
func createInfoTool() (tool.Tool, error) {
	type Input struct {
		Query    string `json:"query" description:"Câu hỏi cần tra cứu"`
		Category string `json:"category" description:"Loại thông tin: price, schedule, location, general"`
	}
	type Output struct {
		Answer string `json:"answer"`
		Source string `json:"source"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		// Mô phỏng tra cứu thông tin
		return Output{
			Answer: fmt.Sprintf("Thông tin về '%s': [Dữ liệu mô phỏng cho %s]", input.Query, input.Category),
			Source: "Travel Database",
		}
	}

	return functiontool.New(functiontool.Config{
		Name:        "lookup_information",
		Description: "Tra cứu thông tin về giá cả, lịch trình, địa điểm",
	}, handler)
}

// createSupportTool - Tool hỗ trợ kỹ thuật
func createSupportTool() (tool.Tool, error) {
	type Input struct {
		IssueType   string `json:"issue_type" description:"Loại vấn đề: technical, complaint, account"`
		Description string `json:"description" description:"Mô tả chi tiết vấn đề"`
		Priority    string `json:"priority" description:"Mức độ ưu tiên: low, medium, high"`
	}
	type Output struct {
		TicketID   string `json:"ticket_id"`
		Status     string `json:"status"`
		Message    string `json:"message"`
		NextSteps  string `json:"next_steps"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		ticketID := fmt.Sprintf("TK-%d", 98765)
		return Output{
			TicketID:  ticketID,
			Status:    "opened",
			Message:   fmt.Sprintf("Đã tạo ticket hỗ trợ cho vấn đề %s", input.IssueType),
			NextSteps: "Đội hỗ trợ sẽ liên hệ trong vòng 24 giờ",
		}
	}

	return functiontool.New(functiontool.Config{
		Name:        "create_support_ticket",
		Description: "Tạo ticket hỗ trợ cho các vấn đề kỹ thuật và khiếu nại",
	}, handler)
}

// ============================================================================
// SUB-AGENTS
// ============================================================================

// createBookingAgent - Agent xử lý đặt vé/phòng
func createBookingAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	bookingTool, err := createBookingTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:  "booking_agent",
		Model: m,
		// QUAN TRỌNG: Description phải rõ ràng để Coordinator biết khi nào route đến
		Description: `Chuyên gia đặt chỗ - Xử lý TẤT CẢ yêu cầu đặt vé và đặt phòng.
Bao gồm: đặt vé máy bay, vé tàu, vé xe, đặt phòng khách sạn, resort.
Từ khóa: đặt, book, reservation, vé, phòng, chuyến bay, khách sạn.`,
		Instruction: `Bạn là chuyên gia đặt chỗ của Travel Assistant.

NHIỆM VỤ:
Xử lý các yêu cầu đặt vé và đặt phòng từ khách hàng.

QUY TRÌNH:
1. Xác định loại đặt chỗ (máy bay/tàu/xe/khách sạn)
2. Thu thập thông tin: điểm đến, ngày, số người
3. Sử dụng tool process_booking để xử lý
4. Xác nhận với khách hàng

NGÔN NGỮ: Tiếng Việt, lịch sự, chuyên nghiệp

VÍ DỤ YÊU CẦU:
- "Đặt vé máy bay đi Đà Nẵng"
- "Book phòng khách sạn 5 sao"
- "Tôi muốn đặt 2 vé tàu đi Sapa"`,
		Tools: []tool.Tool{bookingTool},
	})
}

// createInfoAgent - Agent tra cứu thông tin
func createInfoAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	infoTool, err := createInfoTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:  "info_agent",
		Model: m,
		Description: `Chuyên gia thông tin - Tra cứu và cung cấp thông tin du lịch.
Bao gồm: giá vé, lịch trình, thông tin điểm đến, thời tiết, địa điểm.
Từ khóa: giá, bao nhiêu, lịch trình, khi nào, thông tin, ở đâu.
KHÔNG xử lý: đặt vé, khiếu nại, lỗi kỹ thuật.`,
		Instruction: `Bạn là chuyên gia thông tin của Travel Assistant.

NHIỆM VỤ:
Cung cấp thông tin chính xác về du lịch, giá cả, lịch trình.

QUY TRÌNH:
1. Hiểu câu hỏi của khách hàng
2. Xác định loại thông tin cần tra cứu
3. Sử dụng tool lookup_information
4. Trình bày kết quả rõ ràng

NGÔN NGỮ: Tiếng Việt, thân thiện, dễ hiểu

VÍ DỤ CÂU HỎI:
- "Giá vé máy bay đi Đà Nẵng?"
- "Thời tiết Đà Lạt tháng 12?"
- "Các điểm du lịch ở Phú Quốc?"`,
		Tools: []tool.Tool{infoTool},
	})
}

// createSupportAgent - Agent hỗ trợ kỹ thuật
func createSupportAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	supportTool, err := createSupportTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:  "support_agent",
		Model: m,
		Description: `Chuyên gia hỗ trợ - Xử lý vấn đề kỹ thuật và khiếu nại.
Bao gồm: lỗi hệ thống, vấn đề tài khoản, khiếu nại dịch vụ, hoàn tiền.
Từ khóa: lỗi, không được, hỏng, khiếu nại, complaint, hoàn tiền, delay.
KHÔNG xử lý: đặt vé mới, hỏi thông tin chung.`,
		Instruction: `Bạn là chuyên gia hỗ trợ của Travel Assistant.

NHIỆM VỤ:
Giải quyết các vấn đề kỹ thuật và khiếu nại của khách hàng.

QUY TRÌNH:
1. Lắng nghe và thể hiện sự đồng cảm
2. Xác định loại vấn đề và mức độ nghiêm trọng
3. Sử dụng tool create_support_ticket
4. Hướng dẫn các bước tiếp theo

NGÔN NGỮ: Tiếng Việt, đồng cảm, chuyên nghiệp

QUAN TRỌNG:
- Luôn xin lỗi về sự bất tiện
- Đảm bảo khách hàng cảm thấy được lắng nghe

VÍ DỤ VẤN ĐỀ:
- "Tôi không đăng nhập được"
- "Chuyến bay bị delay, tôi muốn khiếu nại"
- "App bị lỗi khi thanh toán"`,
		Tools: []tool.Tool{supportTool},
	})
}

// ============================================================================
// COORDINATOR AGENT
// ============================================================================

// createCoordinator - Agent điều phối với Auto-Flow routing
func createCoordinator(ctx context.Context, m model.Model, subAgents []agent.Agent) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:  "coordinator",
		Model: m,
		Description: "Điều phối viên chính - Phân tích yêu cầu và chuyển đến agent phù hợp",
		Instruction: `Bạn là điều phối viên chính của Travel Assistant.

VAI TRÒ:
Phân tích yêu cầu của khách hàng và chuyển đến agent chuyên biệt phù hợp.

QUY TẮC ROUTING:

1. ĐẶT CHỖ → booking_agent
   Khi khách hàng muốn:
   - Đặt vé máy bay, tàu, xe
   - Đặt phòng khách sạn
   - Book, reservation
   Ví dụ: "Đặt vé đi Đà Nẵng", "Book phòng khách sạn"

2. THÔNG TIN → info_agent
   Khi khách hàng hỏi về:
   - Giá cả, chi phí
   - Lịch trình, thời gian
   - Thông tin điểm đến, thời tiết
   Ví dụ: "Giá vé bao nhiêu?", "Thời tiết Đà Lạt?"

3. HỖ TRỢ → support_agent
   Khi khách hàng gặp:
   - Lỗi kỹ thuật, sự cố
   - Khiếu nại, phàn nàn
   - Vấn đề tài khoản
   Ví dụ: "Không đăng nhập được", "Muốn khiếu nại"

XỬ LÝ TRƯỜNG HỢP KHÔNG RÕ RÀNG:
- Nếu yêu cầu mơ hồ, hỏi lại để làm rõ
- Nếu chào hỏi đơn giản, phản hồi thân thiện và hỏi có thể giúp gì

LƯU Ý:
- Luôn lịch sự và chuyên nghiệp
- Chuyển đến ĐÚNG agent để tránh làm phiền khách hàng`,
		// SubAgents cho phép Auto-Flow routing
		SubAgents: subAgents,
	})
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	ctx := context.Background()

	// 1. Kiểm tra API Key
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("Vui lòng set GOOGLE_API_KEY environment variable")
	}

	// 2. Tạo Gemini Model
	geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Không thể tạo model: %v", err)
	}

	// 3. Tạo các Sub-Agents
	bookingAgent, err := createBookingAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo booking agent: %v", err)
	}

	infoAgent, err := createInfoAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo info agent: %v", err)
	}

	supportAgent, err := createSupportAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo support agent: %v", err)
	}

	// 4. Tạo Coordinator với Sub-Agents
	coordinator, err := createCoordinator(ctx, geminiModel, []agent.Agent{
		bookingAgent,
		infoAgent,
		supportAgent,
	})
	if err != nil {
		log.Fatalf("Không thể tạo coordinator: %v", err)
	}

	// 5. Cấu hình và khởi chạy
	config := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(coordinator),
	}

	fmt.Println("=== Travel Assistant - Routing Demo ===")
	fmt.Println("Khởi động server...")

	launcher := full.NewLauncher()
	err = launcher.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatalf("Lỗi chạy ứng dụng: %v\n\nCú pháp: %s", err, launcher.CommandLineSyntax())
	}
}
```

### Giải thích Code

**1. Tools (Dòng 24-110)**
- `createBookingTool`: Xử lý đặt vé/phòng, trả về mã đặt chỗ
- `createInfoTool`: Tra cứu thông tin, trả về kết quả
- `createSupportTool`: Tạo ticket hỗ trợ

**2. Sub-Agents (Dòng 116-215)**

Mỗi agent có **Description** rất quan trọng:
```go
Description: `Chuyên gia đặt chỗ - Xử lý TẤT CẢ yêu cầu đặt vé...
Bao gồm: đặt vé máy bay, vé tàu...
Từ khóa: đặt, book, reservation...`
```

- **Description** được LLM sử dụng để quyết định routing
- Liệt kê cụ thể: nhiệm vụ, từ khóa, KHÔNG xử lý gì
- Tránh overlap giữa các agents

**3. Coordinator (Dòng 221-285)**

```go
coordinator := llmagent.New({
    ...
    SubAgents: []agent.Agent{bookingAgent, infoAgent, supportAgent},
})
```

- **SubAgents** field kích hoạt Auto-Flow
- LLM tự động chọn sub-agent dựa trên Description
- Instruction của Coordinator hướng dẫn quy tắc routing

**4. Auto-Flow Mechanism**

```
User Input → Coordinator (phân tích) → Chọn Sub-Agent → Execute → Response
```

ADK-Go tự động:
1. Đọc Description của tất cả SubAgents
2. So sánh với user input
3. Chuyển đến agent phù hợp nhất
4. Trả kết quả về user

**5. Best Practices trong Code**

- Description có "KHÔNG xử lý" để tránh routing sai
- Instruction có ví dụ cụ thể
- Coordinator có hướng dẫn xử lý trường hợp mơ hồ
- Mỗi agent độc lập, có tool riêng

### Chạy và Test

```bash
# Chạy ứng dụng
go run cmd/routing/main.go web

# Truy cập http://localhost:8080
```

**Test inputs:**
| Input | Expected Route |
|-------|----------------|
| "Đặt vé máy bay đi Hà Nội" | booking_agent |
| "Giá phòng khách sạn?" | info_agent |
| "App bị lỗi không thanh toán được" | support_agent |
| "Xin chào" | coordinator (phản hồi trực tiếp) |
