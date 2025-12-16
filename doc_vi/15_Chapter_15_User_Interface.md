# Chương 15: User Interface (UI) (Giao Diện Người Dùng)

## Tổng quan về Mẫu thiết kế User Interface

Các tác nhân AI, dù thông minh hay tinh vi đến đâu, cuối cùng vẫn phải tương tác với con người. Mẫu **User Interface (UI) (Giao Diện Người Dùng)** tập trung vào việc thiết kế và xây dựng các điểm tiếp xúc nơi con người và tác nhân giao tiếp, đảm bảo các tương tác này hiệu quả, trực quan và tự nhiên nhất có thể. Một UI được thiết kế tốt là yếu tố then chốt cho việc áp dụng, sự hài lòng của người dùng và khả năng của tác nhân trong việc hoàn thành các mục tiêu trong thế giới thực một cách thành công.

Giao diện người dùng cho các tác nhân AI không chỉ là các nút và biểu mẫu; chúng là các phương tiện để:

1.  **Chuyển đổi Ý định của Con người:** Cho phép người dùng thể hiện ý định, mục tiêu và ràng buộc của họ một cách rõ ràng cho tác nhân.
2.  **Thông báo Hành động của Tác nhân:** Cung cấp thông tin chi tiết về những gì tác nhân đang làm, lý do tại sao nó đang làm điều đó và kết quả hành động của nó.
3.  **Tạo điều kiện can thiệp:** Cung cấp cho người dùng các cơ chế để sửa lỗi, điều chỉnh hướng của tác nhân hoặc dừng các hoạt động khi cần.
4.  **Xây dựng niềm tin:** Một UI minh bạch và dễ hiểu giúp xây dựng niềm tin vào khả năng và độ tin cậy của tác nhân.

Các loại UI cho tác nhân có thể thay đổi đáng kể tùy thuộc vào trường hợp sử dụng:

*   **UI Hội thoại (Conversational UIs):** Giao diện dựa trên văn bản hoặc giọng nói (chatbot, trợ lý giọng nói) nơi người dùng tương tác thông qua ngôn ngữ tự nhiên. Chúng thường là đơn giản nhất để triển khai nhưng đòi hỏi khả năng xử lý ngôn ngữ tự nhiên mạnh mẽ và quản lý ngữ cảnh.
*   **UI Dựa trên GUI (Graphical User Interface - GUIs):** Giao diện trực quan hơn với các nút, trường nhập liệu, thanh trượt và các thành phần trực quan khác. Chúng có thể cung cấp quyền kiểm soát chặt chẽ hơn và phản hồi trực quan nhưng có thể phức tạp hơn để thiết kế và xây dựng.
*   **UI Kết hợp (Hybrid UIs):** Kết hợp các yếu tố hội thoại và GUI, cho phép người dùng chuyển đổi giữa các phương thức tương tác khi thích hợp. Ví dụ: một chatbot có thể hiển thị các nút để chọn các tùy chọn hoặc các trường để nhập dữ liệu có cấu trúc.
*   **UI Đa phương thức (Multimodal UIs):** Tận dụng các giao diện có thể xử lý và hiển thị thông tin bằng nhiều phương thức, bao gồm văn bản, hình ảnh, âm thanh và video. Các giao diện này cho phép các tương tác phong phú hơn và tự nhiên hơn, phản ánh cách con người tương tác với thế giới.
*   **UI Tích hợp (Integrated UIs):** Nơi tác nhân không có một giao diện độc lập mà thay vào đó tích hợp khả năng của nó vào các ứng dụng hiện có (ví dụ: một plugin trợ lý AI trong một trình soạn thảo văn bản).

Thiết kế UI cho các tác nhân AI phải tính đến các nguyên tắc về khả năng sử dụng, khả năng truy cập và đặc biệt là kỳ vọng của người dùng về hành vi AI. Mục tiêu không chỉ là trình bày thông tin mà còn là tạo ra một trải nghiệm cộng tác, nơi người dùng cảm thấy có quyền kiểm soát và hiểu được tác nhân.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu User Interface là trung tâm của mọi ứng dụng tác nhân AI có người dùng cuối tương tác:

1.  **Trợ lý Ảo và Chatbot:**
    *   **UI:** Giao diện dựa trên văn bản hoặc giọng nói (ví dụ: Google Assistant, ChatGPT) cho phép người dùng đặt câu hỏi, đưa ra lệnh và nhận phản hồi.
    *   **Mục tiêu:** Cung cấp khả năng giao tiếp tự nhiên và hiệu quả, giảm ma sát cho các tác vụ đơn giản và cung cấp hỗ trợ.
2.  **Công cụ Phát triển AI và IDE:**
    *   **UI:** Các trình soạn thảo mã tích hợp với các tác nhân AI để tạo mã, gỡ lỗi và hoàn thành mã (ví dụ: GitHub Copilot trong VS Code, các plugin của Gemini trong IDE).
    *   **Mục tiêu:** Tăng năng suất của nhà phát triển, tự động hóa các tác vụ lặp đi lặp lại và cung cấp hỗ trợ mã hóa thông minh.
3.  **Hệ thống Quản lý Khách hàng (CRM):**
    *   **UI:** Bảng điều khiển nơi tác nhân AI tóm tắt các tương tác của khách hàng, đề xuất các bước tiếp theo hoặc tự động hóa việc nhập dữ liệu.
    *   **Mục tiêu:** Cải thiện hiệu quả của đại lý hỗ trợ, cá nhân hóa tương tác với khách hàng và đảm bảo tính nhất quán của dữ liệu.
4.  **Nền tảng Thiết kế và Sáng tạo:**
    *   **UI:** Các công cụ cho phép người dùng mô tả trực quan các thiết kế hoặc điều chỉnh các thuộc tính sáng tạo thông qua các lệnh ngôn ngữ tự nhiên.
    *   **Mục tiêu:** Dân chủ hóa thiết kế, tăng tốc quá trình sáng tạo và cho phép người dùng không phải là nhà thiết kế tạo ra nội dung chất lượng cao.
5.  **Tác nhân Robot và Tự trị:**
    *   **UI:** Bảng điều khiển cho phép người vận hành giám sát trạng thái của robot, điều chỉnh các tham số hoặc đưa ra các lệnh cấp cao. Các giao diện AR/VR cũng có thể được sử dụng để tương tác trực quan với robot trong không gian vật lý.
    *   **Mục tiêu:** Cung cấp quyền kiểm soát an toàn và trực quan đối với các hệ thống vật lý, cho phép sự can thiệp của con người khi cần thiết.
6.  **Hệ thống Phân tích Dữ liệu và Báo cáo:**
    *   **UI:** Giao diện cho phép người dùng đặt câu hỏi về dữ liệu bằng ngôn ngữ tự nhiên, tạo biểu đồ và báo cáo hoặc yêu cầu phân tích nâng cao.
    *   **Mục tiêu:** Dân chủ hóa phân tích dữ liệu, giúp người dùng không có kỹ thuật dễ dàng hơn trong việc trích xuất thông tin chi tiết và đưa ra quyết định dựa trên dữ liệu.

Một UI được thiết kế tốt không chỉ là giao diện; đó là cầu nối quan trọng giữa khả năng của tác nhân và tiện ích của nó đối với con người.

## Các kỹ thuật triển khai

Thiết kế và triển khai giao diện người dùng hiệu quả cho các tác nhân AI yêu cầu một cách tiếp cận có cân nhắc:

1.  **Giao diện hội thoại:**
    *   **Nhắc nhở Rõ ràng (Clear Prompting):** Thiết kế prompt ban đầu hướng dẫn người dùng về khả năng của tác nhân và cách tương tác hiệu quả.
    *   **Quản lý Ngữ cảnh (Context Management):** Sử dụng bộ nhớ để duy trì ngữ cảnh hội thoại, tránh lặp lại và cho phép các cuộc hội thoại tự nhiên hơn.
    *   **Xử lý Nỗ lực và Không rõ ràng (Handling Ambiguity and Nuance):** Xây dựng các cơ chế cho tác nhân để tìm kiếm sự làm rõ khi ý định của người dùng không rõ ràng.
    *   **Phản hồi Hiển thị (Displaying Feedback):** Cho tác nhân hiển thị tiến trình, các hành động đang diễn ra và lỗi một cách rõ ràng.
2.  **Giao diện GUI:**
    *   **Trạng thái Hình ảnh Rõ ràng (Clear Visual State):** Đảm bảo rằng trạng thái hiện tại của tác nhân, bao gồm các tác vụ đang thực hiện, đầu ra tạm thời và bất kỳ lỗi nào, được hiển thị rõ ràng và trực quan cho người dùng.
    *   **Điều khiển Trực quan (Intuitive Controls):** Cung cấp các nút, thanh trượt và trường nhập liệu dễ hiểu và dễ sử dụng để người dùng kiểm soát hành vi của tác nhân.
    *   **Tùy chỉnh (Customization):** Cho phép người dùng tùy chỉnh các cài đặt, sở thích hoặc thậm chí là nhân cách của tác nhân để phù hợp với nhu cầu của họ.
    *   **Hiển thị Đa phương thức (Multimodal Display):** Khả năng hiển thị các phương thức khác nhau như hình ảnh, video, âm thanh cùng với văn bản.
3.  **Tích hợp (Integration):**
    *   **API và SDK:** Cung cấp các API và SDK mạnh mẽ cho phép các nhà phát triển tích hợp khả năng của tác nhân vào các ứng dụng hoặc nền tảng hiện có.
    *   **Plugin và Tiện ích mở rộng:** Cho phép tạo các plugin hoặc tiện ích mở rộng mở rộng chức năng của tác nhân sang các môi trường khác (ví dụ: trình soạn thảo văn bản, trình duyệt web).
4.  **Khả năng Truy cập (Accessibility):**
    *   **Hỗ trợ Công nghệ Hỗ trợ (Assistive Technology Support):** Đảm bảo rằng giao diện có thể truy cập được thông qua trình đọc màn hình, điều hướng bằng bàn phím hoặc các công nghệ hỗ trợ khác.
    *   **Tùy chọn Ngôn ngữ (Language Options):** Cung cấp hỗ trợ đa ngôn ngữ để phục vụ nhiều đối tượng người dùng hơn.

## Ví dụ Code Thực hành (Google ADK)

Google Agent Developer Kit (ADK) tập trung vào việc xây dựng logic tác nhân cốt lõi và tích hợp nó với các UI khác nhau. ADK được thiết kế để tách biệt logic tác nhân khỏi triển khai UI. Điều này có nghĩa là bạn có thể xây dựng tác nhân của mình một lần và sau đó kết nối nó với nhiều loại giao diện người dùng khác nhau, từ chatbot văn bản đơn giản đến các ứng dụng web phức tạp hoặc thậm chí là giao diện thoại.

Ví dụ này cho thấy một tác nhân ADK rất cơ bản được kết nối với một giao diện dòng lệnh (CLI). Mặc dù bản thân giao diện dòng lệnh là tối thiểu, ví dụ này làm nổi bật cách bạn sẽ nhận đầu vào từ người dùng (trong trường hợp này, qua CLI), chuyển nó đến tác nhân của mình và sau đó hiển thị phản hồi của tác nhân trở lại người dùng.

```python
import uuid
import asyncio
from google.adk.agents import LlmAgent
from google.adk.runners import InMemoryRunner
from google.genai import types
import nest_asyncio

# Áp dụng nest_asyncio để cho phép chạy asyncio trong môi trường này (nếu cần, ví dụ: Jupyter)
nest_asyncio.apply()

# --- Định nghĩa Tác nhân ---
simple_cli_agent = LlmAgent(
    name='cli_assistant_agent',
    model='gemini-2.0-flash',
    instruction=(
        "Bạn là một trợ lý thân thiện và hữu ích, được thiết kế để tương tác qua giao diện dòng lệnh. "
        "Hãy giữ các phản hồi của bạn ngắn gọn và đi thẳng vào vấn đề."
    ),
    description="Một tác nhân được thiết kế để tương tác qua giao diện dòng lệnh."
)

# --- Logic Giao diện Dòng lệnh Đơn giản ---
async def run_cli_interaction():
    print("Chào mừng bạn đến với CLI Assistant Agent. Gõ 'exit' để thoát.")
    runner = InMemoryRunner(simple_cli_agent)
    user_id = "cli_user_1"
    session_id = str(uuid.uuid4())
    
    while True:
        user_input = input("\nBạn: ")
        if user_input.lower() == 'exit':
            print("Tạm biệt!")
            break
        
        response_text = ""
        # Truyền đầu vào của người dùng đến tác nhân và nhận phản hồi
        async for event in runner.run(
            user_id=user_id,
            session_id=session_id,
            new_message=types.Content(
                role='user',
                parts=[types.Part(text=user_input)]
            )
        ):
            if event.is_final_response() and event.content:
                 if hasattr(event.content, 'text') and event.content.text:
                    response_text = event.content.text
                 elif event.content.parts:
                     response_text = "".join([part.text for part in event.content.parts if part.text])
        
        print(f"Tác nhân: {response_text}")

# --- Chạy Giao diện Dòng lệnh ---
if __name__ == "__main__":
    asyncio.run(run_cli_interaction())
```

Mã này trình bày một ví dụ tối thiểu về tác nhân ADK tương tác qua giao diện dòng lệnh. Tác nhân `simple_cli_agent` được khởi tạo với một instruction được điều chỉnh cho các tương tác CLI. Hàm `run_cli_interaction` tạo ra một vòng lặp CLI cơ bản, đọc đầu vào của người dùng, chuyển nó đến tác nhân thông qua `InMemoryRunner` và in phản hồi của tác nhân. Vòng lặp tiếp tục cho đến khi người dùng gõ "exit". Điều này minh họa sự tách biệt rõ ràng giữa logic tác nhân cốt lõi và giao diện người dùng, cho phép tác nhân được tích hợp vào các UI khác một cách dễ dàng.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các tác nhân AI, mặc dù rất mạnh mẽ, thường thiếu một cách hiệu quả và trực quan để người dùng tương tác với chúng. Các giao diện dựa trên văn bản thuần túy hoặc không có giao diện nào cả có thể gây khó khăn cho người dùng trong việc hiểu khả năng của tác nhân, theo dõi tiến độ của nó hoặc can thiệp khi cần thiết. Ma sát này hạn chế tiện ích, khả năng sử dụng và mức độ chấp nhận của tác nhân, đặc biệt là trong các ứng dụng hướng đến người tiêu dùng hoặc các trường hợp sử dụng cần sự tham gia của con người.
*   **Tại sao:** Mẫu User Interface (UI) cung cấp một giải pháp tiêu chuẩn hóa bằng cách tập trung vào việc thiết kế và xây dựng các điểm tiếp xúc nơi con người và tác nhân giao tiếp. Một UI được thiết kế tốt là rất quan trọng để chuyển đổi ý định của con người thành các lệnh có thể thực thi được, thông báo hành động của tác nhân và tạo điều kiện can thiệp của người dùng. UI có thể là hội thoại (chatbot), đồ họa (GUI), kết hợp hoặc đa phương thức. Mục tiêu không chỉ là trình bày thông tin mà còn là tạo ra một trải nghiệm cộng tác, xây dựng niềm tin và cung cấp cho người dùng cảm giác kiểm soát. Các framework như Google ADK được thiết kế để tách biệt logic tác nhân cốt lõi khỏi triển khai UI, cho phép các tác nhân được kết nối với nhiều giao diện khác nhau.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này cho TẤT CẢ các tác nhân AI dự định tương tác với người dùng cuối, hoặc bất kỳ ai cần giám sát hoặc điều khiển tác nhân. Ưu tiên các nguyên tắc thiết kế lấy người dùng làm trung tâm, sự rõ ràng, tính minh bạch và khả năng truy cập để đảm bảo UI không chỉ hấp dẫn về mặt hình ảnh mà còn hiệu quả và đáng tin cậy.

## Những Điểm Chính (Key Takeaways)

*   Thiết kế giao diện người dùng là rất quan trọng cho các tác nhân AI tương tác hiệu quả với con người.
*   UI là phương tiện để chuyển đổi ý định của con người, thông báo hành động của tác nhân, tạo điều kiện can thiệp và xây dựng niềm tin.
*   Các loại UI bao gồm hội thoại, đồ họa, kết hợp và đa phương thức.
*   Các kỹ thuật triển khai bao gồm prompt engineering rõ ràng, quản lý ngữ cảnh, trạng thái trực quan rõ ràng, điều khiển trực quan, tùy chỉnh và khả năng truy cập.
*   ADK thúc đẩy sự tách biệt giữa logic tác nhân và UI, cho phép linh hoạt trong việc lựa chọn giao diện.

## Kết luận

Chương này đã làm rõ tầm quan trọng của mẫu User Interface (UI) trong việc phát triển các tác nhân AI. Chúng ta đã khám phá cách một UI được thiết kế tốt không chỉ là một giao diện mà còn là cầu nối quan trọng giữa khả năng của tác nhân và tiện ích của nó đối với người dùng. Bằng cách tập trung vào sự rõ ràng, tính minh bạch và khả năng kiểm soát của người dùng, chúng ta có thể tạo ra các giao diện cho phép con người thể hiện ý định một cách hiệu quả, hiểu được hành động của tác nhân và can thiệp khi cần thiết. Các loại UI đa dạng, từ hội thoại đến đồ họa và đa phương thức, cung cấp các tùy chọn phong phú để phù hợp với các trường hợp sử dụng khác nhau. Google ADK nhấn mạnh sự tách biệt giữa logic tác nhân và UI, tạo điều kiện thuận lợi cho việc tích hợp các tác nhân vào các giao diện đa dạng. Khi chúng ta tiếp tục khám phá các khía cạnh của thiết kế agentic, chương tiếp theo sẽ đi sâu vào việc giám sát hoạt động của tác nhân, điều này rất quan trọng để đảm bảo hiệu suất, hành vi và độ tin cậy của chúng.

## Tài liệu tham khảo
1.  Google ADK Documentation (Developing Interfaces): https://google.github.io/adk-docs/interfaces/
2.  "Designing Bots" by Amir Shevat: https://www.amazon.com/Designing-Bots-Creating-Conversational-Interfaces/dp/1491976008
3.  "Conversational AI: Dialogue Systems, Theory, Practice, and Research" by Michael McTear: https://www.amazon.com/Conversational-AI-Dialogue-Systems-Practice/dp/178712613X
4.  Nielsen Norman Group - AI User Experience: https://www.nngroup.com/topic/ai-ux/
