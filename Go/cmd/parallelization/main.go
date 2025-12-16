package main

import (
	"context"
	"time"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
)

type AnalysisResult struct {
	AnalystType string        `json:"analyst_type"`
	Analysis    string        `json:"analyst"`
	Confidence  float64       `json:"confidence"`
	Duration    time.Duration `json:"duration"`
}

type CompanyResearch struct {
	CompanyName string `json:"company_name"`
	Industry    string `json:"industry"`
	Description string `json:"description"`
}

type SynthesizeReport struct {
	Company          string           `json:"company"`
	ExecutiveSummary string           `json:"executive_summary"`
	Analyses         []AnalysisResult `json:"analyses"`
	TotalDuration    time.Duration    `json:"total_duration"`
}

func createSearchAggent(ctx context.Context, m model.LLM) (agent.Agent, error) {

}

func createFinancialAnalyst(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "financial_analyst",
		Model:       m,
		Description: "Chuyên gia phân tích tài chính công ty. Sử dụng agent này để đánh giá doanh thu, lợi nhuận, dòng tiền, và sức khỏe tài chính.",
		Instruction: `Bạn là chuyên gia phân tích tài chính với 20 năm kinh nghiệm.

NHIỆM VỤ:
Khi nhận thông tin về công ty, hãy phân tích:
- Tình hình doanh thu và tăng trưởng
- Biên lợi nhuận và hiệu quả hoạt động
- Dòng tiền và thanh khoản
- Cấu trúc nợ và đòn bẩy tài chính
- So sánh với ngành

ĐỊNH DẠNG TRẢ VỀ:
📊 PHÂN TÍCH TÀI CHÍNH
- Điểm mạnh: [liệt kê]
- Điểm yếu: [liệt kê]
- Đánh giá tổng quan: [1-10 điểm]
- Khuyến nghị: [mua/giữ/bán]`,
	})
}

func createMarketAnalyst(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "market_analyst",
		Model:       m,
		Description: "Chuyên gia phân tích thị trường và cạnh tranh. Sử dụng agent này để đánh giá vị thế thị trường, đối thủ, và xu hướng ngành.",
		Instruction: `Bạn là chuyên gia phân tích thị trường với kiến thức sâu rộng về các ngành công nghiệp.

NHIỆM VỤ:
Khi nhận thông tin về công ty, hãy phân tích:
- Thị phần và vị thế cạnh tranh
- Đối thủ chính và ưu thế cạnh tranh
- Xu hướng thị trường và tiềm năng tăng trưởng
- Rào cản gia nhập ngành
- Cơ hội mở rộng

ĐỊNH DẠNG TRẢ VỀ:
🏆 PHÂN TÍCH THỊ TRƯỜNG
- Vị thế hiện tại: [mô tả]
- Đối thủ chính: [danh sách]
- Xu hướng ngành: [tăng/ổn định/giảm]
- Tiềm năng: [cao/trung bình/thấp]`,
	})
}
func createRiskAnalyst(ctx context.Context, m model.LLM) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "risk_analyst",
		Model:       m,
		Description: "Chuyên gia đánh giá rủi ro doanh nghiệp. Sử dụng agent này để xác định và đánh giá các rủi ro tiềm ẩn.",
		Instruction: `Bạn là chuyên gia quản lý rủi ro với kinh nghiệm đánh giá doanh nghiệp.

NHIỆM VỤ:
Khi nhận thông tin về công ty, hãy phân tích:
- Rủi ro hoạt động (operational risk)
- Rủi ro pháp lý và tuân thủ
- Rủi ro thị trường và kinh tế vĩ mô
- Rủi ro danh tiếng
- Rủi ro chuỗi cung ứng

ĐỊNH DẠNG TRẢ VỀ:
⚠️ ĐÁNH GIÁ RỦI RO
- Rủi ro cao: [liệt kê]
- Rủi ro trung bình: [liệt kê]
- Rủi ro thấp: [liệt kê]
- Tổng điểm rủi ro: [1-10]
- Biện pháp giảm thiểu: [khuyến nghị]`,
	})
}
func createResearchOrchestrator(ctx context.Context, m model.LLM, analysts []agent.Agent) (agent.Agent, error) {
	// Wrap các analyst agents thành tools
	var analystTools []tool.Tool
	for _, analyst := range analysts {
		analystTools = append(analystTools, agenttool.New(analyst, nil))
	}

	return llmagent.New(llmagent.Config{
		Name:        "research_orchestrator",
		Model:       m,
		Description: "Điều phối viên nghiên cứu công ty - Thực hiện phân tích song song",
		Instruction: `Bạn là trưởng nhóm nghiên cứu đầu tư. Khi người dùng yêu cầu phân tích công ty,
bạn phải thực hiện QUY TRÌNH sau:

**QUY TRÌNH PHÂN TÍCH SONG SONG:**

BƯỚC 1: PHÂN TÍCH ĐỒNG THỜI
Gọi TẤT CẢ 4 analyst cùng lúc (chúng sẽ chạy song song):
- financial_analyst: Phân tích tài chính
- market_analyst: Phân tích thị trường
- tech_analyst: Phân tích công nghệ
- risk_analyst: Đánh giá rủi ro

BƯỚC 2: TỔNG HỢP
Sau khi nhận đủ kết quả từ 4 analyst, tổng hợp thành báo cáo:

📋 BÁO CÁO NGHIÊN CỨU CÔNG TY: [Tên công ty]
═══════════════════════════════════════════

1. TÓM TẮT ĐIỀU HÀNH
[Tóm tắt ngắn gọn các điểm chính]

2. PHÂN TÍCH CHI TIẾT
[Tổng hợp từ 4 analyst]

3. KHUYẾN NGHỊ ĐẦU TƯ
[Mua/Giữ/Bán + Lý do]

4. RỦI RO CẦN LƯU Ý
[Danh sách rủi ro chính]

**NẾU NGƯỜI DÙNG CHÀO HỎI:**
Giới thiệu bản thân và hướng dẫn họ cung cấp tên công ty để phân tích.

**VÍ DỤ YÊU CẦU:**
"Phân tích công ty VinGroup" hoặc "Nghiên cứu Apple Inc"`,
		Tools: analystTools,
	})
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
}
