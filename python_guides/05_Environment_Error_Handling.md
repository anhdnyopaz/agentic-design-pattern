# Phần 5: Quản lý Môi trường và Xử lý Lỗi

Phần cuối cùng này đảm bảo ứng dụng của bạn an toàn và bền vững (Robustness).

## 1. Biến Môi trường (Environment Variables) - Bảo vệ API Key

Không bao giờ được viết trực tiếp API Key vào code (Hardcode). Nếu bạn upload code lên GitHub, kẻ xấu sẽ lấy trộm key của bạn.

### Cách làm chuẩn:

1.  Tạo file `.env` trong thư mục gốc dự án.
2.  Thêm key vào file `.env`:
    ```
    OPENAI_API_KEY=sk-proj-xxxxxxxxxxxx
    GOOGLE_API_KEY=AIzaSyxxxxxxxxxxxx
    ```
3.  Thêm `.env` vào file `.gitignore` để không bao giờ commit nó.

### Cách đọc trong Python:
Sử dụng thư viện `python-dotenv`.

```python
import os
from dotenv import load_dotenv

# 1. Tải biến từ file .env vào hệ thống
load_dotenv()

# 2. Lấy biến
api_key = os.getenv("OPENAI_API_KEY")

# 3. Kiểm tra an toàn
if not api_key:
    raise ValueError("Chưa cấu hình API Key! Vui lòng kiểm tra file .env")

print("Đã tải API Key thành công.")
```

## 2. Xử lý Lỗi (Try / Except)

AI và API mạng rất dễ gặp lỗi (Time out, Rate limit, AI trả lời sai định dạng). Bạn cần dùng khối `try...except` để chương trình không bị crash (Sập).

```python
def call_ai_agent(input_text):
    # Giả lập lỗi
    if input_text == "error":
        raise ConnectionError("Mất kết nối tới OpenAI!")
    return "AI phản hồi: OK"

try:
    # Khối code có thể gây lỗi
    user_input = "error"
    result = call_ai_agent(user_input)
    print(result)

except ConnectionError as e:
    # Xử lý lỗi kết nối cụ thể
    print(f"Lỗi mạng: {e}. Đang thử lại...")
    # Logic thử lại (Retry) có thể đặt ở đây

except Exception as e:
    # Bắt tất cả các lỗi còn lại
    print(f"Đã xảy ra lỗi không mong muốn: {e}")

finally:
    # Luôn chạy dù có lỗi hay không (ví dụ: đóng file, ngắt DB)
    print("Kết thúc phiên làm việc.")
```

## 3. Debugging (Gỡ lỗi) cơ bản

Khi Agent không hoạt động như ý, `print` là bạn tốt nhất, nhưng `logging` chuyên nghiệp hơn.

```python
import logging

# Cấu hình logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def complex_task():
    logger.info("Bắt đầu tác vụ...")
    x = 10
    logger.debug(f"Giá trị x = {x}") # Chỉ hiện nếu level=DEBUG
    
    try:
        result = 100 / 0
    except ZeroDivisionError:
        logger.error("Lỗi chia cho 0!")

complex_task()
```

---
**Bài tập thực hành:**
1. Tạo file `.env` chứa `MY_SECRET=12345`.
2. Tạo file `env_practice.py`.
3. Dùng `dotenv` tải và in ra giá trị `MY_SECRET`.
4. Viết một khối `try...except` để bắt lỗi khi cố gắng đọc một biến môi trường không tồn tại và in ra thông báo thân thiện.
