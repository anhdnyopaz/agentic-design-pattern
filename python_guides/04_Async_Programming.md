# Phần 4: Async Programming (Lập Trình Bất Đồng Bộ)

Đây là phần khó nhất nhưng quan trọng nhất khi làm việc với các framework như **Google ADK** hoặc khi xây dựng các ứng dụng Chatbot thời gian thực.

## 1. Tại sao cần Async?

Khi Agent gọi LLM (OpenAI/Gemini), nó phải chờ phản hồi qua mạng (Internet).
*   **Synchronous (Đồng bộ):** Chương trình dừng lại hoàn toàn, chờ server trả lời xong mới làm việc khác. (Chậm, phí thời gian chờ).
*   **Asynchronous (Bất đồng bộ):** Trong khi chờ phản hồi, chương trình có thể làm việc khác (ví dụ: nhận tin nhắn mới từ người dùng).

## 2. Cú pháp `async` và `await`

*   `async def`: Định nghĩa một hàm bất đồng bộ (Coroutine).
*   `await`: Tạm dừng hàm này, chờ tác vụ kia xong thì mới chạy tiếp.

```python
import asyncio
import time

# Hàm đồng bộ bình thường
def sync_task():
    time.sleep(2) # Dừng chương trình 2 giây
    print("Xong việc đồng bộ")

# Hàm bất đồng bộ
async def async_task():
    print("Bắt đầu gọi API giả lập...")
    await asyncio.sleep(2) # Giả lập chờ API, nhưng không chặn chương trình
    print("API trả về kết quả!")

# Cách chạy
# async_task() # -> Lỗi! Bạn không thể gọi trực tiếp hàm async
# Bạn phải dùng asyncio.run()
if __name__ == "__main__":
    asyncio.run(async_task())
```

## 3. Streaming (Dữ liệu luồng) với `async for`

Trong các ví dụ Google ADK, bạn thấy rất nhiều `async for`. Đây là kỹ thuật để nhận từng phần của câu trả lời (Streaming) thay vì chờ cả đoạn văn dài.

```python
# Giả lập một hàm trả về dữ liệu từng chút một (Generator)
async def stream_response_from_llm():
    words = ["Xin", "chào", "tôi", "là", "AI", "."]
    for word in words:
        await asyncio.sleep(0.5) # Giả lập độ trễ mạng
        yield word # Trả về từng từ một

async def main():
    print("Agent đang gõ: ", end="", flush=True)
    
    # async for dùng để duyệt qua Async Generator
    async for token in stream_response_from_llm():
        print(token + " ", end="", flush=True)
    
    print("\nĐã xong!")

if __name__ == "__main__":
    asyncio.run(main())
```

## 4. Áp dụng vào Google ADK (Phân tích ví dụ)

Đoạn mã sau trích từ các ví dụ Google ADK trong tài liệu:

```python
# InMemoryRunner chạy agent
runner = InMemoryRunner(agent)

# runner.run() trả về một luồng sự kiện (events)
# Chúng ta dùng async for để bắt từng sự kiện khi nó xảy ra
async for event in runner.run(user_id=..., ...):
    
    # Nếu sự kiện là phản hồi cuối cùng
    if event.is_final_response():
        print(event.content.text)
```

Điều này cho phép ứng dụng của bạn hiển thị trạng thái "Agent đang suy nghĩ...", "Agent đang dùng công cụ...", "Agent đang trả lời..." theo thời gian thực (Real-time).

---
**Bài tập thực hành:**
1. Tạo file `async_practice.py`.
2. Viết một hàm `async def cook_dinner()` mất 3 giây (dùng `await asyncio.sleep(3)`).
3. Viết một hàm `async def wash_dishes()` mất 2 giây.
4. Viết hàm `main` sử dụng `asyncio.gather(cook_dinner(), wash_dishes())` để chạy cả hai việc cùng lúc.
5. So sánh tổng thời gian chạy (nên là khoảng 3 giây thay vì 5 giây).

```