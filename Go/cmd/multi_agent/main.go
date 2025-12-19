// Package main demonstrates Multi-Agent Collaboration pattern using ADK-Go
//
// Multi-Agent Collaboration cho phép nhiều agent chuyên biệt làm việc cùng nhau.
// Ví dụ này demo 2 patterns:
// 1. Sequential: Blog Creation Team (Researcher → Writer → Editor)
// 2. Parallel + Sequential: Market Analysis Team (3 analysts song song → Aggregator)
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
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// ============================================================================
// STATE KEYS
// ============================================================================

// Blog Creation Team
const (
	stateResearchData = "research_data"
	stateDraftContent = "draft_content"
	stateFinalArticle = "final_article"
)

// Market Analysis Team
const (
	stateTechAnalysis = "tech_analysis"
	stateSentiment    = "sentiment_analysis"
	stateFinancial    = "financial_analysis"
	stateMarketReport = "market_report"
)

// ============================================================================
// TOOLS FOR BLOG CREATION TEAM
// ============================================================================

// SearchTopicArgs - Arguments cho tool tìm kiếm
type SearchTopicArgs struct {
	Topic    string `json:"topic" description:"Chủ đề cần tìm kiếm"`
	Keywords string `json:"keywords" description:"Từ khóa bổ sung"`
}

// SearchTopicResult - Kết quả tìm kiếm
type SearchTopicResult struct {
	Findings  string `json:"findings"`
	Trends    string `json:"trends"`
	KeyPoints string `json:"key_points"`
}

// SearchTopic simulates searching for topic information
func SearchTopic(ctx tool.Context, input SearchTopicArgs) (SearchTopicResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[search_topic] Topic: %s, Keywords: %s\n", input.Topic, input.Keywords)
	fmt.Println("───────────────────────────────────────────────────────")

	return SearchTopicResult{
		Findings:  fmt.Sprintf("Nghiên cứu về '%s' cho thấy nhiều phát triển quan trọng trong lĩnh vực này. Các chuyên gia nhận định đây là xu hướng đáng chú ý với tiềm năng lớn.", input.Topic),
		Trends:    "1. Xu hướng tăng trưởng mạnh mẽ\n2. Ứng dụng thực tế ngày càng nhiều\n3. Đầu tư R&D gia tăng đáng kể",
		KeyPoints: "- Điểm chính 1: Tầm quan trọng trong công nghiệp\n- Điểm chính 2: Ảnh hưởng đến người dùng cuối\n- Điểm chính 3: Triển vọng tương lai sáng sủa",
	}, nil
}

// WriteSectionArgs - Arguments cho tool viết
type WriteSectionArgs struct {
	SectionTitle string `json:"section_title" description:"Tiêu đề phần"`
	Content      string `json:"content" description:"Nội dung cần viết"`
}

// WriteSectionResult - Kết quả viết
type WriteSectionResult struct {
	Section   string `json:"section"`
	WordCount int    `json:"word_count"`
}

// WriteSection writes a section of the article
func WriteSection(ctx tool.Context, input WriteSectionArgs) (WriteSectionResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[write_section] Writing section: %s\n", input.SectionTitle)
	fmt.Println("───────────────────────────────────────────────────────")

	section := fmt.Sprintf("## %s\n\n%s\n", input.SectionTitle, input.Content)
	return WriteSectionResult{
		Section:   section,
		WordCount: len(input.Content) / 5,
	}, nil
}

// EditContentArgs - Arguments cho tool chỉnh sửa
type EditContentArgs struct {
	Content  string `json:"content" description:"Nội dung cần chỉnh sửa"`
	EditType string `json:"edit_type" description:"Loại chỉnh sửa: grammar, style, structure"`
}

// EditContentResult - Kết quả chỉnh sửa
type EditContentResult struct {
	EditedContent string `json:"edited_content"`
	Changes       string `json:"changes"`
}

// EditContent edits and improves content
func EditContent(ctx tool.Context, input EditContentArgs) (EditContentResult, error) {
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("[edit_content] Edit type: %s\n", input.EditType)
	fmt.Println("───────────────────────────────────────────────────────")

	return EditContentResult{
		EditedContent: input.Content + "\n\n---\n*Đã được review và chỉnh sửa bởi Editor*",
		Changes:       "- Cải thiện cấu trúc câu\n- Thêm transitions mượt mà\n- Tối ưu hóa cho SEO",
	}, nil
}

// createBlogTools tạo tools cho Blog Creation Team
func createBlogTools() (research, writer, editor []tool.Tool, err error) {
	searchTool, err := functiontool.New(
		functiontool.Config{
			Name:        "search_topic",
			Description: "Tìm kiếm và phân tích thông tin về một chủ đề cụ thể",
		},
		SearchTopic,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create search tool: %w", err)
	}

	writeTool, err := functiontool.New(
		functiontool.Config{
			Name:        "write_section",
			Description: "Viết một phần nội dung bài viết với tiêu đề và nội dung",
		},
		WriteSection,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create write tool: %w", err)
	}

	editTool, err := functiontool.New(
		functiontool.Config{
			Name:        "edit_content",
			Description: "Chỉnh sửa và cải thiện nội dung bài viết",
		},
		EditContent,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create edit tool: %w", err)
	}

	return []tool.Tool{searchTool}, []tool.Tool{writeTool}, []tool.Tool{editTool}, nil
}

// ============================================================================
// BLOG CREATION TEAM (SEQUENTIAL PATTERN)
// ============================================================================

func createResearcherAgent(m model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "ResearcherAgent",
		Model:       m,
		Description: "Chuyên gia nghiên cứu và thu thập thông tin về chủ đề.",
		Instruction: `Bạn là một Senior Research Analyst với 10 năm kinh nghiệm.

**VAI TRÒ:**
Nghiên cứu và thu thập thông tin toàn diện về chủ đề được yêu cầu.

**NHIỆM VỤ:**
1. Nhận yêu cầu về chủ đề cần viết bài
2. Sử dụng tool search_topic để tìm kiếm thông tin
3. Phân tích và tổng hợp các findings
4. Output research data có cấu trúc rõ ràng

**OUTPUT FORMAT:**
Trả về research data bao gồm:
- Tổng quan về chủ đề
- Xu hướng chính (3-5 điểm)
- Key insights quan trọng
- Nguồn tham khảo

**QUY TẮC:**
- Tập trung vào thông tin chính xác và có nguồn
- Highlight các điểm quan trọng nhất
- Chuẩn bị data để Writer có thể sử dụng ngay`,
		Tools:     tools,
		OutputKey: stateResearchData,
	})
}

func createWriterAgent(m model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "WriterAgent",
		Model:       m,
		Description: "Chuyên gia viết nội dung dựa trên research data.",
		Instruction: fmt.Sprintf(`Bạn là một Technical Content Writer chuyên nghiệp.

**VAI TRÒ:**
Viết bài blog chất lượng cao dựa trên research data đã được chuẩn bị.

**RESEARCH DATA TỪ RESEARCHER:**
{%s}

**NHIỆM VỤ:**
1. Đọc và hiểu kỹ research data
2. Tạo outline cho bài viết
3. Sử dụng tool write_section để viết từng phần:
   - Introduction (mở đầu hấp dẫn, hook người đọc)
   - Main content (2-3 sections chính dựa trên research)
   - Conclusion (kết luận và call-to-action)
4. Kết hợp thành bài viết hoàn chỉnh

**YÊU CẦU:**
- Giọng văn engaging và dễ đọc
- Sử dụng ví dụ cụ thể từ research
- Độ dài khoảng 300-500 từ
- Format: Markdown với headers rõ ràng

**OUTPUT:**
Draft bài viết hoàn chỉnh, sẵn sàng cho Editor review.`, stateResearchData),
		Tools:     tools,
		OutputKey: stateDraftContent,
	})
}

func createEditorAgent(m model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "EditorAgent",
		Model:       m,
		Description: "Chuyên gia chỉnh sửa và hoàn thiện bài viết.",
		Instruction: fmt.Sprintf(`Bạn là một Chief Editor với tiêu chuẩn xuất bản cao.

**VAI TRÒ:**
Review và polish bài viết để đạt chất lượng xuất bản.

**DRAFT TỪ WRITER:**
{%s}

**NHIỆM VỤ:**
1. Đọc kỹ draft từ Writer
2. Sử dụng tool edit_content để chỉnh sửa:
   - Grammar và spelling check
   - Style improvement
   - Structure optimization
3. Đảm bảo các tiêu chí:
   - Không có lỗi chính tả/ngữ pháp
   - Flow mạch lạc giữa các phần
   - Tone nhất quán xuyên suốt
   - SEO-friendly

**CHECKLIST CUỐI CÙNG:**
- [ ] Tiêu đề hấp dẫn và rõ ràng
- [ ] Introduction hook người đọc
- [ ] Nội dung cung cấp giá trị thực
- [ ] Kết luận có call-to-action
- [ ] Format Markdown đúng chuẩn

**OUTPUT:**
Bài viết final, sẵn sàng publish.`, stateDraftContent),
		Tools:     tools,
		OutputKey: stateFinalArticle,
	})
}

func createBlogCreationTeam(m model.LLM) (agent.Agent, error) {
	researchTools, writerTools, editorTools, err := createBlogTools()
	if err != nil {
		return nil, fmt.Errorf("failed to create tools: %w", err)
	}

	researcher, err := createResearcherAgent(m, researchTools)
	if err != nil {
		return nil, fmt.Errorf("failed to create researcher: %w", err)
	}

	writer, err := createWriterAgent(m, writerTools)
	if err != nil {
		return nil, fmt.Errorf("failed to create writer: %w", err)
	}

	editor, err := createEditorAgent(m, editorTools)
	if err != nil {
		return nil, fmt.Errorf("failed to create editor: %w", err)
	}

	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "BlogCreationTeam",
			Description: "Team tạo blog chuyên nghiệp: Researcher → Writer → Editor",
			SubAgents:   []agent.Agent{researcher, writer, editor},
		},
	})
}

// ============================================================================
// MARKET ANALYSIS TEAM (PARALLEL + SEQUENTIAL PATTERN)
// ============================================================================

func createTechnicalAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "TechnicalAnalyst",
		Model:       m,
		Description: "Phân tích kỹ thuật thị trường.",
		Instruction: `Bạn là Technical Analyst chuyên về chart patterns và technical indicators.

**NHIỆM VỤ:**
Phân tích kỹ thuật thị trường dựa trên:
- Trend analysis (xu hướng ngắn/trung/dài hạn)
- Support và Resistance levels
- Volume patterns
- Key indicators: RSI, MACD, Moving Averages

**OUTPUT FORMAT:**
## Technical Analysis

**Trend:** Bullish/Bearish/Sideways
**Timeframe:** Short/Medium/Long term

**Key Levels:**
- Support: [levels]
- Resistance: [levels]

**Indicators:**
- RSI: [value và interpretation]
- MACD: [signal]

**Signal:** BUY/SELL/HOLD
**Confidence:** High/Medium/Low`,
		OutputKey: stateTechAnalysis,
	})
}

func createSentimentAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "SentimentAnalyst",
		Model:       m,
		Description: "Phân tích cảm xúc và tâm lý thị trường.",
		Instruction: `Bạn là Sentiment Analyst chuyên về market psychology.

**NHIỆM VỤ:**
Phân tích cảm xúc thị trường dựa trên:
- News sentiment (tích cực/tiêu cực)
- Social media trends
- Fear & Greed index
- Institutional behavior

**OUTPUT FORMAT:**
## Sentiment Analysis

**Overall Sentiment:** Positive/Negative/Neutral
**Fear & Greed:** [level]

**News Impact:**
- Positive factors: [list]
- Negative factors: [list]

**Social Trends:**
[Key observations]

**Contrarian Signals:**
[Nếu có]`,
		OutputKey: stateSentiment,
	})
}

func createFinancialAnalyst(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "FinancialAnalyst",
		Model:       m,
		Description: "Phân tích tài chính và fundamentals.",
		Instruction: `Bạn là Financial Analyst chuyên về fundamental analysis.

**NHIỆM VỤ:**
Phân tích tài chính dựa trên:
- Revenue và profit trends
- Valuation metrics (P/E, P/B, etc.)
- Cash flow analysis
- Competitive position

**OUTPUT FORMAT:**
## Financial Analysis

**Valuation:** Overvalued/Undervalued/Fair Value
**Key Metrics:**
- P/E: [value]
- P/B: [value]
- Revenue Growth: [%]

**Financial Health:**
- Debt level: [assessment]
- Cash position: [assessment]

**Growth Prospects:**
[Analysis]

**Risk Factors:**
[List key risks]`,
		OutputKey: stateFinancial,
	})
}

func createAggregatorAgent(m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "AggregatorAgent",
		Model:       m,
		Description: "Tổng hợp phân tích từ các chuyên gia.",
		Instruction: fmt.Sprintf(`Bạn là Chief Investment Officer, tổng hợp insights từ team analysts.

**PHÂN TÍCH TỪ CÁC CHUYÊN GIA:**

### Technical Analysis:
{%s}

### Sentiment Analysis:
{%s}

### Financial Analysis:
{%s}

**NHIỆM VỤ:**
1. Review và synthesize tất cả phân tích
2. Identify consensus và divergence giữa các analysts
3. Weight các yếu tố theo tầm quan trọng
4. Đưa ra recommendation cuối cùng

**OUTPUT FORMAT:**

# Báo Cáo Phân Tích Thị Trường Tổng Hợp

## Executive Summary
[Tóm tắt ngắn gọn 2-3 câu về overall view]

## Tổng Hợp Phân Tích

### Technical View
[Summary từ Technical Analyst]

### Sentiment View
[Summary từ Sentiment Analyst]

### Fundamental View
[Summary từ Financial Analyst]

## Consensus & Divergence
[Điểm đồng thuận và khác biệt giữa các analysts]

## Final Recommendation
- **Action:** BUY / SELL / HOLD
- **Confidence Level:** High / Medium / Low
- **Risk Level:** High / Medium / Low
- **Time Horizon:** Short / Medium / Long term

## Key Risks & Opportunities
**Risks:**
- [List]

**Opportunities:**
- [List]

## Disclaimer
*Đây là phân tích mang tính tham khảo, không phải lời khuyên đầu tư.*`, stateTechAnalysis, stateSentiment, stateFinancial),
		OutputKey: stateMarketReport,
	})
}

func createMarketAnalysisTeam(m model.LLM) (agent.Agent, error) {
	techAnalyst, err := createTechnicalAnalyst(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create tech analyst: %w", err)
	}

	sentimentAnalyst, err := createSentimentAnalyst(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create sentiment analyst: %w", err)
	}

	financialAnalyst, err := createFinancialAnalyst(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create financial analyst: %w", err)
	}

	aggregator, err := createAggregatorAgent(m)
	if err != nil {
		return nil, fmt.Errorf("failed to create aggregator: %w", err)
	}

	// Phase 1: Parallel analysis - 3 analysts chạy đồng thời
	parallelAnalysis, err := parallelagent.New(parallelagent.Config{
		AgentConfig: agent.Config{
			Name:        "ParallelAnalysisPhase",
			Description: "3 analysts phân tích song song: Technical, Sentiment, Financial",
			SubAgents:   []agent.Agent{techAnalyst, sentimentAnalyst, financialAnalyst},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create parallel agent: %w", err)
	}

	// Phase 2: Sequential - Parallel Analysis → Aggregator
	return sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "MarketAnalysisTeam",
			Description: "Team phân tích thị trường: [Tech, Sentiment, Financial] song song → Aggregator",
			SubAgents:   []agent.Agent{parallelAnalysis, aggregator},
		},
	})
}

// ============================================================================
// MAIN
// ============================================================================

func printBanner() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("  Multi-Agent Collaboration Demo - ADK-Go")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("")
	fmt.Println("  Đang chạy: Blog Creation Team (Sequential)")
	fmt.Println("  ┌─────────────────────────────────────────────────────┐")
	fmt.Println("  │  Researcher  →  Writer  →  Editor                   │")
	fmt.Println("  │  (research)     (draft)     (final)                 │")
	fmt.Println("  └─────────────────────────────────────────────────────┘")
	fmt.Println("")
	fmt.Println("  Thay đổi sang Market Analysis Team bằng cách:")
	fmt.Println("  Uncomment createMarketAnalysisTeam() trong main()")
	fmt.Println("")
	fmt.Println("  Prompts mẫu:")
	fmt.Println("  • 'Viết bài về xu hướng AI Agents năm 2024'")
	fmt.Println("  • 'Viết bài về tương lai của Large Language Models'")
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

	// ========================================
	// CHỌN TEAM ĐỂ CHẠY
	// ========================================

	// Option 1: Blog Creation Team (Sequential)
	team, err := createBlogCreationTeam(geminiModel)

	// Option 2: Market Analysis Team (Parallel + Sequential)
	// Uncomment dòng dưới và comment dòng trên để chạy Market Analysis
	// team, err := createMarketAnalysisTeam(geminiModel)

	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(team),
	}

	lch := full.NewLauncher()
	printBanner()

	err = lch.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
