# Chương 3: Parallelization (Song Song Hóa)

## Tổng quan về Mẫu thiết kế Parallelization

Trong các chương trước, chúng ta đã khám phá Prompt Chaining cho các quy trình làm việc tuần tự và Routing cho việc ra quyết định động và chuyển đổi giữa các đường dẫn khác nhau. Mặc dù các mẫu này rất cần thiết, nhưng nhiều tác vụ agentic phức tạp liên quan đến nhiều tác vụ phụ có thể được thực thi **đồng thời** thay vì tuần tự. Đây là lúc mẫu **Parallelization (Song song hóa)** trở nên quan trọng.

Parallelization là kỹ thuật thực thi nhiều thành phần, chẳng hạn như các lệnh gọi LLM, sử dụng công cụ hoặc thậm chí toàn bộ tác nhân phụ, **cùng lúc**. Thay vì chờ một bước hoàn thành trước khi bắt đầu bước tiếp theo, việc thực thi song song cho phép các tác vụ độc lập chạy cùng một lúc, giảm đáng kể thời gian thực thi tổng thể cho các tác vụ có thể được chia thành các phần độc lập.

Hãy xem xét một tác nhân được thiết kế để nghiên cứu một chủ đề và tóm tắt các phát hiện của nó. Một cách tiếp cận tuần tự có thể là:
1.  Tìm kiếm Nguồn A.
2.  Tóm tắt Nguồn A.
3.  Tìm kiếm Nguồn B.
4.  Tóm tắt Nguồn B.
5.  Tổng hợp câu trả lời cuối cùng từ bản tóm tắt A và B.

Một cách tiếp cận song song có thể thay thế là:
1.  Tìm kiếm Nguồn A và Tìm kiếm Nguồn B đồng thời.
2.  Khi cả hai tìm kiếm hoàn tất, Tóm tắt Nguồn A và Tóm tắt Nguồn B đồng thời.
3.  Tổng hợp câu trả lời cuối cùng từ bản tóm tắt A và B (bước này thường là tuần tự, chờ các bước song song hoàn thành).

Ý tưởng cốt lõi là xác định các phần của quy trình làm việc không phụ thuộc vào đầu ra của các phần khác và thực thi chúng song song. Điều này đặc biệt hiệu quả khi xử lý các dịch vụ bên ngoài (như API hoặc cơ sở dữ liệu) có độ trễ, vì bạn có thể gửi nhiều yêu cầu cùng lúc.

Việc triển khai Parallelization thường yêu cầu các framework hỗ trợ thực thi bất đồng bộ (asynchronous execution) hoặc đa luồng/đa xử lý (multi-threading/multi-processing). Các framework agentic hiện đại được thiết kế với các hoạt động bất đồng bộ, cho phép bạn dễ dàng xác định các bước có thể chạy song song.

Mẫu Parallelization rất quan trọng để cải thiện hiệu quả và khả năng phản hồi của các hệ thống agentic, đặc biệt khi xử lý các tác vụ liên quan đến nhiều lần tra cứu độc lập, tính toán hoặc tương tác với các dịch vụ bên ngoài. Đây là một kỹ thuật quan trọng để tối ưu hóa hiệu suất của các quy trình làm việc agentic phức tạp.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Parallelization là một mẫu mạnh mẽ để tối ưu hóa hiệu suất tác nhân trong nhiều ứng dụng khác nhau:

1.  **Thu thập Thông tin và Nghiên cứu:** Thu thập thông tin từ nhiều nguồn đồng thời là một trường hợp sử dụng cổ điển.
    *   **Trường hợp sử dụng:** Một tác nhân nghiên cứu một công ty.
    *   **Các tác vụ song song:** Tìm kiếm tin tức, lấy dữ liệu chứng khoán, kiểm tra đề cập trên mạng xã hội và truy vấn cơ sở dữ liệu công ty, tất cả cùng một lúc.
    *   **Lợi ích:** Thu thập một cái nhìn toàn diện nhanh hơn nhiều so với các tra cứu tuần tự.
2.  **Xử lý và Phân tích Dữ liệu:** Áp dụng các kỹ thuật phân tích khác nhau hoặc xử lý các phân đoạn dữ liệu khác nhau đồng thời.
    *   **Trường hợp sử dụng:** Một tác nhân phân tích phản hồi của khách hàng.
    *   **Các tác vụ song song:** Chạy phân tích cảm xúc, trích xuất từ khóa, phân loại phản hồi và xác định các vấn đề khẩn cấp đồng thời trên một loạt các mục phản hồi.
    *   **Lợi ích:** Cung cấp phân tích đa diện nhanh chóng.
3.  **Tương tác Đa API hoặc Công cụ:** Gọi nhiều API hoặc công cụ độc lập để thu thập các loại thông tin khác nhau hoặc thực hiện các hành động khác nhau.
    *   **Trường hợp sử dụng:** Một tác nhân lập kế hoạch du lịch.
    *   **Các tác vụ song song:** Kiểm tra giá vé máy bay, tìm kiếm phòng khách sạn, tra cứu các sự kiện địa phương và tìm các đề xuất nhà hàng cùng lúc.
    *   **Lợi ích:** Trình bày một kế hoạch du lịch hoàn chỉnh nhanh hơn.
4.  **Tạo Nội dung với Nhiều Thành phần:** Tạo các phần khác nhau của một nội dung phức tạp song song.
    *   **Trường hợp sử dụng:** Một tác nhân tạo email marketing.
    *   **Các tác vụ song song:** Tạo dòng chủ đề, soạn thảo nội dung email, tìm hình ảnh liên quan và tạo văn bản nút kêu gọi hành động đồng thời.
    *   **Lợi ích:** Lắp ráp email cuối cùng hiệu quả hơn.
5.  **Xác thực và Xác minh:** Thực hiện nhiều kiểm tra hoặc xác thực độc lập đồng thời.
    *   **Trường hợp sử dụng:** Một tác nhân xác minh đầu vào của người dùng.
    *   **Các tác vụ song song:** Kiểm tra định dạng email, xác thực số điện thoại, xác minh địa chỉ so với cơ sở dữ liệu và kiểm tra các từ ngữ thô tục đồng thời.
    *   **Lợi ích:** Cung cấp phản hồi nhanh hơn về tính hợp lệ của đầu vào.
6.  **Xử lý Đa phương thức (Multi-Modal Processing):** Xử lý các phương thức khác nhau (văn bản, hình ảnh, âm thanh) của cùng một đầu vào đồng thời.
    *   **Trường hợp sử dụng:** Một tác nhân phân tích một bài đăng trên mạng xã hội với văn bản và hình ảnh.
    *   **Các tác vụ song song:** Phân tích văn bản để tìm cảm xúc và từ khóa và phân tích hình ảnh để tìm các đối tượng và mô tả cảnh đồng thời.
    *   **Lợi ích:** Tích hợp thông tin từ các phương thức khác nhau nhanh hơn.
7.  **Kiểm tra A/B hoặc Tạo Nhiều Tùy chọn:** Tạo nhiều biến thể của phản hồi hoặc đầu ra song song để chọn biến thể tốt nhất.
    *   **Trường hợp sử dụng:** Một tác nhân tạo các tùy chọn văn bản sáng tạo khác nhau.
    *   **Các tác vụ song song:** Tạo ba tiêu đề khác nhau cho một bài báo đồng thời bằng cách sử dụng các prompt hoặc mô hình hơi khác nhau.
    *   **Lợi ích:** Cho phép so sánh nhanh và lựa chọn tùy chọn tốt nhất.

Parallelization là một kỹ thuật tối ưu hóa cơ bản trong thiết kế agentic, cho phép các nhà phát triển xây dựng các ứng dụng hiệu quả hơn và phản hồi nhanh hơn bằng cách tận dụng việc thực thi đồng thời cho các tác vụ độc lập.

## Ví dụ Code Thực hành (LangChain)

Việc thực thi song song trong framework LangChain được thực hiện bởi LangChain Expression Language (LCEL). Phương pháp chính liên quan đến việc cấu trúc nhiều thành phần có thể chạy (`runnable components`) trong một cấu trúc từ điển hoặc danh sách. Khi tập hợp này được truyền làm đầu vào cho một thành phần tiếp theo trong chuỗi, runtime của LCEL sẽ thực thi các `runnable` có trong đó một cách đồng thời.

Trong ngữ cảnh của LangGraph, nguyên tắc này được áp dụng cho cấu trúc đồ thị. Các quy trình làm việc song song được định nghĩa bằng cách kiến trúc hóa đồ thị sao cho nhiều node, thiếu các phụ thuộc tuần tự trực tiếp, có thể được khởi tạo từ một node chung duy nhất. Các đường dẫn song song này thực thi độc lập trước khi kết quả của chúng có thể được tổng hợp tại một điểm hội tụ tiếp theo trong đồ thị.

Đoạn mã Python dưới đây triển khai một quy trình làm việc xử lý song song được xây dựng bằng framework LangChain. Quy trình làm việc này được thiết kế để thực thi hai hoạt động độc lập đồng thời để phản hồi một truy vấn của người dùng. Các quy trình song song này được khởi tạo dưới dạng các chuỗi hoặc hàm riêng biệt, và các đầu ra tương ứng của chúng sau đó được tổng hợp thành một kết quả thống nhất.

```python
import os
import asyncio
from typing import Optional
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import Runnable, RunnableParallel, RunnablePassthrough

# --- Cấu hình ---
# Đảm bảo biến môi trường OPENAI_API_KEY của bạn đã được thiết lập
# (ví dụ: đặt trong file .env hoặc trực tiếp trong môi trường)
# load_dotenv() # Uncomment dòng này nếu bạn sử dụng file .env

try:
    llm: Optional[ChatOpenAI] = ChatOpenAI(model="gpt-4o-mini", temperature=0.7)
    print(f"Mô hình ngôn ngữ đã khởi tạo: {llm.model}")
except Exception as e:
    print(f"Lỗi khởi tạo mô hình ngôn ngữ: {e}")
    llm = None

# --- Định nghĩa các Chuỗi Độc lập ---
# Ba chuỗi này đại diện cho các tác vụ riêng biệt có thể được thực thi song song.

# Chuỗi 1: Tóm tắt chủ đề
summarize_chain: Runnable = (
    ChatPromptTemplate.from_messages([
        ("system", "Tóm tắt chủ đề sau một cách súc tích:"),
        ("user", "{topic}")
    ])
    | llm
    | StrOutputParser()
)

# Chuỗi 2: Tạo câu hỏi
questions_chain: Runnable = (
    ChatPromptTemplate.from_messages([
        ("system", "Tạo ba câu hỏi thú vị về chủ đề sau:"),
        ("user", "{topic}")
    ])
    | llm
    | StrOutputParser()
)

# Chuỗi 3: Trích xuất thuật ngữ khóa
terms_chain: Runnable = (
    ChatPromptTemplate.from_messages([
        ("system", "Xác định 5-10 thuật ngữ khóa từ chủ đề sau, cách nhau bằng dấu phẩy:"),
        ("user", "{topic}")
    ])
    | llm
    | StrOutputParser()
)

# --- Xây dựng Chuỗi Song song + Tổng hợp ---
# 1. Định nghĩa khối các tác vụ sẽ chạy song song. Kết quả của chúng,
# cùng với chủ đề gốc, sẽ được đưa vào bước tiếp theo.
map_chain = RunnableParallel({
    "summary": summarize_chain,
    "questions": questions_chain,
    "key_terms": terms_chain,
    "topic": RunnablePassthrough() # Truyền chủ đề gốc
})

# 2. Định nghĩa prompt tổng hợp cuối cùng sẽ kết hợp các kết quả song song.
synthesis_prompt = ChatPromptTemplate.from_messages([
    ("system", """Dựa trên thông tin sau:
Tóm tắt: {summary}
Các câu hỏi liên quan: {questions}
Các thuật ngữ khóa: {key_terms}
Tổng hợp một câu trả lời toàn diện.""",),
    ("user", "Chủ đề gốc: {topic}")
])

# 3. Xây dựng chuỗi đầy đủ bằng cách nối các kết quả song song trực tiếp
# vào prompt tổng hợp, sau đó là LLM và trình phân tích cú pháp đầu ra.
full_parallel_chain = map_chain | synthesis_prompt | llm | StrOutputParser()

# --- Chạy Chuỗi ---
async def run_parallel_example(topic: str) -> None:
    """
    Thực thi bất đồng bộ chuỗi xử lý song song với một chủ đề cụ thể
    và in kết quả tổng hợp.
    """
    if not llm:
        print("LLM chưa được khởi tạo. Không thể chạy ví dụ.")
        return

    print(f"\n--- Chạy Ví dụ Parallel LangChain cho Chủ đề: '{topic}' ---")
    try:
        # Đầu vào của `ainvoke` là chuỗi 'topic' duy nhất,
        # sau đó được truyền đến mỗi runnable trong `map_chain`.
        response = await full_parallel_chain.ainvoke(topic)
        print("\n--- Phản hồi Cuối cùng ---")
        print(response)
    except Exception as e:
        print(f"\n Một lỗi đã xảy ra trong quá trình thực thi chuỗi: {e}")

if __name__ == "__main__":
    test_topic = "Lịch sử khám phá vũ trụ"
    # Trong Python 3.7+, asyncio.run là cách chuẩn để chạy một hàm bất đồng bộ.
    asyncio.run(run_parallel_example(test_topic))
```

Mã Python này triển khai một ứng dụng LangChain được thiết kế để xử lý một chủ đề một cách hiệu quả bằng cách tận dụng việc thực thi song song. Lưu ý rằng `asyncio` cung cấp tính đồng thời (concurrency), không phải tính song song thực sự (parallelism). Nó đạt được điều này trên một luồng duy nhất bằng cách sử dụng một vòng lặp sự kiện (event loop) chuyển đổi thông minh giữa các tác vụ khi một tác vụ nhàn rỗi (ví dụ: chờ yêu cầu mạng). Điều này tạo ra hiệu ứng nhiều tác vụ đang tiến triển cùng một lúc, nhưng bản thân mã vẫn được thực thi bởi chỉ một luồng, bị giới hạn bởi Global Interpreter Lock (GIL) của Python.

Về bản chất, mã này thiết lập một quy trình làm việc trong đó nhiều lệnh gọi LLM (để tóm tắt, đặt câu hỏi và tìm thuật ngữ) xảy ra cùng một lúc cho một chủ đề nhất định và kết quả của chúng sau đó được kết hợp bởi một lệnh gọi LLM cuối cùng. Điều này thể hiện ý tưởng cốt lõi của việc song song hóa trong một quy trình làm việc agentic sử dụng LangChain.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Nhiều quy trình làm việc agentic liên quan đến nhiều tác vụ phụ phải được hoàn thành để đạt được mục tiêu cuối cùng. Một quy trình thực thi tuần tự hoàn toàn, trong đó mỗi tác vụ chờ tác vụ trước đó hoàn thành, thường không hiệu quả và chậm. Độ trễ này trở thành một nút thắt cổ chai đáng kể khi các tác vụ phụ thuộc vào các hoạt động I/O bên ngoài, chẳng hạn như gọi các API khác nhau hoặc truy vấn nhiều cơ sở dữ liệu.
*   **Tại sao:** Mẫu Parallelization cung cấp một giải pháp tiêu chuẩn hóa bằng cách cho phép thực thi đồng thời các tác vụ độc lập. Nó hoạt động bằng cách xác định các thành phần của một quy trình làm việc, như việc sử dụng công cụ hoặc lệnh gọi LLM, không phụ thuộc vào các đầu ra tức thời của nhau. Các framework agentic như LangChain và Google ADK cung cấp các cấu trúc tích hợp để xác định và quản lý các hoạt động đồng thời này.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này khi một quy trình làm việc chứa nhiều hoạt động độc lập có thể chạy đồng thời, chẳng hạn như tìm nạp dữ liệu từ nhiều API, xử lý các phần dữ liệu khác nhau hoặc tạo nhiều phần nội dung để tổng hợp sau này.

## Những Điểm Chính (Key Takeaways)

*   Parallelization là một mẫu để thực thi các tác vụ độc lập đồng thời nhằm cải thiện hiệu quả.
*   Nó đặc biệt hữu ích khi các tác vụ liên quan đến việc chờ tài nguyên bên ngoài, chẳng hạn như các lệnh gọi API.
*   Việc áp dụng kiến trúc đồng thời hoặc song song mang lại sự phức tạp và chi phí đáng kể, ảnh hưởng đến các giai đoạn phát triển chính như thiết kế, gỡ lỗi và ghi nhật ký hệ thống.
*   Các framework như LangChain và Google ADK cung cấp hỗ trợ tích hợp để xác định và quản lý việc thực thi song song.
*   Trong LangChain Expression Language (LCEL), `RunnableParallel` là một cấu trúc khóa để chạy nhiều runnables song song.

## Kết luận

Mẫu Parallelization là một phương pháp tối ưu hóa các quy trình làm việc tính toán bằng cách thực thi đồng thời các tác vụ phụ độc lập. Cách tiếp cận này giảm độ trễ tổng thể, đặc biệt trong các hoạt động phức tạp liên quan đến nhiều suy luận mô hình hoặc các lệnh gọi đến các dịch vụ bên ngoài.

Các framework cung cấp các cơ chế riêng biệt để triển khai mẫu này. Trong LangChain, các cấu trúc như `RunnableParallel` được sử dụng để xác định và thực thi rõ ràng nhiều chuỗi xử lý cùng lúc. Ngược lại, các framework như Google Agent Developer Kit (ADK) có thể đạt được tính song song thông qua ủy quyền đa tác nhân, nơi một mô hình điều phối viên chính gán các tác vụ phụ khác nhau cho các tác nhân chuyên biệt có thể hoạt động đồng thời.

Bằng cách tích hợp xử lý song song với các luồng điều khiển tuần tự (chaining) và có điều kiện (routing), có thể xây dựng các hệ thống tính toán tinh vi, hiệu suất cao có khả năng quản lý các tác vụ đa dạng và phức tạp một cách hiệu quả.

## Tài liệu tham khảo
1.  LangChain Expression Language (LCEL) Documentation (Parallelism): https://python.langchain.com/docs/concepts/lcel/
2.  Google Agent Developer Kit (ADK) Documentation (Multi-Agent Systems): https://google.github.io/adk-docs/agents/multi-agents/
3.  Python asyncio Documentation: https://docs.python.org/3/library/asyncio.html