# Chương 8: Memory Management (Quản Lý Bộ Nhớ)

Quản lý bộ nhớ hiệu quả là rất quan trọng để các tác nhân thông minh lưu giữ thông tin. Các tác nhân yêu cầu các loại bộ nhớ khác nhau, giống như con người, để hoạt động hiệu quả. Chương này đi sâu vào quản lý bộ nhớ, giải quyết cụ thể các yêu cầu về bộ nhớ tức thời (ngắn hạn) và bền vững (dài hạn) của các tác nhân.

Trong các hệ thống tác nhân, bộ nhớ đề cập đến khả năng của một tác nhân trong việc lưu giữ và sử dụng thông tin từ các tương tác, quan sát và trải nghiệm học tập trong quá khứ. Khả năng này cho phép các tác nhân đưa ra các quyết định sáng suốt, duy trì ngữ cảnh hội thoại và cải thiện theo thời gian. Bộ nhớ tác nhân thường được phân loại thành hai loại chính:

1.  **Bộ nhớ Ngắn hạn (Short-Term Memory - Contextual Memory):** Tương tự như bộ nhớ làm việc (working memory), loại này chứa thông tin hiện đang được xử lý hoặc mới được truy cập. Đối với các tác nhân sử dụng các mô hình ngôn ngữ lớn (LLM), bộ nhớ ngắn hạn chủ yếu tồn tại trong **cửa sổ ngữ cảnh (context window)**. Cửa sổ này chứa các tin nhắn gần đây, phản hồi của tác nhân, kết quả sử dụng công cụ và phản chiếu của tác nhân từ tương tác hiện tại, tất cả đều thông báo cho các phản hồi và hành động tiếp theo của LLM. Cửa sổ ngữ cảnh có dung lượng hạn chế, hạn chế lượng thông tin gần đây mà một tác nhân có thể truy cập trực tiếp. Quản lý bộ nhớ ngắn hạn hiệu quả bao gồm việc giữ thông tin phù hợp nhất trong không gian hạn chế này, có thể thông qua các kỹ thuật như tóm tắt các đoạn hội thoại cũ hoặc nhấn mạnh các chi tiết chính. Sự ra đời của các mô hình có cửa sổ ngữ cảnh dài ('long context' windows) chỉ đơn giản là mở rộng kích thước của bộ nhớ ngắn hạn này, cho phép lưu giữ nhiều thông tin hơn trong một tương tác duy nhất. Tuy nhiên, ngữ cảnh này vẫn là tạm thời và bị mất khi phiên kết thúc, và có thể tốn kém và không hiệu quả để xử lý mỗi lần. Do đó, các tác nhân yêu cầu các loại bộ nhớ riêng biệt để đạt được sự bền vững thực sự, nhớ lại thông tin từ các tương tác trong quá khứ và xây dựng một cơ sở kiến thức lâu dài.

2.  **Bộ nhớ Dài hạn (Long-Term Memory - Persistent Memory):** Loại này hoạt động như một kho lưu trữ thông tin mà các tác nhân cần giữ lại qua các tương tác, tác vụ hoặc khoảng thời gian dài khác nhau, giống như các cơ sở tri thức dài hạn. Dữ liệu thường được lưu trữ bên ngoài môi trường xử lý tức thời của tác nhân, thường là trong cơ sở dữ liệu, đồ thị tri thức (knowledge graphs) hoặc **cơ sở dữ liệu vector (vector databases)**. Trong cơ sở dữ liệu vector, thông tin được chuyển đổi thành các vector số và được lưu trữ, cho phép các tác nhân truy xuất dữ liệu dựa trên **độ tương đồng ngữ nghĩa (semantic similarity)** thay vì khớp từ khóa chính xác, một quá trình được gọi là tìm kiếm ngữ nghĩa (semantic search). Khi một tác nhân cần thông tin từ bộ nhớ dài hạn, nó truy vấn bộ nhớ ngoài, truy xuất dữ liệu liên quan và tích hợp nó vào ngữ cảnh ngắn hạn để sử dụng ngay lập tức, do đó kết hợp kiến thức trước đó với tương tác hiện tại.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Quản lý bộ nhớ là rất quan trọng để các tác nhân theo dõi thông tin và hoạt động thông minh theo thời gian. Điều này rất cần thiết để các tác nhân vượt qua khả năng trả lời câu hỏi cơ bản. Các ứng dụng bao gồm:

*   **Chatbot và AI Hội thoại:** Duy trì luồng hội thoại dựa vào bộ nhớ ngắn hạn. Chatbot yêu cầu ghi nhớ các đầu vào trước đó của người dùng để cung cấp các phản hồi mạch lạc. Bộ nhớ dài hạn cho phép chatbot nhớ lại sở thích của người dùng, các vấn đề trong quá khứ hoặc các cuộc thảo luận trước đó, cung cấp các tương tác được cá nhân hóa và liên tục.
*   **Tác nhân Định hướng Tác vụ (Task-Oriented Agents):** Các tác nhân quản lý các tác vụ đa bước cần bộ nhớ ngắn hạn để theo dõi các bước trước đó, tiến trình hiện tại và mục tiêu tổng thể. Thông tin này có thể nằm trong ngữ cảnh của tác vụ hoặc bộ nhớ tạm thời. Bộ nhớ dài hạn rất quan trọng để truy cập dữ liệu liên quan đến người dùng cụ thể không có trong ngữ cảnh tức thời.
*   **Trải nghiệm Cá nhân hóa:** Các tác nhân cung cấp các tương tác phù hợp sử dụng bộ nhớ dài hạn để lưu trữ và truy xuất sở thích của người dùng, hành vi trong quá khứ và thông tin cá nhân. Điều này cho phép các tác nhân điều chỉnh phản hồi và đề xuất của họ.
*   **Học tập và Cải thiện:** Các tác nhân có thể tinh chỉnh hiệu suất của mình bằng cách học hỏi từ các tương tác trong quá khứ. Các chiến lược thành công, sai lầm và thông tin mới được lưu trữ trong bộ nhớ dài hạn, tạo điều kiện cho các thích ứng trong tương lai. Các tác nhân học tăng cường (Reinforcement learning agents) lưu trữ các chiến lược hoặc kiến thức đã học theo cách này.
*   **Truy xuất Thông tin (RAG):** Các tác nhân được thiết kế để trả lời các câu hỏi truy cập vào một cơ sở tri thức, bộ nhớ dài hạn của chúng, thường được triển khai trong Retrieval Augmented Generation (RAG). Tác nhân truy xuất các tài liệu hoặc dữ liệu liên quan để thông báo cho các phản hồi của nó.
*   **Hệ thống Tự trị:** Robot hoặc xe tự lái yêu cầu bộ nhớ cho bản đồ, tuyến đường, vị trí đối tượng và các hành vi đã học. Điều này liên quan đến bộ nhớ ngắn hạn cho môi trường xung quanh tức thời và bộ nhớ dài hạn cho kiến thức môi trường chung.

Bộ nhớ cho phép các tác nhân duy trì lịch sử, học hỏi, cá nhân hóa các tương tác và quản lý các vấn đề phức tạp, phụ thuộc vào thời gian.

## Ví dụ Code Thực hành: Quản lý Bộ nhớ trong Google Agent Developer Kit (ADK)

Google Agent Developer Kit (ADK) cung cấp một phương pháp có cấu trúc để quản lý ngữ cảnh và bộ nhớ, bao gồm các thành phần cho ứng dụng thực tế. Việc nắm vững Session (Phiên), State (Trạng thái) và Memory (Bộ nhớ) của ADK là rất quan trọng để xây dựng các tác nhân cần lưu giữ thông tin.

Giống như trong các tương tác của con người, các tác nhân yêu cầu khả năng nhớ lại các trao đổi trước đó để thực hiện các cuộc hội thoại mạch lạc và tự nhiên. ADK đơn giản hóa việc quản lý ngữ cảnh thông qua ba khái niệm cốt lõi và các dịch vụ liên quan của chúng.

Mỗi tương tác với một tác nhân có thể được coi là một luồng hội thoại duy nhất. Các tác nhân có thể cần truy cập dữ liệu từ các tương tác trước đó. ADK cấu trúc điều này như sau:

*   **Session (Phiên):** Một luồng trò chuyện riêng lẻ ghi lại các tin nhắn và hành động (Events) cho tương tác cụ thể đó, cũng lưu trữ dữ liệu tạm thời (State) liên quan đến cuộc hội thoại đó.
*   **State (session.state):** Dữ liệu được lưu trữ trong một Session, chứa thông tin chỉ liên quan đến luồng trò chuyện hiện tại, đang hoạt động.
*   **Memory (Bộ nhớ):** Một kho lưu trữ thông tin có thể tìm kiếm được lấy từ nhiều cuộc trò chuyện trong quá khứ hoặc các nguồn bên ngoài, đóng vai trò là tài nguyên để truy xuất dữ liệu ngoài cuộc hội thoại ngay lập tức.

ADK cung cấp các dịch vụ chuyên dụng để quản lý các thành phần quan trọng cần thiết cho việc xây dựng các hệ thống tác nhân phức tạp, có trạng thái và nhận thức ngữ cảnh. `SessionService` quản lý các luồng trò chuyện (đối tượng Session) bằng cách xử lý việc khởi tạo, ghi lại và chấm dứt của chúng, trong khi `MemoryService` giám sát việc lưu trữ và truy xuất kiến thức dài hạn (Memory).

Cả `SessionService` và `MemoryService` đều cung cấp các tùy chọn cấu hình khác nhau, cho phép người dùng chọn phương thức lưu trữ dựa trên nhu cầu ứng dụng. Các tùy chọn trong bộ nhớ (in-memory) có sẵn cho mục đích thử nghiệm, mặc dù dữ liệu sẽ không tồn tại sau khi khởi động lại. Để lưu trữ bền vững và khả năng mở rộng, ADK cũng hỗ trợ cơ sở dữ liệu và các dịch vụ dựa trên đám mây.

### Session: Theo dõi từng cuộc trò chuyện

Một đối tượng `Session` trong ADK được thiết kế để theo dõi và quản lý các luồng trò chuyện riêng lẻ. Khi bắt đầu một cuộc hội thoại với một tác nhân, `SessionService` tạo ra một đối tượng Session. Đối tượng này đóng gói tất cả dữ liệu liên quan đến một luồng hội thoại cụ thể, bao gồm các định danh duy nhất, bản ghi thời gian của các sự kiện (`Events`), một khu vực lưu trữ cho dữ liệu tạm thời cụ thể của phiên được gọi là `state`.

Ví dụ sử dụng `InMemorySessionService` cho phát triển cục bộ:

```python
# Ví dụ: Sử dụng InMemorySessionService
# Phù hợp cho phát triển và thử nghiệm cục bộ nơi dữ liệu
# không cần bền vững qua các lần khởi động lại ứng dụng.
from google.adk.sessions import InMemorySessionService
session_service = InMemorySessionService()
```

Ví dụ sử dụng `DatabaseSessionService` cho lưu trữ bền vững:

```python
# Ví dụ: Sử dụng DatabaseSessionService
# Phù hợp cho sản xuất hoặc phát triển yêu cầu lưu trữ bền vững.
# Yêu cầu cấu hình URL cơ sở dữ liệu (ví dụ: SQLite, PostgreSQL).
from google.adk.sessions import DatabaseSessionService
# Ví dụ sử dụng file SQLite cục bộ:
db_url = "sqlite:///./my_agent_data.db"
session_service = DatabaseSessionService(db_url=db_url)
```

Ngoài ra, còn có `VertexAiSessionService` sử dụng cơ sở hạ tầng Vertex AI cho sản xuất có thể mở rộng trên Google Cloud.

### State: Bảng nháp của Phiên

Trong ADK, mỗi Session bao gồm một thành phần `state` giống như bộ nhớ làm việc tạm thời của tác nhân trong suốt thời gian của cuộc hội thoại cụ thể đó. Trong khi `session.events` ghi lại toàn bộ lịch sử trò chuyện, `session.state` lưu trữ và cập nhật các điểm dữ liệu động liên quan đến cuộc trò chuyện đang hoạt động.

Về cơ bản, `session.state` hoạt động như một từ điển (dictionary), lưu trữ dữ liệu dưới dạng các cặp khóa-giá trị. Chức năng cốt lõi của nó là cho phép tác nhân giữ lại và quản lý các chi tiết cần thiết cho đối thoại mạch lạc, chẳng hạn như sở thích của người dùng, tiến trình tác vụ, thu thập dữ liệu gia tăng hoặc các cờ điều kiện ảnh hưởng đến các hành động tiếp theo của tác nhân.

Tổ chức trạng thái có thể đạt được bằng cách sử dụng các tiền tố khóa (key prefixes) để xác định phạm vi và tính bền vững của dữ liệu. Các khóa không có tiền tố là dành riêng cho phiên (session-specific).
*   `user:` tiền tố liên kết dữ liệu với một ID người dùng trên tất cả các phiên.
*   `app:` tiền tố chỉ định dữ liệu được chia sẻ giữa tất cả người dùng của ứng dụng.
*   `temp:` tiền tố chỉ ra dữ liệu chỉ hợp lệ cho lượt xử lý hiện tại và không được lưu trữ bền vững.

Có hai cách chính để cập nhật state:

1.  **Cách đơn giản: Sử dụng `output_key` (cho Phản hồi văn bản của Tác nhân):** Đây là phương pháp dễ nhất nếu bạn chỉ muốn lưu phản hồi văn bản cuối cùng của tác nhân trực tiếp vào state. Khi thiết lập `LlmAgent`, bạn chỉ định `output_key`.

    ```python
    # Định nghĩa một LlmAgent với output_key.
    greeting_agent = LlmAgent(
        name="Greeter",
        model="gemini-2.0-flash",
        instruction="Tạo một lời chào ngắn gọn, thân thiện.",
        output_key="last_greeting"
    )
    ```

2.  **Cách tiêu chuẩn: Sử dụng `EventActions.state_delta` (cho các Cập nhật Phức tạp hơn):** Đối với các trường hợp bạn cần thực hiện những việc phức tạp hơn – như cập nhật nhiều khóa cùng lúc, lưu những thứ không chỉ là văn bản, nhắm mục tiêu các phạm vi cụ thể như `user:` hoặc `app:`, hoặc thực hiện các cập nhật không gắn liền với phản hồi văn bản cuối cùng của tác nhân – bạn sẽ xây dựng thủ công một từ điển các thay đổi trạng thái của mình (`state_delta`) và đưa nó vào `EventActions` của Sự kiện bạn đang thêm vào.

    Ví dụ về việc cập nhật state từ bên trong một công cụ:

    ```python
    import time
    from google.adk.tools.tool_context import ToolContext

    def log_user_login(tool_context: ToolContext) -> dict:
        """
        Cập nhật trạng thái phiên khi có sự kiện đăng nhập của người dùng.
        Công cụ này đóng gói tất cả các thay đổi trạng thái liên quan đến đăng nhập người dùng.
        """
        # Truy cập state trực tiếp thông qua context được cung cấp.
        state = tool_context.state
        
        # Lấy giá trị hiện tại hoặc mặc định, sau đó cập nhật state.
        login_count = state.get("user:login_count", 0) + 1
        state["user:login_count"] = login_count
        state["task_status"] = "active"
        state["user:last_login_ts"] = time.time()
        state["temp:validation_needed"] = True
        
        print("State đã được cập nhật từ bên trong công cụ `log_user_login`.")
        
        return {
            "status": "success",
            "message": f"Đã theo dõi đăng nhập người dùng. Tổng số lần đăng nhập: {login_count}."
        }
    ```

    Lưu ý rằng việc sửa đổi trực tiếp từ điển `session.state` sau khi truy xuất phiên không được khuyến khích vì nó bỏ qua cơ chế xử lý sự kiện tiêu chuẩn.

### Memory: Kiến thức Dài hạn với MemoryService

Trong các hệ thống tác nhân, thành phần Session duy trì bản ghi lịch sử trò chuyện hiện tại (sự kiện) và dữ liệu tạm thời (trạng thái) cụ thể cho một cuộc hội thoại. Tuy nhiên, để các tác nhân giữ lại thông tin qua nhiều tương tác hoặc truy cập dữ liệu bên ngoài, cần phải quản lý kiến thức dài hạn. Điều này được tạo điều kiện bởi `MemoryService`.

Session và State có thể được hình dung là bộ nhớ ngắn hạn, trong khi Kiến thức Dài hạn được quản lý bởi `MemoryService` hoạt động như một kho lưu trữ bền vững và có thể tìm kiếm được. Kho lưu trữ này có thể chứa thông tin từ nhiều tương tác trong quá khứ hoặc các nguồn bên ngoài.

ADK cung cấp một số triển khai để tạo kho lưu trữ kiến thức dài hạn này:
*   `InMemoryMemoryService`: Thích hợp cho thử nghiệm, dữ liệu bị mất khi khởi động lại.
*   `VertexAiRagMemoryService`: Thường được sử dụng cho môi trường sản xuất, tận dụng dịch vụ RAG của Google Cloud, cho phép khả năng tìm kiếm ngữ nghĩa, bền vững và có thể mở rộng.

Ví dụ sử dụng `VertexAiRagMemoryService`:

```python
from google.adk.memory import VertexAiRagMemoryService

# Tên tài nguyên của Vertex AI RAG Corpus của bạn
RAG_CORPUS_RESOURCE_NAME = "projects/your-gcp-project-id/locations/us-central1/ragCorpora/your-corpus-id"

memory_service = VertexAiRagMemoryService(
    rag_corpus=RAG_CORPUS_RESOURCE_NAME,
    similarity_top_k=5,
    vector_distance_threshold=0.7
)
# Khi sử dụng dịch vụ này, các phương thức như add_session_to_memory
# và search_memory sẽ tương tác với Vertex AI RAG Corpus được chỉ định.
```

## Ví dụ Code Thực hành: Quản lý Bộ nhớ trong LangChain và LangGraph

Trong LangChain và LangGraph, Bộ nhớ là một thành phần quan trọng để tạo ra các ứng dụng hội thoại thông minh và tự nhiên.

**Bộ nhớ Ngắn hạn:** Có phạm vi luồng (thread-scoped), theo dõi cuộc hội thoại đang diễn ra trong một phiên hoặc luồng duy nhất. LangGraph quản lý bộ nhớ ngắn hạn như một phần của trạng thái tác nhân, được duy trì thông qua một **checkpointer**, cho phép một luồng được tiếp tục bất cứ lúc nào.

**Bộ nhớ Dài hạn:** Lưu trữ dữ liệu cụ thể của người dùng hoặc cấp ứng dụng qua các phiên và được chia sẻ giữa các luồng hội thoại. Nó được lưu trong các "không gian tên" (namespaces) tùy chỉnh và có thể được gọi lại bất cứ lúc nào trong bất kỳ luồng nào. LangGraph cung cấp các cửa hàng (stores) để lưu và gọi lại bộ nhớ dài hạn.

### Quản lý Bộ nhớ Thủ công với `ChatMessageHistory`

```python
from langchain.memory import ChatMessageHistory

# Khởi tạo đối tượng lịch sử
history = ChatMessageHistory()

# Thêm tin nhắn của người dùng và AI
history.add_user_message("Tôi sẽ đến New York vào tuần tới.")
history.add_ai_message("Tuyệt vời! Đó là một thành phố tuyệt vời.")

# Truy cập danh sách tin nhắn
print(history.messages)
```

### Bộ nhớ Tự động cho Chuỗi với `ConversationBufferMemory`

Để tích hợp bộ nhớ trực tiếp vào các chuỗi, `ConversationBufferMemory` là một lựa chọn phổ biến. Nó giữ một bộ đệm của cuộc hội thoại và làm cho nó có sẵn cho prompt của bạn.

```python
from langchain.memory import ConversationBufferMemory

# Khởi tạo bộ nhớ
memory = ConversationBufferMemory()

# Lưu một lượt hội thoại
memory.save_context({"input": "Thời tiết thế nào?"}, {"output": "Hôm nay trời nắng."})

# Tải bộ nhớ dưới dạng chuỗi
print(memory.load_memory_variables({}))
```

Tích hợp bộ nhớ này vào một `LLMChain` cho phép mô hình truy cập lịch sử cuộc hội thoại:

```python
from langchain_openai import OpenAI
from langchain.chains import LLMChain
from langchain.prompts import PromptTemplate
from langchain.memory import ConversationBufferMemory

llm = OpenAI(temperature=0)
template = """Bạn là một đại lý du lịch hữu ích.
Cuộc hội thoại trước đó:
{history}
Câu hỏi mới: {question}
Phản hồi:"""
prompt = PromptTemplate.from_template(template)

# Cấu hình bộ nhớ
# memory_key "history" khớp với biến trong prompt
memory = ConversationBufferMemory(memory_key="history")

# Xây dựng Chuỗi
conversation = LLMChain(llm=llm, prompt=prompt, memory=memory)

# Chạy cuộc hội thoại
response = conversation.predict(question="Tôi muốn đặt một chuyến bay.")
print(response)
```

## Vertex Memory Bank

Memory Bank (Ngân hàng Bộ nhớ), một dịch vụ được quản lý trong Vertex AI Agent Engine, cung cấp cho các tác nhân bộ nhớ dài hạn, bền vững. Dịch vụ sử dụng các mô hình Gemini để phân tích không đồng bộ lịch sử hội thoại để trích xuất các sự kiện chính và sở thích của người dùng.

Thông tin này được lưu trữ bền vững, được tổ chức bởi một phạm vi xác định như ID người dùng, và được cập nhật thông minh để hợp nhất dữ liệu mới và giải quyết các mâu thuẫn. Khi bắt đầu một phiên mới, tác nhân truy xuất các bộ nhớ liên quan thông qua việc gọi lại dữ liệu đầy đủ hoặc tìm kiếm tương tự bằng cách sử dụng embeddings. Quy trình này cho phép một tác nhân duy trì tính liên tục qua các phiên và cá nhân hóa các phản hồi dựa trên thông tin đã nhớ lại.

```python
from google.adk.memory import VertexAiMemoryBankService

agent_engine_id = agent_engine.api_resource.name.split("/")[-1]

memory_service = VertexAiMemoryBankService(
    project="PROJECT_ID",
    location="LOCATION",
    agent_engine_id=agent_engine_id
)

session = await session_service.get_session(
    app_name=app_name,
    user_id="USER_ID",
    session_id=session.id
)

await memory_service.add_session_to_memory(session)
```

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các hệ thống Agentic cần ghi nhớ thông tin từ các tương tác trong quá khứ để thực hiện các tác vụ phức tạp và cung cấp trải nghiệm mạch lạc. Nếu không có cơ chế bộ nhớ, các tác nhân là phi trạng thái (stateless), không thể duy trì ngữ cảnh hội thoại, học hỏi từ kinh nghiệm hoặc cá nhân hóa phản hồi. Vấn đề cốt lõi là làm thế nào để quản lý hiệu quả cả thông tin tức thời, tạm thời của một cuộc trò chuyện duy nhất và kiến thức rộng lớn, bền vững được thu thập theo thời gian.
*   **Tại sao:** Giải pháp tiêu chuẩn hóa là triển khai một hệ thống bộ nhớ hai thành phần phân biệt giữa lưu trữ ngắn hạn và dài hạn. Bộ nhớ ngắn hạn, theo ngữ cảnh giữ dữ liệu tương tác gần đây trong cửa sổ ngữ cảnh của LLM để duy trì luồng hội thoại. Đối với thông tin phải tồn tại lâu dài, các giải pháp bộ nhớ dài hạn sử dụng cơ sở dữ liệu bên ngoài, thường là vector stores, để truy xuất ngữ nghĩa, hiệu quả. Các framework agentic như Google ADK cung cấp các thành phần cụ thể để quản lý điều này, chẳng hạn như `Session` cho luồng hội thoại và `State` cho dữ liệu tạm thời của nó. Một `MemoryService` chuyên dụng được sử dụng để giao tiếp với cơ sở tri thức dài hạn.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này khi một tác nhân cần làm nhiều hơn là trả lời một câu hỏi duy nhất. Nó rất cần thiết cho các tác nhân phải duy trì ngữ cảnh trong suốt cuộc trò chuyện, theo dõi tiến trình trong các tác vụ đa bước hoặc cá nhân hóa các tương tác bằng cách nhớ lại sở thích và lịch sử của người dùng. Triển khai quản lý bộ nhớ bất cứ khi nào tác nhân dự kiến sẽ học hỏi hoặc thích nghi dựa trên những thành công, thất bại trong quá khứ hoặc thông tin mới thu được.

## Những Điểm Chính (Key Takeaways)

*   Bộ nhớ cực kỳ quan trọng để các tác nhân theo dõi mọi thứ, học hỏi và cá nhân hóa các tương tác.
*   AI hội thoại dựa vào cả bộ nhớ ngắn hạn cho ngữ cảnh tức thời trong một cuộc trò chuyện và bộ nhớ dài hạn cho kiến thức bền vững qua nhiều phiên.
*   Bộ nhớ ngắn hạn (những thứ tức thời) là tạm thời, thường bị giới hạn bởi cửa sổ ngữ cảnh của LLM hoặc cách framework truyền ngữ cảnh.
*   Bộ nhớ dài hạn (những thứ tồn tại xung quanh) lưu thông tin qua các cuộc trò chuyện khác nhau bằng cách sử dụng bộ nhớ ngoài như cơ sở dữ liệu vector và được truy cập bằng cách tìm kiếm.
*   Các framework như ADK có các phần cụ thể như Session, State và MemoryService để quản lý bộ nhớ.
*   Trong ADK, bạn nên cập nhật state bằng cách sử dụng `EventActions.state_delta` hoặc `output_key` khi thêm sự kiện, không phải bằng cách thay đổi trực tiếp từ điển state.
*   LangChain cung cấp các công cụ thiết thực như `ConversationBufferMemory` để tự động đưa lịch sử cuộc hội thoại vào prompt.
*   LangGraph cho phép bộ nhớ dài hạn, nâng cao bằng cách sử dụng một store để lưu và truy xuất các sự kiện ngữ nghĩa, trải nghiệm tình huống hoặc thậm chí các quy tắc thủ tục có thể cập nhật qua các phiên người dùng khác nhau.

## Kết luận

Chương này đi sâu vào công việc thực sự quan trọng của quản lý bộ nhớ cho các hệ thống tác nhân, cho thấy sự khác biệt giữa ngữ cảnh ngắn hạn và kiến thức tồn tại trong một thời gian dài. Chúng ta đã nói về cách các loại bộ nhớ này được thiết lập và nơi bạn thấy chúng được sử dụng trong việc xây dựng các tác nhân thông minh hơn có thể ghi nhớ mọi thứ. Chúng ta đã xem xét chi tiết cách Google ADK cung cấp cho bạn các phần cụ thể như Session, State và MemoryService để xử lý việc này. Bây giờ chúng ta đã đề cập đến cách các tác nhân có thể ghi nhớ mọi thứ, cả ngắn hạn và dài hạn, chúng ta có thể chuyển sang cách chúng có thể học hỏi và thích nghi. Mẫu tiếp theo "Học hỏi và Thích nghi" là về việc một tác nhân thay đổi cách nó suy nghĩ, hành động hoặc những gì nó biết, tất cả dựa trên những trải nghiệm hoặc dữ liệu mới.

## Tài liệu tham khảo
1.  ADK Memory, https://google.github.io/adk-docs/sessions/memory/
2.  LangGraph Memory, https://langchain-ai.github.io/langgraph/concepts/memory/
3.  Vertex AI Agent Engine Memory Bank, https://cloud.google.com/blog/products/ai-machine-learning/vertex-ai-memory-bank-in-public-preview