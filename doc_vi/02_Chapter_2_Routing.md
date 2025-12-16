# Chương 2: Routing (Định Tuyến)

## Tổng quan về Mẫu thiết kế Routing

Trong khi xử lý tuần tự thông qua Prompt Chaining là một kỹ thuật nền tảng để thực thi các quy trình làm việc xác định, tuyến tính với các mô hình ngôn ngữ, thì khả năng áp dụng của nó bị hạn chế trong các tình huống đòi hỏi phản ứng thích ứng. Các hệ thống Agentic trong thế giới thực thường phải phân xử giữa nhiều hành động tiềm năng dựa trên các yếu tố ngẫu nhiên, chẳng hạn như trạng thái của môi trường, đầu vào của người dùng hoặc kết quả của một hoạt động trước đó. Khả năng ra quyết định năng động này, chi phối luồng kiểm soát đến các chức năng, công cụ hoặc quy trình phụ chuyên biệt khác nhau, đạt được thông qua một cơ chế được gọi là **Routing (Định tuyến)**.

Routing giới thiệu logic điều kiện vào khuôn khổ hoạt động của một tác nhân, cho phép chuyển đổi từ một đường dẫn thực thi cố định sang một mô hình nơi tác nhân đánh giá linh hoạt các tiêu chí cụ thể để chọn từ một tập hợp các hành động tiếp theo có thể. Điều này cho phép hành vi hệ thống linh hoạt và nhận thức ngữ cảnh hơn.

Ví dụ, một tác nhân được thiết kế cho các yêu cầu của khách hàng, khi được trang bị chức năng định tuyến, trước tiên có thể phân loại một truy vấn đến để xác định ý định của người dùng. Dựa trên phân loại này, sau đó nó có thể hướng truy vấn đến một tác nhân chuyên biệt để trả lời câu hỏi trực tiếp, một công cụ truy xuất cơ sở dữ liệu để lấy thông tin tài khoản, hoặc một quy trình leo thang cho các vấn đề phức tạp, thay vì mặc định theo một lộ trình phản hồi đơn lẻ được xác định trước. Do đó, một tác nhân tinh vi sử dụng routing có thể:

1.  Phân tích truy vấn của người dùng.
2.  **Định tuyến** truy vấn dựa trên *ý định* của nó:
    *   Nếu ý định là "kiểm tra trạng thái đơn hàng", hãy định tuyến đến một tác nhân phụ hoặc chuỗi công cụ tương tác với cơ sở dữ liệu đơn hàng.
    *   Nếu ý định là "thông tin sản phẩm", hãy định tuyến đến một tác nhân phụ hoặc chuỗi tìm kiếm danh mục sản phẩm.
    *   Nếu ý định là "hỗ trợ kỹ thuật", hãy định tuyến đến một chuỗi khác truy cập hướng dẫn khắc phục sự cố hoặc leo thang cho con người.
    *   Nếu ý định không rõ ràng, hãy định tuyến đến một tác nhân phụ làm rõ hoặc chuỗi prompt để hỏi lại người dùng.

## Các cơ chế Routing

Thành phần cốt lõi của mẫu Routing là một cơ chế thực hiện việc đánh giá và chỉ đạo luồng. Cơ chế này có thể được triển khai theo nhiều cách:

1.  **LLM-based Routing (Định tuyến dựa trên LLM):** Bản thân mô hình ngôn ngữ có thể được nhắc (prompted) để phân tích đầu vào và đưa ra một mã định danh hoặc hướng dẫn cụ thể cho biết bước hoặc đích đến tiếp theo. Ví dụ: một prompt có thể yêu cầu LLM "Phân tích truy vấn của người dùng sau và chỉ đưa ra danh mục: 'Trạng thái đơn hàng', 'Thông tin sản phẩm', 'Hỗ trợ kỹ thuật', hoặc 'Khác'.". Hệ thống agentic sau đó đọc đầu ra này và chỉ đạo quy trình làm việc tương ứng.
2.  **Embedding-based Routing (Định tuyến dựa trên nhúng - Semantic Router):** Truy vấn đầu vào có thể được chuyển đổi thành một vector embedding. Embedding này sau đó được so sánh với các embedding đại diện cho các tuyến (routes) hoặc khả năng khác nhau. Truy vấn được định tuyến đến tuyến có embedding tương tự nhất. Điều này hữu ích cho định tuyến ngữ nghĩa (semantic routing), nơi quyết định dựa trên ý nghĩa của đầu vào thay vì chỉ là từ khóa.
3.  **Rule-based Routing (Định tuyến dựa trên quy tắc):** Điều này liên quan đến việc sử dụng các quy tắc hoặc logic được xác định trước (ví dụ: câu lệnh if-else, switch cases) dựa trên từ khóa, mẫu (regex) hoặc dữ liệu có cấu trúc được trích xuất từ đầu vào. Cách này có thể nhanh hơn và xác định hơn (deterministic) so với định tuyến dựa trên LLM, nhưng kém linh hoạt hơn trong việc xử lý các đầu vào sắc thái hoặc mới lạ.
4.  **Machine Learning Model-Based Routing (Định tuyến dựa trên mô hình ML):** Nó sử dụng một mô hình phân biệt (discriminative model), chẳng hạn như một bộ phân loại (classifier), đã được huấn luyện cụ thể trên một kho dữ liệu được gán nhãn nhỏ để thực hiện nhiệm vụ định tuyến. Mặc dù chia sẻ những điểm tương đồng về mặt khái niệm với các phương pháp dựa trên embedding, đặc điểm chính của nó là quy trình tinh chỉnh có giám sát (supervised fine-tuning). Kỹ thuật này khác biệt với định tuyến dựa trên LLM vì thành phần ra quyết định không phải là một mô hình tạo sinh (generative model) thực thi prompt tại thời điểm suy luận. Thay vào đó, logic định tuyến được mã hóa trong các trọng số đã học của mô hình được tinh chỉnh.

Cơ chế định tuyến có thể được triển khai tại nhiều thời điểm trong chu kỳ hoạt động của một tác nhân. Chúng có thể được áp dụng ngay từ đầu để phân loại nhiệm vụ chính, tại các điểm trung gian trong chuỗi xử lý để xác định hành động tiếp theo, hoặc trong một quy trình con để chọn công cụ phù hợp nhất từ một tập hợp.

Việc triển khai routing cho phép hệ thống vượt ra ngoài xử lý tuần tự tất định. Nó tạo điều kiện cho sự phát triển của các luồng thực thi thích ứng hơn, có thể phản ứng linh hoạt và phù hợp với nhiều loại đầu vào và thay đổi trạng thái hơn.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu routing là một cơ chế kiểm soát quan trọng trong thiết kế các hệ thống agentic thích ứng, cho phép chúng thay đổi linh hoạt đường dẫn thực thi để phản ứng với các đầu vào biến đổi và trạng thái bên trong. Tiện ích của nó trải rộng trên nhiều miền bằng cách cung cấp một lớp logic điều kiện cần thiết.

*   **Tương tác Người-Máy (HCI):** Trong các trợ lý ảo hoặc gia sư AI, routing được sử dụng để diễn giải ý định của người dùng. Một phân tích ban đầu về truy vấn ngôn ngữ tự nhiên xác định hành động tiếp theo phù hợp nhất, cho dù đó là gọi một công cụ truy xuất thông tin cụ thể, leo thang cho người vận hành con người, hay chọn mô-đun tiếp theo trong chương trình giảng dạy dựa trên hiệu suất của người dùng. Điều này cho phép hệ thống vượt ra ngoài các luồng đối thoại tuyến tính và phản hồi theo ngữ cảnh.
*   **Quy trình Xử lý Dữ liệu và Tài liệu Tự động:** Routing đóng vai trò là chức năng phân loại và phân phối. Dữ liệu đến, chẳng hạn như email, phiếu hỗ trợ (support tickets), hoặc tải trọng API, được phân tích dựa trên nội dung, siêu dữ liệu hoặc định dạng. Hệ thống sau đó hướng từng mục đến quy trình làm việc tương ứng, chẳng hạn như quy trình nhập khách hàng tiềm năng, hàm chuyển đổi dữ liệu cụ thể cho các định dạng JSON hoặc CSV, hoặc đường dẫn leo thang vấn đề khẩn cấp.
*   **Hệ thống Phức tạp Đa Tác nhân (Multi-Agent):** Trong các hệ thống bao gồm nhiều công cụ hoặc tác nhân chuyên biệt, routing hoạt động như một bộ điều phối cấp cao (dispatcher). Một hệ thống nghiên cứu bao gồm các tác nhân riêng biệt để tìm kiếm, tóm tắt và phân tích thông tin sẽ sử dụng bộ định tuyến để giao nhiệm vụ cho tác nhân phù hợp nhất dựa trên mục tiêu hiện tại. Tương tự, một trợ lý mã hóa AI sử dụng routing để xác định ngôn ngữ lập trình và ý định của người dùng—để gỡ lỗi, giải thích hoặc dịch—trước khi chuyển đoạn mã đến công cụ chuyên biệt chính xác.

Cuối cùng, routing cung cấp khả năng phân xử logic cần thiết để tạo ra các hệ thống đa dạng về chức năng và nhận thức ngữ cảnh. Nó biến một tác nhân từ một người thực thi tĩnh các trình tự được xác định trước thành một hệ thống động có thể đưa ra quyết định về phương pháp hiệu quả nhất để hoàn thành nhiệm vụ trong các điều kiện thay đổi.

## Ví dụ Code Thực hành (LangChain)

Triển khai routing trong mã liên quan đến việc xác định các đường dẫn có thể có và logic quyết định đường dẫn nào sẽ đi. Các framework như LangChain và LangGraph cung cấp các thành phần và cấu trúc cụ thể cho việc này. Cấu trúc đồ thị dựa trên trạng thái của LangGraph đặc biệt trực quan để hình dung và triển khai logic định tuyến.

Mã này minh họa một hệ thống giống như tác nhân đơn giản sử dụng LangChain và Google Generative AI. Nó thiết lập một "điều phối viên" (coordinator) định tuyến các yêu cầu của người dùng đến các trình xử lý "tác nhân phụ" (sub-agent) mô phỏng khác nhau dựa trên ý định của yêu cầu (đặt chỗ, thông tin, hoặc không rõ ràng). Hệ thống sử dụng một mô hình ngôn ngữ để phân loại yêu cầu và sau đó ủy quyền cho hàm xử lý thích hợp, mô phỏng một mẫu ủy quyền cơ bản thường thấy trong các kiến trúc đa tác nhân.

```python
# Cài đặt các thư viện cần thiết:
# pip install langchain langgraph langchain-google-genai google-adk

from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough, RunnableBranch
import os

# --- Cấu hình ---
# Đảm bảo biến môi trường GOOGLE_API_KEY của bạn đã được thiết lập
# os.environ["GOOGLE_API_KEY"] = "YOUR_API_KEY"

try:
    llm = ChatGoogleGenerativeAI(model="gemini-2.0-flash", temperature=0)
    print(f"Mô hình ngôn ngữ đã được khởi tạo: {llm.model}")
except Exception as e:
    print(f"Lỗi khởi tạo mô hình ngôn ngữ: {e}")
    llm = None

# --- Định nghĩa các Trình xử lý Tác nhân phụ Mô phỏng ---

def booking_handler(request: str) -> str:
    """Mô phỏng Tác nhân Đặt chỗ xử lý yêu cầu."""
    print("\n--- ĐANG ỦY QUYỀN CHO TRÌNH XỬ LÝ ĐẶT CHỖ ---")
    return f"Trình xử lý Đặt chỗ đã xử lý yêu cầu: '{request}'. Kết quả: Đã mô phỏng hành động đặt chỗ."

def info_handler(request: str) -> str:
    """Mô phỏng Tác nhân Thông tin xử lý yêu cầu."""
    print("\n--- ĐANG ỦY QUYỀN CHO TRÌNH XỬ LÝ THÔNG TIN ---")
    return f"Trình xử lý Thông tin đã xử lý yêu cầu: '{request}'. Kết quả: Đã mô phỏng truy xuất thông tin."

def unclear_handler(request: str) -> str:
    """Xử lý các yêu cầu không thể ủy quyền."""
    print("\n--- ĐANG XỬ LÝ YÊU CẦU KHÔNG RÕ RÀNG ---")
    return f"Điều phối viên không thể ủy quyền yêu cầu: '{request}'. Vui lòng làm rõ."

# --- Định nghĩa Chuỗi Bộ định tuyến (Router Chain) của Điều phối viên ---
# Chuỗi này quyết định trình xử lý nào sẽ được ủy quyền.

coordinator_router_prompt = ChatPromptTemplate.from_messages([
    ("system", """Phân tích yêu cầu của người dùng và xác định trình xử lý chuyên biệt nào nên xử lý nó.
    - Nếu yêu cầu liên quan đến việc đặt chuyến bay hoặc khách sạn, hãy xuất ra 'booker'.
    - Đối với tất cả các câu hỏi thông tin chung khác, hãy xuất ra 'info'.
    - Nếu yêu cầu không rõ ràng hoặc không phù hợp với cả hai danh mục, hãy xuất ra 'unclear'.
    CHỈ xuất ra một từ: 'booker', 'info', hoặc 'unclear'."""),
    ("user", "{request}")
])

if llm:
    coordinator_router_chain = coordinator_router_prompt | llm | StrOutputParser()

    # --- Định nghĩa Logic Ủy quyền ---
    # Sử dụng RunnableBranch để định tuyến dựa trên đầu ra của chuỗi bộ định tuyến.

    # Định nghĩa các nhánh cho RunnableBranch
    branches = {
        "booker": RunnablePassthrough.assign(output=lambda x: booking_handler(x['request']['request'])),
        "info": RunnablePassthrough.assign(output=lambda x: info_handler(x['request']['request'])),
        "unclear": RunnablePassthrough.assign(output=lambda x: unclear_handler(x['request']['request']))
    }

    # Tạo RunnableBranch. Nó lấy đầu ra của chuỗi bộ định tuyến ('decision')
    # và định tuyến đầu vào ban đầu ('request') đến trình xử lý tương ứng.
    delegation_branch = RunnableBranch(
        (lambda x: x['decision'].strip() == 'booker', branches["booker"]),
        (lambda x: x['decision'].strip() == 'info', branches["info"]),
        branches["unclear"] # Nhánh mặc định
    )

    # Kết hợp chuỗi bộ định tuyến và nhánh ủy quyền thành một runnable duy nhất
    coordinator_agent = {
        "decision": coordinator_router_chain,
        "request": RunnablePassthrough()
    } | delegation_branch | (lambda x: x['output']) # Trích xuất đầu ra cuối cùng

    # --- Ví dụ Sử dụng ---
    def main():
        print("--- Chạy với yêu cầu đặt chỗ ---")
        request_a = "Đặt cho tôi một chuyến bay đến London."
        result_a = coordinator_agent.invoke({"request": request_a})
        print(f"Kết quả cuối cùng A: {result_a}")

        print("\n--- Chạy với yêu cầu thông tin ---")
        request_b = "Thủ đô của Ý là gì?"
        result_b = coordinator_agent.invoke({"request": request_b})
        print(f"Kết quả cuối cùng B: {result_b}")

        print("\n--- Chạy với yêu cầu không rõ ràng ---")
        request_c = "Kể cho tôi nghe về vật lý lượng tử." # Có thể đi vào info hoặc unclear tùy thuộc vào logic
        # Trong ví dụ này, nó có thể đi vào 'info' nếu prompt đủ rộng,
        # nhưng hãy giả sử prompt chặt chẽ hơn hoặc đây là một yêu cầu nằm ngoài phạm vi.
        # Hãy thử một cái gì đó thực sự mơ hồ.
        request_d = "Blah blah blah."
        result_d = coordinator_agent.invoke({"request": request_d})
        print(f"Kết quả cuối cùng D: {result_d}")

    if __name__ == "__main__":
        main()
```

Một thành phần cốt lõi là `coordinator_router_chain`, sử dụng `ChatPromptTemplate` để hướng dẫn mô hình ngôn ngữ phân loại các yêu cầu người dùng đến thành một trong ba danh mục: 'booker', 'info', hoặc 'unclear'. Đầu ra của chuỗi bộ định tuyến này sau đó được sử dụng bởi một `RunnableBranch` để ủy quyền yêu cầu ban đầu cho hàm xử lý tương ứng. `RunnableBranch` kiểm tra quyết định từ mô hình ngôn ngữ và chuyển dữ liệu yêu cầu đến `booking_handler`, `info_handler`, hoặc `unclear_handler`. `coordinator_agent` kết hợp các thành phần này, trước tiên định tuyến yêu cầu để đưa ra quyết định và sau đó chuyển yêu cầu đến trình xử lý đã chọn.

## Ví dụ Code Thực hành (Google ADK)

Bộ công cụ phát triển tác nhân (ADK) của Google cung cấp một môi trường có cấu trúc để xác định các khả năng và hành vi của tác nhân. Trái ngược với các kiến trúc dựa trên đồ thị tính toán rõ ràng, việc định tuyến trong mô hình ADK thường được thực hiện bằng cách xác định một tập hợp các "công cụ" (tools) rời rạc đại diện cho các chức năng của tác nhân. Việc lựa chọn công cụ thích hợp để phản hồi truy vấn của người dùng được quản lý bởi logic nội bộ của framework, tận dụng một mô hình cơ bản để khớp ý định của người dùng với trình xử lý chức năng chính xác.

Trong ví dụ dưới đây, một "Coordinator" (Điều phối viên) được thiết lập để định tuyến các yêu cầu người dùng đến các tác nhân phụ chuyên biệt ("Booker" cho đặt chỗ và "Info" cho thông tin chung) dựa trên các hướng dẫn đã xác định.

```python
# Cài đặt: pip install google-adk google-genai nest_asyncio

import uuid
import os
from google.adk.agents import Agent
from google.adk.runners import InMemoryRunner
from google.adk.tools import FunctionTool
from google.genai import types
import asyncio
import nest_asyncio

# Áp dụng nest_asyncio để cho phép chạy asyncio trong môi trường này (nếu cần, ví dụ: Jupyter)
nest_asyncio.apply()

# --- Định nghĩa các Hàm Công cụ ---
def booking_handler(request: str) -> str:
    """Xử lý các yêu cầu đặt chỗ cho các chuyến bay và khách sạn."""
    print(f"\n[Booker] Đang xử lý đặt chỗ: {request}")
    return f"Đã mô phỏng hành động đặt chỗ cho '{request}'."

def info_handler(request: str) -> str:
    """Xử lý các yêu cầu thông tin chung."""
    print(f"\n[Info] Đang truy xuất thông tin: {request}")
    return f"Đã mô phỏng truy xuất thông tin cho '{request}'."

# Tạo Tools từ Functions
booking_tool = FunctionTool(booking_handler)
info_tool = FunctionTool(info_handler)

# --- Định nghĩa các Tác nhân Phụ chuyên biệt ---
booking_agent = Agent(
    name="Booker",
    model="gemini-2.0-flash",
    description="Một tác nhân chuyên biệt xử lý tất cả các yêu cầu đặt chuyến bay và khách sạn.",
    tools=[booking_tool]
)

info_agent = Agent(
    name="Info",
    model="gemini-2.0-flash",
    description="Một tác nhân chuyên biệt cung cấp thông tin chung và trả lời câu hỏi.",
    tools=[info_tool]
)

# --- Định nghĩa Tác nhân Điều phối viên (Parent Agent) ---
coordinator = Agent(
    name="Coordinator",
    model="gemini-2.0-flash",
    instruction=(
        "Bạn là điều phối viên chính. Nhiệm vụ duy nhất của bạn là phân tích các yêu cầu "
        "của người dùng và ủy quyền chúng cho tác nhân chuyên biệt thích hợp.\n"
        "- Đối với các yêu cầu liên quan đến đặt chuyến bay hoặc khách sạn, hãy ủy quyền cho tác nhân 'Booker'.\n"
        "- Đối với tất cả các câu hỏi thông tin chung khác, hãy ủy quyền cho tác nhân 'Info'."
    ),
    description="Một điều phối viên định tuyến các yêu cầu của người dùng đến tác nhân chuyên biệt chính xác.",
    # Sự hiện diện của sub_agents cho phép ủy quyền dựa trên LLM (Auto-Flow) theo mặc định.
    sub_agents=[booking_agent, info_agent]
)

# --- Logic Thực thi ---
async def run_coordinator(request: str):
    print(f"\n--- Chạy Điều phối viên với yêu cầu: '{request}' ---")
    runner = InMemoryRunner(coordinator)
    user_id = "user_123"
    session_id = str(uuid.uuid4())
    
    # Khởi tạo session (cần thiết cho InMemoryRunner trong một số phiên bản ADK hoặc logic tùy chỉnh)
    # Trong ví dụ đơn giản này, runner.run sẽ tự xử lý hoặc chúng ta giả định môi trường đã sẵn sàng.
    
    response_text = ""
    async for event in runner.run(
        user_id=user_id,
        session_id=session_id,
        new_message=types.Content(
            role='user',
            parts=[types.Part(text=request)]
        )
    ):
        if event.is_final_response() and event.content:
             if hasattr(event.content, 'text') and event.content.text:
                response_text = event.content.text
             elif event.content.parts:
                 response_text = "".join([part.text for part in event.content.parts if part.text])
    
    print(f"Phản hồi cuối cùng: {response_text}")

# --- Chạy Ví dụ ---
async def main():
    # Cần set GOOGLE_API_KEY trong môi trường
    if not os.environ.get("GOOGLE_API_KEY"):
        print("Vui lòng thiết lập biến môi trường GOOGLE_API_KEY.")
        return

    await run_coordinator("Đặt cho tôi một khách sạn ở Paris.")
    await run_coordinator("Tháp Eiffel cao bao nhiêu?")

if __name__ == "__main__":
    asyncio.run(main())
```

Tác nhân `Coordinator` có vai trò chính là phân tích các tin nhắn người dùng đến và ủy quyền chúng cho tác nhân `Booker` hoặc `Info`. Sự ủy quyền này được xử lý tự động bởi cơ chế `Auto-Flow` của ADK vì `Coordinator` đã xác định `sub_agents`. `InMemoryRunner` được sử dụng để xử lý yêu cầu của người dùng thông qua tác nhân điều phối.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các hệ thống Agentic thường phải phản ứng với nhiều loại đầu vào và tình huống không thể xử lý bằng một quy trình tuyến tính, duy nhất. Một quy trình làm việc tuần tự đơn giản thiếu khả năng đưa ra quyết định dựa trên bối cảnh. Nếu không có cơ chế để chọn đúng công cụ hoặc quy trình phụ cho một tác vụ cụ thể, hệ thống vẫn cứng nhắc và không thích ứng.
*   **Tại sao:** Mẫu Routing cung cấp một giải pháp tiêu chuẩn hóa bằng cách giới thiệu logic điều kiện vào khuôn khổ hoạt động của tác nhân. Nó cho phép hệ thống phân tích truy vấn đến để xác định ý định hoặc bản chất của nó. Dựa trên phân tích này, tác nhân tự động hướng luồng kiểm soát đến công cụ, chức năng hoặc tác nhân phụ chuyên biệt phù hợp nhất.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu Routing khi một tác nhân phải quyết định giữa nhiều quy trình làm việc, công cụ hoặc tác nhân phụ riêng biệt dựa trên đầu vào của người dùng hoặc trạng thái hiện tại. Nó rất cần thiết cho các ứng dụng cần phân loại hoặc xử lý các yêu cầu đến để xử lý các loại tác vụ khác nhau, chẳng hạn như bot hỗ trợ khách hàng phân biệt giữa các yêu cầu bán hàng, hỗ trợ kỹ thuật và quản lý tài khoản.

## Những Điểm Chính (Key Takeaways)

*   Routing cho phép các tác nhân đưa ra quyết định năng động về bước tiếp theo trong quy trình làm việc dựa trên các điều kiện.
*   Nó cho phép các tác nhân xử lý các đầu vào đa dạng và điều chỉnh hành vi của chúng, vượt ra ngoài việc thực thi tuyến tính.
*   Logic định tuyến có thể được triển khai bằng cách sử dụng LLM, các hệ thống dựa trên quy tắc hoặc độ tương đồng của embedding.
*   Các framework như LangGraph và Google ADK cung cấp các cách có cấu trúc để xác định và quản lý định tuyến trong các quy trình làm việc của tác nhân, mặc dù với các cách tiếp cận kiến trúc khác nhau.

## Kết luận

Mẫu Routing là một bước quan trọng trong việc xây dựng các hệ thống agentic thực sự năng động và phản hồi. Bằng cách triển khai routing, chúng ta vượt ra ngoài các luồng thực thi tuyến tính đơn giản và trao quyền cho các tác nhân của mình đưa ra các quyết định thông minh về cách xử lý thông tin, phản hồi đầu vào của người dùng và sử dụng các công cụ hoặc tác nhân phụ có sẵn.

Chúng ta đã thấy cách routing có thể được áp dụng trong nhiều lĩnh vực khác nhau, từ chatbot dịch vụ khách hàng đến các đường ống xử lý dữ liệu phức tạp. Khả năng phân tích đầu vào và điều hướng quy trình làm việc có điều kiện là nền tảng để tạo ra các tác nhân có thể xử lý sự thay đổi vốn có của các tác vụ trong thế giới thực. Việc thành thạo mẫu Routing là điều cần thiết để xây dựng các tác nhân có thể điều hướng thông minh các tình huống khác nhau và cung cấp các phản hồi hoặc hành động phù hợp dựa trên ngữ cảnh.

## Tài liệu tham khảo
1. LangGraph Documentation: https://www.langchain.com/
2. Google Agent Developer Kit Documentation: https://google.github.io/adk-docs/