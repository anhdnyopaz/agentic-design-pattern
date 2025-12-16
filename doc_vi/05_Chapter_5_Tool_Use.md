# Chương 5: Tool Use (Sử Dụng Công Cụ / Function Calling)

## Tổng quan về Mẫu thiết kế Tool Use

Cho đến nay, chúng ta đã thảo luận về các mẫu Agentic chủ yếu liên quan đến việc điều phối các tương tác giữa các mô hình ngôn ngữ và quản lý luồng thông tin trong quy trình làm việc nội bộ của tác nhân (Chaining, Routing, Parallelization, Reflection). Tuy nhiên, để các tác nhân thực sự hữu ích và tương tác với thế giới thực hoặc các hệ thống bên ngoài, chúng cần có khả năng sử dụng **Công cụ (Tools)**.

Mẫu Tool Use, thường được triển khai thông qua một cơ chế gọi là **Function Calling**, cho phép một tác nhân tương tác với các API, cơ sở dữ liệu, dịch vụ bên ngoài hoặc thậm chí thực thi mã. Nó cho phép LLM là cốt lõi của tác nhân quyết định khi nào và cách sử dụng một chức năng bên ngoài cụ thể dựa trên yêu cầu của người dùng hoặc trạng thái hiện tại của tác vụ.

Quá trình này thường bao gồm:

1.  **Định nghĩa Công cụ (Tool Definition):** Các hàm hoặc khả năng bên ngoài được định nghĩa và mô tả cho LLM. Mô tả này bao gồm mục đích, tên và các tham số mà nó chấp nhận, cùng với các kiểu và mô tả của chúng.
2.  **Quyết định của LLM (LLM Decision):** LLM nhận yêu cầu của người dùng và các định nghĩa công cụ có sẵn. Dựa trên sự hiểu biết của nó về yêu cầu và các công cụ, LLM quyết định xem có cần gọi một hoặc nhiều công cụ để hoàn thành yêu cầu hay không.
3.  **Tạo Lời gọi Hàm (Function Call Generation):** Nếu LLM quyết định sử dụng một công cụ, nó sẽ tạo ra một đầu ra có cấu trúc (thường là một đối tượng JSON) chỉ định tên công cụ cần gọi và các đối số (tham số) để truyền cho nó, được trích xuất từ yêu cầu của người dùng.
4.  **Thực thi Công cụ (Tool Execution):** Framework Agentic hoặc lớp điều phối chặn đầu ra có cấu trúc này. Nó xác định công cụ được yêu cầu và thực thi hàm bên ngoài thực tế với các đối số được cung cấp.
5.  **Quan sát/Kết quả (Observation/Result):** Đầu ra hoặc kết quả từ việc thực thi công cụ được trả về cho tác nhân.
6.  **Xử lý của LLM (LLM Processing) (Tùy chọn nhưng phổ biến):** LLM nhận đầu ra của công cụ làm ngữ cảnh và sử dụng nó để tạo ra phản hồi cuối cùng cho người dùng hoặc quyết định bước tiếp theo trong quy trình làm việc (có thể liên quan đến việc gọi một công cụ khác, phản chiếu hoặc cung cấp câu trả lời cuối cùng).

Mẫu này là nền tảng vì nó phá vỡ các giới hạn của dữ liệu đào tạo của LLM và cho phép nó truy cập thông tin cập nhật, thực hiện các phép tính mà nó không thể tự thực hiện nội bộ, tương tác với dữ liệu dành riêng cho người dùng hoặc kích hoạt các hành động trong thế giới thực. Function Calling là cơ chế kỹ thuật bắc cầu khoảng cách giữa khả năng suy luận của LLM và vô số chức năng bên ngoài có sẵn.

Mặc dù "function calling" mô tả một cách thích hợp việc gọi các hàm mã được xác định trước, nhưng việc xem xét khái niệm "tool calling" rộng lớn hơn là hữu ích. Thuật ngữ rộng hơn này thừa nhận rằng các khả năng của tác nhân có thể mở rộng rất xa ngoài việc thực thi hàm đơn giản. Một "công cụ" có thể là một hàm truyền thống, nhưng nó cũng có thể là một điểm cuối API phức tạp, một yêu cầu đến cơ sở dữ liệu, hoặc thậm chí là một hướng dẫn được gửi đến một tác nhân chuyên biệt khác. Quan điểm này cho phép chúng ta hình dung các hệ thống phức tạp hơn, nơi, ví dụ, một tác nhân chính có thể ủy quyền một tác vụ phân tích dữ liệu phức tạp cho một "tác nhân phân tích" chuyên dụng hoặc truy vấn một cơ sở tri thức bên ngoài thông qua API của nó.

Các framework như LangChain, LangGraph và Google Agent Developer Kit (ADK) cung cấp hỗ trợ mạnh mẽ để định nghĩa công cụ và tích hợp chúng vào các quy trình làm việc của tác nhân, thường tận dụng các khả năng gọi hàm gốc của các LLM hiện đại như trong dòng Gemini hoặc OpenAI. Trên "canvas" của các framework này, bạn xác định các công cụ và sau đó cấu hình các tác nhân (thường là LLM Agents) để nhận biết và có khả năng sử dụng các công cụ này.

Tool Use là một mẫu nền tảng để xây dựng các tác nhân mạnh mẽ, tương tác và nhận thức bên ngoài.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu Tool Use có thể áp dụng trong hầu hết mọi trường hợp tác nhân cần vượt ra ngoài việc tạo văn bản để thực hiện hành động hoặc truy xuất thông tin cụ thể, động:

1.  **Truy xuất Thông tin từ Nguồn Bên ngoài:** Truy cập dữ liệu thời gian thực hoặc thông tin không có trong dữ liệu đào tạo của LLM.
    *   **Trường hợp sử dụng:** Một tác nhân thời tiết.
    *   **Công cụ:** Một API thời tiết nhận một vị trí và trả về điều kiện thời tiết hiện tại.
    *   **Luồng Tác nhân:** Người dùng hỏi, "Thời tiết ở London thế nào?", LLM xác định cần công cụ thời tiết, gọi công cụ với "London", công cụ trả về dữ liệu, LLM định dạng dữ liệu thành phản hồi thân thiện với người dùng.
2.  **Tương tác với Cơ sở dữ liệu và API:** Thực hiện các truy vấn, cập nhật hoặc các hoạt động khác trên dữ liệu có cấu trúc.
    *   **Trường hợp sử dụng:** Một tác nhân thương mại điện tử.
    *   **Công cụ:** Các lệnh gọi API để kiểm tra hàng tồn kho sản phẩm, lấy trạng thái đơn hàng hoặc xử lý thanh toán.
    *   **Luồng Tác nhân:** Người dùng hỏi "Sản phẩm X có trong kho không?", LLM gọi API hàng tồn kho, công cụ trả về số lượng hàng tồn kho, LLM cho người dùng biết trạng thái hàng tồn kho.
3.  **Thực hiện Tính toán và Phân tích Dữ liệu:** Sử dụng các máy tính bên ngoài, thư viện phân tích dữ liệu hoặc công cụ thống kê.
    *   **Trường hợp sử dụng:** Một tác nhân tài chính.
    *   **Công cụ:** Một hàm máy tính, một API dữ liệu thị trường chứng khoán, một công cụ bảng tính.
    *   **Luồng Tác nhân:** Người dùng hỏi "Giá hiện tại của AAPL là bao nhiêu và tính toán lợi nhuận tiềm năng nếu tôi mua 100 cổ phiếu với giá 150$?", LLM gọi API chứng khoán, lấy giá hiện tại, sau đó gọi công cụ máy tính, lấy kết quả, định dạng phản hồi.
4.  **Gửi Thông tin liên lạc:** Gửi email, tin nhắn hoặc thực hiện các lệnh gọi API đến các dịch vụ liên lạc bên ngoài.
    *   **Trường hợp sử dụng:** Một tác nhân trợ lý cá nhân.
    *   **Công cụ:** Một API gửi email.
    *   **Luồng Tác nhân:** Người dùng nói, "Gửi email cho John về cuộc họp ngày mai.", LLM gọi công cụ email với người nhận, chủ đề và nội dung được trích xuất từ yêu cầu.
5.  **Thực thi Mã:** Chạy các đoạn mã trong một môi trường an toàn để thực hiện các tác vụ cụ thể.
    *   **Trường hợp sử dụng:** Một tác nhân trợ lý mã hóa.
    *   **Công cụ:** Một trình thông dịch mã (code interpreter).
    *   **Luồng Tác nhân:** Người dùng cung cấp một đoạn Python và hỏi, "Đoạn mã này làm gì?", LLM sử dụng công cụ trình thông dịch để chạy mã và phân tích đầu ra của nó.
6.  **Điều khiển các Hệ thống hoặc Thiết bị khác:** Tương tác với các thiết bị nhà thông minh, nền tảng IoT hoặc các hệ thống được kết nối khác.
    *   **Trường hợp sử dụng:** Một tác nhân nhà thông minh.
    *   **Công cụ:** Một API để điều khiển đèn thông minh.
    *   **Luồng Tác nhân:** Người dùng nói, "Tắt đèn phòng khách.", LLM gọi công cụ nhà thông minh với lệnh và thiết bị mục tiêu.

Tool Use là yếu tố biến mô hình ngôn ngữ từ một trình tạo văn bản thành một tác nhân có khả năng cảm nhận, suy luận và hành động trong thế giới kỹ thuật số hoặc vật lý.

## Ví dụ Code Thực hành (LangChain)

Việc triển khai sử dụng công cụ trong framework LangChain là một quy trình hai giai đoạn. Ban đầu, một hoặc nhiều công cụ được định nghĩa, thường bằng cách đóng gói các hàm Python hiện có hoặc các thành phần có thể chạy khác. Sau đó, các công cụ này được liên kết với một mô hình ngôn ngữ, do đó cấp cho mô hình khả năng tạo ra yêu cầu sử dụng công cụ có cấu trúc khi nó xác định rằng cần một lệnh gọi hàm bên ngoài để thực hiện truy vấn của người dùng.

Việc thực thi ví dụ này yêu cầu cài đặt các thư viện LangChain cốt lõi và một gói nhà cung cấp cụ thể cho mô hình. Hơn nữa, việc xác thực đúng cách với dịch vụ mô hình ngôn ngữ đã chọn, thường thông qua API key được cấu hình trong môi trường cục bộ, là một điều kiện tiên quyết cần thiết.

```python
# Cài đặt: pip install langchain langchain-community langchain-openai

import os
import getpass
import asyncio
from typing import List
from dotenv import load_dotenv
import logging

from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.tools import tool as langchain_tool
from langchain.agents import create_tool_calling_agent, AgentExecutor

# --- Cấu hình API ---
# Tải biến môi trường
load_dotenv()

# Yêu cầu người dùng nhập API key một cách an toàn và đặt chúng làm biến môi trường
if not os.environ.get("GOOGLE_API_KEY"):
    os.environ["GOOGLE_API_KEY"] = getpass.getpass("Nhập Google API Key của bạn: ")

# --- Khởi tạo LLM ---
try:
    # Cần một mô hình có khả năng gọi hàm/công cụ.
    llm = ChatGoogleGenerativeAI(model="gemini-pro", temperature=0) # Sử dụng gemini-pro cho khả năng gọi hàm
    print(f"✓ Mô hình ngôn ngữ đã khởi tạo: {llm.model}")
except Exception as e:
    print(f"✗ Lỗi khởi tạo mô hình ngôn ngữ: {e}")
    llm = None

# --- Định nghĩa một Công cụ ---
@langchain_tool
def search_information(query: str) -> str:
    """
    Cung cấp thông tin thực tế về một chủ đề nhất định. Sử dụng công cụ này để tìm câu trả lời cho các cụm từ
    như 'thủ đô của Pháp' hoặc 'thời tiết ở London?'
    """
    print(f"\n--- ☑ Công cụ được gọi: search_information với truy vấn: '{query}' ---")
    # Mô phỏng một công cụ tìm kiếm với một từ điển các kết quả được xác định trước.
    simulated_results = {
        "thời tiết ở london": "Thời tiết ở London hiện đang có mây với nhiệt độ 15°C.",
        "thủ đô của pháp": "Thủ đô của Pháp là Paris.",
        "dân số trái đất": "Dân số ước tính của Trái đất là khoảng 8 tỷ người.",
        "ngọn núi cao nhất": "Núi Everest là ngọn núi cao nhất trên mực nước biển.",
        "default": f"Kết quả tìm kiếm mô phỏng cho '{query}': Không tìm thấy thông tin cụ thể, nhưng chủ đề có vẻ thú vị."
    }
    result = simulated_results.get(query.lower(), simulated_results["default"])
    print(f"--- KẾT QUẢ CÔNG CỤ: {result} ---")
    return result

tools = [search_information]

# --- Tạo một Tác nhân Gọi Công cụ ---
if llm:
    # Prompt template này yêu cầu một placeholder `agent_scratchpad`
    # cho các bước nội bộ của tác nhân.
    agent_prompt = ChatPromptTemplate.from_messages([
        ("system", "Bạn là một trợ lý hữu ích. Hãy trả lời các câu hỏi bằng cách sử dụng các công cụ có sẵn."),
        ("human", "{input}"),
        ("placeholder", "{agent_scratchpad}"),
    ])

    # Tạo tác nhân, liên kết LLM, công cụ và prompt lại với nhau.
    agent = create_tool_calling_agent(llm, tools, agent_prompt)

    # AgentExecutor là runtime gọi tác nhân và thực thi các công cụ đã chọn.
    # Đối số 'tools' không cần thiết ở đây vì chúng đã được liên kết với tác nhân.
    agent_executor = AgentExecutor(agent=agent, tools=tools, verbose=True)

    async def run_agent_with_tool(query: str):
        """Gọi agent executor với một truy vấn và in phản hồi cuối cùng."""
        print(f"\n--- Chạy Agent với Truy vấn: '{query}' ---")
        try:
            response = await agent_executor.ainvoke({"input": query})
            print("\n--- ☑ Phản hồi Agent Cuối cùng ---")
            print(response["output"])
        except Exception as e:
            print(f"\n ✗ Một lỗi đã xảy ra trong quá trình thực thi agent: {e}")

    async def main():
        # Chạy tất cả các truy vấn agent đồng thời.
        tasks = [
            run_agent_with_tool("Thủ đô của Pháp là gì?"),
            run_agent_with_tool("Thời tiết ở London thế nào?"),
            run_agent_with_tool("Kể cho tôi nghe điều gì đó về chó.") # Sẽ kích hoạt phản hồi công cụ mặc định
        ]
        await asyncio.gather(*tasks)

    if __name__ == "__main__":
        asyncio.run(main())
```

Mã này thiết lập một tác nhân gọi công cụ sử dụng thư viện LangChain và mô hình Google Gemini. Nó định nghĩa một công cụ `search_information` mô phỏng việc cung cấp các câu trả lời thực tế cho các truy vấn cụ thể. Công cụ này có các phản hồi được xác định trước cho "thời tiết ở london", "thủ đô của pháp" và "dân số trái đất", và một phản hồi mặc định cho các truy vấn khác. Một mô hình `ChatGoogleGenerativeAI` được khởi tạo, đảm bảo nó có khả năng gọi công cụ. `ChatPromptTemplate` được tạo để hướng dẫn tương tác của tác nhân. Hàm `create_tool_calling_agent` được sử dụng để kết hợp mô hình ngôn ngữ, công cụ và prompt thành một tác nhân. `AgentExecutor` sau đó được thiết lập để quản lý việc thực thi tác nhân và gọi công cụ. Hàm bất đồng bộ `run_agent_with_tool` được định nghĩa để gọi tác nhân với một truy vấn nhất định và in kết quả. Hàm bất đồng bộ chính chuẩn bị nhiều truy vấn để chạy đồng thời. Các truy vấn này được thiết kế để kiểm tra cả các phản hồi cụ thể và mặc định của công cụ `search_information`. Cuối cùng, lệnh gọi `asyncio.run(main())` thực thi tất cả các tác vụ của tác nhân. Mã bao gồm các kiểm tra để khởi tạo LLM thành công trước khi thiết lập và thực thi tác nhân.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** LLM là các trình tạo văn bản mạnh mẽ, nhưng về cơ bản chúng bị ngắt kết nối khỏi thế giới bên ngoài. Kiến thức của chúng là tĩnh, giới hạn trong dữ liệu mà chúng đã được đào tạo, và chúng thiếu khả năng thực hiện hành động hoặc truy xuất thông tin thời gian thực. Hạn chế cố hữu này ngăn chúng hoàn thành các tác vụ yêu cầu tương tác với các API, cơ sở dữ liệu hoặc dịch vụ bên ngoài. Nếu không có cầu nối đến các hệ thống bên ngoài này, tiện ích của chúng để giải quyết các vấn đề trong thế giới thực bị hạn chế nghiêm trọng.
*   **Tại sao:** Mẫu Tool Use, thường được triển khai thông qua function calling, cung cấp một giải pháp tiêu chuẩn hóa cho vấn đề này. Nó hoạt động bằng cách mô tả các chức năng bên ngoài có sẵn, hoặc "công cụ", cho LLM theo cách mà nó có thể hiểu. Dựa trên yêu cầu của người dùng, LLM Agentic sau đó có thể quyết định xem có cần công cụ hay không và tạo ra một đối tượng dữ liệu có cấu trúc (như JSON) chỉ định chức năng cần gọi và các đối số. Một lớp điều phối thực thi lệnh gọi hàm này, truy xuất kết quả và đưa nó trở lại LLM. Điều này cho phép LLM kết hợp thông tin bên ngoài cập nhật hoặc kết quả của một hành động vào phản hồi cuối cùng của nó, giúp nó có khả năng hành động.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu Tool Use bất cứ khi nào tác nhân cần vượt ra ngoài kiến thức nội bộ của LLM và tương tác với thế giới bên ngoài. Điều này rất cần thiết cho các tác vụ yêu cầu dữ liệu thời gian thực (ví dụ: kiểm tra thời tiết, giá cổ phiếu), truy cập thông tin riêng tư hoặc độc quyền (ví dụ: truy vấn cơ sở dữ liệu của công ty), thực hiện các phép tính chính xác, thực thi mã hoặc kích hoạt các hành động trong các hệ thống khác (ví dụ: gửi email, điều khiển thiết bị thông minh).

## Những Điểm Chính (Key Takeaways)

*   Tool Use (Function Calling) cho phép các tác nhân tương tác với các hệ thống bên ngoài và truy cập thông tin động.
*   Nó liên quan đến việc định nghĩa các công cụ với các mô tả và tham số rõ ràng mà LLM có thể hiểu.
*   LLM quyết định khi nào sử dụng một công cụ và tạo các lệnh gọi hàm có cấu trúc.
*   Các framework Agentic thực thi các lệnh gọi công cụ thực tế và trả về kết quả cho LLM.
*   Tool Use rất cần thiết để xây dựng các tác nhân có thể thực hiện các hành động trong thế giới thực và cung cấp thông tin cập nhật.
*   LangChain đơn giản hóa việc định nghĩa công cụ bằng cách sử dụng decorator `@tool` và cung cấp `create_tool_calling_agent` và `AgentExecutor` để xây dựng các tác nhân sử dụng công cụ.

## Kết luận

Mẫu Tool Use là một nguyên tắc kiến trúc quan trọng để mở rộng phạm vi chức năng của các mô hình ngôn ngữ lớn vượt ra ngoài khả năng tạo văn bản nội tại của chúng. Bằng cách trang bị cho một mô hình khả năng giao tiếp với phần mềm và nguồn dữ liệu bên ngoài, mô hình này cho phép tác nhân thực hiện các hành động, thực thi tính toán và truy xuất thông tin từ các hệ thống khác. Quá trình này liên quan đến việc mô hình tạo ra một yêu cầu có cấu trúc để gọi một công cụ bên ngoài khi nó xác định rằng việc làm như vậy là cần thiết để thực hiện truy vấn của người dùng. Các framework như LangChain, Google ADK và CrewAI cung cấp các trừu tượng và thành phần có cấu trúc tạo điều kiện thuận lợi cho việc tích hợp các công cụ bên ngoài này. Các framework này quản lý quá trình hiển thị các thông số kỹ thuật công cụ cho mô hình và phân tích các yêu cầu sử dụng công cụ tiếp theo của nó. Điều này đơn giản hóa việc phát triển các hệ thống agentic tinh vi có thể tương tác và thực hiện hành động trong các môi trường kỹ thuật số bên ngoài.

## Tài liệu tham khảo

1.  LangChain Documentation (Tools): https://python.langchain.com/docs/integrations/tools/
2.  Google Agent Developer Kit (ADK) Documentation (Tools): https://google.github.io/adk-docs/tools/
3.  OpenAI Function Calling Documentation: https://platform.openai.com/docs/guides/function-calling
4.  CrewAI Documentation (Tools): https://docs.crewai.com/concepts/tools