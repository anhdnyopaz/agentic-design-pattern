# Hướng dẫn Python Cơ Bản cho Agentic Design Patterns

Chào mừng bạn! Thư mục này chứa bộ tài liệu hướng dẫn lập trình Python "cấp tốc" nhưng chuyên sâu, được thiết kế riêng để bạn có đủ kiến thức nền tảng thực hành các ví dụ trong tài liệu **Agentic Design Patterns**.

Bộ tài liệu này không dạy Python từ con số 0, mà tập trung vào **các kỹ thuật cụ thể được sử dụng trong CrewAI, LangChain, và Google ADK**.

## Danh sách bài học

1.  **[01_Python_Basics_Setup.md](./01_Python_Basics_Setup.md)**
    *   Cách tạo môi trường ảo (venv) để cài đặt thư viện không bị lỗi.
    *   Làm việc với Dictionary và List (Cấu trúc dữ liệu chính của AI).
    *   F-Strings cho Prompt Engineering.

2.  **[02_Functions_Decorators_Typing.md](./02_Functions_Decorators_Typing.md)**
    *   **Type Hinting**: Tại sao cần viết `def func(a: int) -> str:`?
    *   **Decorators**: Hiểu về `@tool` để tạo công cụ cho Agent.
    *   **Docstrings**: Cách viết tài liệu để AI hiểu được code của bạn.

3.  **[03_OOP_Classes.md](./03_OOP_Classes.md)**
    *   Hiểu về Class và Object (`Agent`, `Task`, `Crew`).
    *   Cách khởi tạo và cấu hình các đối tượng này.
    *   Giới thiệu Pydantic (Data Validation).

4.  **[04_Async_Programming.md](./04_Async_Programming.md)** (Rất quan trọng cho Google ADK)
    *   Khác biệt giữa Đồng bộ (Sync) và Bất đồng bộ (Async).
    *   Cách dùng `async`, `await`.
    *   Streaming dữ liệu với `async for`.

5.  **[05_Environment_Error_Handling.md](./05_Environment_Error_Handling.md)**
    *   Bảo mật API Key với `.env`.
    *   Xử lý lỗi (Robustness) với `try...except`.

## Lộ trình học tập

1.  Đọc lần lượt từ bài 01 đến 05.
2.  Thực hiện các **Bài tập thực hành** nhỏ ở cuối mỗi file.
3.  Sau khi nắm vững, quay lại thư mục `doc_vi` và chạy thử các đoạn code ví dụ trong các chương (ví dụ: Chương 5 về Tool Use, Chương 13 về Safety).

Chúc bạn học tốt và xây dựng được những AI Agent mạnh mẽ!
