package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/parallelagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

// createInputProcessor tạo agent xử lý input và extract tên công ty
func createInputProcessor(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "InputProcessor",
		Model:       m,
		Description: "Xử lý input từ user và extract tên công ty.",
		Instruction: `Bạn là một processor đơn giản. Nhiệm vụ của bạn là extract tên công ty từ câu hỏi của user.

Ví dụ:
- Input: "tôi muốn biết thông tin về công ty vinamilk" → Output: "Vinamilk"
- Input: "phân tích Tesla" → Output: "Tesla"
- Input: "nghiên cứu VinGroup" → Output: "VinGroup"

CHỈ output tên công ty, không thêm gì khác. Nếu không tìm thấy tên công ty, output "UNKNOWN".`,
		OutputKey: "company_name",
	})
}

// createFinancialAnalyst tạo agent phân tích tài chính
func createFinancialAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "FinancialAnalyst",
		Model:       m,
		Description: "Chuyên gia phân tích tài chính công ty.",
		Instruction: `Bạn là chuyên gia phân tích tài chính với 20 năm kinh nghiệm.

**CÔNG TY CẦN PHÂN TÍCH: {company_name}**

Sử dụng Google Search để tìm thông tin tài chính MỚI NHẤT về công ty {company_name}.

NHIỆM VỤ:
- Tìm kiếm dữ liệu doanh thu, lợi nhuận, tăng trưởng của {company_name}
- Phân tích biên lợi nhuận và hiệu quả hoạt động
- Đánh giá dòng tiền và cấu trúc nợ

ĐỊNH DẠNG TRẢ VỀ:
📊 PHÂN TÍCH TÀI CHÍNH - {company_name}
• Dữ liệu tìm được: [tóm tắt từ search]
• Điểm mạnh: [liệt kê]
• Điểm yếu: [liệt kê]
• Đánh giá: [1-10]/10
• Khuyến nghị: [Mua/Giữ/Bán]

Output CHỈ phần phân tích, không thêm lời mở đầu hay kết thúc.`,
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
		OutputKey: "financial_analysis_result",
	})
}

// createMarketAnalyst tạo agent phân tích thị trường
func createMarketAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "MarketAnalyst",
		Model:       m,
		Description: "Chuyên gia phân tích thị trường và cạnh tranh.",
		Instruction: `Bạn là chuyên gia phân tích thị trường với kiến thức sâu rộng.

**CÔNG TY CẦN PHÂN TÍCH: {company_name}**

Sử dụng Google Search để tìm thông tin thị trường MỚI NHẤT về công ty {company_name}.

NHIỆM VỤ:
- Tìm kiếm thị phần và vị thế cạnh tranh của {company_name}
- Phân tích đối thủ chính
- Đánh giá xu hướng ngành và tiềm năng tăng trưởng

ĐỊNH DẠNG TRẢ VỀ:
🏆 PHÂN TÍCH THỊ TRƯỜNG - {company_name}
• Dữ liệu tìm được: [tóm tắt từ search]
• Vị thế: [Dẫn đầu/Top 3/Trung bình/Theo sau]
• Đối thủ chính: [danh sách]
• Xu hướng ngành: [Tăng/Ổn định/Giảm]
• Tiềm năng: [Cao/Trung bình/Thấp]

Output CHỈ phần phân tích, không thêm lời mở đầu hay kết thúc.`,
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
		OutputKey: "market_analysis_result",
	})
}

// createRiskAnalyst tạo agent đánh giá rủi ro
func createRiskAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "RiskAnalyst",
		Model:       m,
		Description: "Chuyên gia đánh giá rủi ro doanh nghiệp.",
		Instruction: `Bạn là chuyên gia quản lý rủi ro với kinh nghiệm đánh giá doanh nghiệp.

**CÔNG TY CẦN PHÂN TÍCH: {company_name}**

Sử dụng Google Search để tìm thông tin về rủi ro và tin tức tiêu cực về công ty {company_name}.

NHIỆM VỤ:
- Tìm kiếm tin tức về rủi ro, vấn đề pháp lý, scandal của {company_name}
- Đánh giá rủi ro hoạt động, pháp lý, thị trường
- Đề xuất biện pháp giảm thiểu

ĐỊNH DẠNG TRẢ VỀ:
⚠️ ĐÁNH GIÁ RỦI RO - {company_name}
• Dữ liệu tìm được: [tóm tắt từ search]
• Rủi ro CAO: [liệt kê nếu có]
• Rủi ro TRUNG BÌNH: [liệt kê]
• Rủi ro THẤP: [liệt kê]
• Điểm rủi ro: [1-10]/10
• Biện pháp giảm thiểu: [khuyến nghị]

Output CHỈ phần phân tích, không thêm lời mở đầu hay kết thúc.`,
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
		OutputKey: "risk_analysis_result",
	})
}

// createSynthesisAgent tạo agent tổng hợp báo cáo
func createSynthesisAgent(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "SynthesisAgent",
		Model:       m,
		Description: "Tổng hợp kết quả phân tích từ các analysts thành báo cáo hoàn chỉnh.",
		Instruction: `Bạn là trưởng nhóm nghiên cứu đầu tư, chịu trách nhiệm tổng hợp báo cáo cuối cùng.

**QUAN TRỌNG**: Toàn bộ response PHẢI dựa HOÀN TOÀN vào thông tin từ các Input bên dưới.
KHÔNG thêm bất kỳ thông tin nào từ bên ngoài.

**Input từ các Analysts:**

📊 **Phân tích Tài chính:**
{financial_analysis_result}

🏆 **Phân tích Thị trường:**
{market_analysis_result}

⚠️ **Đánh giá Rủi ro:**
{risk_analysis_result}

**Output Format:**

═══════════════════════════════════════════════════════
📋 BÁO CÁO NGHIÊN CỨU TỔNG HỢP
(Dựa trên dữ liệu thực từ internet)
═══════════════════════════════════════════════════════

📌 TÓM TẮT ĐIỀU HÀNH
[3-4 câu tóm tắt điểm quan trọng nhất từ 3 phân tích trên]

📊 PHÂN TÍCH TÀI CHÍNH
[Tổng hợp từ Financial Analyst - CHỈ dựa trên input ở trên]

🏆 PHÂN TÍCH THỊ TRƯỜNG
[Tổng hợp từ Market Analyst - CHỈ dựa trên input ở trên]

⚠️ ĐÁNH GIÁ RỦI RO
[Tổng hợp từ Risk Analyst - CHỈ dựa trên input ở trên]

💡 KHUYẾN NGHỊ ĐẦU TƯ
• Đánh giá tổng hợp: [điểm/10 - trung bình từ 3 analysts]
• Khuyến nghị: [MUA/GIỮ/BÁN]
• Lý do: [dựa trên các phân tích ở trên]

═══════════════════════════════════════════════════════

Output CHỈ báo cáo theo format trên, không thêm gì khác.`,
	})
}

// createParallelAnalysts tạo parallel agent chạy các analysts đồng thời
func createParallelAnalysts(analysts ...agent.Agent) (agent.Agent, error) {
	return parallelagent.New(parallelagent.Config{
		AgentConfig: agent.Config{
			Name:        "ParallelAnalystsAgent",
			Description: "Chạy tất cả analysts song song để thu thập thông tin từ nhiều nguồn cùng lúc.",
			SubAgents:   analysts,
		},
	})
}

// createResearchPipeline tạo sequential pipeline orchestrate toàn bộ workflow
func createResearchPipeline(subAgents ...agent.Agent) (agent.Agent, error) {
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "ResearchPipeline",
			Description: "Pipeline nghiên cứu: extract input → chạy parallel analysts → tổng hợp báo cáo.",
			SubAgents:   subAgents,
		},
	})
}

// PipelineAgents chứa các agent cần thiết để build pipeline
type PipelineAgents struct {
	InputProcessor   agent.Agent
	FinancialAnalyst agent.Agent
	MarketAnalyst    agent.Agent
	RiskAnalyst      agent.Agent
	SynthesisAgent   agent.Agent
}

// buildPipeline assemble các agents thành pipeline hoàn chỉnh
func buildPipeline(agents PipelineAgents) (agent.Agent, error) {
	parallelAnalysts, err := createParallelAnalysts(
		agents.FinancialAnalyst,
		agents.MarketAnalyst,
		agents.RiskAnalyst,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parallel agent: %w", err)
	}

	return createResearchPipeline(agents.InputProcessor, parallelAnalysts, agents.SynthesisAgent)
}

func printBanner() {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  Company Research Agent - TRUE Parallelization")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  Architecture:")
	fmt.Println("  ┌─────────────────────────────────────────────────┐")
	fmt.Println("  │            Sequential Pipeline                  │")
	fmt.Println("  │                                                 │")
	fmt.Println("  │  ┌─────────────────────────────────────────┐   │")
	fmt.Println("  │  │  1. Input Processor                     │   │")
	fmt.Println("  │  │     Extract company name → {company_name}│   │")
	fmt.Println("  │  └─────────────────────────────────────────┘   │")
	fmt.Println("  │                      ↓                         │")
	fmt.Println("  │  ┌─────────────────────────────────────────┐   │")
	fmt.Println("  │  │  2. Parallel Analysts (CONCURRENT)      │   │")
	fmt.Println("  │  │  ┌───────────┬───────────┬───────────┐  │   │")
	fmt.Println("  │  │  │ Financial │  Market   │   Risk    │  │   │")
	fmt.Println("  │  │  │  Analyst  │  Analyst  │  Analyst  │  │   │")
	fmt.Println("  │  │  └───────────┴───────────┴───────────┘  │   │")
	fmt.Println("  │  └─────────────────────────────────────────┘   │")
	fmt.Println("  │                      ↓                         │")
	fmt.Println("  │  ┌─────────────────────────────────────────┐   │")
	fmt.Println("  │  │  3. Synthesis Agent                     │   │")
	fmt.Println("  │  │     Combines all results into report    │   │")
	fmt.Println("  │  └─────────────────────────────────────────┘   │")
	fmt.Println("  └─────────────────────────────────────────────────┘")
	fmt.Println("")
	fmt.Println("Khởi động server...")
	fmt.Println("Nhập tên công ty để phân tích (ví dụ: 'Tesla', 'VinGroup')")
}

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatalln("GOOGLE_API_KEY environment variable not set")
	}

	// Tạo model - dễ dàng tuỳ chỉnh config tại đây
	geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Không thể tạo Model: %v", err)
	}

	// Tạo các agents - dễ dàng tuỳ chỉnh từng agent tại đây
	inputProcessor, err := createInputProcessor(geminiModel)
	if err != nil {
		log.Fatalf("Failed to create input processor: %v", err)
	}

	financialAnalyst, err := createFinancialAnalyst(geminiModel)
	if err != nil {
		log.Fatalf("Failed to create financial analyst: %v", err)
	}

	marketAnalyst, err := createMarketAnalyst(geminiModel)
	if err != nil {
		log.Fatalf("Failed to create market analyst: %v", err)
	}

	riskAnalyst, err := createRiskAnalyst(geminiModel)
	if err != nil {
		log.Fatalf("Failed to create risk analyst: %v", err)
	}

	synthesisAgent, err := createSynthesisAgent(geminiModel)
	if err != nil {
		log.Fatalf("Failed to create synthesis agent: %v", err)
	}

	// Assemble pipeline từ các agents
	pipeline, err := buildPipeline(PipelineAgents{
		InputProcessor:   inputProcessor,
		FinancialAnalyst: financialAnalyst,
		MarketAnalyst:    marketAnalyst,
		RiskAnalyst:      riskAnalyst,
		SynthesisAgent:   synthesisAgent,
	})
	if err != nil {
		log.Fatalf("Failed to build pipeline: %v", err)
	}

	cfg := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(pipeline),
	}

	printBanner()

	lch := full.NewLauncher()
	if err := lch.Execute(ctx, cfg, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
