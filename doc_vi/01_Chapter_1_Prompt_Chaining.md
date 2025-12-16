# Chương 1: Prompt Chaining (Chuỗi Lời Nhắc)

## Tổng quan về Mẫu thiết kế Prompt Chaining

Prompt Chaining (Chuỗi lời nhắc), đôi khi được gọi là mẫu Pipeline (Đường ống), đại diện cho một mô hình mạnh mẽ để xử lý các tác vụ phức tạp khi tận dụng các mô hình ngôn ngữ lớn (LLM). Thay vì mong đợi một LLM giải quyết một vấn đề phức tạp chỉ trong một bước đơn lẻ và nguyên khối, Prompt Chaining ủng hộ chiến lược "chia để trị". Ý tưởng cốt lõi là chia nhỏ vấn đề ban đầu, khó khăn thành một chuỗi các vấn đề con nhỏ hơn, dễ quản lý hơn. Mỗi vấn đề con được giải quyết riêng biệt thông qua một lời nhắc (prompt) được thiết kế cụ thể, và đầu ra (output) tạo ra từ một prompt sẽ được sử dụng một cách chiến lược làm đầu vào (input) cho prompt tiếp theo trong chuỗi.

Kỹ thuật xử lý tuần tự này vốn dĩ mang lại tính mô-đun và sự rõ ràng trong việc tương tác với các LLM. Bằng cách phân rã một tác vụ phức tạp, việc hiểu và gỡ lỗi (debug) từng bước riêng lẻ trở nên dễ dàng hơn, làm cho quy trình tổng thể trở nên mạnh mẽ và dễ diễn giải hơn. Mỗi bước trong chuỗi có thể được chế tác và tối ưu hóa tỉ mỉ để tập trung vào một khía cạnh cụ thể của vấn đề lớn hơn, dẫn đến các kết quả đầu ra chính xác và tập trung hơn.

Việc đầu ra của bước này đóng vai trò là đầu vào cho bước tiếp theo là rất quan trọng. Việc chuyển giao thông tin này thiết lập một chuỗi phụ thuộc (do đó có tên là chaining), nơi bối cảnh và kết quả của các hoạt động trước đó hướng dẫn quá trình xử lý tiếp theo. Điều này cho phép LLM xây dựng dựa trên công việc trước đó của nó, tinh chỉnh sự hiểu biết và tiến dần đến giải pháp mong muốn.

Hơn nữa, Prompt Chaining không chỉ là về việc chia nhỏ vấn đề; nó còn cho phép tích hợp kiến thức và công cụ bên ngoài. Tại mỗi bước, LLM có thể được hướng dẫn để tương tác với các hệ thống bên ngoài, API hoặc cơ sở dữ liệu, làm giàu kiến thức và khả năng của nó vượt ra ngoài dữ liệu huấn luyện nội bộ. Khả năng này mở rộng đáng kể tiềm năng của các LLM, cho phép chúng hoạt động không chỉ như các mô hình biệt lập mà còn là các thành phần tích hợp của các hệ thống thông minh rộng lớn hơn.

Tầm quan trọng của Prompt Chaining vượt xa việc giải quyết vấn đề đơn giản. Nó đóng vai trò là một kỹ thuật nền tảng để xây dựng các tác nhân AI (AI Agents) tinh vi. Các tác nhân này có thể sử dụng các chuỗi prompt để tự chủ lập kế hoạch, suy luận và hành động trong các môi trường động. Bằng cách cấu trúc chiến lược chuỗi các prompt, một tác nhân có thể tham gia vào các tác vụ đòi hỏi suy luận, lập kế hoạch và ra quyết định đa bước. Các quy trình làm việc của tác nhân như vậy có thể mô phỏng các quy trình suy nghĩ của con người chặt chẽ hơn, cho phép các tương tác tự nhiên và hiệu quả hơn với các miền và hệ thống phức tạp.

### Hạn chế của các lời nhắc đơn lẻ (Single Prompts)
Đối với các tác vụ đa diện, việc sử dụng một prompt đơn lẻ, phức tạp cho một LLM có thể không hiệu quả, khiến mô hình phải vật lộn với các ràng buộc và hướng dẫn. Điều này có thể dẫn đến việc bỏ qua hướng dẫn (instruction neglect) nơi các phần của prompt bị bỏ qua, trôi dạt ngữ cảnh (contextual drift) nơi mô hình mất dấu ngữ cảnh ban đầu, lan truyền lỗi (error propagation) nơi các lỗi sớm bị khuếch đại, các prompt yêu cầu cửa sổ ngữ cảnh dài hơn nơi mô hình nhận không đủ thông tin để phản hồi lại, và ảo giác (hallucination) nơi tải nhận thức tăng làm tăng khả năng thông tin không chính xác. Ví dụ, một truy vấn yêu cầu phân tích báo cáo nghiên cứu thị trường, tóm tắt các phát hiện, xác định xu hướng với các điểm dữ liệu và soạn thảo email có nguy cơ thất bại vì mô hình có thể tóm tắt tốt nhưng không trích xuất được dữ liệu hoặc soạn thảo email không đúng cách.

### Tăng cường độ tin cậy thông qua phân rã tuần tự
Prompt Chaining giải quyết những thách thức này bằng cách chia nhỏ tác vụ phức tạp thành một quy trình làm việc tuần tự, tập trung, giúp cải thiện đáng kể độ tin cậy và khả năng kiểm soát. Với ví dụ trên, một cách tiếp cận đường ống hoặc chuỗi có thể được mô tả như sau:

1.  **Prompt Ban đầu (Tóm tắt):** "Tóm tắt những phát hiện chính của báo cáo nghiên cứu thị trường sau: [văn bản]." Trọng tâm duy nhất của mô hình là tóm tắt, làm tăng độ chính xác của bước đầu tiên này.
2.  **Prompt Thứ hai (Xác định Xu hướng):** "Sử dụng bản tóm tắt, hãy xác định ba xu hướng mới nổi hàng đầu và trích xuất các điểm dữ liệu cụ thể hỗ trợ mỗi xu hướng: [kết quả từ bước 1]." Prompt này giờ đây bị ràng buộc hơn và xây dựng trực tiếp dựa trên một đầu ra đã được xác thực.
3.  **Prompt Thứ ba (Soạn thảo Email):** "Soạn thảo một email ngắn gọn gửi cho đội ngũ marketing phác thảo các xu hướng sau và dữ liệu hỗ trợ của chúng: [kết quả từ bước 2]."

Sự phân rã này cho phép kiểm soát chi tiết hơn đối với quy trình. Mỗi bước đơn giản hơn và ít mơ hồ hơn, làm giảm tải nhận thức cho mô hình và dẫn đến kết quả cuối cùng chính xác và đáng tin cậy hơn. Tính mô-đun này tương tự như một đường ống tính toán nơi mỗi chức năng thực hiện một hoạt động cụ thể trước khi chuyển kết quả của nó cho chức năng tiếp theo. Để đảm bảo phản hồi chính xác cho từng tác vụ cụ thể, mô hình có thể được gán một vai trò riêng biệt ở mỗi giai đoạn. Trong kịch bản đã cho, prompt ban đầu có thể được chỉ định là "Chuyên gia phân tích thị trường", prompt tiếp theo là "Chuyên gia phân tích thương mại", và prompt thứ ba là "Người viết tài liệu chuyên nghiệp", v.v.

### Vai trò của Đầu ra có Cấu trúc (Structured Output)
Độ tin cậy của một chuỗi prompt phụ thuộc rất nhiều vào tính toàn vẹn của dữ liệu được truyền giữa các bước. Nếu đầu ra của một prompt mơ hồ hoặc định dạng kém, prompt tiếp theo có thể thất bại do đầu vào bị lỗi. Để giảm thiểu điều này, việc chỉ định định dạng đầu ra có cấu trúc, chẳng hạn như JSON hoặc XML, là rất quan trọng.

Ví dụ, đầu ra từ bước xác định xu hướng có thể được định dạng dưới dạng đối tượng JSON:
```json
{
"trends": [
{
"trend_name": "Cá nhân hóa được hỗ trợ bởi AI",
"supporting_data": "73% người tiêu dùng thích làm việc với các thương hiệu sử dụng thông tin cá nhân để làm cho trải nghiệm mua sắm của họ phù hợp hơn."
},
{
"trend_name": "Thương hiệu Bền vững và Có đạo đức",
"supporting_data": "Doanh số của các sản phẩm có tuyên bố liên quan đến ESG tăng 28% trong 5 năm qua, so với 20% cho các sản phẩm không có."
}
]
}
```
Định dạng có cấu trúc này đảm bảo rằng dữ liệu có thể đọc được bằng máy và có thể được phân tích cú pháp chính xác và chèn vào prompt tiếp theo mà không gây mơ hồ. Thực tiễn này giảm thiểu các lỗi có thể phát sinh từ việc diễn giải ngôn ngữ tự nhiên và là một thành phần quan trọng trong việc xây dựng các hệ thống dựa trên LLM đa bước mạnh mẽ.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Prompt Chaining là một mẫu linh hoạt có thể áp dụng trong nhiều tình huống khi xây dựng các hệ thống agentic. Tiện ích cốt lõi của nó nằm ở việc chia nhỏ các vấn đề phức tạp thành các bước tuần tự, dễ quản lý. Dưới đây là một số ứng dụng thực tế và trường hợp sử dụng:

### 1. Quy trình Xử lý Thông tin (Information Processing Workflows)
Nhiều tác vụ liên quan đến việc xử lý thông tin thô qua nhiều biến đổi. Ví dụ: tóm tắt tài liệu, trích xuất các thực thể chính, và sau đó sử dụng các thực thể đó để truy vấn cơ sở dữ liệu hoặc tạo báo cáo. Một chuỗi prompt có thể trông như sau:
*   **Prompt 1:** Trích xuất nội dung văn bản từ một URL hoặc tài liệu nhất định.
*   **Prompt 2:** Tóm tắt văn bản đã được làm sạch.
*   **Prompt 3:** Trích xuất các thực thể cụ thể (ví dụ: tên, ngày tháng, địa điểm) từ bản tóm tắt hoặc văn bản gốc.
*   **Prompt 4:** Sử dụng các thực thể để tìm kiếm trong cơ sở tri thức nội bộ.
*   **Prompt 5:** Tạo báo cáo cuối cùng kết hợp bản tóm tắt, các thực thể và kết quả tìm kiếm.

Phương pháp này được áp dụng trong các lĩnh vực như phân tích nội dung tự động, phát triển trợ lý nghiên cứu dựa trên AI và tạo báo cáo phức tạp.

### 2. Trả lời Truy vấn Phức tạp (Complex Query Answering)
Trả lời các câu hỏi phức tạp đòi hỏi nhiều bước suy luận hoặc truy xuất thông tin là một trường hợp sử dụng chính. Ví dụ: "Nguyên nhân chính của sự sụp đổ thị trường chứng khoán năm 1929 là gì, và chính sách của chính phủ đã phản ứng như thế nào?"
*   **Prompt 1:** Xác định các câu hỏi phụ cốt lõi trong truy vấn của người dùng (nguyên nhân sụp đổ, phản ứng của chính phủ).
*   **Prompt 2:** Nghiên cứu hoặc truy xuất thông tin cụ thể về nguyên nhân của sự sụp đổ năm 1929.
*   **Prompt 3:** Nghiên cứu hoặc truy xuất thông tin cụ thể về phản ứng chính sách của chính phủ đối với sự sụp đổ thị trường chứng khoán năm 1929.
*   **Prompt 4:** Tổng hợp thông tin từ bước 2 và 3 thành một câu trả lời mạch lạc cho truy vấn ban đầu.

Phương pháp xử lý tuần tự này là không thể thiếu để phát triển các hệ thống AI có khả năng suy luận đa bước và tổng hợp thông tin.

### 3. Trích xuất và Chuyển đổi Dữ liệu (Data Extraction and Transformation)
Việc chuyển đổi văn bản phi cấu trúc thành định dạng có cấu trúc thường đạt được thông qua một quy trình lặp đi lặp lại, yêu cầu các sửa đổi tuần tự để cải thiện độ chính xác và tính đầy đủ của đầu ra.
*   **Prompt 1:** Cố gắng trích xuất các trường cụ thể (ví dụ: tên, địa chỉ, số tiền) từ tài liệu hóa đơn.
*   **Xử lý:** Kiểm tra xem tất cả các trường bắt buộc đã được trích xuất chưa và chúng có đáp ứng các yêu cầu định dạng không.
*   **Prompt 2 (Có điều kiện):** Nếu các trường bị thiếu hoặc sai định dạng, hãy tạo một prompt mới yêu cầu mô hình tìm cụ thể thông tin bị thiếu/sai định dạng, có thể cung cấp ngữ cảnh từ lần thử thất bại.
*   **Xử lý:** Xác thực lại kết quả. Lặp lại nếu cần thiết.
*   **Đầu ra:** Cung cấp dữ liệu có cấu trúc đã được trích xuất và xác thực.

### 4. Quy trình Tạo Nội dung (Content Generation Workflows)
Việc soạn thảo nội dung phức tạp là một tác vụ thủ tục thường được phân rã thành các giai đoạn riêng biệt, bao gồm lên ý tưởng ban đầu, lập dàn ý cấu trúc, viết nháp và sửa đổi sau đó.
*   **Prompt 1:** Tạo 5 ý tưởng chủ đề dựa trên sở thích chung của người dùng.
*   **Xử lý:** Cho phép người dùng chọn một ý tưởng hoặc tự động chọn ý tưởng tốt nhất.
*   **Prompt 2:** Dựa trên chủ đề đã chọn, tạo một dàn ý chi tiết.
*   **Prompt 3:** Viết bản nháp phần dựa trên điểm đầu tiên trong dàn ý.
*   **Prompt 4:** Viết bản nháp phần dựa trên điểm thứ hai trong dàn ý, cung cấp phần trước để làm ngữ cảnh. Tiếp tục điều này cho tất cả các điểm trong dàn ý.
*   **Prompt 5:** Xem xét và tinh chỉnh bản nháp hoàn chỉnh để đảm bảo tính mạch lạc, giọng điệu và ngữ pháp.

### 5. Tác nhân Hội thoại có Trạng thái (Conversational Agents with State)
Mặc dù các kiến trúc quản lý trạng thái toàn diện sử dụng các phương pháp phức tạp hơn là liên kết tuần tự, prompt chaining cung cấp một cơ chế nền tảng để duy trì tính liên tục của cuộc hội thoại. Kỹ thuật này duy trì ngữ cảnh bằng cách xây dựng mỗi lượt hội thoại như một prompt mới kết hợp một cách có hệ thống thông tin hoặc các thực thể đã trích xuất từ các tương tác trước đó trong chuỗi đối thoại.

### 6. Tạo và Tinh chỉnh Mã (Code Generation and Refinement)
Việc tạo mã chức năng thường là một quy trình đa giai đoạn, yêu cầu một vấn đề được phân rã thành một chuỗi các hoạt động logic rời rạc được thực thi dần dần.
*   **Prompt 1:** Hiểu yêu cầu của người dùng về một hàm mã. Tạo mã giả hoặc dàn ý.
*   **Prompt 2:** Viết bản nháp mã ban đầu dựa trên dàn ý.
*   **Prompt 3:** Xác định các lỗi tiềm ẩn hoặc các khu vực cần cải thiện trong mã (có thể sử dụng công cụ phân tích tĩnh hoặc một lệnh gọi LLM khác).
*   **Prompt 4:** Viết lại hoặc tinh chỉnh mã dựa trên các vấn đề đã xác định.
*   **Prompt 5:** Thêm tài liệu hoặc các trường hợp kiểm thử (test cases).

### 7. Suy luận Đa phương thức và Đa bước (Multimodal and multi-step reasoning)
Phân tích các bộ dữ liệu với các phương thức đa dạng đòi hỏi phải chia nhỏ vấn đề thành các tác vụ nhỏ hơn dựa trên prompt. Ví dụ: diễn giải một hình ảnh có chứa hình ảnh với văn bản nhúng, các nhãn làm nổi bật các đoạn văn bản cụ thể và dữ liệu bảng giải thích từng nhãn, đòi hỏi một cách tiếp cận như vậy.
*   **Prompt 1:** Trích xuất và hiểu văn bản từ yêu cầu hình ảnh của người dùng.
*   **Prompt 2:** Liên kết văn bản hình ảnh đã trích xuất với các nhãn tương ứng của nó.
*   **Prompt 3:** Diễn giải thông tin đã thu thập bằng cách sử dụng một bảng để xác định đầu ra cần thiết.

## Ví dụ Code Thực hành (LangChain)

Việc triển khai prompt chaining bao gồm từ các lệnh gọi hàm tuần tự, trực tiếp trong một tập lệnh đến việc sử dụng các framework chuyên biệt. LangChain và LangGraph là những lựa chọn phù hợp vì các API cốt lõi của chúng được thiết kế rõ ràng để soạn thảo các chuỗi và đồ thị hoạt động.

Mã sau đây triển khai một chuỗi prompt hai bước hoạt động như một đường ống xử lý dữ liệu. Giai đoạn đầu được thiết kế để phân tích cú pháp văn bản phi cấu trúc và trích xuất thông tin cụ thể. Giai đoạn tiếp theo sau đó nhận đầu ra đã trích xuất này và chuyển đổi nó thành định dạng dữ liệu có cấu trúc.

```python
import os
# from dotenv import load_dotenv # Để bảo mật tốt hơn, tải biến môi trường từ file .env
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser

# load_dotenv()
# Đảm bảo OPENAI_API_KEY của bạn được thiết lập trong file .env hoặc môi trường
# Khởi tạo Mô hình Ngôn ngữ (khuyến nghị sử dụng ChatOpenAI)
llm = ChatOpenAI(temperature=0)

# --- Prompt 1: Trích xuất Thông tin ---
prompt_extract = ChatPromptTemplate.from_template(
    "Trích xuất các thông số kỹ thuật từ văn bản sau:\n\n{text_input}"
)

# --- Prompt 2: Chuyển đổi sang JSON ---
prompt_transform = ChatPromptTemplate.from_template(
    "Chuyển đổi các thông số kỹ thuật sau thành một đối tượng JSON với các khóa 'cpu', 'memory', và 'storage':\n\n{specifications}"
)

# --- Xây dựng Chuỗi sử dụng LCEL (LangChain Expression Language) ---
# StrOutputParser() chuyển đổi đầu ra tin nhắn của LLM thành một chuỗi đơn giản.
extraction_chain = prompt_extract | llm | StrOutputParser()

# Chuỗi đầy đủ chuyển đầu ra của chuỗi trích xuất vào biến 'specifications'
# cho prompt chuyển đổi.
full_chain = (
    {"specifications": extraction_chain}
    | prompt_transform
    | llm
    | StrOutputParser()
)

# --- Chạy Chuỗi ---
input_text = "Mẫu laptop mới có bộ vi xử lý 3.5 GHz octa-core, 16GB RAM, và ổ cứng 1TB NVMe SSD."

# Thực thi chuỗi với từ điển văn bản đầu vào.
final_result = full_chain.invoke({"text_input": input_text})

print("\n--- Kết quả JSON Cuối cùng ---")
print(final_result)
```

Mã Python này minh họa cách sử dụng thư viện LangChain để xử lý văn bản. Nó sử dụng hai prompt riêng biệt: một để trích xuất thông số kỹ thuật từ chuỗi đầu vào và một để định dạng các thông số này thành đối tượng JSON. `StrOutputParser` đảm bảo đầu ra là một chuỗi có thể sử dụng được. `LangChain Expression Language (LCEL)` được sử dụng để nối các prompt và mô hình ngôn ngữ lại với nhau một cách thanh lịch. `full_chain` sau đó lấy đầu ra của quá trình trích xuất và sử dụng nó làm đầu vào cho prompt chuyển đổi.

## Context Engineering và Prompt Engineering

**Context Engineering** (Kỹ thuật Ngữ cảnh) là kỷ luật có hệ thống về việc thiết kế, xây dựng và cung cấp một môi trường thông tin hoàn chỉnh cho một mô hình AI trước khi tạo token. Phương pháp này khẳng định rằng chất lượng đầu ra của một mô hình ít phụ thuộc vào kiến trúc của chính mô hình đó và phụ thuộc nhiều hơn vào sự phong phú của ngữ cảnh được cung cấp.

Nó đại diện cho một sự tiến hóa đáng kể từ **Prompt Engineering** (Kỹ thuật Lời nhắc) truyền thống, vốn tập trung chủ yếu vào việc tối ưu hóa cách diễn đạt truy vấn ngay lập tức của người dùng. Context Engineering mở rộng phạm vi này để bao gồm nhiều lớp thông tin, chẳng hạn như **system prompt** (lời nhắc hệ thống), là một tập hợp các hướng dẫn nền tảng xác định các tham số hoạt động của AI (ví dụ: "Bạn là một người viết kỹ thuật; giọng điệu của bạn phải trang trọng và chính xác").

Ngữ cảnh được làm giàu thêm với dữ liệu bên ngoài. Điều này bao gồm các tài liệu được truy xuất (RAG), nơi AI chủ động lấy thông tin từ cơ sở tri thức để thông báo cho phản hồi của nó. Nó cũng kết hợp các đầu ra của công cụ, là kết quả từ việc AI sử dụng API bên ngoài để lấy dữ liệu thời gian thực. Dữ liệu rõ ràng này được kết hợp với dữ liệu ngầm quan trọng, chẳng hạn như danh tính người dùng, lịch sử tương tác và trạng thái môi trường. Nguyên tắc cốt lõi là ngay cả các mô hình tiên tiến cũng hoạt động kém khi được cung cấp một cái nhìn hạn chế hoặc được xây dựng kém về môi trường hoạt động.

Do đó, thực hành này định hình lại nhiệm vụ từ việc chỉ trả lời một câu hỏi sang việc xây dựng một bức tranh hoạt động toàn diện cho tác nhân. Cuối cùng, Context Engineering là một phương pháp quan trọng để nâng cao các chatbot phi trạng thái thành các hệ thống có khả năng cao, nhận thức được tình huống.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các tác vụ phức tạp thường làm quá tải các LLM khi được xử lý trong một prompt duy nhất, dẫn đến các vấn đề hiệu suất đáng kể. Tải nhận thức lên mô hình làm tăng khả năng xảy ra lỗi như bỏ qua hướng dẫn, mất ngữ cảnh và tạo ra thông tin không chính xác.
*   **Tại sao:** Prompt Chaining cung cấp một giải pháp tiêu chuẩn hóa bằng cách chia nhỏ một vấn đề phức tạp thành một chuỗi các tác vụ phụ nhỏ hơn, được kết nối với nhau. Mỗi bước trong chuỗi sử dụng một prompt tập trung để thực hiện một hoạt động cụ thể, cải thiện đáng kể độ tin cậy và khả năng kiểm soát. Đầu ra từ một prompt được chuyển làm đầu vào cho prompt tiếp theo, tạo ra một quy trình làm việc logic tiến dần đến giải pháp cuối cùng.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này khi một tác vụ quá phức tạp cho một prompt duy nhất, liên quan đến nhiều giai đoạn xử lý riêng biệt, yêu cầu tương tác với các công cụ bên ngoài giữa các bước, hoặc khi xây dựng các hệ thống Agentic cần thực hiện suy luận đa bước và duy trì trạng thái.

## Những Điểm Chính (Key Takeaways)

*   Prompt Chaining chia nhỏ các tác vụ phức tạp thành một chuỗi các bước nhỏ hơn, tập trung. Điều này đôi khi được gọi là mẫu Pipeline.
*   Mỗi bước trong một chuỗi liên quan đến một lệnh gọi LLM hoặc logic xử lý, sử dụng đầu ra của bước trước đó làm đầu vào.
*   Mẫu này cải thiện độ tin cậy và khả năng quản lý của các tương tác phức tạp với các mô hình ngôn ngữ.
*   Các framework như LangChain/LangGraph và Google ADK cung cấp các công cụ mạnh mẽ để xác định, quản lý và thực thi các trình tự đa bước này.

## Kết luận

Bằng cách giải cấu trúc các vấn đề phức tạp thành một trình tự các tác vụ phụ đơn giản hơn, dễ quản lý hơn, prompt chaining cung cấp một khuôn khổ mạnh mẽ để hướng dẫn các mô hình ngôn ngữ lớn. Chiến lược "chia để trị" này nâng cao đáng kể độ tin cậy và khả năng kiểm soát đầu ra bằng cách tập trung mô hình vào một hoạt động cụ thể tại một thời điểm. Là một mẫu nền tảng, nó cho phép phát triển các tác nhân AI tinh vi có khả năng suy luận đa bước, tích hợp công cụ và quản lý trạng thái. Cuối cùng, việc thành thạo prompt chaining là rất quan trọng để xây dựng các hệ thống mạnh mẽ, nhận thức ngữ cảnh có thể thực thi các quy trình làm việc phức tạp vượt xa khả năng của một prompt đơn lẻ.

## Tài liệu tham khảo
1. LangChain Documentation on LCEL: https://python.langchain.com/v0.2/docs/core_modules/expression_language/
2. LangGraph Documentation: https://langchain-ai.github.io/langgraph/
3. Prompt Engineering Guide - Chaining Prompts: https://www.promptingguide.ai/techniques/chaining
4. OpenAI API Documentation (General Prompting Concepts): https://platform.openai.com/docs/guides/gpt/prompting
5. Crew AI Documentation (Tasks and Processes): https://docs.crewai.com/
6. Google AI for Developers (Prompting Guides): https://cloud.google.com/discover/what-is-prompt-engineering?hl=en
7. Vertex Prompt Optimizer: https://cloud.google.com/vertex-ai/generative-ai/docs/learn/prompts/prompt-optimizer