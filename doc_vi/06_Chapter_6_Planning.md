# Chương 6: Planning (Lập Kế Hoạch)

## Tổng quan về Mẫu thiết kế Planning

Hành vi thông minh thường đòi hỏi nhiều hơn là chỉ phản ứng lại đầu vào tức thời. Nó yêu cầu tầm nhìn xa, khả năng phân rã các tác vụ phức tạp thành các bước nhỏ hơn, dễ quản lý hơn, và lập chiến lược để đạt được kết quả mong muốn. Đây là lúc mẫu **Planning (Lập Kế Hoạch)** phát huy tác dụng. Cốt lõi của nó, lập kế hoạch là khả năng của một tác nhân hoặc một hệ thống tác nhân để xây dựng một chuỗi các hành động nhằm chuyển từ trạng thái ban đầu sang trạng thái mục tiêu.

Trong bối cảnh AI, việc coi một tác nhân lập kế hoạch như một chuyên gia mà bạn ủy thác một mục tiêu phức tạp là hữu ích. Khi bạn yêu cầu nó "tổ chức một buổi họp mặt công ty", bạn đang xác định *cái gì* - mục tiêu và các ràng buộc của nó - nhưng không phải *làm thế nào*. Nhiệm vụ cốt lõi của tác nhân là tự động vạch ra một lộ trình đến mục tiêu đó. Trước tiên, nó phải hiểu trạng thái ban đầu (ví dụ: ngân sách, số lượng người tham gia, ngày mong muốn) và trạng thái mục tiêu (một buổi họp mặt đã được đặt thành công), và sau đó khám phá chuỗi hành động tối ưu để kết nối chúng. Kế hoạch không được biết trước; nó được tạo ra để đáp ứng yêu cầu.

Một đặc điểm nổi bật của quá trình này là khả năng thích ứng. Một kế hoạch ban đầu chỉ là một điểm khởi đầu, không phải là một kịch bản cứng nhắc. Sức mạnh thực sự của tác nhân là khả năng kết hợp thông tin mới và điều khiển dự án vượt qua các trở ngại. Ví dụ, nếu địa điểm ưa thích không còn trống hoặc nhà cung cấp dịch vụ ăn uống đã kín lịch, một tác nhân có khả năng sẽ không thất bại đơn giản. Nó sẽ thích nghi. Nó ghi nhận ràng buộc mới, đánh giá lại các lựa chọn của nó và lập một kế hoạch mới, có thể bằng cách đề xuất các địa điểm hoặc ngày thay thế.

Huy nhiên, điều quan trọng là phải nhận ra sự đánh đổi giữa tính linh hoạt và khả năng dự đoán. Lập kế hoạch động là một công cụ cụ thể, không phải là một giải pháp phổ quát. Khi giải pháp của một vấn đề đã được hiểu rõ và có thể lặp lại, việc ràng buộc tác nhân vào một quy trình làm việc cố định, được xác định trước sẽ hiệu quả hơn. Cách tiếp cận này giới hạn quyền tự chủ của tác nhân để giảm sự không chắc chắn và rủi ro hành vi không thể đoán trước, đảm bảo kết quả đáng tin cậy và nhất quán. Do đó, quyết định sử dụng tác nhân lập kế hoạch so với tác nhân thực thi tác vụ đơn giản phụ thuộc vào một câu hỏi duy nhất: *cách thực hiện* có cần được khám phá hay đã biết?

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Mẫu Planning là một quy trình tính toán cốt lõi trong các hệ thống tự trị, cho phép một tác nhân tổng hợp một chuỗi các hành động để đạt được một mục tiêu cụ thể, đặc biệt trong các môi trường động hoặc phức tạp. Quá trình này biến một mục tiêu cấp cao thành một kế hoạch có cấu trúc gồm các bước riêng biệt, có thể thực thi được.

*   **Tự động hóa tác vụ thủ tục:** Trong các lĩnh vực như tự động hóa tác vụ thủ tục, lập kế hoạch được sử dụng để điều phối các quy trình làm việc phức tạp. Ví dụ, một quy trình kinh doanh như giới thiệu nhân viên mới có thể được phân rã thành một chuỗi các tác vụ phụ có hướng, chẳng hạn như tạo tài khoản hệ thống, giao các mô-đun đào tạo và phối hợp với các phòng ban khác nhau. Tác nhân tạo ra một kế hoạch để thực thi các bước này theo thứ tự logic, gọi các công cụ cần thiết hoặc tương tác với các hệ thống khác nhau để quản lý các phụ thuộc.
*   **Robot học và Điều hướng Tự trị:** Trong robot học và điều hướng tự trị, lập kế hoạch là nền tảng cho việc di chuyển trong không gian trạng thái. Một hệ thống, dù là robot vật lý hay thực thể ảo, phải tạo ra một đường dẫn hoặc chuỗi hành động để chuyển đổi từ trạng thái ban đầu sang trạng thái mục tiêu. Điều này liên quan đến việc tối ưu hóa các số liệu như thời gian hoặc mức tiêu thụ năng lượng trong khi tuân thủ các ràng buộc môi trường, như tránh chướng ngại vật hoặc tuân thủ các quy định giao thông.
*   **Tổng hợp Thông tin có Cấu trúc:** Mẫu này cũng rất quan trọng để tổng hợp thông tin có cấu trúc. Khi được giao nhiệm vụ tạo ra một đầu ra phức tạp như báo cáo nghiên cứu, tác nhân có thể xây dựng một kế hoạch bao gồm các giai đoạn riêng biệt để thu thập thông tin, tóm tắt dữ liệu, cấu trúc nội dung và tinh chỉnh lặp lại. Tương tự, trong các kịch bản hỗ trợ khách hàng liên quan đến việc giải quyết vấn đề đa bước, tác nhân có thể tạo và tuân theo một kế hoạch có hệ thống để chẩn đoán, triển khai giải pháp và leo thang.

Về bản chất, mẫu Planning cho phép tác nhân vượt ra ngoài các hành động phản ứng đơn giản để có hành vi định hướng mục tiêu. Nó cung cấp khuôn khổ logic cần thiết để giải quyết các vấn đề yêu cầu một chuỗi các hoạt động phụ thuộc lẫn nhau một cách mạch lạc.

## Ví dụ Code Thực hành (CrewAI)

Phần sau đây sẽ minh họa việc triển khai mẫu Planner sử dụng framework CrewAI. Mẫu này liên quan đến một tác nhân trước tiên xây dựng một kế hoạch đa bước để giải quyết một truy vấn phức tạp và sau đó thực thi kế hoạch đó theo trình tự.

```python
# Cài đặt: pip install crewai langchain-openai python-dotenv

import os
from dotenv import load_dotenv
from crewai import Agent, Task, Crew, Process
from langchain_openai import ChatOpenAI # Sử dụng langchain_openai nếu có, hoặc một LLM khác

# Tải biến môi trường từ file .env để bảo mật
load_dotenv()

# 1. Định nghĩa rõ ràng mô hình ngôn ngữ
# LƯU Ý: Đảm bảo có GOOGLE_API_KEY được đặt trong biến môi trường
llm = ChatOpenAI(model="gpt-4o-mini", temperature=0.7) # Hoặc bất kỳ LLM nào bạn chọn

# 2. Định nghĩa một tác nhân rõ ràng và tập trung
planner_writer_agent = Agent(
    role='Article Planner and Writer',
    goal='Lập kế hoạch và sau đó viết một bản tóm tắt súc tích, hấp dẫn về một chủ đề cụ thể.',
    backstory=(
        'Bạn là một người viết kỹ thuật và chiến lược nội dung chuyên nghiệp. '
        'Điểm mạnh của bạn nằm ở việc tạo ra một kế hoạch rõ ràng, có thể thực hiện được trước khi viết, '
        'đảm bảo bản tóm tắt cuối cùng vừa mang tính thông tin vừa dễ hiểu.'
    ),
    verbose=True,
    allow_delegation=False, # Không cho phép ủy quyền trong tác vụ này
    llm=llm # Gán LLM cụ thể cho tác nhân
)

# 3. Định nghĩa một tác vụ với đầu ra có cấu trúc và cụ thể hơn
topic = "Tầm quan trọng của Học tăng cường trong AI"
high_level_task = Task(
    description=(
        f"1. Tạo một kế hoạch dạng gạch đầu dòng cho bản tóm tắt về chủ đề: '{topic}'.\n"
        f"2. Viết bản tóm tắt dựa trên kế hoạch của bạn, giữ nó khoảng 200 từ."
    ),
    expected_output=(
        "Một báo cáo cuối cùng chứa hai phần riêng biệt:\n\n"
        "### Kế hoạch\n"
        "- Một danh sách gạch đầu dòng phác thảo các điểm chính của bản tóm tắt.\n\n"
        "### Tóm tắt\n"
        "- Một bản tóm tắt súc tích và có cấu trúc rõ ràng về chủ đề."
    ),
    agent=planner_writer_agent,
)

# 4. Xây dựng Crew với một quy trình rõ ràng
crew = Crew(
    agents=[planner_writer_agent],
    tasks=[high_level_task],
    process=Process.sequential, # Thực thi các tác vụ theo trình tự
    verbose=2 # Đặt độ chi tiết nhật ký thành 2 để có nhật ký thực thi crew chi tiết
)

# Thực thi tác vụ
print("## Đang chạy quy trình tạo kế hoạch và viết bài của Crew ##")
result = crew.kickoff()

print("\n\n---\n## Kết quả Tác vụ ##\n---")
print(result)

```

Mã này sử dụng thư viện CrewAI để tạo một tác nhân AI lập kế hoạch và viết bản tóm tắt về một chủ đề nhất định. Nó bắt đầu bằng cách xác định rõ ràng một mô hình ngôn ngữ (`ChatOpenAI`) và một tác nhân (`planner_writer_agent`) với vai trò, mục tiêu và backstory cụ thể nhấn mạnh chuyên môn của nó trong lập kế hoạch và viết kỹ thuật. Một tác vụ (`high_level_task`) được định nghĩa với mô tả rõ ràng và định dạng cụ thể cho đầu ra mong đợi, yêu cầu tác nhân trước tiên tạo một kế hoạch và sau đó viết một bản tóm tắt về chủ đề. Một `Crew` được tập hợp với tác nhân và tác vụ, được đặt để xử lý chúng theo trình tự (`Process.sequential`). Cuối cùng, phương thức `crew.kickoff()` được gọi để thực thi tác vụ đã xác định và in kết quả.

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các tác nhân AI hoạt động trong các môi trường phức tạp thường phải đối mặt với vô số hành động tiềm năng, các mục tiêu xung đột và tài nguyên hữu hạn. Nếu không có một phương pháp rõ ràng để xác định hành động tiếp theo, các tác nhân này có nguy cơ trở nên kém hiệu quả và không hiệu quả. Điều này có thể dẫn đến sự chậm trễ đáng kể trong hoạt động hoặc thất bại hoàn toàn trong việc đạt được các mục tiêu chính. Thách thức cốt lõi là quản lý số lượng lớn các lựa chọn này để đảm bảo tác nhân hành động có mục đích và logic.
*   **Tại sao:** Mẫu Planning cung cấp một giải pháp tiêu chuẩn hóa cho vấn đề này bằng cách yêu cầu hệ thống Agentic trước tiên tạo ra một kế hoạch mạch lạc để đạt được mục tiêu. Nó liên quan đến việc phân rã một mục tiêu cấp cao thành một chuỗi các bước hoặc mục tiêu phụ nhỏ hơn, có thể thực hiện được. Điều này cho phép hệ thống quản lý các quy trình làm việc phức tạp, điều phối các công cụ khác nhau và xử lý các phụ thuộc theo thứ tự logic. LLM đặc biệt phù hợp cho việc này, vì chúng có thể tạo ra các kế hoạch hợp lý và hiệu quả dựa trên dữ liệu đào tạo rộng lớn và sự hiểu biết về các tác vụ.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này khi yêu cầu của người dùng quá phức tạp để được xử lý bằng một hành động hoặc công cụ duy nhất. Nó lý tưởng để tự động hóa các quy trình đa bước, chẳng hạn như tạo báo cáo nghiên cứu chi tiết, giới thiệu nhân viên mới hoặc thực hiện phân tích cạnh tranh. Áp dụng mẫu Planning bất cứ khi nào một tác vụ yêu cầu một chuỗi các hoạt động phụ thuộc lẫn nhau để đạt được kết quả cuối cùng, tổng hợp.

## Những Điểm Chính (Key Takeaways)

*   Planning cho phép các tác nhân phân rã các mục tiêu phức tạp thành các bước tuần tự, có thể thực hiện được.
*   Nó rất cần thiết để xử lý các tác vụ đa bước, tự động hóa quy trình làm việc và điều hướng các môi trường phức tạp.
*   LLM có thể thực hiện lập kế hoạch bằng cách tạo ra các cách tiếp cận từng bước dựa trên mô tả tác vụ.
*   Việc nhắc nhở rõ ràng hoặc thiết kế các tác vụ để yêu cầu các bước lập kế hoạch khuyến khích hành vi này trong các framework tác nhân.
*   Google Deep Research là một tác nhân phân tích các nguồn được lấy thay mặt chúng ta bằng cách sử dụng Google Search làm công cụ. Nó phản ánh, lập kế hoạch và thực thi.

## Kết luận

Tóm lại, mẫu Planning là một thành phần nền tảng nâng cấp các hệ thống Agentic từ những công cụ phản ứng đơn giản thành những người thực thi chiến lược, định hướng mục tiêu. Các mô hình ngôn ngữ lớn hiện đại cung cấp khả năng cốt lõi cho việc này, tự động phân rã các mục tiêu cấp cao thành các bước mạch lạc, có thể thực hiện được. Mẫu này có thể mở rộng từ việc thực thi tác vụ tuần tự đơn giản, như được minh họa bởi tác nhân CrewAI tạo và tuân theo kế hoạch viết, đến các hệ thống phức tạp và năng động hơn. Tác nhân Google DeepResearch minh họa ứng dụng nâng cao này, tạo ra các kế hoạch nghiên cứu lặp lại thích nghi và phát triển dựa trên việc thu thập thông tin liên tục. Cuối cùng, lập kế hoạch cung cấp cầu nối thiết yếu giữa ý định của con người và việc thực thi tự động cho các vấn đề phức tạp. Bằng cách cấu trúc cách tiếp cận giải quyết vấn đề, mẫu này cho phép các tác nhân quản lý các quy trình làm việc phức tạp và mang lại kết quả toàn diện, tổng hợp.

## Tài liệu tham khảo
1. Google DeepResearch (Gemini Feature): [gemini.google.com](https://gemini.google.com/)
2. OpenAI, Introducing deep research: [https://openai.com/index/introducing-deep-research/](https://openai.com/index/introducing-deep-research/)
3. Perplexity, Introducing Perplexity Deep Research: [https://www.perplexity.ai/hub/blog/introducing-perplexity-deep-research](https://www.perplexity.ai/hub/blog/introducing-perplexity-deep-research)