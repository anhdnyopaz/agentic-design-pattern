// Package main demonstrates Planning pattern using SequentialAgent in Google ADK-Go
//
// Planning cho phép agent phân rã mục tiêu phức tạp thành các bước có thể thực thi.
// Sử dụng SequentialAgent để điều phối: Planner → Executor
// - Planner: Phân tích goal và tạo kế hoạch JSON
// - Executor: Thực thi từng bước trong kế hoạch
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
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
	statePlan        = "research_plan"
	stateFinalReport = "final_report"
)

// ============================================================================
// DATA TYPES
// ============================================================================

// PlanStep đại diện cho một bước trong kế hoạch
type PlanStep struct {
	StepNumber  int    `json:"step_number"`
	Action      string `json:"action"`
	Description string `json:"description"`
	Expected    string `json:"expected_output"`
	ToolToUse   string `json:"tool_to_use,omitempty"`
}

// ResearchPlan đại diện cho kế hoạch nghiên cứu hoàn chỉnh
type ResearchPlan struct {
	Goal                string     `json:"goal"`
	TotalSteps          int        `json:"total_steps"`
	Steps               []PlanStep `json:"steps"`
	EstimatedComplexity string     `json:"estimated_complexity"`
}

// ============================================================================
// RESEARCH TOOLS
// ============================================================================

// SearchKnowledgeArgs - Arguments cho tool tìm kiếm
type SearchKnowledgeArgs struct {
	Query string `json:"query" description:"Truy vấn tìm kiếm thông tin"`
	Topic string `json:"topic" description:"Chủ đề liên quan"`
}

// SearchKnowledgeResult - Kết quả tìm kiếm
type SearchKnowledgeResult struct {
	Findings   string `json:"findings"`
	Sources    string `json:"sources"`
	Confidence string `json:"confidence"`
}

// SearchKnowledge simulates searching for information
func SearchKnowledge(ctx tool.Context, input SearchKnowledgeArgs) (SearchKnowledgeResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[search_knowledge] Query: %s\n", input.Query)
	fmt.Printf("[search_knowledge] Topic: %s\n", input.Topic)
	fmt.Println("───────────────────────────────────────────────────────")

	return SearchKnowledgeResult{
		Findings:   fmt.Sprintf("Kết quả nghiên cứu về '%s': Thông tin chi tiết về %s bao gồm các khía cạnh quan trọng, xu hướng hiện tại, và các ứng dụng thực tế.", input.Query, input.Topic),
		Sources:    "Academic papers, Industry reports, Expert interviews",
		Confidence: "high",
	}, nil
}

// AnalyzeContentArgs - Arguments cho tool phân tích
type AnalyzeContentArgs struct {
	Content   string `json:"content" description:"Nội dung cần phân tích"`
	Objective string `json:"objective" description:"Mục tiêu phân tích"`
}

// AnalyzeContentResult - Kết quả phân tích
type AnalyzeContentResult struct {
	KeyPoints string `json:"key_points"`
	Insights  string `json:"insights"`
	Gaps      string `json:"gaps"`
}

// AnalyzeContent phân tích nội dung và trích xuất insights
func AnalyzeContent(ctx tool.Context, input AnalyzeContentArgs) (AnalyzeContentResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[analyze_content] Objective: %s\n", input.Objective)
	fmt.Println("───────────────────────────────────────────────────────")

	return AnalyzeContentResult{
		KeyPoints: "1. Điểm quan trọng về kiến trúc\n2. Xu hướng ứng dụng\n3. Thách thức và giải pháp",
		Insights:  "Phân tích cho thấy xu hướng tích hợp AI vào các quy trình nghiệp vụ đang tăng mạnh",
		Gaps:      "Cần nghiên cứu thêm về khả năng mở rộng và bảo mật",
	}, nil
}

// WriteReportArgs - Arguments cho tool viết báo cáo
type WriteReportArgs struct {
	Title    string `json:"title" description:"Tiêu đề báo cáo"`
	Sections string `json:"sections" description:"Các phần nội dung cần viết"`
	Data     string `json:"data" description:"Dữ liệu và findings để đưa vào báo cáo"`
}

// WriteReportResult - Kết quả viết báo cáo
type WriteReportResult struct {
	Report string `json:"report"`
	Status string `json:"status"`
}

// WriteReport tạo báo cáo từ dữ liệu nghiên cứu
func WriteReport(ctx tool.Context, input WriteReportArgs) (WriteReportResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[write_report] Title: %s\n", input.Title)
	fmt.Println("───────────────────────────────────────────────────────")

	report := fmt.Sprintf(`
# %s

## Tóm tắt điều hành
Báo cáo này tổng hợp kết quả nghiên cứu dựa trên phân tích đa chiều.

## Các phần chính
%s

## Dữ liệu và Phân tích
%s

## Kết luận và Khuyến nghị
Dựa trên nghiên cứu, chúng tôi khuyến nghị:
1. Tiếp tục theo dõi xu hướng phát triển
2. Đầu tư vào các giải pháp có tính mở rộng
3. Xây dựng năng lực nội bộ

## Nguồn tham khảo
- Các nguồn học thuật và công nghiệp đã được tham chiếu
`, input.Title, input.Sections, input.Data)

	return WriteReportResult{
		Report: report,
		Status: "completed",
	}, nil
}

// createResearchTools tạo tất cả các tools cho research
func createResearchTools() ([]tool.Tool, error) {
	searchTool, err := functiontool.New(
		functiontool.Config{
			Name:        "search_knowledge",
			Description: "Tìm kiếm thông tin và kiến thức về một chủ đề cụ thể. Sử dụng khi cần thu thập dữ liệu ban đầu.",
		},
		SearchKnowledge,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create search tool: %w", err)
	}

	analyzeTool, err := functiontool.New(
		functiontool.Config{
			Name:        "analyze_content",
			Description: "Phân tích nội dung đã thu thập và trích xuất insights, key points. Sử dụng sau khi có dữ liệu từ search.",
		},
		AnalyzeContent,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create analyze tool: %w", err)
	}

	writeTool, err := functiontool.New(
		functiontool.Config{
			Name:        "write_report",
			Description: "Viết báo cáo tổng hợp từ dữ liệu và phân tích. Sử dụng ở bước cuối để tạo output.",
		},
		WriteReport,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create write tool: %w", err)
	}

	return []tool.Tool{searchTool, analyzeTool, writeTool}, nil
}

// ============================================================================
// PLANNER AGENT
// ============================================================================

func createPlannerAgent(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "ResearchPlanner",
		Model:       m,
		Description: "Lập kế hoạch nghiên cứu chi tiết dựa trên mục tiêu của người dùng.",
		Instruction: `Bạn là một Research Strategist chuyên nghiệp với 15 năm kinh nghiệm.

**VAI TRÒ:**
Phân tích yêu cầu nghiên cứu và tạo kế hoạch chi tiết, có cấu trúc.

**NHIỆM VỤ:**
Khi nhận được yêu cầu nghiên cứu từ người dùng:

1. **Phân tích mục tiêu:** Hiểu rõ người dùng muốn nghiên cứu gì
2. **Xác định phạm vi:** Giới hạn những gì sẽ và không được đề cập
3. **Phân rã thành 3-4 bước cụ thể**
4. **Xác định tool phù hợp cho mỗi bước**

**TOOLS CÓ SẴN:**
- search_knowledge: Tìm kiếm thông tin về chủ đề
- analyze_content: Phân tích và trích xuất insights
- write_report: Viết báo cáo tổng hợp

**OUTPUT FORMAT:**
Trả về kế hoạch dưới dạng JSON:

{
  "goal": "Mục tiêu nghiên cứu cụ thể",
  "total_steps": 3,
  "estimated_complexity": "low/medium/high",
  "steps": [
    {
      "step_number": 1,
      "action": "Thu thập thông tin",
      "description": "Mô tả chi tiết bước này",
      "expected_output": "Dữ liệu thô về chủ đề",
      "tool_to_use": "search_knowledge"
    },
    {
      "step_number": 2,
      "action": "Phân tích dữ liệu",
      "description": "Phân tích thông tin đã thu thập",
      "expected_output": "Insights và key findings",
      "tool_to_use": "analyze_content"
    },
    {
      "step_number": 3,
      "action": "Tổng hợp báo cáo",
      "description": "Viết báo cáo cuối cùng",
      "expected_output": "Báo cáo nghiên cứu hoàn chỉnh",
      "tool_to_use": "write_report"
    }
  ]
}

**QUY TẮC:**
- Tối đa 4 bước
- Mỗi bước có tool cụ thể
- Bước cuối luôn là write_report
- CHỈ output JSON, không giải thích thêm`,
		OutputKey: statePlan,
	})
}

// ============================================================================
// EXECUTOR AGENT
// ============================================================================

func createExecutorAgent(m model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "ResearchExecutor",
		Model:       m,
		Description: "Thực thi kế hoạch nghiên cứu theo từng bước.",
		Instruction: fmt.Sprintf(`Bạn là một Research Analyst chuyên nghiệp.

**VAI TRÒ:**
Thực thi kế hoạch nghiên cứu đã được Planner tạo ra.

**KẾ HOẠCH CẦN THỰC THI:**
{%s}

**QUY TRÌNH THỰC THI:**

1. Đọc kế hoạch từ context (JSON ở trên)
2. Với MỖI BƯỚC trong "steps":
   - Gọi tool được chỉ định trong "tool_to_use"
   - Truyền parameters phù hợp
   - Lưu kết quả để dùng cho bước sau

3. Sau khi thực thi TẤT CẢ các bước:
   - Tổng hợp kết quả
   - Gọi write_report để tạo báo cáo cuối cùng

**TOOLS SỬ DỤNG:**
- search_knowledge(query, topic): Thu thập thông tin
- analyze_content(content, objective): Phân tích dữ liệu
- write_report(title, sections, data): Viết báo cáo

**YÊU CẦU:**
- Thực thi ĐÚNG THỨ TỰ các bước
- Sử dụng output bước trước làm input bước sau
- Output cuối cùng là báo cáo hoàn chỉnh từ write_report`, statePlan),
		Tools:     tools,
		OutputKey: stateFinalReport,
	})
}

// ============================================================================
// PLANNING PIPELINE
// ============================================================================

func createPlanningPipeline(m model.LLM) (agent.Agent, error) {
	// 1. Tạo Planner Agent
	planner, err := createPlannerAgent(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create planner: %w", err)
	}

	// 2. Tạo Research Tools
	tools, err := createResearchTools()
	if err != nil {
		return nil, fmt.Errorf("failed to create tools: %w", err)
	}

	// 3. Tạo Executor Agent với tools
	executor, err := createExecutorAgent(m, tools)
	if err != nil {
		return nil, fmt.Errorf("failed to create executor: %w", err)
	}

	// 4. Tạo Sequential Pipeline: Planner → Executor
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "ResearchPipeline",
			Description: "Pipeline nghiên cứu: Lập kế hoạch → Thực thi → Báo cáo",
			SubAgents:   []agent.Agent{planner, executor},
		},
	})
}

// ============================================================================
// MAIN
// ============================================================================

func printBanner() {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("  Research Assistant - Planning Pattern với ADK-Go")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("")
	fmt.Println("  Pipeline:")
	fmt.Println("  ┌─────────────────────────────────────────────────┐")
	fmt.Println("  │  1. Planner Agent                               │")
	fmt.Println("  │     - Phân tích yêu cầu nghiên cứu              │")
	fmt.Println("  │     - Tạo kế hoạch JSON với các bước            │")
	fmt.Println("  │     → {research_plan}                           │")
	fmt.Println("  │                 ↓                               │")
	fmt.Println("  │  2. Executor Agent                              │")
	fmt.Println("  │     - Đọc kế hoạch từ state                     │")
	fmt.Println("  │     - Thực thi từng bước với tools              │")
	fmt.Println("  │     - Tổng hợp báo cáo cuối cùng                │")
	fmt.Println("  │     → {final_report}                            │")
	fmt.Println("  └─────────────────────────────────────────────────┘")
	fmt.Println("")
	fmt.Println("Ví dụ prompts:")
	fmt.Println("  • 'Nghiên cứu về xu hướng AI Agents năm 2024'")
	fmt.Println("  • 'Phân tích thị trường xe điện tại Việt Nam'")
	fmt.Println("  • 'Tìm hiểu về kiến trúc Microservices'")
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
	pipeline, err := createPlanningPipeline(geminiModel)
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
