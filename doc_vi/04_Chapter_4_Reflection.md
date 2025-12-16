# Chương 4: Reflection (Phản Chiếu / Tự Phê Bình)

## Tổng quan về Mẫu thiết kế Reflection

Trong các chương trước, chúng ta đã khám phá các mẫu Agentic cơ bản: Chaining cho việc thực thi tuần tự, Routing cho việc lựa chọn đường dẫn động và Parallelization cho việc thực thi tác vụ đồng thời. Các mẫu này cho phép các tác nhân thực hiện các tác vụ phức tạp một cách hiệu quả và linh hoạt hơn. Tuy nhiên, ngay cả với các quy trình làm việc tinh vi, đầu ra hoặc kế hoạch ban đầu của một tác nhân có thể không tối ưu, chính xác hoặc hoàn chỉnh. Đây là lúc mẫu **Reflection (Phản Chiếu)** phát huy tác dụng.

Mẫu Reflection liên quan đến việc một tác nhân tự đánh giá công việc, đầu ra hoặc trạng thái nội bộ của mình và sử dụng đánh giá đó để cải thiện hiệu suất hoặc tinh chỉnh phản hồi của nó. Đây là một hình thức **tự sửa lỗi (self-correction)** hoặc **tự cải thiện (self-improvement)**, cho phép tác nhân liên tục tinh chỉnh đầu ra hoặc điều chỉnh cách tiếp cận của mình dựa trên phản hồi, phê bình nội bộ hoặc so sánh với các tiêu chí mong muốn. Reflection đôi khi có thể được thực hiện bởi một tác nhân riêng biệt có vai trò cụ thể là phân tích đầu ra của tác nhân ban đầu.

Không giống như một chuỗi tuần tự đơn giản nơi đầu ra được chuyển trực tiếp đến bước tiếp theo, hoặc định tuyến chọn một đường dẫn, Reflection giới thiệu một **vòng lặp phản hồi (feedback loop)**. Tác nhân không chỉ tạo ra một đầu ra; nó sau đó kiểm tra đầu ra đó (hoặc quy trình đã tạo ra nó), xác định các vấn đề tiềm ẩn hoặc các lĩnh vực cần cải thiện, và sử dụng những hiểu biết đó để tạo ra một phiên bản tốt hơn hoặc sửa đổi các hành động trong tương lai của nó.

Quy trình điển hình bao gồm:
1.  **Thực thi (Execution):** Tác nhân thực hiện một tác vụ hoặc tạo ra một đầu ra ban đầu.
2.  **Đánh giá/Phê bình (Evaluation/Critique):** Tác nhân (thường sử dụng một lệnh gọi LLM khác hoặc một tập hợp các quy tắc) phân tích kết quả từ bước trước đó. Đánh giá này có thể kiểm tra tính chính xác về mặt thực tế, tính mạch lạc, phong cách, tính đầy đủ, tuân thủ các hướng dẫn hoặc các tiêu chí liên quan khác.
3.  **Phản chiếu/Tinh chỉnh (Reflection/Refinement):** Dựa trên lời phê bình, tác nhân xác định cách cải thiện. Điều này có thể liên quan đến việc tạo ra một đầu ra đã tinh chỉnh, điều chỉnh các tham số cho bước tiếp theo hoặc thậm chí sửa đổi kế hoạch tổng thể.
4.  **Lặp lại (Iterative):** Đầu ra đã tinh chỉnh hoặc cách tiếp cận đã điều chỉnh sau đó có thể được thực thi, và quá trình phản chiếu có thể lặp lại cho đến khi đạt được kết quả thỏa đáng hoặc một điều kiện dừng được đáp ứng.

Một triển khai quan trọng và hiệu quả cao của mẫu Reflection là tách quy trình thành hai vai trò logic riêng biệt: **Người tạo (Producer)** và **Người phê bình (Critic)**. Điều này thường được gọi là mô hình "Generator-Critic" hoặc "Producer-Reviewer". Mặc dù một tác nhân đơn lẻ có thể tự phản chiếu, việc sử dụng hai tác nhân chuyên biệt (hoặc hai lệnh gọi LLM riêng biệt với các prompt hệ thống riêng biệt) thường mang lại kết quả mạnh mẽ và khách quan hơn.

1.  **Tác nhân Người tạo (The Producer Agent):** Trách nhiệm chính của tác nhân này là thực hiện việc thực thi ban đầu của tác vụ. Nó tập trung hoàn toàn vào việc tạo ra nội dung, cho dù đó là viết mã, soạn thảo một bài đăng trên blog, hoặc tạo một kế hoạch. Nó nhận prompt ban đầu và tạo ra phiên bản đầu tiên của đầu ra.
2.  **Tác nhân Người phê bình (The Critic Agent):** Mục đích duy nhất của tác nhân này là đánh giá đầu ra được tạo bởi Người tạo. Nó được cung cấp một bộ hướng dẫn khác, thường là một nhân cách riêng biệt (ví dụ: "Bạn là một kỹ sư phần mềm cao cấp," "Bạn là một người kiểm tra lỗi tỉ mỉ"). Các hướng dẫn của Critic hướng dẫn nó phân tích công việc của Producer dựa trên các tiêu chí cụ thể, chẳng hạn như độ chính xác về mặt thực tế, chất lượng mã, yêu cầu về phong cách hoặc tính đầy đủ. Nó được thiết kế để tìm lỗi, đề xuất cải tiến và cung cấp phản hồi có cấu trúc.

Sự tách biệt các mối quan tâm này là mạnh mẽ vì nó ngăn chặn "thiên vị nhận thức" của một tác nhân khi xem xét công việc của chính nó. Tác nhân Critic tiếp cận đầu ra với một góc nhìn mới mẻ, chuyên tâm hoàn toàn vào việc tìm lỗi và các lĩnh vực cần cải thiện. Phản hồi từ Critic sau đó được chuyển trở lại cho tác nhân Producer, tác nhân này sử dụng nó làm hướng dẫn để tạo ra một phiên bản mới, tinh chỉnh hơn của đầu ra. Các ví dụ mã LangChain và ADK đều triển khai mô hình hai tác nhân này: ví dụ LangChain sử dụng một "reflector_prompt" cụ thể để tạo ra nhân cách phê bình, trong khi ví dụ ADK định nghĩa rõ ràng một tác nhân producer và một tác nhân reviewer.

Việc triển khai Reflection thường yêu cầu cấu trúc quy trình làm việc của tác nhân để bao gồm các vòng lặp phản hồi này. Điều này có thể đạt được thông qua các vòng lặp lặp đi lặp lại trong mã, hoặc sử dụng các framework hỗ trợ quản lý trạng thái và chuyển đổi có điều kiện dựa trên kết quả đánh giá.

Giao điểm giữa reflection với thiết lập mục tiêu và giám sát (xem Chương 11) đáng chú ý. Một mục tiêu cung cấp thước đo cuối cùng để tác nhân tự đánh giá, trong khi giám sát theo dõi tiến trình của nó. Trong một số trường hợp thực tế, Reflection có thể hoạt động như một công cụ điều chỉnh, sử dụng phản hồi được giám sát để phân tích các sai lệch và điều chỉnh chiến lược của nó. Sự kết hợp này biến tác nhân từ một người thực thi thụ động thành một hệ thống có mục đích, thích ứng để đạt được mục tiêu của mình.

Hơn nữa, hiệu quả của mẫu Reflection được tăng cường đáng kể khi LLM giữ một bộ nhớ về cuộc hội thoại (xem Chương 8). Lịch sử hội thoại này cung cấp ngữ cảnh quan trọng cho giai đoạn đánh giá, cho phép tác nhân đánh giá đầu ra của nó không chỉ riêng lẻ mà còn dựa trên các tương tác trước đó, phản hồi của người dùng và các mục tiêu đang phát triển.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu Reflection có giá trị trong các trường hợp mà chất lượng, độ chính xác của đầu ra hoặc sự tuân thủ các ràng buộc phức tạp là rất quan trọng:

1.  **Viết sáng tạo và Tạo nội dung:** Tinh chỉnh văn bản, câu chuyện, bài thơ hoặc nội dung marketing được tạo ra.
    *   **Trường hợp sử dụng:** Một tác nhân viết bài đăng trên blog.
    *   **Phản chiếu:** Tạo bản nháp, phê bình nó về luồng, giọng điệu và sự rõ ràng, sau đó viết lại dựa trên lời phê bình. Lặp lại cho đến khi bài đăng đáp ứng các tiêu chuẩn chất lượng.
    *   **Lợi ích:** Tạo ra nội dung trau chuốt và hiệu quả hơn.
2.  **Tạo và Gỡ lỗi Mã:** Viết mã, xác định lỗi và sửa lỗi.
    *   **Trường hợp sử dụng:** Một tác nhân viết một hàm Python.
    *   **Phản chiếu:** Viết mã ban đầu, chạy thử nghiệm hoặc phân tích tĩnh, xác định lỗi hoặc sự kém hiệu quả, sau đó sửa đổi mã dựa trên các phát hiện.
    *   **Lợi ích:** Tạo ra mã mạnh mẽ và chức năng hơn.
3.  **Giải quyết vấn đề phức tạp:** Đánh giá các bước trung gian hoặc các giải pháp được đề xuất trong các tác vụ suy luận đa bước.
    *   **Trường hợp sử dụng:** Một tác nhân giải một câu đố logic.
    *   **Phản chiếu:** Đề xuất một bước, đánh giá xem nó có dẫn đến giải pháp hoặc gây ra mâu thuẫn không, quay lại hoặc chọn một bước khác nếu cần.
    *   **Lợi ích:** Cải thiện khả năng của tác nhân trong việc điều hướng không gian vấn đề phức tạp.
4.  **Tóm tắt và Tổng hợp Thông tin:** Tinh chỉnh các bản tóm tắt để đảm bảo độ chính xác, đầy đủ và súc tích.
    *   **Trường hợp sử dụng:** Một tác nhân tóm tắt một tài liệu dài.
    *   **Phản chiếu:** Tạo bản tóm tắt ban đầu, so sánh nó với các điểm chính trong tài liệu gốc, tinh chỉnh bản tóm tắt để bao gồm thông tin còn thiếu hoặc cải thiện độ chính xác.
    *   **Lợi ích:** Tạo ra các bản tóm tắt chính xác và toàn diện hơn.
5.  **Lập kế hoạch và Chiến lược:** Đánh giá một kế hoạch được đề xuất và xác định các lỗi hoặc cải tiến tiềm năng.
    *   **Trường hợp sử dụng:** Một tác nhân lập kế hoạch một loạt các hành động để đạt được mục tiêu.
    *   **Phản chiếu:** Tạo một kế hoạch, mô phỏng việc thực thi nó hoặc đánh giá tính khả thi của nó dựa trên các ràng buộc, sửa đổi kế hoạch dựa trên đánh giá.
    *   **Lợi ích:** Phát triển các kế hoạch hiệu quả và thực tế hơn.
6.  **Tác nhân Hội thoại:** Xem xét các lượt hội thoại trước đó để duy trì ngữ cảnh, sửa lỗi hiểu lầm hoặc cải thiện chất lượng phản hồi.
    *   **Trường hợp sử dụng:** Một chatbot hỗ trợ khách hàng.
    *   **Phản chiếu:** Sau khi người dùng phản hồi, xem xét lịch sử hội thoại và tin nhắn được tạo gần đây nhất để đảm bảo tính mạch lạc và xử lý chính xác đầu vào mới nhất của người dùng.
    *   **Lợi ích:** Dẫn đến các cuộc hội thoại tự nhiên và hiệu quả hơn.

Reflection bổ sung một lớp nhận thức cấp cao (meta-cognition) cho các hệ thống agentic, cho phép chúng học hỏi từ các đầu ra và quy trình của chính mình, dẫn đến kết quả thông minh, đáng tin cậy và chất lượng cao hơn.

## Ví dụ Code Thực hành (LangChain)

Việc triển khai một quy trình reflection lặp lại, hoàn chỉnh yêu cầu các cơ chế quản lý trạng thái và thực thi theo chu kỳ. Mặc dù những điều này được xử lý nguyên bản trong các framework dựa trên đồ thị như LangGraph hoặc thông qua mã thủ tục tùy chỉnh, nguyên tắc cơ bản của một chu kỳ reflection đơn lẻ có thể được minh họa một cách hiệu quả bằng cách sử dụng cú pháp kết hợp của LCEL (LangChain Expression Language).

Ví dụ này triển khai một vòng lặp reflection bằng cách sử dụng thư viện LangChain và mô hình GPT-4o của OpenAI để tạo và tinh chỉnh lặp lại một hàm Python tính giai thừa của một số. Quy trình bắt đầu với một prompt tác vụ, tạo mã ban đầu, và sau đó liên tục phản chiếu mã dựa trên các phê bình từ một vai trò kỹ sư phần mềm cao cấp được mô phỏng, tinh chỉnh mã trong mỗi lần lặp cho đến khi giai đoạn phê bình xác định rằng mã hoàn hảo hoặc đạt đến số lần lặp tối đa. Cuối cùng, nó in mã đã tinh chỉnh.

```python
import os
from dotenv import load_dotenv
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.messages import SystemMessage, HumanMessage

# --- Cấu hình ---
# Tải biến môi trường từ file .env (cho OPENAI_API_KEY)
load_dotenv()

# Kiểm tra xem API key đã được thiết lập chưa
if not os.getenv("OPENAI_API_KEY"):
    raise ValueError("OPENAI_API_KEY không tìm thấy trong file .env. Vui lòng thêm nó.")

# Khởi tạo Chat LLM. Chúng ta sử dụng gpt-4o để suy luận tốt hơn.
# Nhiệt độ thấp hơn được sử dụng để đầu ra xác định hơn.
llm = ChatOpenAI(model="gpt-4o", temperature=0.1)

def run_reflection_loop():
    """
    Minh họa một vòng lặp reflection AI đa bước để cải thiện dần dần một hàm Python.
    """
    # --- Tác vụ Cốt lõi ---
    task_prompt = """
    Nhiệm vụ của bạn là tạo một hàm Python có tên `calculate_factorial`.
    Hàm này phải thực hiện các điều sau:
    1. Chấp nhận một số nguyên `n` làm đầu vào duy nhất.
    2. Tính giai thừa của nó (n!).
    3. Bao gồm một docstring rõ ràng giải thích chức năng của hàm.
    4. Xử lý các trường hợp đặc biệt: Giai thừa của 0 là 1.
    5. Xử lý đầu vào không hợp lệ: Nâng ValueError nếu đầu vào là số âm.
    """

    # --- Vòng lặp Reflection ---
    max_iterations = 3
    current_code = ""

    # Chúng ta sẽ xây dựng lịch sử hội thoại để cung cấp ngữ cảnh trong mỗi bước.
    message_history = [HumanMessage(content=task_prompt)]

    for i in range(max_iterations):
        print("\n" + "=" * 25 + f" VÒNG LẶP REFLECTION: LẶP LẠI {i + 1} " + "=" * 25)

        # --- 1. GIAI ĐOẠN TẠO / TINH CHỈNH ---
        # Trong lần lặp đầu tiên, nó tạo ra. Trong các lần lặp tiếp theo, nó tinh chỉnh.
        if i == 0:
            print("\n>>> GIAI ĐOẠN 1: TẠO mã ban đầu...")
            # Tin nhắn đầu tiên chỉ là prompt tác vụ.
            response = llm.invoke(message_history)
            current_code = response.content
        else:
            print("\n>>> GIAI ĐOẠN 1: TINH CHỈNH mã dựa trên phê bình trước đó...")
            # Lịch sử tin nhắn hiện chứa tác vụ, mã cuối cùng và phê bình cuối cùng.
            # Chúng ta hướng dẫn mô hình áp dụng các phê bình.
            message_history.append(HumanMessage(content="Vui lòng tinh chỉnh mã bằng cách sử dụng các phê bình được cung cấp."))
            response = llm.invoke(message_history)
            current_code = response.content

        print(f"\n--- Mã được tạo (phiên bản {i + 1}) ---\n" + current_code)
        message_history.append(HumanMessage(content=current_code)) # Thêm mã được tạo vào lịch sử

        # --- 2. GIAI ĐOẠN PHẢN CHIẾU ---
        print("\n>>> GIAI ĐOẠN 2: PHẢN CHIẾU trên mã đã tạo...")
        # Tạo một prompt cụ thể cho tác nhân phản chiếu.
        # Điều này yêu cầu mô hình hoạt động như một người đánh giá mã cấp cao.
        reflector_prompt = [
            SystemMessage(content="""
            Bạn là một kỹ sư phần mềm cấp cao và là một chuyên gia về Python.
            Vai trò của bạn là thực hiện đánh giá mã tỉ mỉ.
            Đánh giá phê bình mã Python được cung cấp dựa trên các yêu cầu tác vụ gốc.
            Tìm kiếm lỗi, vấn đề về phong cách, trường hợp đặc biệt bị thiếu và các lĩnh vực cần cải thiện.
            Nếu mã hoàn hảo và đáp ứng tất cả các yêu cầu, hãy trả lời bằng cụm từ 'CODE_IS_PERFECT' duy nhất.
            Nếu không, hãy cung cấp danh sách các phê bình có dấu đầu dòng.
            """),
            HumanMessage(content=f"Tác vụ gốc:\n{task_prompt}\n\nMã cần xem xét:\n{current_code}")
        ]

        critique_response = llm.invoke(reflector_prompt)
        critique = critique_response.content

        # --- 3. ĐIỀU KIỆN DỪNG ---
        if "CODE_IS_PERFECT" in critique:
            print("\n--- Phê bình ---\nKhông tìm thấy phê bình nào thêm. Mã đạt yêu cầu.")
            break
        print("\n--- Phê bình ---\n" + critique)
        message_history.append(HumanMessage(content=f"Phê bình mã trước đó:\n{critique}")) # Thêm phê bình vào lịch sử để tinh chỉnh tiếp theo

    print("\n" + "=" * 30 + " KẾT QUẢ CUỐI CÙNG " + "=" * 30)
    print("\nMã đã tinh chỉnh cuối cùng sau quá trình reflection:\n")
    print(current_code)

if __name__ == "__main__":
    run_reflection_loop()
```

Mã này bắt đầu bằng cách thiết lập môi trường, tải API key và khởi tạo một mô hình ngôn ngữ mạnh mẽ như GPT-4o với nhiệt độ thấp để có đầu ra tập trung. Nhiệm vụ cốt lõi được định nghĩa bởi một prompt yêu cầu tạo một hàm Python để tính giai thừa của một số, bao gồm các yêu cầu cụ thể về docstring, trường hợp đặc biệt (giai thừa của 0) và xử lý lỗi cho đầu vào âm. Hàm `run_reflection_loop` điều phối quy trình tinh chỉnh lặp đi lặp lại. Trong vòng lặp, ở lần lặp đầu tiên, mô hình ngôn ngữ tạo mã ban đầu dựa trên prompt tác vụ. Trong các lần lặp tiếp theo, nó tinh chỉnh mã dựa trên các phê bình từ bước trước đó. Một vai trò "reflector" riêng biệt, cũng do mô hình ngôn ngữ đảm nhận nhưng với một prompt hệ thống khác, đóng vai trò là một kỹ sư phần mềm cao cấp để phê bình mã được tạo dựa trên các yêu cầu tác vụ gốc. Lời phê bình này được cung cấp dưới dạng danh sách các vấn đề có dấu đầu dòng hoặc cụm từ 'CODE_IS_PERFECT' nếu không tìm thấy vấn đề nào. Vòng lặp tiếp tục cho đến khi lời phê bình chỉ ra rằng mã hoàn hảo hoặc đạt đến số lần lặp tối đa. Lịch sử hội thoại được duy trì và truyền đến mô hình ngôn ngữ trong mỗi bước để cung cấp ngữ cảnh cho cả giai đoạn tạo/tinh chỉnh và reflection.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Đầu ra ban đầu của tác nhân thường không tối ưu, mắc lỗi, không đầy đủ hoặc không đáp ứng các yêu cầu phức tạp. Các quy trình làm việc Agentic cơ bản thiếu một quy trình tích hợp để tác nhân nhận biết và tự sửa lỗi. Điều này được giải quyết bằng cách để tác nhân tự đánh giá công việc của mình hoặc, mạnh mẽ hơn, bằng cách giới thiệu một tác nhân logic riêng biệt đóng vai trò là nhà phê bình, ngăn chặn phản hồi ban đầu trở thành phản hồi cuối cùng bất kể chất lượng.
*   **Tại sao:** Mẫu Reflection đưa ra một giải pháp bằng cách giới thiệu một cơ chế tự sửa lỗi và tinh chỉnh. Nó thiết lập một vòng lặp phản hồi nơi tác nhân "người tạo" (producer) tạo ra một đầu ra, và sau đó tác nhân "người phê bình" (critic) (hoặc chính người tạo) đánh giá nó dựa trên các tiêu chí được xác định trước. Lời phê bình này sau đó được sử dụng để tạo ra một phiên bản cải tiến. Quy trình lặp lại gồm tạo, đánh giá và tinh chỉnh này liên tục nâng cao chất lượng của kết quả cuối cùng, dẫn đến các kết quả chính xác, mạch lạc và đáng tin cậy hơn.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu Reflection khi chất lượng, độ chính xác và chi tiết của đầu ra cuối cùng quan trọng hơn tốc độ và chi phí. Nó đặc biệt hiệu quả cho các tác vụ như tạo nội dung dài, trau chuốt, viết và gỡ lỗi mã, và tạo các kế hoạch chi tiết. Sử dụng một tác nhân phê bình riêng biệt khi các tác vụ yêu cầu tính khách quan cao hoặc đánh giá chuyên biệt mà một tác nhân người tạo tổng quát có thể bỏ lỡ.

## Những Điểm Chính (Key Takeaways)

*   Ưu điểm chính của mẫu Reflection là khả năng tự sửa lỗi và tinh chỉnh đầu ra lặp đi lặp lại, dẫn đến chất lượng, độ chính xác cao hơn đáng kể và tuân thủ các hướng dẫn phức tạp.
*   Nó liên quan đến một vòng lặp phản hồi gồm thực thi, đánh giá/phê bình và tinh chỉnh. Reflection là điều cần thiết cho các tác vụ đòi hỏi đầu ra chất lượng cao, chính xác hoặc sắc thái.
*   Một triển khai mạnh mẽ là mô hình Producer-Critic, nơi một tác nhân riêng biệt (hoặc vai trò được nhắc) đánh giá đầu ra ban đầu. Sự tách biệt các mối quan tâm này nâng cao tính khách quan và cho phép phản hồi chuyên biệt, có cấu trúc hơn.

## Kết luận

Mẫu reflection cung cấp một cơ chế quan trọng để tự sửa lỗi trong quy trình làm việc của một tác nhân, cho phép cải thiện lặp đi lặp lại vượt ra ngoài việc thực thi một bước. Điều này đạt được bằng cách tạo ra một vòng lặp nơi hệ thống tạo ra một đầu ra, đánh giá nó dựa trên các tiêu chí cụ thể và sau đó sử dụng đánh giá đó để tạo ra một kết quả đã tinh chỉnh. Sự đánh giá này có thể được thực hiện bởi chính tác nhân (tự phản chiếu) hoặc, thường hiệu quả hơn, bởi một tác nhân phê bình riêng biệt, đại diện cho một lựa chọn kiến trúc chính trong mẫu.

Trong khi một quy trình reflection đa bước, hoàn toàn tự động đòi hỏi một kiến trúc mạnh mẽ để quản lý trạng thái, nguyên tắc cốt lõi của nó được thể hiện một cách hiệu quả trong một chu kỳ tạo-phê bình-tinh chỉnh đơn lẻ. Là một cấu trúc kiểm soát, reflection có thể được tích hợp với các mẫu nền tảng khác để xây dựng các hệ thống agentic mạnh mẽ và phức tạp hơn về chức năng.

## Tài liệu tham khảo
1.  Training Language Models to Self-Correct via Reinforcement Learning, https://arxiv.org/abs/2409.12917
2.  LangChain Expression Language (LCEL) Documentation: https://python.langchain.com/docs/introduction/
3.  LangGraph Documentation: https://www.langchain.com/langgraph
4.  Google Agent Developer Kit (ADK) Documentation (Multi-Agent Systems): https://google.github.io/adk-docs/agents/multi-agents/