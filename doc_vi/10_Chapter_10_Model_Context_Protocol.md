# Chương 10: Model Context Protocol (MCP)

Để cho phép các LLM hoạt động hiệu quả như các tác nhân, khả năng của chúng phải vượt ra ngoài việc tạo đa phương thức. Sự tương tác với môi trường bên ngoài là cần thiết, bao gồm quyền truy cập vào dữ liệu hiện tại, sử dụng phần mềm bên ngoài và thực hiện các tác vụ hoạt động cụ thể. **Model Context Protocol (MCP)** giải quyết nhu cầu này bằng cách cung cấp một giao diện tiêu chuẩn hóa để các LLM giao tiếp với các tài nguyên bên ngoài. Giao thức này đóng vai trò là cơ chế chính để tạo điều kiện cho sự tích hợp nhất quán và có thể dự đoán được.

## Tổng quan về Mẫu thiết kế MCP

Hãy tưởng tượng một bộ chuyển đổi đa năng (universal adapter) cho phép bất kỳ LLM nào cắm vào bất kỳ hệ thống, cơ sở dữ liệu hoặc công cụ bên ngoài nào mà không cần tích hợp tùy chỉnh cho từng cái. Đó về cơ bản là Model Context Protocol (MCP). Đây là một tiêu chuẩn mở được thiết kế để chuẩn hóa cách các LLM như Gemini, GPT của OpenAI, Mixtral và Claude giao tiếp với các ứng dụng, nguồn dữ liệu và công cụ bên ngoài. Hãy nghĩ về nó như một cơ chế kết nối chung giúp đơn giản hóa cách các LLM lấy bối cảnh, thực hiện hành động và tương tác với các hệ thống khác nhau.

MCP hoạt động dựa trên kiến trúc **client-server**. Nó định nghĩa cách các yếu tố khác nhau—dữ liệu (được gọi là **Resources** - Tài nguyên), các mẫu tương tác (về cơ bản là **Prompts**), và các chức năng có thể hành động (được gọi là **Tools** - Công cụ)—được hiển thị bởi một **MCP Server**. Những thứ này sau đó được tiêu thụ bởi một **MCP Client**, có thể là một ứng dụng máy chủ LLM hoặc chính một tác nhân AI. Cách tiếp cận tiêu chuẩn hóa này làm giảm đáng kể độ phức tạp của việc tích hợp các LLM vào các môi trường hoạt động đa dạng.

Tuy nhiên, MCP là một hợp đồng cho một "giao diện agentic", và hiệu quả của nó phụ thuộc rất nhiều vào thiết kế của các API cơ bản mà nó hiển thị. Có một rủi ro là các nhà phát triển chỉ đơn giản bao bọc các API cũ, có sẵn mà không sửa đổi, điều này có thể không tối ưu cho một tác nhân. Ví dụ, nếu API của hệ thống bán vé chỉ cho phép truy xuất chi tiết vé từng cái một, một tác nhân được yêu cầu tóm tắt các vé ưu tiên cao sẽ chậm và không chính xác ở khối lượng lớn. Để thực sự hiệu quả, API cơ bản nên được cải thiện với các tính năng xác định (deterministic) như lọc và sắp xếp để giúp tác nhân không xác định (non-deterministic) hoạt động hiệu quả. Điều này nhấn mạnh rằng các tác nhân không thay thế các quy trình làm việc xác định một cách kỳ diệu; chúng thường yêu cầu sự hỗ trợ xác định mạnh mẽ hơn để thành công.

Hơn nữa, MCP có thể bao bọc một API mà đầu vào hoặc đầu ra của nó vẫn không vốn dĩ dễ hiểu đối với tác nhân. Một API chỉ hữu ích nếu định dạng dữ liệu của nó thân thiện với tác nhân, một sự đảm bảo mà bản thân MCP không thực thi. Ví dụ, việc tạo một máy chủ MCP cho kho tài liệu trả về các tệp dưới dạng PDF hầu như vô dụng nếu tác nhân tiêu thụ không thể phân tích nội dung PDF. Cách tiếp cận tốt hơn là trước tiên tạo một API trả về phiên bản văn bản của tài liệu, chẳng hạn như Markdown, mà tác nhân thực sự có thể đọc và xử lý. Điều này chứng minh rằng các nhà phát triển phải xem xét không chỉ kết nối, mà còn bản chất của dữ liệu được trao đổi để đảm bảo khả năng tương thích thực sự.

## MCP so với Tool Function Calling

Model Context Protocol (MCP) và tool function calling (gọi hàm công cụ) là các cơ chế riêng biệt cho phép LLM tương tác với các khả năng bên ngoài (bao gồm cả công cụ) và thực hiện các hành động. Mặc dù cả hai đều phục vụ để mở rộng khả năng của LLM vượt ra ngoài việc tạo văn bản, chúng khác nhau về cách tiếp cận và mức độ trừu tượng.

**Tool function calling** có thể được coi là một yêu cầu trực tiếp từ một LLM đến một công cụ hoặc hàm cụ thể, được xác định trước. Trong bối cảnh này, chúng ta sử dụng các từ "công cụ" và "hàm" thay thế cho nhau. Tương tác này được đặc trưng bởi mô hình giao tiếp một-một, trong đó LLM định dạng một yêu cầu dựa trên sự hiểu biết của nó về ý định của người dùng yêu cầu hành động bên ngoài. Mã ứng dụng sau đó thực thi yêu cầu này và trả lại kết quả cho LLM. Quá trình này thường là độc quyền và thay đổi giữa các nhà cung cấp LLM khác nhau.

Ngược lại, **Model Context Protocol (MCP)** hoạt động như một giao diện tiêu chuẩn hóa để các LLM khám phá, giao tiếp và sử dụng các khả năng bên ngoài. Nó hoạt động như một giao thức mở tạo điều kiện thuận lợi cho sự tương tác với một loạt các công cụ và hệ thống, nhằm mục đích thiết lập một hệ sinh thái nơi bất kỳ công cụ tuân thủ nào cũng có thể được truy cập bởi bất kỳ LLM tuân thủ nào. Điều này thúc đẩy khả năng tương tác, khả năng kết hợp và khả năng tái sử dụng trên các hệ thống và triển khai khác nhau. Bằng cách áp dụng mô hình liên kết, chúng ta cải thiện đáng kể khả năng tương tác và mở khóa giá trị của các tài sản hiện có. Chiến lược này cho phép chúng ta đưa các dịch vụ khác biệt và cũ vào một hệ sinh thái hiện đại chỉ bằng cách bao bọc chúng trong một giao diện tuân thủ MCP. Các dịch vụ này tiếp tục hoạt động độc lập, nhưng giờ đây có thể được kết hợp vào các ứng dụng và quy trình làm việc mới, với sự hợp tác của chúng được điều phối bởi các LLM. Điều này thúc đẩy sự nhanh nhẹn và khả năng tái sử dụng mà không yêu cầu viết lại tốn kém các hệ thống nền tảng.

Dưới đây là bảng phân tích các sự phân biệt cơ bản giữa MCP và tool function calling:

| Đặc điểm | Tool Function Calling | Model Context Protocol (MCP) |
| :--- | :--- | :--- |
| **Tiêu chuẩn hóa** | Độc quyền và dành riêng cho nhà cung cấp. Định dạng và triển khai khác nhau giữa các nhà cung cấp LLM. | Một giao thức mở, được tiêu chuẩn hóa, thúc đẩy khả năng tương tác giữa các LLM và công cụ khác nhau. |
| **Phạm vi** | Một cơ chế trực tiếp để LLM yêu cầu thực thi một hàm cụ thể, được xác định trước. | Một khuôn khổ rộng hơn về cách các LLM và các công cụ bên ngoài khám phá và giao tiếp với nhau. |
| **Kiến trúc** | Tương tác một-một giữa LLM và logic xử lý công cụ của ứng dụng. | Kiến trúc client-server nơi các ứng dụng hỗ trợ LLM (client) có thể kết nối và sử dụng các máy chủ MCP khác nhau (tools). |
| **Khám phá** | LLM được thông báo rõ ràng công cụ nào có sẵn trong ngữ cảnh của một cuộc trò chuyện cụ thể. | Cho phép khám phá động các công cụ có sẵn. Một ứng dụng khách MCP có thể truy vấn một máy chủ để xem nó cung cấp những khả năng gì. |
| **Khả năng tái sử dụng** | Tích hợp công cụ thường được ghép nối chặt chẽ với ứng dụng cụ thể và LLM đang được sử dụng. | Thúc đẩy sự phát triển của các "máy chủ MCP" độc lập, có thể tái sử dụng, có thể được truy cập bởi bất kỳ ứng dụng tuân thủ nào. |

Hãy nghĩ về *tool function calling* như việc đưa cho AI một bộ công cụ tùy chỉnh cụ thể, giống như một cờ lê và tuốc nơ vít cụ thể. Điều này hiệu quả cho một xưởng với một bộ nhiệm vụ cố định. *MCP (Model Context Protocol)*, mặt khác, giống như việc tạo ra một hệ thống ổ cắm điện tiêu chuẩn hóa, phổ quát. Nó không tự cung cấp các công cụ, nhưng nó cho phép bất kỳ công cụ tuân thủ nào từ bất kỳ nhà sản xuất nào cắm vào và hoạt động, cho phép một xưởng làm việc năng động và không ngừng mở rộng.

Tóm lại, function calling cung cấp quyền truy cập trực tiếp vào một vài chức năng cụ thể, trong khi MCP là khung giao tiếp tiêu chuẩn hóa cho phép các LLM khám phá và sử dụng một loạt các tài nguyên bên ngoài. Đối với các ứng dụng đơn giản, các công cụ cụ thể là đủ; đối với các hệ thống AI phức tạp, liên kết với nhau cần thích ứng, một tiêu chuẩn phổ quát như MCP là điều cần thiết.

## Các cân nhắc bổ sung cho MCP

Mặc dù MCP trình bày một khuôn khổ mạnh mẽ, một đánh giá kỹ lưỡng đòi hỏi phải xem xét một số khía cạnh quan trọng ảnh hưởng đến sự phù hợp của nó cho một trường hợp sử dụng nhất định.

*   **Tool vs. Resource vs. Prompt:** Điều quan trọng là phải hiểu vai trò cụ thể của các thành phần này.
    *   **Resource (Tài nguyên):** Là dữ liệu tĩnh (ví dụ: tệp PDF, bản ghi cơ sở dữ liệu).
    *   **Tool (Công cụ):** Là một hàm có thể thực thi thực hiện một hành động (ví dụ: gửi email, truy vấn API).
    *   **Prompt:** Là một mẫu hướng dẫn LLM cách tương tác với tài nguyên hoặc công cụ, đảm bảo sự tương tác có cấu trúc và hiệu quả.
*   **Khả năng khám phá (Discoverability):** Một lợi thế chính của MCP là ứng dụng khách MCP có thể truy vấn động một máy chủ để tìm hiểu những công cụ và tài nguyên nào nó cung cấp. Cơ chế khám phá "kịp thời" (just-in-time) này rất mạnh mẽ cho các tác nhân cần thích nghi với các khả năng mới mà không cần triển khai lại.
*   **Bảo mật:** Việc hiển thị các công cụ và dữ liệu qua bất kỳ giao thức nào cũng đòi hỏi các biện pháp bảo mật mạnh mẽ. Một triển khai MCP phải bao gồm xác thực và ủy quyền để kiểm soát ứng dụng khách nào có thể truy cập máy chủ nào và những hành động cụ thể nào họ được phép thực hiện.
*   **Triển khai:** Mặc dù MCP là một tiêu chuẩn mở, việc triển khai nó có thể phức tạp. Tuy nhiên, các nhà cung cấp đang bắt đầu đơn giản hóa quá trình này. Ví dụ: một số nhà cung cấp mô hình như Anthropic hoặc **FastMCP** cung cấp SDK tóm tắt đi nhiều mã soạn sẵn (boilerplate code), giúp các nhà phát triển tạo và kết nối các máy khách và máy chủ MCP dễ dàng hơn.
*   **Xử lý lỗi:** Một chiến lược xử lý lỗi toàn diện là rất quan trọng. Giao thức phải xác định cách các lỗi (ví dụ: lỗi thực thi công cụ, máy chủ không khả dụng, yêu cầu không hợp lệ) được truyền đạt lại cho LLM để nó có thể hiểu lỗi và có khả năng thử một cách tiếp cận thay thế.
*   **Máy chủ Cục bộ so với Từ xa:** Máy chủ MCP có thể được triển khai cục bộ trên cùng một máy với tác nhân hoặc từ xa trên một máy chủ khác. Máy chủ cục bộ có thể được chọn vì tốc độ và bảo mật với dữ liệu nhạy cảm, trong khi kiến trúc máy chủ từ xa cho phép truy cập chia sẻ, có thể mở rộng vào các công cụ chung trên toàn tổ chức.
*   **Theo yêu cầu so với Hàng loạt:** MCP có thể hỗ trợ cả các phiên tương tác theo yêu cầu và xử lý hàng loạt quy mô lớn hơn. Sự lựa chọn phụ thuộc vào ứng dụng, từ một tác nhân hội thoại thời gian thực cần truy cập công cụ ngay lập tức đến một đường ống phân tích dữ liệu xử lý các bản ghi theo lô.
*   **Cơ chế Vận chuyển:** Giao thức cũng định nghĩa các lớp vận chuyển cơ bản để giao tiếp. Đối với các tương tác cục bộ, nó sử dụng JSON-RPC qua STDIO (đầu vào/đầu ra tiêu chuẩn) để giao tiếp giữa các quy trình hiệu quả. Đối với các kết nối từ xa, nó tận dụng các giao thức thân thiện với web như Streamable HTTP và Server-Sent Events (SSE) để cho phép giao tiếp client-server bền vững và hiệu quả.

Model Context Protocol sử dụng mô hình client-server để chuẩn hóa luồng thông tin. Hiểu sự tương tác của các thành phần là chìa khóa cho hành vi agentic tiên tiến của MCP:

1.  **Large Language Model (LLM):** Trí thông minh cốt lõi. Nó xử lý các yêu cầu của người dùng, lập kế hoạch và quyết định khi nào nó cần truy cập thông tin bên ngoài hoặc thực hiện một hành động.
2.  **MCP Client:** Đây là một ứng dụng hoặc trình bao bọc xung quanh LLM. Nó hoạt động như người trung gian, dịch ý định của LLM thành một yêu cầu chính thức tuân theo tiêu chuẩn MCP. Nó chịu trách nhiệm khám phá, kết nối và giao tiếp với các MCP Server.
3.  **MCP Server:** Đây là cửa ngõ vào thế giới bên ngoài. Nó hiển thị một tập hợp các công cụ, tài nguyên và prompt cho bất kỳ MCP Client nào được ủy quyền. Mỗi máy chủ thường chịu trách nhiệm cho một miền cụ thể, chẳng hạn như kết nối với cơ sở dữ liệu nội bộ của công ty, dịch vụ email hoặc API công khai.
4.  **Dịch vụ Bên thứ ba (3P) Tùy chọn:** Điều này đại diện cho công cụ, ứng dụng hoặc nguồn dữ liệu bên ngoài thực tế mà MCP Server quản lý và hiển thị. Nó là điểm cuối cùng thực hiện hành động được yêu cầu, chẳng hạn như truy vấn cơ sở dữ liệu độc quyền, tương tác với nền tảng SaaS hoặc gọi API thời tiết công cộng.

Luồng tương tác diễn ra như sau:
1.  **Khám phá:** MCP Client, thay mặt cho LLM, truy vấn một MCP Server để hỏi những khả năng nào nó cung cấp. Máy chủ phản hồi bằng một bảng kê khai liệt kê các công cụ có sẵn (ví dụ: `send_email`), tài nguyên (ví dụ: `customer_database`) và prompt.
2.  **Xây dựng Yêu cầu:** LLM xác định rằng nó cần sử dụng một trong những công cụ đã khám phá. Ví dụ: nó quyết định gửi email. Nó xây dựng một yêu cầu, chỉ định công cụ sẽ sử dụng (`send_email`) và các tham số cần thiết (người nhận, chủ đề, nội dung).
3.  **Giao tiếp Client:** MCP Client lấy yêu cầu đã xây dựng của LLM và gửi nó dưới dạng một cuộc gọi tiêu chuẩn hóa đến MCP Server thích hợp.
4.  **Thực thi Server:** MCP Server nhận yêu cầu. Nó xác thực client, xác nhận yêu cầu, và sau đó thực thi hành động đã chỉ định bằng cách giao tiếp với phần mềm cơ bản (ví dụ: gọi hàm `send()` của một API email).
5.  **Phản hồi và Cập nhật Ngữ cảnh:** Sau khi thực thi, MCP Server gửi một phản hồi tiêu chuẩn hóa trở lại cho MCP Client. Phản hồi này cho biết hành động có thành công hay không và bao gồm bất kỳ đầu ra liên quan nào (ví dụ: ID xác nhận cho email đã gửi). Client sau đó chuyển kết quả này trở lại cho LLM, cập nhật ngữ cảnh của nó và cho phép nó tiến hành bước tiếp theo của tác vụ.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

MCP mở rộng đáng kể khả năng của AI/LLM, làm cho chúng linh hoạt và mạnh mẽ hơn. Dưới đây là chín trường hợp sử dụng chính:

*   **Tích hợp Cơ sở dữ liệu:** MCP cho phép các LLM và tác nhân truy cập liền mạch và tương tác với dữ liệu có cấu trúc trong cơ sở dữ liệu. Ví dụ, sử dụng MCP Toolbox for Databases, một tác nhân có thể truy vấn các bộ dữ liệu Google BigQuery để lấy thông tin thời gian thực, tạo báo cáo hoặc cập nhật hồ sơ, tất cả đều được điều khiển bởi các lệnh ngôn ngữ tự nhiên.
*   **Điều phối Phương tiện Tạo sinh (Generative Media Orchestration):** MCP cho phép các tác nhân tích hợp với các dịch vụ phương tiện tạo sinh tiên tiến. Thông qua MCP Tools for Genmedia Services, một tác nhân có thể điều phối các quy trình làm việc liên quan đến Imagen của Google để tạo hình ảnh, Veo để tạo video, Chirp 3 HD cho giọng nói thực tế hoặc Lyria để sáng tác nhạc, cho phép tạo nội dung động trong các ứng dụng AI.
*   **Tương tác API Bên ngoài:** MCP cung cấp một cách tiêu chuẩn hóa để các LLM gọi và nhận phản hồi từ bất kỳ API bên ngoài nào. Điều này có nghĩa là một tác nhân có thể lấy dữ liệu thời tiết trực tiếp, lấy giá cổ phiếu, gửi email hoặc tương tác với các hệ thống CRM, mở rộng khả năng của nó vượt xa mô hình ngôn ngữ cốt lõi.
*   **Trích xuất Thông tin Dựa trên Suy luận:** Tận dụng các kỹ năng suy luận mạnh mẽ của LLM, MCP tạo điều kiện cho việc trích xuất thông tin hiệu quả, phụ thuộc vào truy vấn vượt qua các hệ thống tìm kiếm và truy xuất thông thường. Thay vì tìm kiếm truyền thống trả về toàn bộ tài liệu, một tác nhân có thể phân tích văn bản và trích xuất điều khoản, số liệu hoặc tuyên bố chính xác trả lời trực tiếp câu hỏi phức tạp của người dùng.
*   **Phát triển Công cụ Tùy chỉnh:** Các nhà phát triển có thể xây dựng các công cụ tùy chỉnh và hiển thị chúng qua máy chủ MCP (ví dụ: sử dụng FastMCP). Điều này cho phép các chức năng nội bộ chuyên biệt hoặc các hệ thống độc quyền được cung cấp cho các LLM và các tác nhân khác ở định dạng tiêu chuẩn hóa, dễ tiêu thụ mà không cần sửa đổi trực tiếp LLM.
*   **Giao tiếp LLM-đến-Ứng dụng Tiêu chuẩn hóa:** MCP đảm bảo một lớp giao tiếp nhất quán giữa các LLM và các ứng dụng mà chúng tương tác. Điều này làm giảm chi phí tích hợp, thúc đẩy khả năng tương tác giữa các nhà cung cấp LLM và ứng dụng máy chủ khác nhau, và đơn giản hóa việc phát triển các hệ thống agentic phức tạp.
*   **Điều phối Quy trình Làm việc Phức tạp:** Bằng cách kết hợp các công cụ và nguồn dữ liệu khác nhau được hiển thị qua MCP, các tác nhân có thể điều phối các quy trình làm việc đa bước, phức tạp cao. Một tác nhân có thể, ví dụ, truy xuất dữ liệu khách hàng từ cơ sở dữ liệu, tạo hình ảnh tiếp thị được cá nhân hóa, soạn thảo email phù hợp và sau đó gửi nó, tất cả bằng cách tương tác với các dịch vụ MCP khác nhau.
*   **Kiểm soát Thiết bị IoT:** MCP có thể tạo điều kiện cho sự tương tác của LLM với các thiết bị Internet vạn vật (IoT). Một tác nhân có thể sử dụng MCP để gửi lệnh đến các thiết bị gia dụng thông minh, cảm biến công nghiệp hoặc robot, cho phép kiểm soát và tự động hóa ngôn ngữ tự nhiên của các hệ thống vật lý.
*   **Tự động hóa Dịch vụ Tài chính:** Trong các dịch vụ tài chính, MCP có thể cho phép các LLM tương tác với các nguồn dữ liệu tài chính, nền tảng giao dịch hoặc hệ thống tuân thủ khác nhau. Một tác nhân có thể phân tích dữ liệu thị trường, thực hiện giao dịch, tạo tư vấn tài chính cá nhân hóa hoặc tự động hóa báo cáo quy định, tất cả trong khi duy trì giao tiếp an toàn và được tiêu chuẩn hóa.

Tóm lại, Model Context Protocol (MCP) cho phép các tác nhân truy cập thông tin thời gian thực từ cơ sở dữ liệu, API và tài nguyên web. Nó cũng cho phép các tác nhân thực hiện các hành động như gửi email, cập nhật hồ sơ, kiểm soát thiết bị và thực hiện các tác vụ phức tạp bằng cách tích hợp và xử lý dữ liệu từ nhiều nguồn khác nhau. Ngoài ra, MCP hỗ trợ các công cụ tạo phương tiện cho các ứng dụng AI.

## Ví dụ Code Thực hành với ADK

Phần này phác thảo cách kết nối với một máy chủ MCP cục bộ cung cấp các hoạt động hệ thống tệp, cho phép một tác nhân ADK tương tác với hệ thống tệp cục bộ.

### Thiết lập Tác nhân với MCPToolset

Để cấu hình một tác nhân cho tương tác hệ thống tệp, một tệp `agent.py` phải được tạo. `MCPToolset` được khởi tạo trong danh sách `tools` của đối tượng `LlmAgent`. Điều quan trọng là thay thế `"/path/to/your/folder"` trong danh sách `args` bằng đường dẫn tuyệt đối đến một thư mục trên hệ thống cục bộ mà máy chủ MCP có thể truy cập. Thư mục này sẽ là thư mục gốc cho các hoạt động hệ thống tệp được thực hiện bởi tác nhân.

```python
import os
from google.adk.agents import LlmAgent
from google.adk.tools.mcp_tool.mcp_toolset import MCPToolset, StdioServerParameters

# Tạo một đường dẫn tuyệt đối đáng tin cậy đến một thư mục có tên 'mcp_managed_files'
# trong cùng thư mục với script tác nhân này.
# Điều này đảm bảo tác nhân hoạt động ngay lập tức cho việc trình diễn.
# Đối với sản xuất, bạn sẽ trỏ đến một vị trí bền vững và an toàn hơn.
TARGET_FOLDER_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), "mcp_managed_files")

# Đảm bảo thư mục mục tiêu tồn tại trước khi tác nhân cần nó.
os.makedirs(TARGET_FOLDER_PATH, exist_ok=True)

root_agent = LlmAgent(
    model='gemini-2.0-flash',
    name='filesystem_assistant_agent',
    instruction=(
        'Giúp người dùng quản lý các tệp của họ. Bạn có thể liệt kê các tệp, đọc tệp và ghi tệp. '
        f'Bạn đang hoạt động trong thư mục sau: {TARGET_FOLDER_PATH}'
    ),
    tools=[
        MCPToolset(
            connection_params=StdioServerParameters(
                command='npx',
                args=[
                    "-y", # Đối số cho npx để tự động xác nhận cài đặt
                    "@modelcontextprotocol/server-filesystem",
                    TARGET_FOLDER_PATH, # Đây PHẢI là một đường dẫn tuyệt đối đến một thư mục.
                ],
            ),
            # Tùy chọn: Bạn có thể lọc các công cụ nào từ máy chủ MCP được hiển thị.
            # Ví dụ, để chỉ cho phép đọc:
            # tool_filter=['list_directory', 'read_file']
        )
    ],
)
```

`npx` (Node Package Execute), được đóng gói cùng với npm (Node Package Manager) phiên bản 5.2.0 trở lên, là một tiện ích cho phép thực thi trực tiếp các gói Node.js từ registry npm. Điều này loại bỏ nhu cầu cài đặt toàn cầu. Về bản chất, `npx` đóng vai trò là một trình chạy gói npm, và nó thường được sử dụng để chạy nhiều máy chủ MCP cộng đồng, được phân phối dưới dạng các gói Node.js.

Tạo một tệp `__init__.py` là cần thiết để đảm bảo tệp `agent.py` được nhận dạng là một phần của gói Python có thể khám phá cho Agent Development Kit (ADK). Tệp này nên nằm trong cùng thư mục với `agent.py`.

```python
# ./adk_agent_samples/mcp_agent/__init__.py
from . import agent
```

### Tạo một Máy chủ MCP với FastMCP

**FastMCP** là một framework Python cấp cao được thiết kế để hợp lý hóa việc phát triển các máy chủ MCP. Nó cung cấp một lớp trừu tượng đơn giản hóa sự phức tạp của giao thức, cho phép các nhà phát triển tập trung vào logic cốt lõi.

Thư viện cho phép định nghĩa nhanh các công cụ, tài nguyên và prompt bằng cách sử dụng các decorator Python đơn giản. Một lợi thế đáng kể là khả năng tự động tạo lược đồ (schema generation), diễn giải thông minh các chữ ký hàm Python, gợi ý kiểu (type hints) và chuỗi tài liệu (documentation strings) để xây dựng các thông số kỹ thuật giao diện mô hình AI cần thiết. Sự tự động hóa này giảm thiểu cấu hình thủ công và giảm lỗi của con người.

Để minh họa, hãy xem xét một công cụ "greet" cơ bản được cung cấp bởi máy chủ. Các tác nhân ADK và các ứng dụng khách MCP khác có thể tương tác với công cụ này bằng cách sử dụng HTTP sau khi nó hoạt động.

```python
# fastmcp_server.py
# Script này minh họa cách tạo một máy chủ MCP đơn giản bằng FastMCP.
# Nó hiển thị một công cụ duy nhất tạo ra lời chào.

# 1. Đảm bảo bạn đã cài đặt FastMCP:
# pip install fastmcp

from fastmcp import FastMCP

# Khởi tạo máy chủ FastMCP.
mcp_server = FastMCP()

# Định nghĩa một hàm công cụ đơn giản.
# Decorator `@mcp_server.tool` đăng ký hàm Python này như một công cụ MCP.
# Chuỗi tài liệu trở thành mô tả của công cụ cho LLM.
@mcp_server.tool
def greet(name: str) -> str:
    """
    Tạo một lời chào cá nhân hóa.

    Args:
        name: Tên của người cần chào.

    Returns:
        Một chuỗi lời chào.
    """
    return f"Xin chào, {name}! Rất vui được gặp bạn."

# Hoặc nếu bạn muốn chạy nó từ script:
if __name__ == "__main__":
    mcp_server.run(
        transport="sse", # Sử dụng Server-Sent Events cho kết nối HTTP
        host="127.0.0.1",
        port=8000
    )
```

Script Python này định nghĩa một hàm duy nhất có tên `greet`, nhận tên của một người và trả về một lời chào được cá nhân hóa. Decorator `@mcp_server.tool` bên trên hàm này tự động đăng ký nó như một công cụ mà AI hoặc chương trình khác có thể sử dụng. Chuỗi tài liệu và gợi ý kiểu của hàm được FastMCP sử dụng để cho Tác nhân biết cách công cụ hoạt động, những đầu vào nào nó cần và những gì nó sẽ trả về.

Khi script được thực thi, nó khởi động máy chủ FastMCP, lắng nghe các yêu cầu trên `localhost:8000`. Điều này làm cho hàm `greet` có sẵn như một dịch vụ mạng. Một tác nhân sau đó có thể được cấu hình để kết nối với máy chủ này và sử dụng công cụ `greet` để tạo lời chào như một phần của tác vụ lớn hơn.

### Tiêu thụ Máy chủ FastMCP với một Tác nhân ADK

Một tác nhân ADK có thể được thiết lập như một ứng dụng khách MCP để sử dụng máy chủ FastMCP đang chạy. Điều này yêu cầu cấu hình `HttpServerParameters` với địa chỉ mạng của máy chủ FastMCP, thường là `http://localhost:8000/sse`.

Một tham số `tool_filter` có thể được bao gồm để hạn chế việc sử dụng công cụ của tác nhân đối với các công cụ cụ thể được cung cấp bởi máy chủ, chẳng hạn như 'greet'. Khi được nhắc với một yêu cầu như "Greet John Doe", LLM nhúng của tác nhân xác định công cụ 'greet' có sẵn qua MCP, gọi nó với đối số "John Doe" và trả về phản hồi của máy chủ. Quá trình này chứng minh sự tích hợp của các công cụ do người dùng xác định được hiển thị qua MCP với một tác nhân ADK.

```python
# ./adk_agent_samples/fastmcp_client_agent/agent.py
import os
from google.adk.agents import LlmAgent
from google.adk.tools.mcp_tool.mcp_toolset import MCPToolset, HttpServerParameters

# Định nghĩa địa chỉ của máy chủ FastMCP.
# Đảm bảo fastmcp_server.py của bạn (được định nghĩa trước đó) đang chạy trên cổng này.
# Lưu ý: FastMCP sử dụng SSE endpoint tại /sse theo mặc định cho transport="sse"
FASTMCP_SERVER_URL = "http://localhost:8000/sse"

root_agent = LlmAgent(
    model='gemini-2.0-flash', # Hoặc mô hình ưa thích của bạn
    name='fastmcp_greeter_agent',
    instruction='Bạn là một trợ lý thân thiện có thể chào mọi người bằng tên của họ. Sử dụng công cụ "greet".',
    tools=[
        MCPToolset(
            connection_params=HttpServerParameters(
                url=FASTMCP_SERVER_URL,
            ),
            # Tùy chọn: Lọc các công cụ nào từ máy chủ MCP được hiển thị
            # Đối với ví dụ này, chúng ta chỉ mong đợi 'greet'
            tool_filter=['greet']
        )
    ],
)
```

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Để hoạt động như các tác nhân hiệu quả, LLM phải vượt ra ngoài việc tạo văn bản đơn giản. Chúng yêu cầu khả năng tương tác với môi trường bên ngoài để truy cập dữ liệu hiện tại và sử dụng phần mềm bên ngoài. Nếu không có phương pháp giao tiếp tiêu chuẩn hóa, mỗi tích hợp giữa LLM và công cụ bên ngoài hoặc nguồn dữ liệu sẽ trở thành một nỗ lực tùy chỉnh, phức tạp và không thể tái sử dụng. Cách tiếp cận đặc biệt này cản trở khả năng mở rộng và làm cho việc xây dựng các hệ thống AI phức tạp, liên kết với nhau trở nên khó khăn và không hiệu quả.
*   **Tại sao:** Model Context Protocol (MCP) cung cấp một giải pháp tiêu chuẩn hóa bằng cách hoạt động như một giao diện chung giữa LLM và các hệ thống bên ngoài. Nó thiết lập một giao thức mở, được tiêu chuẩn hóa xác định cách các khả năng bên ngoài được khám phá và sử dụng. Hoạt động trên mô hình client-server, MCP cho phép các máy chủ hiển thị các công cụ, tài nguyên dữ liệu và prompt tương tác cho bất kỳ ứng dụng khách tuân thủ nào. Các ứng dụng được hỗ trợ bởi LLM hoạt động như các ứng dụng khách này, khám phá động và tương tác với các tài nguyên có sẵn theo cách có thể dự đoán được. Cách tiếp cận tiêu chuẩn hóa này thúc đẩy một hệ sinh thái các thành phần có thể tương tác và tái sử dụng, đơn giản hóa đáng kể việc phát triển các quy trình làm việc agentic phức tạp.
*   **Quy tắc ngón tay cái:** Sử dụng Model Context Protocol (MCP) khi xây dựng các hệ thống agentic phức tạp, có thể mở rộng hoặc cấp doanh nghiệp cần tương tác với một tập hợp đa dạng và đang phát triển các công cụ bên ngoài, nguồn dữ liệu và API. Nó lý tưởng khi khả năng tương tác giữa các LLM và công cụ khác nhau là ưu tiên, và khi các tác nhân yêu cầu khả năng khám phá động các khả năng mới mà không cần triển khai lại. Đối với các ứng dụng đơn giản hơn với số lượng hàm được xác định trước cố định và hạn chế, việc gọi hàm công cụ trực tiếp có thể là đủ.

## Những Điểm Chính (Key Takeaways)

*   **Model Context Protocol (MCP)** là một tiêu chuẩn mở tạo điều kiện thuận lợi cho giao tiếp tiêu chuẩn hóa giữa LLM và các ứng dụng bên ngoài, nguồn dữ liệu và công cụ.
*   Nó sử dụng kiến trúc **client-server**, xác định các phương pháp để hiển thị và tiêu thụ tài nguyên, prompt và công cụ.
*   **Agent Development Kit (ADK)** hỗ trợ cả việc sử dụng các máy chủ MCP hiện có và hiển thị các công cụ ADK thông qua một máy chủ MCP.
*   **FastMCP** đơn giản hóa việc phát triển và quản lý các máy chủ MCP, đặc biệt là để hiển thị các công cụ được triển khai bằng Python.
*   **MCP Tools for Genmedia Services** cho phép các tác nhân tích hợp với các khả năng phương tiện tạo sinh của Google Cloud (Imagen, Veo, Chirp 3 HD, Lyria).
*   MCP cho phép các LLM và tác nhân tương tác với các hệ thống thế giới thực, truy cập thông tin động và thực hiện các hành động vượt ra ngoài việc tạo văn bản.

## Kết luận

Model Context Protocol (MCP) là một tiêu chuẩn mở tạo điều kiện thuận lợi cho giao tiếp giữa các Mô hình Ngôn ngữ Lớn (LLM) và các hệ thống bên ngoài. Nó sử dụng kiến trúc client-server, cho phép LLM truy cập tài nguyên, sử dụng prompt và thực thi các hành động thông qua các công cụ được tiêu chuẩn hóa. MCP cho phép LLM tương tác với cơ sở dữ liệu, quản lý quy trình làm việc phương tiện tạo sinh, kiểm soát thiết bị IoT và tự động hóa các dịch vụ tài chính. Các ví dụ thực tế minh họa việc thiết lập các tác nhân để giao tiếp với các máy chủ MCP, bao gồm các máy chủ hệ thống tệp và máy chủ được xây dựng bằng FastMCP, minh họa sự tích hợp của nó với Agent Development Kit (ADK). MCP là một thành phần quan trọng để phát triển các tác nhân AI tương tác mở rộng vượt ra ngoài các khả năng ngôn ngữ cơ bản.

## Tài liệu tham khảo
1.  Model Context Protocol (MCP) Documentation (Latest): https://google.github.io/adk-docs/mcp/
2.  FastMCP Documentation: https://github.com/jlowin/fastmcp
3.  MCP Tools for Genmedia Services: https://google.github.io/adk-docs/mcp/#mcp-servers-for-google-cloud-genmedia
4.  MCP Toolbox for Databases Documentation (Latest): https://google.github.io/adk-docs/mcp/databases/