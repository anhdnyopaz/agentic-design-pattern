# Phần 1: Thiết lập Môi trường và Cấu trúc Dữ liệu Cơ bản cho AI

Để chạy được các ví dụ về AI Agent, bước đầu tiên là thiết lập môi trường Python sạch sẽ và nắm vững các cấu trúc dữ liệu mà LLM thường xuyên thao tác (đặc biệt là Dictionary và JSON).

## 1. Thiết lập Môi trường Ảo (Virtual Environment)

Trong các dự án AI, chúng ta cài đặt rất nhiều thư viện (crewai, langchain, google-generativeai...). Để tránh xung đột phiên bản, bạn **luôn luôn** nên sử dụng môi trường ảo.

### Cách tạo và kích hoạt

**Trên macOS/Linux:**
```bash
# 1. Tạo môi trường ảo tên là 'venv'
python3 -m venv venv

# 2. Kích hoạt môi trường
source venv/bin/activate
```

**Trên Windows:**
```cmd
# 1. Tạo môi trường
python -m venv venv

# 2. Kích hoạt
venv\Scripts\activate
```

Sau khi kích hoạt, bạn sẽ thấy `(venv)` ở đầu dòng lệnh. Bây giờ bạn có thể cài đặt thư viện an toàn:
```bash
pip install crewai langchain-openai python-dotenv
```

## 2. Dictionary (Từ điển) - Ngôn ngữ của AI

Hầu hết các LLM trả về dữ liệu dưới dạng JSON, trong Python nó tương ứng với `dict` (Dictionary). Bạn cần thành thạo việc truy xuất và lồng ghép dict.

```python
# Ví dụ: Một cấu trúc phản hồi giả lập từ Agent
agent_response = {
    "task_id": "task_123",
    "status": "completed",
    "output": {
        "summary": "AI Agent là tương lai.",
        "confidence_score": 0.95,
        "tags": ["ai", "tech", "automation"]
    },
    "meta_data": None
}

# 1. Truy cập dữ liệu
print(agent_response["status"]) # -> completed

# 2. Truy cập dữ liệu lồng nhau (Nested)
# Lấy 'confidence_score' nằm bên trong 'output'
score = agent_response["output"]["confidence_score"]
print(f"Độ tin cậy: {score}")

# 3. Sử dụng .get() để tránh lỗi
# Nếu key không tồn tại, chương trình sẽ không crash mà trả về None hoặc giá trị mặc định
error_log = agent_response.get("errors", "Không có lỗi") 
print(error_log) # -> Không có lỗi

# 4. Duyệt qua Dictionary
for key, value in agent_response.items():
    print(f"Key: {key} - Value Type: {type(value)}")
```

## 3. List (Danh sách) và Vòng lặp

Agent thường xử lý danh sách các tác vụ (Tasks) hoặc danh sách các công cụ (Tools).

```python
# Danh sách các công cụ (chuỗi tên)
tools = ["search_tool", "calculator_tool", "weather_tool"]

# 1. Thêm phần tử
tools.append("translation_tool")

# 2. List Comprehension (Cách viết tắt mạnh mẽ thường dùng trong xử lý dữ liệu AI)
# Ví dụ: Muốn viết hoa tất cả tên tool
upper_tools = [t.upper() for t in tools]
print(upper_tools) 
# -> ['SEARCH_TOOL', 'CALCULATOR_TOOL', 'WEATHER_TOOL', 'TRANSLATION_TOOL']

# 3. Lọc dữ liệu
# Chỉ lấy tool có chữ 'tool'
filtered = [t for t in tools if "search" in t]
print(filtered) # -> ['search_tool']
```

## 4. F-Strings (Định dạng chuỗi)

Trong Prompt Engineering, bạn sẽ ghép biến vào chuỗi văn bản rất nhiều. F-string là cách hiện đại và dễ đọc nhất.

```python
user_name = "Alice"
topic = "Machine Learning"

# Prompt template
prompt = f"""
Xin chào {user_name},
Hãy đóng vai một chuyên gia và giải thích cho tôi về {topic}.
"""

print(prompt)
```

---
**Bài tập thực hành:**
1. Tạo một file `setup_practice.py`.
2. Khai báo một dictionary mô tả một `Crew` (gồm tên, danh sách agent, và trạng thái).
3. Sử dụng f-string để in ra câu: "Crew [Tên] có [Số lượng] agents và đang ở trạng thái [Trạng thái]".
