# Chương 18: Feedback and Human-in-the-Loop (Phản hồi và Con người trong Vòng lặp)

## Tổng quan về Mẫu thiết kế Feedback and Human-in-the-Loop

Mặc dù mục tiêu cuối cùng của các tác nhân AI là hoạt động tự chủ, nhưng trong thực tế, các hệ thống mạnh mẽ và đáng tin cậy nhất thường liên quan đến sự cộng tác với con người. Mẫu **Feedback and Human-in-the-Loop (Con người trong Vòng lặp - HitL)** tập trung vào việc thiết kế các cơ chế cho phép con người cung cấp phản hồi cho tác nhân, giám sát hoạt động của nó và can thiệp khi cần thiết. Đây là một khuôn khổ quan trọng để xây dựng niềm tin, cải thiện hiệu suất của tác nhân theo thời gian và đảm bảo các hệ thống AI hoạt động phù hợp với các giá trị và mục tiêu của con người.

HitL không phải là một sự thừa nhận thất bại của AI; mà là một chiến lược thiết kế có chủ ý để khai thác điểm mạnh của cả AI và con người. Các tác nhân vượt trội trong việc xử lý lượng lớn dữ liệu, phát hiện các mẫu và thực hiện các tác vụ lặp đi lặp lại. Con người xuất sắc trong việc hiểu ngữ cảnh phức tạp, đưa ra các quyết định đạo đức, xử lý sự không chắc chắn và thích nghi với các tình huống mới lạ. Khi kết hợp hiệu quả, kết quả là một hệ thống siêu việt khả năng của cả hai thành phần riêng lẻ.

Các thành phần cốt lõi của mẫu Feedback and HitL bao gồm:

1.  **Thu thập Phản hồi (Feedback Collection):**
    *   **Phản hồi rõ ràng (Explicit Feedback):** Người dùng chủ động cung cấp phản hồi thông qua các giao diện UI (ví dụ: nút "Thích/Không thích", biểu mẫu khảo sát, xếp hạng).
    *   **Phản hồi ngầm (Implicit Feedback):** Tác nhân suy ra phản hồi từ hành vi của người dùng (ví dụ: thời gian dành cho một câu trả lời, sửa đổi đầu ra của tác nhân, việc hoàn thành tác vụ).
    *   **Phản hồi của chuyên gia (Expert Feedback):** Các chuyên gia về miền đánh giá các quyết định hoặc đầu ra của tác nhân.
2.  **Giám sát và Cảnh báo (Monitoring and Alerting):**
    *   **Ngưỡng Tin cậy (Confidence Thresholds):** Tác nhân báo hiệu cho con người khi nó không chắc chắn về một quyết định hoặc khi độ tin cậy của nó giảm xuống dưới một ngưỡng nhất định.
    *   **Phát hiện Bất thường (Anomaly Detection):** Hệ thống phát hiện hành vi không mong muốn hoặc các điều kiện vượt quá các tham số hoạt động an toàn.
    *   **Cảnh báo dựa trên Rủi ro (Risk-Based Alerting):** Can thiệp của con người được ưu tiên cho các tình huống có rủi ro cao hoặc tác động cao.
3.  **Điểm can thiệp của Con người (Human Intervention Points):**
    *   **Ủy quyền (Delegation):** Tác nhân có thể ủy quyền một phần hoặc toàn bộ tác vụ cho con người khi nó không thể tự xử lý.
    *   **Sửa lỗi (Correction):** Con người có thể sửa trực tiếp đầu ra hoặc hành động của tác nhân.
    *   **Override/Điều khiển (Override/Control):** Con người có thể tiếp quản quyền kiểm soát tác nhân hoặc dừng các hoạt động của nó.
    *   **Giải thích (Explanation):** Tác nhân cung cấp các giải thích về lý do đằng sau các quyết định hoặc hành động của nó để tạo điều kiện cho sự can thiệp của con người.
4.  **Học hỏi từ Phản hồi (Learning from Feedback):**
    *   **Tinh chỉnh mô hình (Model Retraining/Fine-tuning):** Phản hồi được sử dụng để tinh chỉnh các LLM hoặc các mô hình phụ của tác nhân.
    *   **Cập nhật quy tắc (Rule Updates):** Phản hồi có thể dẫn đến các quy tắc hoặc logic được xác định rõ ràng hơn cho tác nhân.
    *   **Cải thiện prompt (Prompt Improvement):** Phản hồi giúp tinh chỉnh các prompt được sử dụng để điều khiển LLM của tác nhân.

Tích hợp con người vào vòng lặp đảm bảo các hệ thống AI vẫn có trách nhiệm giải trình, công bằng và hiệu quả, phát triển và cải thiện liên tục theo thời gian.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu Feedback and Human-in-the-Loop là rất quan trọng trong nhiều ứng dụng:

1.  **Kiểm duyệt Nội dung:**
    *   **HitL:** Các tác nhân AI gắn cờ nội dung có khả năng vi phạm chính sách (ví dụ: ngôn từ kích động thù địch, nội dung người lớn). Người kiểm duyệt con người sau đó xem xét các trường hợp được gắn cờ, đưa ra quyết định cuối cùng và cung cấp phản hồi để huấn luyện thêm cho AI.
    *   **Lợi ích:** Cải thiện độ chính xác, giảm căng thẳng cho con người và đảm bảo tính nhất quán của chính sách.
2.  **Y tế và Chẩn đoán:**
    *   **HitL:** Các tác nhân AI phân tích hình ảnh y tế (X-quang, MRI) và đề xuất các chẩn đoán hoặc các khu vực đáng lo ngại. Các bác sĩ con người xem xét các phát hiện của AI, xác nhận chẩn đoán và cung cấp dữ liệu được dán nhãn cho việc huấn luyện thêm.
    *   **Lợi ích:** Tăng tốc độ chẩn đoán, giảm lỗi của con người và cải thiện kết quả của bệnh nhân.
3.  **Dịch vụ Khách hàng và Hỗ trợ:**
    *   **HitL:** Chatbot AI xử lý các truy vấn thông thường. Đối với các câu hỏi phức tạp hoặc nhạy cảm, chatbot sẽ leo thang cho một đại lý con người, đồng thời cung cấp ngữ cảnh của cuộc trò chuyện. Đại lý có thể điều chỉnh phản ứng của chatbot hoặc sử dụng nó như một trợ lý.
    *   **Lợi ích:** Giảm tải cho đại lý con người, cải thiện thời gian phản hồi và duy trì chất lượng dịch vụ cao.
4.  **Xe Tự lái:**
    *   **HitL:** Mặc dù AI điều khiển xe, hệ thống vẫn bao gồm sự giám sát của con người. Người lái xe con người có thể tiếp quản quyền kiểm soát khi cần thiết (ví dụ: trong các tình huống nguy hiểm, điều kiện thời tiết khắc nghiệt). AI học hỏi từ các can thiệp này.
    *   **Lợi ích:** Đảm bảo an toàn, cho phép AI học hỏi từ các tình huống mà nó không được đào tạo và xây dựng niềm tin của người dùng.
5.  **Tạo Nội dung Sáng tạo:**
    *   **HitL:** Các tác nhân AI tạo ra các bản nháp ban đầu của bài viết, tác phẩm nghệ thuật hoặc thiết kế. Người sáng tạo con người sau đó tinh chỉnh, biên tập và chọn các tác phẩm tốt nhất, cung cấp phản hồi cho AI về những gì hiệu quả và những gì không.
    *   **Lợi ích:** Tăng cường sự sáng tạo của con người, tăng năng suất và khám phá các khả năng sáng tạo mới.
6.  **Tác nhân Lập trình:**
    *   **HitL:** Các tác nhân AI tạo các đoạn mã hoặc đề xuất sửa lỗi. Các nhà phát triển con người xem xét, chỉnh sửa và chấp nhận hoặc từ chối các đề xuất, cung cấp phản hồi có giá trị để cải thiện khả năng mã hóa của AI.
    *   **Lợi ích:** Tăng tốc độ phát triển, giảm lỗi của con người và tạo điều kiện thuận lợi cho việc học hỏi liên tục cho tác nhân.

Trong mỗi trường hợp, sự cộng tác giữa AI và con người dẫn đến một hệ thống mạnh mẽ hơn, hiệu quả hơn và đáng tin cậy hơn so với việc một trong hai hoạt động riêng lẻ.

## Các kỹ thuật triển khai

Triển khai mẫu Feedback and Human-in-the-Loop đòi hỏi phải suy nghĩ kỹ lưỡng về các điểm tích hợp của con người và các cơ chế phản hồi:

1.  **Giao diện Người dùng trực quan (Intuitive User Interfaces - UI):** Thiết kế các UI rõ ràng, dễ hiểu cho phép người dùng hoặc chuyên gia cung cấp phản hồi một cách dễ dàng. Điều này có thể bao gồm các nút "Thích/Không thích", trường văn bản tự do, hoặc các công cụ chú thích.
2.  **Điểm can thiệp được xác định (Defined Intervention Points):** Xác định rõ ràng các kịch bản mà tác nhân nên tự động leo thang cho con người (ví dụ: độ tin cậy thấp, rủi ro cao, các trường hợp ngoại lệ).
3.  **Ngữ cảnh hóa các Yêu cầu can thiệp (Contextualizing Intervention Requests):** Khi một tác nhân yêu cầu sự can thiệp của con người, nó phải cung cấp tất cả ngữ cảnh liên quan (ví dụ: lịch sử hội thoại, thông tin đầu vào, lý do tại sao nó cần trợ giúp) để con người đưa ra quyết định sáng suốt.
4.  **Tùy chỉnh vai trò và quyền hạn (Customizable Roles and Permissions):** Xác định các vai trò khác nhau cho con người trong vòng lặp (ví dụ: người đánh giá, người hiệu đính, người ra quyết định cuối cùng) với các quyền và trách nhiệm tương ứng.
5.  **Cơ chế Học hỏi từ Phản hồi (Learning from Feedback Mechanisms):** Thiết lập một quy trình để tích hợp phản hồi của con người trở lại hệ thống (ví dụ: cập nhật cơ sở tri thức của tác nhân, tinh chỉnh mô hình, sửa đổi prompt).
6.  **Tính minh bạch và Khả năng Giải thích (Transparency and Explainability - XAI):** Cung cấp các giải thích về lý do tác nhân đưa ra một quyết định hoặc đề xuất cụ thể. Điều này giúp con người hiểu được hành vi của tác nhân và đưa ra phản hồi có thông tin hơn.
7.  **Đánh giá của con người định kỳ (Periodic Human Review):** Ngay cả đối với các tác vụ không yêu cầu can thiệp trực tiếp, việc xem xét định kỳ các tương tác của tác nhân bởi con người có thể phát hiện các vấn đề tiềm ẩn và cung cấp phản hồi để cải thiện.
8.  **Thử nghiệm A/B và Đánh giá Phát triển (A/B Testing and Canary Releases):** Sử dụng các kỹ thuật này để thử nghiệm các phiên bản tác nhân mới với một tập hợp con người dùng nhỏ trước khi triển khai rộng rãi, cho phép con người kiểm tra hiệu suất của tác nhân trong các điều kiện thực tế.

## Ví dụ Code Thực hành (Google ADK)

Google Agent Developer Kit (ADK) có thể được sử dụng để xây dựng các cơ chế Feedback và Human-in-the-Loop. Ví dụ này mô tả một tác nhân ADK được thiết kế để hỗ trợ đặt lịch hẹn. Nếu tác nhân không thể xác nhận hoặc xử lý yêu cầu đặt lịch hẹn, nó sẽ leo thang vấn đề cho một tác nhân con người, cung cấp ngữ cảnh cần thiết để con người tiếp quản.

```python
import uuid
import asyncio
from google.adk.agents import LlmAgent, FunctionAgent
from google.adk.runners import InMemoryRunner
from google.adk.tools import FunctionTool
from google.genai import types
import nest_asyncio

# Áp dụng nest_asyncio để cho phép chạy asyncio trong môi trường này (nếu cần, ví dụ: Jupyter)
nest_asyncio.apply()

# --- 1. Định nghĩa Công cụ ---
@FunctionTool
def book_appointment(date: str, time: str, service: str) -> str:
    """
    Đặt lịch hẹn cho một dịch vụ cụ thể vào ngày và giờ được chỉ định.
    Trả về ID xác nhận lịch hẹn nếu thành công, hoặc thông báo lỗi nếu thất bại.
    """
    # Mô phỏng một API đặt lịch hẹn.
    if "mai" in date.lower() and "massage" in service.lower():
        return "Lỗi: Dịch vụ massage không có sẵn vào ngày mai."
    if "10h sáng" in time.lower() and "thứ 7" in date.lower():
        return "Lỗi: Không có lịch hẹn nào vào thứ 7 lúc 10h sáng."
    
    confirmation_id = f"APPT-{uuid.uuid4().hex[:8]}"
    return f"Đã đặt lịch hẹn {service} vào {date} lúc {time}. ID xác nhận: {confirmation_id}"

# --- 2. Định nghĩa Tác nhân Con người (Human Agent) ---
# Tác nhân này là một placeholder cho một tác nhân con người thực sự,
# nhưng nó minh họa cách tác nhân AI có thể ủy quyền cho một vai trò con người.
# Trong một triển khai thực tế, điều này có thể gửi một thông báo đến một người thật.
human_escalation_agent = FunctionAgent(
    name="HumanAgent",
    description="Một tác nhân con người xử lý các yêu cầu phức tạp hoặc cần sự can thiệp của con người.",
    execute_function=lambda request: f"Yêu cầu được ủy quyền cho con người: {request}. Con người sẽ xử lý điều này."
)

# --- 3. Định nghĩa Tác nhân Đặt lịch hẹn (AI Agent) ---
appointment_agent = LlmAgent(
    name='appointment_booking_agent',
    model='gemini-2.0-flash',
    instruction=(
        "Bạn là một trợ lý đặt lịch hẹn. Sử dụng công cụ `book_appointment` để giúp người dùng đặt lịch hẹn. "
        "Nếu bạn không thể hoàn thành việc đặt lịch hẹn hoặc gặp lỗi, hãy leo thang cho `HumanAgent` "
        "và cung cấp tất cả ngữ cảnh liên quan để họ có thể tiếp quản. "
        "Luôn xác nhận các chi tiết với người dùng trước khi đặt lịch."
    ),
    description="Một tác nhân đặt lịch hẹn có thể leo thang cho con người.",
    tools=[book_appointment],
    # Liên kết tác nhân con người để leo thang
    sub_agents=[human_escalation_agent]
)

# --- Logic Thực thi ---
def run_appointment_agent(user_message: str):
    print(f"\n--- Chạy Tác nhân Đặt lịch hẹn với Tin nhắn: '{user_message}' ---")
    runner = InMemoryRunner(appointment_agent)
    user_id = "booking_user_1"
    session_id = str(uuid.uuid4())
    
    response_text = ""
    async for event in runner.run(
        user_id=user_id,
        session_id=session_id,
        new_message=types.Content(
            role='user',
            parts=[types.Part(text=user_message)]
        )
    ):
        if event.is_final_response() and event.content:
             if hasattr(event.content, 'text') and event.content.text:
                response_text = event.content.text
             elif event.content.parts:
                 response_text = "".join([part.text for part in event.content.parts if part.text])
    
    print(f"Phản hồi cuối cùng của Tác nhân: {response_text}")

# --- Chạy Ví dụ ---
def main():
    # Trường hợp đặt lịch hẹn thành công
    run_appointment_agent("Tôi muốn đặt lịch cắt tóc vào thứ Tư tuần sau lúc 3 giờ chiều.")

    # Trường hợp thất bại và cần leo thang
    run_appointment_agent("Đặt lịch massage cho tôi vào ngày mai.")
    
    # Một trường hợp thất bại khác
    run_appointment_agent("Tôi muốn đặt lịch khám bác sĩ vào thứ 7 lúc 10h sáng.")


if __name__ == "__main__":
    asyncio.run(main())
```

Mã này sử dụng Google ADK để tạo một tác nhân đặt lịch hẹn AI với khả năng leo thang cho một tác nhân con người. Công cụ `book_appointment` mô phỏng một API đặt lịch, với một số logic lỗi tích hợp. `human_escalation_agent` là một `FunctionAgent` đóng vai trò là điểm tiếp quản của con người. `appointment_agent` là một `LlmAgent` chính được hướng dẫn sử dụng công cụ `book_appointment` và leo thang cho `HumanAgent` khi gặp lỗi hoặc không thể hoàn thành tác vụ. Hàm `run_appointment_agent` xử lý việc thực thi tác nhân và thu thập phản hồi. Hàm `main` minh họa ba kịch bản: một đặt lịch thành công và hai kịch bản mà tác nhân AI không thể tự xử lý và do đó ủy quyền cho tác nhân con người, chứng minh một cơ chế Human-in-the-Loop cơ bản.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các tác nhân AI thường hoạt động trong các môi trường phức tạp và không thể đoán trước, nơi chúng có thể gặp phải các tình huống không xác định, đưa ra các quyết định phụ tối ưu hoặc thực hiện các hành động không mong muốn. Nếu không có các cơ chế để con người cung cấp phản hồi, giám sát hoạt động của tác nhân và can thiệp khi cần thiết, các hệ thống AI có thể trở nên không đáng tin cậy, không hiệu quả hoặc thậm chí nguy hiểm, đặc biệt là trong các trường hợp sử dụng quan trọng. Điều này làm xói mòn niềm tin và hạn chế tiện ích tổng thể của tác nhân. Điều này làm xói mòn niềm tin và hạn chế tiện ích tổng thể của tác nhân.
*   **Tại sao:** Mẫu Feedback and Human-in-the-Loop (HitL) cung cấp một giải pháp tiêu chuẩn hóa bằng cách tích hợp con người vào thiết kế và hoạt động của các hệ thống AI. Nó khai thác điểm mạnh của cả AI (xử lý dữ liệu quy mô lớn, nhận dạng mẫu) và con người (ngữ cảnh phức tạp, đạo đức, thích ứng với tình huống mới lạ). Các thành phần cốt lõi bao gồm thu thập phản hồi (rõ ràng, ngầm, chuyên gia), giám sát và cảnh báo (ngưỡng tin cậy, phát hiện bất thường), điểm can thiệp của con người (ủy quyền, sửa lỗi, ghi đè), và học hỏi từ phản hồi (tinh chỉnh mô hình, cập nhật quy tắc).
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này cho TẤT CẢ các tác nhân AI được triển khai trong môi trường sản xuất, đặc biệt là những tác nhân tương tác với con người, đưa ra các quyết định quan trọng, xử lý dữ liệu nhạy cảm hoặc kiểm soát các hệ thống vật lý. HitL là điều kiện tiên quyết để xây dựng các hệ thống AI có trách nhiệm, đáng tin cậy và hiệu quả về mặt xã hội. Nó không phải là sự thừa nhận thất bại của AI mà là một chiến lược thiết kế có chủ ý để xây dựng niềm tin và liên tục cải thiện.

## Những Điểm Chính (Key Takeaways)

*   Feedback and Human-in-the-Loop (HitL) là một mẫu thiết kế quan trọng để xây dựng các tác nhân AI đáng tin cậy và hiệu quả.
*   Nó kết hợp điểm mạnh của AI (hiệu suất quy mô) và con người (ngữ cảnh, đạo đức).
*   Các thành phần chính bao gồm thu thập phản hồi (rõ ràng/ngầm/chuyên gia), giám sát/cảnh báo và điểm can thiệp của con người.
*   Các điểm can thiệp của con người bao gồm ủy quyền tác vụ, sửa lỗi và ghi đè hoạt động của tác nhân.
*   Phản hồi được sử dụng để tinh chỉnh mô hình, cập nhật quy tắc và cải thiện prompt.
*   Các ứng dụng thực tế bao gồm kiểm duyệt nội dung, y tế, dịch vụ khách hàng, xe tự lái và tạo nội dung.
*   ADK tạo điều kiện cho HitL thông qua các tác nhân chức năng có thể đóng vai trò là điểm leo thang của con người.

## Kết luận

Chương này đã làm sáng tỏ tầm quan trọng của mẫu Feedback and Human-in-the-Loop (Con người trong Vòng lặp), một nguyên tắc thiết kế quan trọng cho các tác nhân AI trong thế giới thực. Chúng ta đã khám phá cách tích hợp con người vào vòng lặp không phải là một sự thừa nhận hạn chế của AI mà là một chiến lược có chủ ý để khai thác điểm mạnh của cả AI và con người, dẫn đến các hệ thống hiệu quả, đáng tin cậy và đáng tin cậy hơn. Các cơ chế thu thập phản hồi, giám sát hoạt động và can thiệp của con người, cùng với việc học hỏi liên tục từ phản hồi đó, là nền tảng để xây dựng các tác nhân trưởng thành. Ví dụ ADK về một tác nhân đặt lịch hẹn có khả năng leo thang cho con người đã minh họa một cách tiếp cận thực tế. Bằng cách ưu tiên sự cộng tác của con người, chúng ta đảm bảo rằng các tác nhân AI hoạt động hài hòa với các giá trị của con người, cung cấp các giải pháp có trách nhiệm và hữu ích trong các lĩnh vực khác nhau. Chương tiếp theo sẽ giới thiệu tầm quan trọng của việc xây dựng các tác nhân có khả năng mở rộng để đáp ứng nhu cầu ngày càng tăng.

## Tài liệu tham khảo
1.  Google AI Principles (Human Oversight): https://ai.google/responsibility/principles/
2.  Human-in-the-Loop Machine Learning by Robert Monarch: https://www.manning.com/books/human-in-the-loop-machine-learning
3.  Google ADK Documentation (Human-in-the-Loop): https://google.github.io/adk-docs/agents/#human-in-the-loop
4.  Responsible AI Toolkit: https://www.microsoft.com/en-us/ai/responsible-ai-resources
