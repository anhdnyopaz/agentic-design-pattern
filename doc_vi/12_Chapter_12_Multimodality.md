# Chương 12: Multimodality (Nhận thức và Tạo Sinh Đa Phương thức)

## Tổng quan về Mẫu thiết kế Multimodality

Các mô hình ngôn ngữ lớn (LLM) ban đầu được thiết kế để xử lý và tạo văn bản. Tuy nhiên, thế giới thực của con người rất phong phú và đa dạng, chúng ta giao tiếp và hiểu thông tin thông qua nhiều "phương thức" (modalities) khác nhau—văn bản, hình ảnh, âm thanh, video, và thậm chí cả cảm ứng và mùi vị. Để các tác nhân AI thực sự nhận thức được thế giới và tương tác với nó một cách tự nhiên và hiệu quả, chúng phải vượt ra ngoài các tương tác chỉ dựa trên văn bản và đón nhận **Multimodality (Đa phương thức)**.

Mẫu Multimodality đề cập đến khả năng của một tác nhân AI để **nhận thức** (hiểu) thông tin từ nhiều phương thức đầu vào và **tạo sinh** (tạo ra) thông tin hoặc hành động dưới nhiều phương thức đầu ra khác nhau.

### Nhận thức Đa phương thức (Multi-Modal Perception)

Điều này liên quan đến việc xử lý và hiểu dữ liệu từ nhiều nguồn khác nhau đồng thời. Ví dụ, khi một tác nhân được cung cấp một hình ảnh với văn bản nhúng, các nhãn làm nổi bật các đoạn văn bản cụ thể và dữ liệu bảng giải thích từng nhãn, nó cần khả năng:

1.  **Phân tích Hình ảnh:** Hiểu được nội dung trực quan, các đối tượng, cảnh và bố cục chung.
2.  **Đọc Văn bản (OCR):** Trích xuất văn bản từ hình ảnh và hiểu ý nghĩa ngữ nghĩa của nó.
3.  **Xử lý Dữ liệu Cấu trúc:** Kết hợp thông tin từ bảng, liên kết các nhãn với dữ liệu và hiểu mối quan hệ của chúng.
4.  **Tổng hợp:** Kết hợp tất cả các phương thức này để xây dựng một sự hiểu biết toàn diện về thông tin được trình bày.

### Tạo Sinh Đa phương thức (Multi-Modal Generation)

Sau khi hiểu thông tin đa phương thức, một tác nhân đa phương thức có thể tạo ra các phản hồi không chỉ bằng văn bản mà còn bằng hình ảnh, âm thanh hoặc các dạng khác. Ví dụ:

1.  **Tạo Hình ảnh/Video:** Tạo hình ảnh hoặc video dựa trên mô tả văn bản.
2.  **Tổng hợp Giọng nói (TTS):** Chuyển đổi văn bản thành lời nói tự nhiên.
3.  **Tạo Âm thanh:** Tạo nhạc, hiệu ứng âm thanh hoặc giọng nói.
4.  **Tạo Văn bản từ phương thức khác:** Mô tả hình ảnh hoặc video bằng văn bản.

Việc tích hợp các phương thức này cho phép các tác nhân có trải nghiệm phong phú hơn, tương tự như con người trong thế giới. Một tác nhân đa phương thức có thể xem một video, nghe lời nói, đọc phụ đề và tổng hợp tất cả để trả lời một câu hỏi, và sau đó tạo ra một hình ảnh hoặc một đoạn âm thanh làm phản hồi.

Các LLM mới nhất như Gemini của Google được thiết kế với nhận thức và tạo sinh đa phương thức, có thể xử lý các phương thức khác nhau đồng thời, dẫn đến sự hiểu biết và phản ứng mạch lạc hơn. Các framework agentic tận dụng các LLM này để xây dựng các tác nhân có thể tương tác với các phương thức đa dạng này một cách tự nhiên.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu Multimodality mở ra một loạt các trường hợp sử dụng mới cho các tác nhân AI:

1.  **Trợ lý Ảo Nâng cao:**
    *   **Nhận thức:** Một trợ lý có thể "nhìn" một hình ảnh về một thiết bị bị hỏng, "nghe" mô tả giọng nói của người dùng về vấn đề, và "đọc" mã lỗi trong hướng dẫn sử dụng.
    *   **Tạo sinh:** Nó có thể trả lời bằng cách "tạo" một video hướng dẫn sửa chữa, "nói" hướng dẫn từng bước, hoặc "tạo" một sơ đồ có chú thích hiển thị các bộ phận cần thay thế.
2.  **Kiểm duyệt Nội dung:**
    *   **Nhận thức:** Một tác nhân có thể phân tích một bài đăng trên mạng xã hội bao gồm văn bản, hình ảnh và video. Nó có thể xác định ngôn ngữ độc hại, hình ảnh không phù hợp hoặc nội dung video vi phạm chính sách cùng một lúc.
    *   **Tạo sinh:** Nó có thể "tạo" một báo cáo tóm tắt các vi phạm, "gắn cờ" nội dung để xem xét thủ công và "tạo" thông báo cho người dùng về hành vi của họ.
3.  **Học tập và Đào tạo Tương tác:**
    *   **Nhận thức:** Một tác nhân AI giáo dục có thể "xem" bài tập về nhà viết tay của học sinh (hình ảnh), "nghe" học sinh đọc câu trả lời của họ (âm thanh) và "đọc" nội dung sách giáo khoa liên quan (văn bản).
    *   **Tạo sinh:** Nó có thể "tạo" phản hồi được cá nhân hóa bằng giọng nói, "vẽ" các khái niệm khó hiểu trên bảng trắng ảo và "tạo" các câu hỏi luyện tập bổ sung bằng văn bản.
4.  **Tạo Nội dung Sáng tạo:**
    *   **Nhận thức:** Một nhà thiết kế game có thể cung cấp mô tả văn bản cho một cảnh, một hình ảnh tham chiếu cho phong cách và một đoạn âm thanh cho tâm trạng.
    *   **Tạo sinh:** Tác nhân có thể "tạo" một hình ảnh concept phù hợp, "tạo" một đoạn nhạc nền và "viết" một kịch bản ngắn cho cảnh đó.
5.  **Robot học và Tương tác Thực tế Tăng cường (AR):**
    *   **Nhận thức:** Một robot có thể "nhìn" môi trường của nó thông qua camera, "nghe" lệnh của con người và "cảm nhận" vật thể thông qua cảm biến xúc giác.
    *   **Tạo sinh:** Nó có thể "di chuyển" các chi của nó để thực hiện một tác vụ, "nói" xác nhận hoàn thành, và "hiển thị" thông tin trên màn hình AR.
6.  **Chẩn đoán Y tế:**
    *   **Nhận thức:** Một tác nhân AI y tế có thể "phân tích" hình ảnh X-quang hoặc MRI, "đọc" bệnh án của bệnh nhân (văn bản) và "nghe" mô tả triệu chứng của bác sĩ (âm thanh).
    *   **Tạo sinh:** Nó có thể "tạo" một báo cáo chẩn đoán, "hiển thị" các khu vực đáng lo ngại trên hình ảnh y tế và "đề xuất" các xét nghiệm bổ sung bằng văn bản.

Multimodality cho phép các tác nhân thu hẹp khoảng cách giữa các thế giới kỹ thuật số và vật lý, dẫn đến các tương tác tự nhiên và mạnh mẽ hơn nhiều.

## Ví dụ Code Thực hành (Google ADK)

Google Agent Developer Kit (ADK) được thiết kế để tận dụng các khả năng đa phương thức vốn có của các mô hình Gemini. Dưới đây là cách bạn có thể định nghĩa các tác nhân trong ADK để phản hồi các yêu cầu đa phương thức và tạo ra các phản hồi đa phương thức.

Ví dụ này cho thấy cách thiết lập một tác nhân ADK đơn giản để phản hồi các yêu cầu đa phương thức. Tác nhân này, được gọi là `multi_modal_agent`, sử dụng mô hình `gemini-2.0-flash`. Hướng dẫn của nó chỉ dẫn nó phản hồi theo phong cách thân thiện và hỗ trợ, và nếu một hình ảnh được cung cấp, nó sẽ mô tả hình ảnh bằng văn bản.

```python
import uuid
import asyncio
from google.adk.agents import LlmAgent
from google.adk.runners import InMemoryRunner
from google.genai import types
import nest_asyncio

# Áp dụng nest_asyncio để cho phép chạy asyncio trong môi trường này (nếu cần, ví dụ: Jupyter)
nest_asyncio.apply()

# --- Định nghĩa Tác nhân Đa phương thức ---
multi_modal_agent = LlmAgent(
    name='multi_modal_agent',
    model='gemini-2.0-flash', # Mô hình Gemini hỗ trợ đa phương thức
    instruction=(
        "Bạn là một trợ lý thân thiện và hữu ích. "
        "Phản hồi các câu hỏi của người dùng. "
        "Nếu người dùng cung cấp một hình ảnh, hãy mô tả hình ảnh đó bằng văn bản. "
        "Nếu không có hình ảnh, hãy trả lời truy vấn của họ."
    ),
    description="Một tác nhân có thể hiểu và phản hồi cả văn bản và hình ảnh."
)

# --- Logic Thực thi ---
async def run_multi_modal_agent(user_message: str, image_uri: str = None):
    print(f"\n--- Chạy Tác nhân với Tin nhắn: '{user_message}' (Hình ảnh: {image_uri or 'Không có'}) ---")
    runner = InMemoryRunner(multi_modal_agent)
    user_id = "user_123"
    session_id = str(uuid.uuid4())
    
    parts = [types.Part(text=user_message)]
    if image_uri:
        parts.append(types.Part(file_uri=image_uri, mime_type="image/jpeg")) # Giả định là JPEG

    response_text = ""
    async for event in runner.run(
        user_id=user_id,
        session_id=session_id,
        new_message=types.Content(
            role='user',
            parts=parts
        )
    ):
        if event.is_final_response() and event.content:
             if hasattr(event.content, 'text') and event.content.text:
                response_text = event.content.text
             elif event.content.parts:
                 response_text = "".join([part.text for part in event.content.parts if part.text])
    
    print(f"Phản hồi cuối cùng của Tác nhân: {response_text}")

# --- Chạy Ví dụ ---
async def main():
    # Ví dụ 1: Chỉ văn bản
    await run_multi_modal_agent("Xin chào, bạn có khỏe không?")

    # Ví dụ 2: Văn bản và Hình ảnh (sử dụng một hình ảnh placeholder)
    # LƯU Ý: Đối với ví dụ thực tế, bạn sẽ cần một URI hình ảnh hợp lệ có thể truy cập được.
    # Đây là một URI hình ảnh mẫu (ảnh mèo)
    sample_image_uri = "https://www.gstatic.com/gemini/reference/images/cat.jpeg" 
    await run_multi_modal_agent("Bạn thấy gì trong hình ảnh này?", image_uri=sample_image_uri)
    
    # Ví dụ 3: Chỉ hình ảnh (yêu cầu mô tả)
    await run_multi_modal_agent("Mô tả hình ảnh này.", image_uri=sample_image_uri)


if __name__ == "__main__":
    asyncio.run(main())
```

Mã này trình bày một tác nhân ADK tận dụng khả năng đa phương thức của mô hình Gemini 2.0 Flash. Tác nhân được khởi tạo với một instruction đơn giản hướng dẫn nó phản hồi thân thiện và mô tả hình ảnh nếu được cung cấp. Hàm `run_multi_modal_agent` được thiết kế để xử lý cả đầu vào văn bản và hình ảnh. Nó tạo một `types.Content` với các `parts` tương ứng cho văn bản và/hoặc URI hình ảnh. `InMemoryRunner` sau đó được sử dụng để thực thi tác nhân và xử lý phản hồi. Hàm `main` minh họa ba trường hợp sử dụng: một truy vấn chỉ văn bản, một truy vấn kết hợp văn bản và hình ảnh, và một truy vấn chỉ hình ảnh (yêu cầu mô tả).

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các mô hình AI ban đầu bị giới hạn trong một phương thức duy nhất (chủ yếu là văn bản), ngăn cản chúng hiểu hoặc tạo sinh trong thế giới phong phú, đa phương thức mà con người chúng ta trải nghiệm. Hạn chế này dẫn đến các tương tác nhân tạo, thiếu sự hiểu biết về ngữ cảnh từ các phương thức không phải văn bản và không có khả năng tạo ra các phản hồi đa dạng, tự nhiên. Nếu không có đa phương thức, các tác nhân AI vẫn còn xa lạ với cách con người giao tiếp và cảm nhận.
*   **Tại sao:** Mẫu Multimodality cung cấp một giải pháp tiêu chuẩn hóa bằng cách cho phép các tác nhân AI nhận thức (hiểu) thông tin từ nhiều phương thức đầu vào (văn bản, hình ảnh, âm thanh, video, v.v.) và tạo sinh (tạo ra) phản hồi dưới nhiều phương thức đầu ra. Các LLM hiện đại như Gemini được thiết kế với các khả năng đa phương thức gốc, có thể đồng thời xử lý và tổng hợp thông tin từ các phương thức khác nhau. Điều này cho phép các tác nhân AI xây dựng sự hiểu biết toàn diện hơn về môi trường, tạo ra các phản hồi phong phú và tự nhiên hơn, và tương tác với con người theo cách trực quan hơn.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu Multimodality bất cứ khi nào tác nhân AI của bạn cần tương tác với người dùng hoặc môi trường theo cách tự nhiên và toàn diện hơn, vượt ra ngoài các tương tác chỉ dựa trên văn bản. Điều này rất cần thiết cho các ứng dụng yêu cầu nhận thức phong phú (ví dụ: mô tả một hình ảnh phức tạp hoặc phân tích một video), tạo sinh biểu cảm (ví dụ: tạo hình ảnh dựa trên văn bản hoặc phản hồi bằng giọng nói), hoặc giao tiếp trong các môi trường đa phương thức hỗn hợp (ví dụ: trợ lý ảo, robot, kiểm duyệt nội dung).

## Những Điểm Chính (Key Takeaways)

*   Đa phương thức là khả năng của một tác nhân để xử lý và tạo thông tin bằng cách sử dụng nhiều phương thức (văn bản, hình ảnh, âm thanh, video).
*   Nó cho phép các tác nhân nhận thức và tương tác với thế giới một cách tự nhiên và phong phú hơn.
*   Các mô hình như Google Gemini được xây dựng với khả năng đa phương thức gốc, có thể xử lý các phương thức khác nhau đồng thời.
*   ADK đơn giản hóa việc xây dựng các tác nhân đa phương thức bằng cách tận dụng các mô hình Gemini hỗ trợ đa phương thức.
*   Các ví dụ về ứng dụng thực tế bao gồm trợ lý ảo nâng cao, kiểm duyệt nội dung, học tập tương tác, tạo nội dung sáng tạo, robot học và chẩn đoán y tế.

## Kết luận

Chương này đã làm nổi bật vai trò biến đổi của Multimodality, khả năng của các tác nhân AI để nhận thức và tạo sinh thông tin qua nhiều phương thức khác nhau. Bằng cách vượt ra ngoài các tương tác chỉ dựa trên văn bản, các tác nhân có thể đạt được sự hiểu biết toàn diện hơn về môi trường và giao tiếp theo cách tự nhiên, phong phú hơn. Các LLM đa phương thức như Gemini và các framework như Google ADK cung cấp các công cụ cần thiết để xây dựng các tác nhân mạnh mẽ này. Việc tích hợp Multimodality là một bước quan trọng trong việc phát triển các tác nhân AI có thể tương tác liền mạch với thế giới thực, thu hẹp khoảng cách giữa các tương tác kỹ thuật số và trải nghiệm của con người. Chương tiếp theo sẽ giới thiệu tầm quan trọng của việc đảm bảo rằng các tác nhân AI hoạt động trong các ranh giới đạo đức.

## Tài liệu tham khảo
1.  Google Gemini API: https://ai.google.dev/models/gemini
2.  Google ADK Documentation (Multimodal Agents): https://google.github.io/adk-docs/agents/#multimodal-agents
3.  OpenAI GPT-4V (Vision) Documentation: https://openai.com/research/gpt-4v-system-card
