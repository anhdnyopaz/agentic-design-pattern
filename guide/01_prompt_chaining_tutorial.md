# Hướng Dẫn Thực Hành: Prompt Chaining với ADK-Go

## Mục Tiêu Bài Học

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu rõ mẫu thiết kế Prompt Chaining
- Biết cách sử dụng Google ADK-Go để xây dựng chuỗi prompt
- Tự tay xây dựng được một pipeline xử lý dữ liệu hoàn chỉnh

---

## Phần 1: Tổng Quan Về Prompt Chaining

### 1.1 Prompt Chaining là gì?

Prompt Chaining (Chuỗi Lời Nhắc) là kỹ thuật **chia nhỏ tác vụ phức tạp** thành nhiều bước tuần tự, trong đó:

```
[Input] → [Bước 1] → [Output 1] → [Bước 2] → [Output 2] → [Bước 3] → [Kết quả cuối]
```

**Ví dụ thực tế:**
- Bước 1: Trích xuất thông số từ văn bản mô tả sản phẩm
- Bước 2: Chuyển đổi thành JSON có cấu trúc
- Bước 3: Tạo email tóm tắt gửi đội marketing

### 1.2 Tại sao cần Prompt Chaining?

| Vấn đề với Single Prompt | Giải pháp với Prompt Chaining |
|--------------------------|-------------------------------|
| Bỏ qua hướng dẫn (instruction neglect) | Mỗi bước tập trung một nhiệm vụ |
| Mất ngữ cảnh (context drift) | Ngữ cảnh được truyền rõ ràng giữa các bước |
| Lan truyền lỗi | Dễ debug từng bước riêng biệt |
| Ảo giác (hallucination) | Giảm tải nhận thức cho model |

---

## Phần 2: Chuẩn Bị Môi Trường

### 2.1 Cài đặt Go

```bash
# Kiểm tra phiên bản Go (yêu cầu >= 1.21)
go version
```

### 2.2 Khởi tạo project

```bash
mkdir prompt-chaining-demo
cd prompt-chaining-demo
go mod init prompt-chaining-demo
```

### 2.3 Cài đặt ADK-Go

```bash
go get google.golang.org/adk@latest
go get google.golang.org/genai
```

### 2.4 Thiết lập API Key

```bash
export GOOGLE_API_KEY="your-api-key-here"
```

---

## Phần 3: Cấu Trúc Project

Tạo cấu trúc thư mục như sau:

```
prompt-chaining-demo/
├── go.mod
├── go.sum
├── main.go
└── internal/
    ├── agents/
    │   ├── extraction.go    # Agent trích xuất
    │   ├── transform.go     # Agent chuyển đổi
    │   └── email.go         # Agent tạo email
    ├── models/
    │   └── types.go         # Định nghĩa struct
    └── tools/
        └── tools.go         # Định nghĩa tools
```

---

## Phần 4: Bài Tập Thực Hành

### Bài 1: Định nghĩa Data Types

**Yêu cầu:** Tạo file `internal/models/types.go` với các struct sau:

```go
// TODO: Hoàn thành các struct

// ProductSpecs - Thông số sản phẩm
// Cần có: CPU, Memory, Storage, Display (optional), Battery (optional), Price (optional)
type ProductSpecs struct {
    // Viết code ở đây
}

// EmailContent - Nội dung email
// Cần có: Subject, Body
type EmailContent struct {
    // Viết code ở đây
}
```

**Gợi ý:**
- Sử dụng JSON tags: `json:"field_name"`
- Sử dụng `omitempty` cho các trường optional

**Kết quả mong đợi:**
```go
type ProductSpecs struct {
    CPU     string `json:"cpu"`
    Memory  string `json:"memory"`
    Storage string `json:"storage"`
    Display string `json:"display,omitempty"`
    Battery string `json:"battery,omitempty"`
    Price   string `json:"price,omitempty"`
}
```

---

### Bài 2: Tạo Function Tool

**Yêu cầu:** Tạo một function tool để trích xuất thông số

**Kiến thức cần biết:**

```go
import "google.golang.org/adk/tool/functiontool"

// Cấu trúc tạo tool:
tool, err := functiontool.New(functiontool.Config{
    Name:        "tên_tool",
    Description: "Mô tả tool",
}, handlerFunction)
```

**Bài tập:**

```go
// TODO: Tạo function tool extract_specifications
// Input: Text string (văn bản mô tả sản phẩm)
// Output: ProductSpecs struct

func createExtractionTool() (tool.Tool, error) {
    // Định nghĩa Input struct
    type Input struct {
        // Viết code ở đây
    }

    // Định nghĩa Output struct
    type Output struct {
        // Viết code ở đây
    }

    // Tạo handler function
    handler := func(ctx tool.Context, input Input) Output {
        // Viết logic xử lý ở đây
    }

    // Tạo và return tool
    return functiontool.New(/* config */, handler)
}
```

---

### Bài 3: Tạo LLM Agent

**Yêu cầu:** Tạo agent với instruction tiếng Việt

**Kiến thức cần biết:**

```go
import "google.golang.org/adk/agent/llmagent"

agent, err := llmagent.New(llmagent.Config{
    Name:        "agent_name",
    Model:       model,              // Gemini model
    Description: "Mô tả agent",
    Instruction: "Hướng dẫn cho agent",
    Tools:       []tool.Tool{...},   // Danh sách tools
})
```

**Bài tập:**

```go
// TODO: Tạo 3 agents

// Agent 1: Extraction Agent
// - Vai trò: "Chuyên gia phân tích sản phẩm"
// - Nhiệm vụ: Trích xuất CPU, Memory, Storage từ văn bản

// Agent 2: Transform Agent
// - Vai trò: "Chuyên gia chuyển đổi dữ liệu"
// - Nhiệm vụ: Chuyển dữ liệu sang JSON

// Agent 3: Email Agent
// - Vai trò: "Chuyên gia soạn thảo email"
// - Nhiệm vụ: Tạo email tóm tắt chuyên nghiệp
```

**Template instruction mẫu:**

```go
instruction := `Bạn là [VAI TRÒ].
Nhiệm vụ: [MÔ TẢ NHIỆM VỤ]

Yêu cầu:
- [Yêu cầu 1]
- [Yêu cầu 2]

Ví dụ đầu vào: [INPUT MẪU]
Ví dụ đầu ra: [OUTPUT MẪU]`
```

---

### Bài 4: Kết Nối Chuỗi với Sequential Agent

**Đây là phần quan trọng nhất!**

**Kiến thức cần biết:**

```go
import "google.golang.org/adk/agent/workflowagents/sequentialagent"

// Sequential Agent thực thi các sub-agents theo thứ tự
chainedAgent, err := sequentialagent.New(sequentialagent.Config{
    AgentConfig: agent.Config{
        Name:        "chain_name",
        Description: "Mô tả chuỗi",
        SubAgents:   []agent.Agent{agent1, agent2, agent3},
    },
})
```

**Bài tập:**

```go
// TODO: Hoàn thành hàm main

func main() {
    ctx := context.Background()

    // 1. Lấy API key từ environment
    apiKey := os.Getenv("GOOGLE_API_KEY")

    // 2. Tạo Gemini model
    model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
        APIKey: apiKey,
    })

    // 3. Tạo 3 agents (gọi các hàm đã tạo ở Bài 3)
    // extractionAgent := ...
    // transformAgent := ...
    // emailAgent := ...

    // 4. Tạo Sequential Agent để nối chuỗi
    // chainedAgent := ...

    // 5. Cấu hình và khởi chạy
    config := &adk.Config{
        AgentLoader: services.NewSingleAgentLoader(chainedAgent),
    }

    launcher := full.NewLauncher()
    launcher.Execute(ctx, config, os.Args[1:])
}
```

---

## Phần 5: Chạy Và Test

### 5.1 Chạy ứng dụng

```bash
go run main.go web
```

Truy cập: `http://localhost:8080`

### 5.2 Test với input mẫu

```
Laptop Dell XPS 15 mới ra mắt với cấu hình mạnh mẽ:
CPU Intel Core i9-13900H 14 nhân, RAM 32GB DDR5 4800MHz,
SSD 1TB NVMe PCIe 4.0, màn hình OLED 15.6 inch 3.5K,
pin 86Wh sử dụng lên đến 13 tiếng. Giá bán 45.990.000 VNĐ.
```

### 5.3 Kết quả mong đợi

```
=== Bước 1: Trích xuất ===
- CPU: Intel Core i9-13900H 14 nhân
- Memory: 32GB DDR5 4800MHz
- Storage: 1TB NVMe PCIe 4.0
- Display: OLED 15.6 inch 3.5K
- Battery: 86Wh, 13 tiếng
- Price: 45.990.000 VNĐ

=== Bước 2: JSON ===
{
  "cpu": "Intel Core i9-13900H 14 nhân",
  "memory": "32GB DDR5 4800MHz",
  "storage": "1TB NVMe PCIe 4.0",
  "display": "OLED 15.6 inch 3.5K",
  "battery": "86Wh, 13 tiếng",
  "price": "45.990.000 VNĐ"
}

=== Bước 3: Email ===
Subject: Phân Tích Sản Phẩm: Dell XPS 15
Body: Kính gửi đội Marketing, ...
```

---

## Phần 6: Các Import Cần Thiết

```go
import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strings"

    "google.golang.org/adk/agent"
    "google.golang.org/adk/agent/llmagent"
    "google.golang.org/adk/agent/workflowagents/sequentialagent"
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

## Phần 7: Bài Tập Nâng Cao

### Nâng cao 1: Thêm Error Handling

Thêm xử lý lỗi giữa các bước:
- Nếu bước 1 không trích xuất được CPU → dừng và báo lỗi
- Nếu JSON không hợp lệ → retry hoặc báo lỗi

### Nâng cao 2: Thêm Validation Step

Thêm bước 2.5: Kiểm tra dữ liệu trước khi tạo email
- Validate các trường bắt buộc
- Chuẩn hóa định dạng

### Nâng cao 3: Parallel Processing

Sử dụng `parallelagent` thay vì `sequentialagent` cho các bước độc lập:

```go
import "google.golang.org/adk/agent/workflowagents/parallelagent"
```

### Nâng cao 4: Callback Functions

Thêm callback để log hoặc lưu kết quả mỗi bước:

```go
AfterModelCallbacks: []llmagent.AfterModelCallback{
    func(ctx agent.CallbackContext, resp *model.LLMResponse, err error) (*model.LLMResponse, error) {
        // Log kết quả
        return resp, err
    },
},
```

---

## Phần 8: Tài Liệu Tham Khảo

1. **ADK-Go GitHub**: https://github.com/google/adk-go
2. **Gemini API**: https://ai.google.dev/docs
3. **Go Modules**: https://go.dev/ref/mod

---

## Checklist Hoàn Thành

- [ ] Hiểu khái niệm Prompt Chaining
- [ ] Tạo được struct ProductSpecs và EmailContent
- [ ] Tạo được function tool với functiontool.New()
- [ ] Tạo được LLM Agent với instruction tiếng Việt
- [ ] Kết nối 3 agents bằng sequentialagent
- [ ] Chạy và test thành công
- [ ] Hoàn thành ít nhất 1 bài tập nâng cao

---

## Ghi Chú

- Đọc kỹ error message khi gặp lỗi
- Kiểm tra API key đã được set chưa
- Sử dụng `go mod tidy` nếu gặp lỗi dependency

---

## Phần 9: Code Mẫu Đầy Đủ (Solution)

> **Lưu ý:** Hãy tự làm bài tập trước khi xem code mẫu!

```go
// cmd/prompt_chaining/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
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
// DATA TYPES
// ============================================================================

// ProductSpecs - Thông số sản phẩm được trích xuất
type ProductSpecs struct {
	CPU     string `json:"cpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
	Display string `json:"display,omitempty"`
	Battery string `json:"battery,omitempty"`
	Price   string `json:"price,omitempty"`
}

// EmailContent - Nội dung email đầu ra
type EmailContent struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ============================================================================
// TOOLS
// ============================================================================

// createExtractionTool - Tool trích xuất thông số từ văn bản
func createExtractionTool() (tool.Tool, error) {
	type Input struct {
		Text string `json:"text" description:"Văn bản mô tả sản phẩm cần trích xuất"`
	}
	type Output struct {
		CPU     string `json:"cpu"`
		Memory  string `json:"memory"`
		Storage string `json:"storage"`
		Display string `json:"display"`
		Battery string `json:"battery"`
		Price   string `json:"price"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		// Tool này chủ yếu để LLM biết cấu trúc output mong muốn
		// Logic trích xuất thực tế do LLM thực hiện
		return Output{}
	}

	return functiontool.New(functiontool.Config{
		Name:        "extract_specifications",
		Description: "Trích xuất thông số kỹ thuật từ văn bản mô tả sản phẩm",
	}, handler)
}

// createTransformTool - Tool chuyển đổi sang JSON
func createTransformTool() (tool.Tool, error) {
	type Input struct {
		CPU     string `json:"cpu"`
		Memory  string `json:"memory"`
		Storage string `json:"storage"`
		Display string `json:"display"`
		Battery string `json:"battery"`
		Price   string `json:"price"`
	}
	type Output struct {
		JSONData string `json:"json_data"`
		Valid    bool   `json:"valid"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		specs := ProductSpecs{
			CPU:     input.CPU,
			Memory:  input.Memory,
			Storage: input.Storage,
			Display: input.Display,
			Battery: input.Battery,
			Price:   input.Price,
		}
		jsonBytes, err := json.MarshalIndent(specs, "", "  ")
		if err != nil {
			return Output{Valid: false}
		}
		return Output{JSONData: string(jsonBytes), Valid: true}
	}

	return functiontool.New(functiontool.Config{
		Name:        "transform_to_json",
		Description: "Chuyển đổi thông số đã trích xuất thành định dạng JSON",
	}, handler)
}

// createEmailTool - Tool tạo email
func createEmailTool() (tool.Tool, error) {
	type Input struct {
		ProductName string `json:"product_name"`
		CPU         string `json:"cpu"`
		Memory      string `json:"memory"`
		Storage     string `json:"storage"`
		Display     string `json:"display"`
		Battery     string `json:"battery"`
		Price       string `json:"price"`
		Recipient   string `json:"recipient"`
	}
	type Output struct {
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	handler := func(ctx tool.Context, input Input) Output {
		subject := fmt.Sprintf("Phân Tích Sản Phẩm: %s", input.ProductName)

		var body strings.Builder
		body.WriteString(fmt.Sprintf("Kính gửi %s,\n\n", input.Recipient))
		body.WriteString(fmt.Sprintf("Dưới đây là phân tích kỹ thuật cho %s:\n\n", input.ProductName))
		body.WriteString("THÔNG SỐ KỸ THUẬT:\n")
		body.WriteString(fmt.Sprintf("• CPU: %s\n", input.CPU))
		body.WriteString(fmt.Sprintf("• RAM: %s\n", input.Memory))
		body.WriteString(fmt.Sprintf("• Ổ cứng: %s\n", input.Storage))
		if input.Display != "" {
			body.WriteString(fmt.Sprintf("• Màn hình: %s\n", input.Display))
		}
		if input.Battery != "" {
			body.WriteString(fmt.Sprintf("• Pin: %s\n", input.Battery))
		}
		if input.Price != "" {
			body.WriteString(fmt.Sprintf("• Giá: %s\n", input.Price))
		}
		body.WriteString("\nTrân trọng,\nĐội Phân Tích Sản Phẩm")

		return Output{Subject: subject, Body: body.String()}
	}

	return functiontool.New(functiontool.Config{
		Name:        "generate_email",
		Description: "Tạo email tóm tắt thông số sản phẩm",
	}, handler)
}

// ============================================================================
// AGENTS
// ============================================================================

// createExtractionAgent - Bước 1: Trích xuất thông số
func createExtractionAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	extractTool, err := createExtractionTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "extraction_agent",
		Model:       m,
		Description: "Chuyên gia phân tích sản phẩm - Trích xuất thông số kỹ thuật",
		Instruction: `Bạn là chuyên gia phân tích thông số kỹ thuật sản phẩm.

NHIỆM VỤ:
Đọc văn bản mô tả sản phẩm và trích xuất các thông số sau:
- CPU/Bộ xử lý
- Memory/RAM
- Storage/Ổ cứng
- Display/Màn hình (nếu có)
- Battery/Pin (nếu có)
- Price/Giá (nếu có)

YÊU CẦU:
- Trích xuất chính xác giá trị từ văn bản
- Nếu không tìm thấy thông số nào, để trống
- Trả về dưới dạng danh sách rõ ràng

VÍ DỤ:
Input: "Laptop có CPU Intel i7, RAM 16GB, SSD 512GB"
Output:
- CPU: Intel i7
- Memory: 16GB
- Storage: 512GB SSD`,
		Tools: []tool.Tool{extractTool},
	})
}

// createTransformAgent - Bước 2: Chuyển đổi JSON
func createTransformAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	transformTool, err := createTransformTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "transform_agent",
		Model:       m,
		Description: "Chuyên gia chuyển đổi dữ liệu - Định dạng JSON",
		Instruction: `Bạn là chuyên gia chuyển đổi dữ liệu.

NHIỆM VỤ:
Nhận thông số kỹ thuật đã trích xuất từ bước trước và chuyển thành JSON.

SCHEMA JSON:
{
  "cpu": "string",
  "memory": "string",
  "storage": "string",
  "display": "string (optional)",
  "battery": "string (optional)",
  "price": "string (optional)"
}

YÊU CẦU:
- JSON phải hợp lệ và có thể parse được
- Giữ nguyên giá trị đã trích xuất
- Sử dụng tool transform_to_json để tạo JSON`,
		Tools: []tool.Tool{transformTool},
	})
}

// createEmailAgent - Bước 3: Tạo email
func createEmailAgent(ctx context.Context, m model.Model) (agent.Agent, error) {
	emailTool, err := createEmailTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "email_agent",
		Model:       m,
		Description: "Chuyên gia soạn thảo email chuyên nghiệp",
		Instruction: `Bạn là chuyên gia soạn thảo email chuyên nghiệp.

NHIỆM VỤ:
Tạo email tóm tắt thông số sản phẩm gửi cho đội Marketing.

YÊU CẦU:
- Tiêu đề ngắn gọn, chuyên nghiệp
- Nội dung rõ ràng, dễ đọc
- Liệt kê đầy đủ thông số đã có
- Ngôn ngữ: Tiếng Việt
- Kết thúc với lời chào phù hợp

SỬ DỤNG:
Tool generate_email với:
- product_name: Tên sản phẩm (lấy từ ngữ cảnh)
- recipient: "Đội Marketing"
- Các thông số từ JSON ở bước trước`,
		Tools: []tool.Tool{emailTool},
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

	// 3. Tạo các agents cho mỗi bước
	extractionAgent, err := createExtractionAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo extraction agent: %v", err)
	}

	transformAgent, err := createTransformAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo transform agent: %v", err)
	}

	emailAgent, err := createEmailAgent(ctx, geminiModel)
	if err != nil {
		log.Fatalf("Không thể tạo email agent: %v", err)
	}

	// 4. Tạo Sequential Agent - Kết nối chuỗi
	chainedAgent, err := sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "product_analysis_pipeline",
			Description: "Pipeline phân tích sản phẩm - Prompt Chaining Pattern",
			SubAgents: []agent.Agent{
				extractionAgent, // Bước 1: Trích xuất
				transformAgent,  // Bước 2: Chuyển đổi JSON
				emailAgent,      // Bước 3: Tạo email
			},
		},
	})
	if err != nil {
		log.Fatalf("Không thể tạo sequential agent: %v", err)
	}

	// 5. Cấu hình và khởi chạy
	config := &adk.Config{
		AgentLoader: services.NewSingleAgentLoader(chainedAgent),
	}

	launcher := full.NewLauncher()
	err = launcher.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatalf("Lỗi chạy ứng dụng: %v\n\nCú pháp: %s", err, launcher.CommandLineSyntax())
	}
}
```

### Giải thích Code

**1. Data Types (Dòng 24-37)**
- `ProductSpecs`: Struct lưu thông số sản phẩm với JSON tags
- `omitempty`: Bỏ qua trường nếu rỗng khi serialize JSON

**2. Tools (Dòng 43-130)**
- Mỗi tool có Input/Output struct riêng
- `functiontool.New()` tạo tool từ handler function
- Handler có thể chứa logic xử lý hoặc chỉ định nghĩa schema

**3. Agents (Dòng 136-220)**
- Mỗi agent có vai trò và instruction riêng
- `llmagent.New()` tạo agent với model và tools
- Instruction viết rõ ràng: NHIỆM VỤ, YÊU CẦU, VÍ DỤ

**4. Sequential Agent (Dòng 250-265)**
- `sequentialagent.New()` tạo chuỗi từ các sub-agents
- SubAgents được thực thi theo thứ tự trong array
- Output của agent trước là context cho agent sau

**5. Launcher (Dòng 267-280)**
- `full.NewLauncher()` tạo web server với UI
- `launcher.Execute()` chạy với các CLI arguments
- Chạy `go run main.go web` để start web interface
