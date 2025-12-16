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
