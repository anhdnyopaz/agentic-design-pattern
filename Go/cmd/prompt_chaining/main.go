// Package main demonstrates Prompt Chaining pattern using Google ADK-Go
//
// QUAN TRỌNG: Trong ADK-Go, để đạt được Prompt Chaining thực sự (output của bước trước
// là input của bước sau), cần sử dụng MỘT LLM Agent chính (orchestrator) với
// các sub-agents được wrap bằng agenttool.New() làm tools.
//
// Sequential Agent trong ADK-Go chạy các sub-agents tuần tự nhưng mỗi agent
// nhận CÙNG user input - không phải output của agent trước.
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
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// ============================================================================
// DATA TYPES
// ============================================================================

type ProductSpecs struct {
	CPU     string `json:"cpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
	Display string `json:"display,omitempty"`
	Battery string `json:"battery,omitempty"`
	Price   string `json:"price,omitempty"`
}

type EmailContent struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ============================================================================
// TOOLS
// ============================================================================

func createExtractionTool() (tool.Tool, error) {
	type Input struct {
		Text string `json:"text" description:"Văn bản mô tả sản phẩm cần trích xuất"`
	}
	type Output struct {
		ProductSpecs
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
		return Output{}, nil
	}

	return functiontool.New(functiontool.Config{
		Name:        "extract_specifications",
		Description: "Trích xuất thông số kỹ thuật từ văn bản mô tả sản phẩm",
	}, handler)
}

func createTransformTool() (tool.Tool, error) {
	type Input struct {
		ProductSpecs
	}
	type Output struct {
		JSONData string `json:"json_data"`
		Valid    bool   `json:"valid"`
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
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
			return Output{Valid: false}, err
		}
		return Output{Valid: true, JSONData: string(jsonBytes)}, nil
	}

	return functiontool.New(functiontool.Config{
		Name:        "transform_to_json",
		Description: "Chuyển đổi thông số đã trích xuất thành định dạng JSON",
	}, handler)
}

func createEmailTool() (tool.Tool, error) {
	type Input struct {
		ProductName string `json:"product_name"`
		Recipient   string `json:"recipient"`
		ProductSpecs
	}
	type Output struct {
		EmailContent
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
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

		return Output{
			EmailContent: EmailContent{
				Subject: subject,
				Body:    body.String(),
			},
		}, nil
	}

	return functiontool.New(functiontool.Config{
		Name:        "generate_email",
		Description: "Tạo email tóm tắt thông số sản phẩm",
	}, handler)
}

// ============================================================================
// SUB-AGENTS (sẽ được wrap thành tools cho orchestrator)
// ============================================================================

func createExtractionAgent(ctx context.Context, m model.LLM) (agent.Agent, error) {
	extractTool, err := createExtractionTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "extraction_agent",
		Model:       m,
		Description: "Chuyên gia trích xuất thông số kỹ thuật từ văn bản mô tả sản phẩm. Sử dụng agent này để phân tích và trích xuất CPU, RAM, Storage, Display, Battery, Price từ văn bản.",
		Instruction: `Bạn là chuyên gia phân tích thông số kỹ thuật sản phẩm.

NHIỆM VỤ:
Khi nhận văn bản mô tả sản phẩm, hãy trích xuất các thông số:
- CPU/Bộ xử lý
- Memory/RAM
- Storage/Ổ cứng
- Display/Màn hình (nếu có)
- Battery/Pin (nếu có)
- Price/Giá (nếu có)

ĐỊNH DẠNG TRẢ VỀ:
Trả về kết quả dưới dạng danh sách rõ ràng:
- CPU: [giá trị]
- Memory: [giá trị]
- Storage: [giá trị]
- Display: [giá trị hoặc N/A]
- Battery: [giá trị hoặc N/A]
- Price: [giá trị hoặc N/A]`,
		Tools: []tool.Tool{extractTool},
	})
}

func createTransformAgent(ctx context.Context, m model.LLM) (agent.Agent, error) {
	transformTool, err := createTransformTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "transform_agent",
		Model:       m,
		Description: "Chuyên gia chuyển đổi thông số kỹ thuật thành định dạng JSON. Sử dụng agent này sau khi đã có thông số được trích xuất.",
		Instruction: `Bạn là chuyên gia chuyển đổi dữ liệu.

NHIỆM VỤ:
Nhận danh sách thông số kỹ thuật và chuyển thành JSON.

SỬ DỤNG TOOL transform_to_json với các tham số:
- cpu: giá trị CPU
- memory: giá trị RAM
- storage: giá trị ổ cứng
- display: giá trị màn hình (nếu có)
- battery: giá trị pin (nếu có)
- price: giá trị giá (nếu có)`,
		Tools: []tool.Tool{transformTool},
	})
}

func createEmailAgent(ctx context.Context, m model.LLM) (agent.Agent, error) {
	emailTool, err := createEmailTool()
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "email_agent",
		Model:       m,
		Description: "Chuyên gia tạo email tóm tắt thông số sản phẩm. Sử dụng agent này để tạo email gửi đội marketing sau khi đã có JSON thông số.",
		Instruction: `Bạn là chuyên gia soạn thảo email chuyên nghiệp.

NHIỆM VỤ:
Tạo email tóm tắt thông số sản phẩm gửi cho đội Marketing.

SỬ DỤNG TOOL generate_email với:
- product_name: Tên sản phẩm
- recipient: "Đội Marketing"
- cpu, memory, storage, display, battery, price: các thông số từ JSON`,
		Tools: []tool.Tool{emailTool},
	})
}

// ============================================================================
// ORCHESTRATOR AGENT - Agent chính điều phối Prompt Chaining
// ============================================================================

func createOrchestratorAgent(ctx context.Context, m model.LLM, subAgents []agent.Agent) (agent.Agent, error) {
	// Wrap các sub-agents thành tools
	var agentTools []tool.Tool
	for _, subAgent := range subAgents {
		agentTools = append(agentTools, agenttool.New(subAgent, nil))
	}

	return llmagent.New(llmagent.Config{
		Name:        "product_analysis_orchestrator",
		Model:       m,
		Description: "Điều phối viên phân tích sản phẩm - Thực hiện Prompt Chaining",
		Instruction: `Bạn là điều phối viên phân tích sản phẩm. Khi người dùng cung cấp mô tả sản phẩm,
bạn phải thực hiện QUY TRÌNH TUẦN TỰ sau:

**QUY TRÌNH PROMPT CHAINING (BẮT BUỘC THEO THỨ TỰ):**

BƯỚC 1: TRÍCH XUẤT
- Gọi extraction_agent với văn bản mô tả sản phẩm
- Đợi kết quả trả về danh sách thông số

BƯỚC 2: CHUYỂN ĐỔI JSON
- Gọi transform_agent với thông số đã trích xuất từ Bước 1
- Đợi kết quả trả về JSON

BƯỚC 3: TẠO EMAIL
- Gọi email_agent với thông số từ JSON ở Bước 2
- Đợi kết quả trả về email hoàn chỉnh

**QUAN TRỌNG:**
- Phải thực hiện đúng thứ tự: Bước 1 → Bước 2 → Bước 3
- Output của bước trước là input cho bước sau
- Cuối cùng, trình bày kết quả của cả 3 bước cho người dùng

**NẾU NGƯỜI DÙNG CHÀO HỎI:**
Phản hồi thân thiện và hướng dẫn họ cung cấp mô tả sản phẩm để phân tích.

**VÍ DỤ MÔ TẢ SẢN PHẨM:**
"Laptop Dell XPS 15 có CPU Intel Core i9-13900H, RAM 32GB DDR5, SSD 1TB NVMe, màn hình OLED 15.6 inch 3.5K, pin 86Wh. Giá 45.990.000 VNĐ"`,
		Tools: agentTools,
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

	// Tạo các sub-agents
	extractionAgent, err := createExtractionAgent(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	transformAgent, err := createTransformAgent(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	emailAgent, err := createEmailAgent(ctx, geminiModel)
	if err != nil {
		log.Fatal(err)
	}

	// Tạo Orchestrator Agent - đây là cách đúng để làm Prompt Chaining trong ADK-Go
	orchestrator, err := createOrchestratorAgent(ctx, geminiModel, []agent.Agent{
		extractionAgent,
		transformAgent,
		emailAgent,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(orchestrator),
	}

	lch := full.NewLauncher()
	fmt.Println("=== Product Analysis Pipeline - Prompt Chaining Demo ===")
	fmt.Println("Khởi động server...")

	err = lch.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
