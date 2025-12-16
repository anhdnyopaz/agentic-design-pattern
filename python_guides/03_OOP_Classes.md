# Phần 3: OOP (Lập Trình Hướng Đối Tượng) cho Agent

Hầu hết các framework Agent (CrewAI, Google ADK) đều được xây dựng dựa trên Class (Lớp). Bạn cần hiểu cách khởi tạo Object (Đối tượng) để cấu hình Agent.

## 1. Class và Object cơ bản

Hãy tưởng tượng `Class` là bản vẽ thiết kế (Blueprint), còn `Object` là ngôi nhà thực sự được xây từ bản vẽ đó.

```python
class Agent:
    # __init__ là hàm khởi tạo (Constructor).
    # Nó chạy ngay khi bạn tạo một Object mới.
    def __init__(self, role: str, goal: str):
        self.role = role  # Lưu trữ thuộc tính (attribute)
        self.goal = goal

    def introduce(self):
        # self đại diện cho chính đối tượng đang được gọi
        return f"Tôi là {self.role}, mục tiêu của tôi là {self.goal}"

# Tạo Object (Instantiation)
researcher = Agent(role="Nhà nghiên cứu", goal="Tìm thông tin mới")
writer = Agent(role="Nhà văn", goal="Viết bài blog")

print(researcher.introduce()) 
# -> Tôi là Nhà nghiên cứu, mục tiêu của tôi là Tìm thông tin mới
```

## 2. Kế thừa (Inheritance) - Tùy chỉnh Agent

Đôi khi bạn muốn tạo một loại Agent đặc biệt kế thừa các tính năng cơ bản nhưng có thêm chức năng riêng.

```python
# BaseAgent là lớp cha
class BaseAgent:
    def execute(self):
        print("Đang thực thi tác vụ cơ bản...")

# SuperAgent kế thừa từ BaseAgent
class SuperAgent(BaseAgent):
    def execute(self):
        # Gọi hàm của lớp cha
        super().execute() 
        print("...Và thêm sức mạnh siêu nhiên!")

my_agent = SuperAgent()
my_agent.execute()
# Output:
# Đang thực thi tác vụ cơ bản...
# ...Và thêm sức mạnh siêu nhiên!
```

## 3. Áp dụng vào CrewAI

Trong CrewAI, bạn không cần định nghĩa Class `Agent` (họ đã làm rồi), bạn chỉ cần tạo Object từ Class đó.

```python
from crewai import Agent, Task

# Bạn đang tạo một Instance (Object) của Class Agent
my_agent = Agent(
    role='Analyst',
    goal='Analyze data',
    backstory='You are an expert...',
    verbose=True,  # Các tham số này được truyền vào __init__ của lớp Agent
    allow_delegation=False
)

# Tương tự với Task
my_task = Task(
    description='Analyze the stock market',
    agent=my_agent # Bạn truyền object my_agent vào trong object my_task
)
```

## 4. Pydantic (OOP nâng cao cho dữ liệu)

Các thư viện AI hiện đại sử dụng `Pydantic` để định nghĩa cấu trúc dữ liệu (Data Models) thay vì Class thông thường. Nó giúp kiểm tra dữ liệu đầu vào/đầu ra của LLM.

```python
from pydantic import BaseModel, Field

# Định nghĩa cấu trúc đầu ra mong muốn từ LLM
class UserProfile(BaseModel):
    name: str = Field(description="Tên của người dùng")
    age: int = Field(description="Tuổi của người dùng")
    interests: list[str] = Field(description="Danh sách sở thích")

# Khi LLM trả về dữ liệu, Pydantic sẽ đảm bảo nó khớp với khuôn mẫu này.
# Nếu LLM trả về age="hai mươi" (chuỗi), Pydantic sẽ báo lỗi hoặc cố gắng sửa thành 20 (int).
```

---
**Bài tập thực hành:**
1. Tạo file `oop_practice.py`.
2. Định nghĩa một class `CustomTool`. Hàm `__init__` nhận vào `name` và `description`.
3. Thêm một hàm `run(self, input_data)` in ra "Tool [name] đang chạy với dữ liệu: [input_data]".
4. Tạo một object từ class này và gọi hàm `run`.
