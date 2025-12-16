# Phần 2: Hàm, Type Hinting và Decorators (Định nghĩa Tools)

Trong các framework như CrewAI hay LangChain, việc định nghĩa "Công cụ" (Tools) cho Agent thực chất là viết các hàm Python. Tuy nhiên, để AI hiểu được cách dùng hàm đó, chúng ta cần hai kỹ thuật nâng cao: **Type Hinting** và **Decorators**.

## 1. Type Hinting (Gợi ý kiểu dữ liệu)

Python là ngôn ngữ kiểu động (không bắt buộc khai báo kiểu), nhưng các thư viện AI cần biết chính xác đầu vào là gì để tạo ra Schema cho LLM.

```python
# Cách viết cũ (Không rõ ràng cho AI)
def calculate_area(width, height):
    return width * height

# Cách viết chuẩn cho Agent (Type Hinting)
def calculate_area(width: float, height: float) -> float:
    """
    Tính diện tích hình chữ nhật.
    """
    return width * height
```

**Tại sao quan trọng?** 
Khi bạn dùng `@tool`, thư viện sẽ đọc `width: float` để nói với LLM rằng: "Khi gọi hàm này, hãy chắc chắn tham số `width` là một số thực, không phải chuỗi văn bản".

### Các kiểu dữ liệu phổ biến từ thư viện `typing`:

```python
from typing import List, Dict, Optional, Union

# List: Danh sách các chuỗi
def process_names(names: List[str]) -> bool:
    pass

# Optional: Có thể là int hoặc None (tham số không bắt buộc)
def search_product(query: str, limit: Optional[int] = 10) -> str:
    pass
```

## 2. Decorators (A.K.A `@tool`)

Decorator là một cú pháp đặc biệt bắt đầu bằng `@`. Nó cho phép bạn "bọc" một hàm bằng một lớp chức năng khác mà không cần sửa đổi mã nguồn của hàm đó.

Trong Agent, Decorator biến một hàm Python bình thường thành một công cụ mà AI có thể nhìn thấy.

```python
from langchain.tools import tool

# @tool là decorator
@tool("SearchTool")
def search_internet(query: str) -> str:
    """
    Tìm kiếm thông tin trên internet.
    Hữu ích khi cần thông tin thời gian thực.
    """
    # Giả lập logic tìm kiếm
    return f"Kết quả tìm kiếm cho: {query}"

# Khi bạn dùng @tool:
# 1. Nó đọc tên hàm và docstring ("Tìm kiếm thông tin...") để làm mô tả cho AI.
# 2. Nó đọc type hinting (query: str) để biết cách gọi hàm.
# 3. Nó đăng ký hàm này vào danh sách công cụ của framework.
```

## 3. Docstrings (Chuỗi tài liệu)

Docstring là đoạn văn bản nằm ngay dưới tên hàm, thường được bao bởi ba dấu ngoặc kép `"""`.
Đối với lập trình bình thường, nó là tài liệu cho con người.
**Đối với lập trình Agent, nó là Prompt cho AI.**

```python
@tool
def heavy_computation(x: int) -> int:
    """
    QUAN TRỌNG: Chỉ sử dụng công cụ này khi x > 1000.
    Nếu x nhỏ hơn, hãy tự tính toán.
    Công cụ này thực hiện phép tính phức tạp ABC.
    """
    return x * x
```
*AI sẽ đọc docstring này và tuân theo hướng dẫn "Chỉ sử dụng khi x > 1000".*

---
**Bài tập thực hành:**
1. Tạo file `tools_practice.py`.
2. Import `tool` từ `langchain.tools` (hoặc `crewai_tools`).
3. Viết một hàm `get_weather(city: str) -> str`.
4. Thêm Type Hinting và Docstring mô tả rõ ràng.
5. Gắn decorator `@tool` cho nó.
6. In ra `get_weather.name` và `get_weather.description` để xem thư viện đã "hiểu" hàm của bạn chưa.
