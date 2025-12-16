# Hướng dẫn Thực hành: Tool Use Pattern với ADK-Go

## Mục tiêu

Sau khi hoàn thành bài hướng dẫn này, bạn sẽ:
- Hiểu mẫu Tool Use (Function Calling) trong thiết kế Agentic
- Biết cách định nghĩa Tools trong ADK-Go
- Triển khai các loại Tools khác nhau
- Xây dựng Agent có khả năng tương tác với thế giới bên ngoài

---

## Phần 1: Giới thiệu Tool Use Pattern

### 1.1 Tool Use là gì?

**Tool Use (Function Calling)** cho phép LLM Agent tương tác với:
- API bên ngoài
- Cơ sở dữ liệu
- Dịch vụ web
- Thực thi code
- Các hệ thống khác

Đây là cầu nối giữa **khả năng suy luận** của LLM và **hành động thực tế**.

### 1.2 Quy trình Tool Use

```
┌─────────────────────────────────────────────────────────────────┐
│                     TOOL USE WORKFLOW                           │
│                                                                 │
│  ┌──────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │   USER   │───▶│     LLM      │───▶│  TOOL CALL   │          │
│  │  INPUT   │    │   DECIDES    │    │  GENERATION  │          │
│  └──────────┘    └──────────────┘    └──────────────┘          │
│                                             │                   │
│                                             ▼                   │
│  ┌──────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │  FINAL   │◀───│     LLM      │◀───│    TOOL      │          │
│  │ RESPONSE │    │  PROCESSES   │    │  EXECUTION   │          │
│  └──────────┘    └──────────────┘    └──────────────┘          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### 1.3 Các bước chi tiết

1. **Định nghĩa Tool (Tool Definition)**
   - Tên, mô tả, parameters
   - Input/Output types

2. **LLM Quyết định (LLM Decision)**
   - Phân tích user request
   - Xác định tool cần gọi

3. **Tạo Tool Call (Function Call Generation)**
   - LLM tạo JSON với tool name + arguments

4. **Thực thi Tool (Tool Execution)**
   - Framework gọi function thực tế

5. **Xử lý kết quả (Result Processing)**
   - LLM nhận kết quả và tạo response

---

## Phần 2: Các loại Tools trong ADK-Go

### 2.1 Function Tool

Tool đơn giản nhất - wrap một Go function:

```go
// Cấu trúc cơ bản
functiontool.New(functiontool.Config{
    Name:        "tool_name",
    Description: "Mô tả tool làm gì",
}, handlerFunction)
```

### 2.2 Agent Tool

Wrap một Agent khác thành Tool:

```go
// Agent như là Tool
agenttool.New(subAgent, nil)
```

### 2.3 Built-in Tools (ADK-Go)

ADK-Go cung cấp một số built-in tools:
- Google Search
- Code Execution
- Vertex AI Extensions

---

## Phần 3: Định nghĩa Tool đúng cách

### 3.1 Cấu trúc Input/Output

```go
// Input struct - các parameters tool nhận vào
type ToolInput struct {
    ParamA string  `json:"param_a" description:"Mô tả param A"`
    ParamB int     `json:"param_b" description:"Mô tả param B"`
    OptionalParam string `json:"optional,omitempty" description:"Tham số tùy chọn"`
}

// Output struct - kết quả tool trả về
type ToolOutput struct {
    Result string `json:"result"`
    Status bool   `json:"status"`
}
```

### 3.2 Handler Function

```go
// Handler signature
func handler(ctx tool.Context, input ToolInput) (ToolOutput, error) {
    // Logic xử lý
    return ToolOutput{
        Result: "...",
        Status: true,
    }, nil
}
```

### 3.3 Best Practices cho Description

**❌ Không tốt:**
```go
Description: "Lấy thời tiết"
```

**✅ Tốt:**
```go
Description: "Lấy thông tin thời tiết hiện tại cho một thành phố. Sử dụng tool này khi người dùng hỏi về thời tiết, nhiệt độ, hoặc điều kiện khí hậu của một địa điểm cụ thể."
```

### 3.4 JSON Tags quan trọng

```go
type Input struct {
    // Bắt buộc - LLM phải cung cấp
    City string `json:"city" description:"Tên thành phố cần tra cứu"`

    // Tùy chọn - có thể bỏ qua
    Unit string `json:"unit,omitempty" description:"Đơn vị nhiệt độ: celsius hoặc fahrenheit"`
}
```

---

## Phần 4: Bài tập - Weather & Calculator Agent

### Mô tả bài tập

Xây dựng một **Personal Assistant Agent** với các tools:

1. **Weather Tool** - Tra cứu thời tiết
2. **Calculator Tool** - Thực hiện phép tính
3. **Time Tool** - Lấy thời gian hiện tại
4. **Unit Converter Tool** - Chuyển đổi đơn vị

### Cấu trúc file

```
cmd/tool_use/
├── main.go
```

---

## Phần 5: Hướng dẫn từng bước

### Bước 1: Import packages

```go
package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "os"
    "strings"
    "time"

    "google.golang.org/adk/agent"
    "google.golang.org/adk/agent/llmagent"
    "google.golang.org/adk/cmd/launcher"
    "google.golang.org/adk/cmd/launcher/full"
    "google.golang.org/adk/model/gemini"
    "google.golang.org/adk/tool"
    "google.golang.org/adk/tool/functiontool"
    "google.golang.org/genai"
)
```

### Bước 2: Tạo Weather Tool

```go
func createWeatherTool() (tool.Tool, error) {
    type Input struct {
        City string `json:"city" description:"Tên thành phố cần tra cứu thời tiết (ví dụ: Hanoi, Tokyo, London)"`
        Unit string `json:"unit,omitempty" description:"Đơn vị nhiệt độ: celsius (mặc định) hoặc fahrenheit"`
    }

    type Output struct {
        City        string  `json:"city"`
        Temperature float64 `json:"temperature"`
        Unit        string  `json:"unit"`
        Condition   string  `json:"condition"`
        Humidity    int     `json:"humidity"`
        Description string  `json:"description"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        // Mô phỏng dữ liệu thời tiết
        weatherData := map[string]struct {
            temp      float64
            condition string
            humidity  int
        }{
            "hanoi":    {28, "Nắng", 75},
            "hochiminh": {32, "Nắng nóng", 80},
            "danang":   {30, "Có mây", 70},
            "tokyo":    {22, "Mát mẻ", 60},
            "london":   {15, "Có mưa", 85},
            "newyork":  {25, "Quang đãng", 55},
            "paris":    {18, "Nhiều mây", 65},
        }

        city := strings.ToLower(strings.ReplaceAll(input.City, " ", ""))
        data, exists := weatherData[city]
        if !exists {
            data = struct {
                temp      float64
                condition string
                humidity  int
            }{25, "Không có dữ liệu chi tiết", 50}
        }

        unit := "°C"
        temp := data.temp
        if strings.ToLower(input.Unit) == "fahrenheit" {
            temp = data.temp*9/5 + 32
            unit = "°F"
        }

        return Output{
            City:        input.City,
            Temperature: temp,
            Unit:        unit,
            Condition:   data.condition,
            Humidity:    data.humidity,
            Description: fmt.Sprintf("Thời tiết tại %s: %.1f%s, %s, độ ẩm %d%%",
                input.City, temp, unit, data.condition, data.humidity),
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "get_weather",
        Description: `Tra cứu thông tin thời tiết hiện tại của một thành phố.
Sử dụng tool này khi người dùng hỏi về:
- Thời tiết của một thành phố
- Nhiệt độ hiện tại
- Điều kiện thời tiết (nắng, mưa, mây...)
- Độ ẩm

Ví dụ câu hỏi: "Thời tiết Hà Nội thế nào?", "Tokyo có nóng không?"`,
    }, handler)
}
```

### Bước 3: Tạo Calculator Tool

```go
func createCalculatorTool() (tool.Tool, error) {
    type Input struct {
        Expression string `json:"expression" description:"Biểu thức toán học cần tính (ví dụ: 2+3*4, sqrt(16), 10^2)"`
    }

    type Output struct {
        Expression string  `json:"expression"`
        Result     float64 `json:"result"`
        Formatted  string  `json:"formatted"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        // Parser đơn giản cho các phép tính cơ bản
        expr := strings.TrimSpace(input.Expression)
        var result float64

        // Xử lý các hàm đặc biệt
        switch {
        case strings.HasPrefix(expr, "sqrt("):
            var num float64
            fmt.Sscanf(expr, "sqrt(%f)", &num)
            result = math.Sqrt(num)
        case strings.HasPrefix(expr, "pow(") || strings.Contains(expr, "^"):
            var base, exp float64
            if strings.Contains(expr, "^") {
                fmt.Sscanf(expr, "%f^%f", &base, &exp)
            } else {
                fmt.Sscanf(expr, "pow(%f,%f)", &base, &exp)
            }
            result = math.Pow(base, exp)
        case strings.HasPrefix(expr, "sin("):
            var num float64
            fmt.Sscanf(expr, "sin(%f)", &num)
            result = math.Sin(num * math.Pi / 180) // Degrees
        case strings.HasPrefix(expr, "cos("):
            var num float64
            fmt.Sscanf(expr, "cos(%f)", &num)
            result = math.Cos(num * math.Pi / 180)
        default:
            // Phép tính cơ bản với 2 số
            var a, b float64
            var op rune
            for _, r := range "+-*/" {
                if strings.Contains(expr, string(r)) {
                    op = r
                    break
                }
            }
            parts := strings.Split(expr, string(op))
            if len(parts) == 2 {
                fmt.Sscanf(parts[0], "%f", &a)
                fmt.Sscanf(parts[1], "%f", &b)
                switch op {
                case '+':
                    result = a + b
                case '-':
                    result = a - b
                case '*':
                    result = a * b
                case '/':
                    if b != 0 {
                        result = a / b
                    }
                }
            }
        }

        return Output{
            Expression: input.Expression,
            Result:     result,
            Formatted:  fmt.Sprintf("%s = %.4f", input.Expression, result),
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "calculator",
        Description: `Thực hiện các phép tính toán học.
Sử dụng tool này khi người dùng yêu cầu:
- Tính toán số học (cộng, trừ, nhân, chia)
- Tính căn bậc hai: sqrt(number)
- Tính lũy thừa: number^power hoặc pow(base,exp)
- Tính sin, cos (độ)

Ví dụ: "Tính 15*7", "Căn bậc hai của 144", "2 mũ 10"`,
    }, handler)
}
```

### Bước 4: Tạo Time Tool

```go
func createTimeTool() (tool.Tool, error) {
    type Input struct {
        Timezone string `json:"timezone,omitempty" description:"Múi giờ (ví dụ: Asia/Ho_Chi_Minh, America/New_York). Mặc định là UTC"`
    }

    type Output struct {
        Time     string `json:"time"`
        Date     string `json:"date"`
        Timezone string `json:"timezone"`
        Unix     int64  `json:"unix_timestamp"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        loc := time.UTC
        tzName := "UTC"

        if input.Timezone != "" {
            if parsedLoc, err := time.LoadLocation(input.Timezone); err == nil {
                loc = parsedLoc
                tzName = input.Timezone
            }
        }

        now := time.Now().In(loc)

        return Output{
            Time:     now.Format("15:04:05"),
            Date:     now.Format("02/01/2006"),
            Timezone: tzName,
            Unix:     now.Unix(),
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "get_current_time",
        Description: `Lấy thời gian và ngày hiện tại.
Sử dụng tool này khi người dùng hỏi:
- Bây giờ là mấy giờ?
- Hôm nay ngày bao nhiêu?
- Thời gian ở múi giờ khác

Timezone phổ biến:
- Việt Nam: Asia/Ho_Chi_Minh
- Nhật Bản: Asia/Tokyo
- Mỹ (NY): America/New_York
- Anh: Europe/London`,
    }, handler)
}
```

### Bước 5: Tạo Unit Converter Tool

```go
func createUnitConverterTool() (tool.Tool, error) {
    type Input struct {
        Value    float64 `json:"value" description:"Giá trị cần chuyển đổi"`
        FromUnit string  `json:"from_unit" description:"Đơn vị gốc (km, m, kg, lb, celsius, fahrenheit...)"`
        ToUnit   string  `json:"to_unit" description:"Đơn vị đích"`
    }

    type Output struct {
        OriginalValue float64 `json:"original_value"`
        FromUnit      string  `json:"from_unit"`
        ConvertedValue float64 `json:"converted_value"`
        ToUnit        string  `json:"to_unit"`
        Formula       string  `json:"formula"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        var result float64
        var formula string
        from := strings.ToLower(input.FromUnit)
        to := strings.ToLower(input.ToUnit)

        switch {
        // Độ dài
        case from == "km" && to == "m":
            result = input.Value * 1000
            formula = "km × 1000 = m"
        case from == "m" && to == "km":
            result = input.Value / 1000
            formula = "m ÷ 1000 = km"
        case from == "m" && to == "cm":
            result = input.Value * 100
            formula = "m × 100 = cm"
        case from == "mile" && to == "km":
            result = input.Value * 1.60934
            formula = "mile × 1.60934 = km"
        case from == "km" && to == "mile":
            result = input.Value / 1.60934
            formula = "km ÷ 1.60934 = mile"

        // Khối lượng
        case from == "kg" && to == "lb":
            result = input.Value * 2.20462
            formula = "kg × 2.20462 = lb"
        case from == "lb" && to == "kg":
            result = input.Value / 2.20462
            formula = "lb ÷ 2.20462 = kg"
        case from == "kg" && to == "g":
            result = input.Value * 1000
            formula = "kg × 1000 = g"

        // Nhiệt độ
        case from == "celsius" && to == "fahrenheit":
            result = input.Value*9/5 + 32
            formula = "°C × 9/5 + 32 = °F"
        case from == "fahrenheit" && to == "celsius":
            result = (input.Value - 32) * 5 / 9
            formula = "(°F - 32) × 5/9 = °C"

        default:
            result = input.Value
            formula = "Không hỗ trợ chuyển đổi này"
        }

        return Output{
            OriginalValue:  input.Value,
            FromUnit:       input.FromUnit,
            ConvertedValue: result,
            ToUnit:         input.ToUnit,
            Formula:        formula,
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "convert_unit",
        Description: `Chuyển đổi giữa các đơn vị đo lường.
Hỗ trợ chuyển đổi:
- Độ dài: km ↔ m, m ↔ cm, mile ↔ km
- Khối lượng: kg ↔ lb, kg ↔ g
- Nhiệt độ: celsius ↔ fahrenheit

Ví dụ: "Đổi 100km sang mile", "5kg bằng bao nhiêu pound?"`,
    }, handler)
}
```

### Bước 6: Tạo Assistant Agent

```go
func createAssistantAgent(ctx context.Context, m model.LLM, tools []tool.Tool) (agent.Agent, error) {
    return llmagent.New(llmagent.Config{
        Name:        "personal_assistant",
        Model:       m,
        Description: "Trợ lý cá nhân với khả năng tra cứu thời tiết, tính toán, xem giờ và chuyển đổi đơn vị",
        Instruction: `Bạn là trợ lý cá nhân thông minh và hữu ích.

**CÔNG CỤ CÓ SẴN:**

1. 🌤️ get_weather - Tra cứu thời tiết
   Sử dụng khi người dùng hỏi về thời tiết, nhiệt độ của một thành phố

2. 🧮 calculator - Máy tính
   Sử dụng khi người dùng cần tính toán số học

3. 🕐 get_current_time - Xem giờ
   Sử dụng khi người dùng hỏi thời gian, ngày tháng

4. 📐 convert_unit - Chuyển đổi đơn vị
   Sử dụng khi người dùng cần đổi đơn vị đo lường

**NGUYÊN TẮC SỬ DỤNG TOOL:**

1. Phân tích câu hỏi để xác định tool cần dùng
2. Trích xuất đúng parameters từ câu hỏi
3. Gọi tool và chờ kết quả
4. Trình bày kết quả một cách thân thiện

**VÍ DỤ:**

Câu hỏi: "Thời tiết Hà Nội hôm nay thế nào?"
→ Gọi get_weather với city="Hanoi"

Câu hỏi: "15 nhân 27 bằng bao nhiêu?"
→ Gọi calculator với expression="15*27"

Câu hỏi: "Bây giờ là mấy giờ ở Tokyo?"
→ Gọi get_current_time với timezone="Asia/Tokyo"

Câu hỏi: "100 độ F là bao nhiêu độ C?"
→ Gọi convert_unit với value=100, from_unit="fahrenheit", to_unit="celsius"

**KHI CHÀO HỎI:**
Giới thiệu bản thân và các khả năng có sẵn.`,
        Tools: tools,
    })
}
```

### Bước 7: Main function

```go
func main() {
    ctx := context.Background()

    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        log.Fatal("Vui lòng set GOOGLE_API_KEY environment variable")
    }

    geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
        APIKey: apiKey,
    })
    if err != nil {
        log.Fatalf("Không thể tạo model: %v", err)
    }

    // Tạo các tools
    weatherTool, err := createWeatherTool()
    if err != nil {
        log.Fatal(err)
    }

    calculatorTool, err := createCalculatorTool()
    if err != nil {
        log.Fatal(err)
    }

    timeTool, err := createTimeTool()
    if err != nil {
        log.Fatal(err)
    }

    converterTool, err := createUnitConverterTool()
    if err != nil {
        log.Fatal(err)
    }

    // Tạo Assistant Agent với tools
    assistant, err := createAssistantAgent(ctx, geminiModel, []tool.Tool{
        weatherTool,
        calculatorTool,
        timeTool,
        converterTool,
    })
    if err != nil {
        log.Fatal(err)
    }

    config := &launcher.Config{
        AgentLoader: agent.NewSingleLoader(assistant),
    }

    lch := full.NewLauncher()
    fmt.Println("=== Personal Assistant - Tool Use Demo ===")
    fmt.Println("Tools: Weather, Calculator, Time, Unit Converter")
    fmt.Println("Khởi động server...")

    err = lch.Execute(ctx, config, os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## Phần 6: Tool Types nâng cao

### 6.1 Tool với Side Effects

Tool thực hiện hành động (gửi email, ghi database):

```go
func createEmailTool() (tool.Tool, error) {
    type Input struct {
        To      string `json:"to" description:"Địa chỉ email người nhận"`
        Subject string `json:"subject" description:"Tiêu đề email"`
        Body    string `json:"body" description:"Nội dung email"`
    }

    type Output struct {
        Success   bool   `json:"success"`
        MessageID string `json:"message_id,omitempty"`
        Error     string `json:"error,omitempty"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        // Thực hiện gửi email thực tế ở đây
        // smtp.SendMail(...)

        return Output{
            Success:   true,
            MessageID: "msg_12345",
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "send_email",
        Description: "Gửi email đến người nhận. CHÚ Ý: Tool này thực hiện hành động thực tế.",
    }, handler)
}
```

### 6.2 Tool với External API

```go
func createStockPriceTool() (tool.Tool, error) {
    type Input struct {
        Symbol string `json:"symbol" description:"Mã cổ phiếu (ví dụ: AAPL, GOOGL, VNM)"`
    }

    type Output struct {
        Symbol string  `json:"symbol"`
        Price  float64 `json:"price"`
        Change float64 `json:"change_percent"`
        Error  string  `json:"error,omitempty"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        // Gọi API thực tế
        // resp, err := http.Get("https://api.stock.com/price/" + input.Symbol)

        // Mô phỏng
        prices := map[string]float64{
            "AAPL":  178.50,
            "GOOGL": 141.20,
            "VNM":   72000,
        }

        price, exists := prices[strings.ToUpper(input.Symbol)]
        if !exists {
            return Output{
                Symbol: input.Symbol,
                Error:  "Không tìm thấy mã cổ phiếu",
            }, nil
        }

        return Output{
            Symbol: input.Symbol,
            Price:  price,
            Change: 2.5, // Mô phỏng
        }, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "get_stock_price",
        Description: "Tra cứu giá cổ phiếu hiện tại theo mã chứng khoán.",
    }, handler)
}
```

### 6.3 Tool với Validation

```go
func createValidatedTool() (tool.Tool, error) {
    type Input struct {
        Email string `json:"email" description:"Địa chỉ email cần validate"`
    }

    type Output struct {
        Valid   bool   `json:"valid"`
        Message string `json:"message"`
    }

    handler := func(ctx tool.Context, input Input) (Output, error) {
        // Validation logic
        if input.Email == "" {
            return Output{Valid: false, Message: "Email không được để trống"}, nil
        }

        if !strings.Contains(input.Email, "@") {
            return Output{Valid: false, Message: "Email không hợp lệ"}, nil
        }

        return Output{Valid: true, Message: "Email hợp lệ"}, nil
    }

    return functiontool.New(functiontool.Config{
        Name: "validate_email",
        Description: "Kiểm tra tính hợp lệ của địa chỉ email.",
    }, handler)
}
```

---

## Phần 7: Xử lý lỗi trong Tools

### 7.1 Error Handling Pattern

```go
handler := func(ctx tool.Context, input Input) (Output, error) {
    // Validation
    if input.RequiredField == "" {
        return Output{
            Success: false,
            Error:   "required_field is required",
        }, nil // Trả về output với error message, không return error
    }

    // Business logic có thể fail
    result, err := someExternalCall(input)
    if err != nil {
        // Log internal error
        log.Printf("External call failed: %v", err)

        // Trả về user-friendly message
        return Output{
            Success: false,
            Error:   "Không thể xử lý yêu cầu. Vui lòng thử lại sau.",
        }, nil
    }

    return Output{
        Success: true,
        Data:    result,
    }, nil
}
```

### 7.2 Retry Logic

```go
func withRetry(fn func() (interface{}, error), maxRetries int) (interface{}, error) {
    var lastErr error
    for i := 0; i < maxRetries; i++ {
        result, err := fn()
        if err == nil {
            return result, nil
        }
        lastErr = err
        time.Sleep(time.Duration(i+1) * time.Second) // Exponential backoff
    }
    return nil, fmt.Errorf("after %d retries: %w", maxRetries, lastErr)
}
```

---

## Phần 8: Bài tập mở rộng

### Bài tập 1: Database Tool

Tạo tool CRUD cho một entity (User, Product...):
- `create_user` - Tạo user mới
- `get_user` - Lấy thông tin user
- `update_user` - Cập nhật user
- `delete_user` - Xóa user

### Bài tập 2: File System Tool

Tạo các tools thao tác file:
- `read_file` - Đọc nội dung file
- `write_file` - Ghi nội dung file
- `list_directory` - Liệt kê files trong thư mục

### Bài tập 3: Multi-step Tool Chain

Kết hợp Tool Use với Prompt Chaining:
1. Tool 1: Lấy dữ liệu
2. Tool 2: Xử lý dữ liệu
3. Tool 3: Lưu kết quả

### Bài tập 4: Tool với Authentication

Tạo tool yêu cầu xác thực:
- Validate API key
- Rate limiting
- Permission checking

---

## Phần 9: Best Practices

### 9.1 Thiết kế Tool

1. **Single Responsibility:** Mỗi tool làm một việc cụ thể
2. **Clear Description:** Mô tả chi tiết khi nào sử dụng
3. **Typed I/O:** Sử dụng struct với JSON tags
4. **Error Messages:** Thông báo lỗi rõ ràng, actionable

### 9.2 Security

```go
// ❌ Không nên
type Input struct {
    Query string `json:"query"` // SQL injection risk
}

// ✅ Nên
type Input struct {
    UserID int    `json:"user_id"`
    Status string `json:"status"`
}
// Validate và build query an toàn trong handler
```

### 9.3 Performance

- Cache kết quả khi có thể
- Set timeout cho external calls
- Limit data returned

### 9.4 Testing

```go
func TestWeatherTool(t *testing.T) {
    tool, _ := createWeatherTool()

    // Test với valid input
    result, err := tool.Execute(ctx, map[string]interface{}{
        "city": "Hanoi",
    })
    assert.NoError(t, err)
    assert.Contains(t, result, "temperature")

    // Test với invalid input
    result, err = tool.Execute(ctx, map[string]interface{}{
        "city": "",
    })
    assert.Error(t, err)
}
```

---

## Phần 10: Code Mẫu Đầy Đủ (Solution)

Dưới đây là code hoàn chỉnh cho file `cmd/tool_use/main.go`:

```go
// Package main demonstrates Tool Use pattern using Google ADK-Go
//
// Tool Use cho phép Agent tương tác với thế giới bên ngoài thông qua
// các function được định nghĩa trước.
//
// Trong ví dụ này:
// - Weather Tool: Tra cứu thời tiết
// - Calculator Tool: Thực hiện phép tính
// - Time Tool: Lấy thời gian hiện tại
// - Unit Converter Tool: Chuyển đổi đơn vị
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// ============================================================================
// WEATHER TOOL
// ============================================================================

func createWeatherTool() (tool.Tool, error) {
	type Input struct {
		City string `json:"city" description:"Tên thành phố cần tra cứu thời tiết"`
		Unit string `json:"unit,omitempty" description:"Đơn vị nhiệt độ: celsius hoặc fahrenheit"`
	}

	type Output struct {
		City        string  `json:"city"`
		Temperature float64 `json:"temperature"`
		Unit        string  `json:"unit"`
		Condition   string  `json:"condition"`
		Humidity    int     `json:"humidity"`
		Description string  `json:"description"`
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
		weatherData := map[string]struct {
			temp      float64
			condition string
			humidity  int
		}{
			"hanoi":     {28, "Nắng", 75},
			"hochiminh": {32, "Nắng nóng", 80},
			"danang":    {30, "Có mây", 70},
			"tokyo":     {22, "Mát mẻ", 60},
			"london":    {15, "Có mưa", 85},
			"newyork":   {25, "Quang đãng", 55},
			"paris":     {18, "Nhiều mây", 65},
			"singapore": {31, "Nóng ẩm", 85},
			"seoul":     {20, "Se lạnh", 50},
		}

		city := strings.ToLower(strings.ReplaceAll(input.City, " ", ""))
		data, exists := weatherData[city]
		if !exists {
			data = struct {
				temp      float64
				condition string
				humidity  int
			}{25, "Không có dữ liệu chi tiết", 50}
		}

		unit := "°C"
		temp := data.temp
		if strings.ToLower(input.Unit) == "fahrenheit" {
			temp = data.temp*9/5 + 32
			unit = "°F"
		}

		return Output{
			City:        input.City,
			Temperature: temp,
			Unit:        unit,
			Condition:   data.condition,
			Humidity:    data.humidity,
			Description: fmt.Sprintf("🌤️ Thời tiết tại %s: %.1f%s, %s, độ ẩm %d%%",
				input.City, temp, unit, data.condition, data.humidity),
		}, nil
	}

	return functiontool.New(functiontool.Config{
		Name: "get_weather",
		Description: `Tra cứu thông tin thời tiết hiện tại của một thành phố.
Sử dụng khi người dùng hỏi về thời tiết, nhiệt độ, độ ẩm của một địa điểm.
Ví dụ: "Thời tiết Hà Nội", "Tokyo có nóng không?", "Trời London thế nào?"`,
	}, handler)
}

// ============================================================================
// CALCULATOR TOOL
// ============================================================================

func createCalculatorTool() (tool.Tool, error) {
	type Input struct {
		Expression string `json:"expression" description:"Biểu thức toán học (ví dụ: 2+3*4, sqrt(16), 10^2)"`
	}

	type Output struct {
		Expression string  `json:"expression"`
		Result     float64 `json:"result"`
		Formatted  string  `json:"formatted"`
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
		expr := strings.TrimSpace(input.Expression)
		var result float64

		switch {
		case strings.HasPrefix(expr, "sqrt("):
			var num float64
			fmt.Sscanf(expr, "sqrt(%f)", &num)
			result = math.Sqrt(num)
		case strings.HasPrefix(expr, "pow(") || strings.Contains(expr, "^"):
			var base, exp float64
			if strings.Contains(expr, "^") {
				fmt.Sscanf(expr, "%f^%f", &base, &exp)
			} else {
				fmt.Sscanf(expr, "pow(%f,%f)", &base, &exp)
			}
			result = math.Pow(base, exp)
		case strings.HasPrefix(expr, "sin("):
			var num float64
			fmt.Sscanf(expr, "sin(%f)", &num)
			result = math.Sin(num * math.Pi / 180)
		case strings.HasPrefix(expr, "cos("):
			var num float64
			fmt.Sscanf(expr, "cos(%f)", &num)
			result = math.Cos(num * math.Pi / 180)
		case strings.HasPrefix(expr, "log("):
			var num float64
			fmt.Sscanf(expr, "log(%f)", &num)
			result = math.Log10(num)
		default:
			var a, b float64
			var op rune
			for _, r := range "+-*/" {
				if strings.Contains(expr, string(r)) {
					op = r
					break
				}
			}
			parts := strings.Split(expr, string(op))
			if len(parts) == 2 {
				fmt.Sscanf(strings.TrimSpace(parts[0]), "%f", &a)
				fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &b)
				switch op {
				case '+':
					result = a + b
				case '-':
					result = a - b
				case '*':
					result = a * b
				case '/':
					if b != 0 {
						result = a / b
					}
				}
			}
		}

		return Output{
			Expression: input.Expression,
			Result:     result,
			Formatted:  fmt.Sprintf("🧮 %s = %.4g", input.Expression, result),
		}, nil
	}

	return functiontool.New(functiontool.Config{
		Name: "calculator",
		Description: `Thực hiện các phép tính toán học.
Hỗ trợ: +, -, *, /, sqrt(), pow(), sin(), cos(), log()
Ví dụ: "15*7", "sqrt(144)", "2^10", "sin(30)"`,
	}, handler)
}

// ============================================================================
// TIME TOOL
// ============================================================================

func createTimeTool() (tool.Tool, error) {
	type Input struct {
		Timezone string `json:"timezone,omitempty" description:"Múi giờ (Asia/Ho_Chi_Minh, America/New_York...)"`
	}

	type Output struct {
		Time     string `json:"time"`
		Date     string `json:"date"`
		Timezone string `json:"timezone"`
		Full     string `json:"full_description"`
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
		loc := time.UTC
		tzName := "UTC"

		if input.Timezone != "" {
			if parsedLoc, err := time.LoadLocation(input.Timezone); err == nil {
				loc = parsedLoc
				tzName = input.Timezone
			}
		}

		now := time.Now().In(loc)
		weekdays := []string{"Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"}

		return Output{
			Time:     now.Format("15:04:05"),
			Date:     now.Format("02/01/2006"),
			Timezone: tzName,
			Full: fmt.Sprintf("🕐 %s, %s - %s (%s)",
				weekdays[now.Weekday()], now.Format("02/01/2006"), now.Format("15:04:05"), tzName),
		}, nil
	}

	return functiontool.New(functiontool.Config{
		Name: "get_current_time",
		Description: `Lấy thời gian và ngày hiện tại.
Timezone: Asia/Ho_Chi_Minh (VN), Asia/Tokyo, America/New_York, Europe/London
Ví dụ: "Mấy giờ rồi?", "Bây giờ ở Tokyo là mấy giờ?"`,
	}, handler)
}

// ============================================================================
// UNIT CONVERTER TOOL
// ============================================================================

func createUnitConverterTool() (tool.Tool, error) {
	type Input struct {
		Value    float64 `json:"value" description:"Giá trị cần chuyển đổi"`
		FromUnit string  `json:"from_unit" description:"Đơn vị gốc"`
		ToUnit   string  `json:"to_unit" description:"Đơn vị đích"`
	}

	type Output struct {
		Original  string `json:"original"`
		Converted string `json:"converted"`
		Formula   string `json:"formula"`
	}

	handler := func(ctx tool.Context, input Input) (Output, error) {
		var result float64
		var formula string
		from := strings.ToLower(input.FromUnit)
		to := strings.ToLower(input.ToUnit)

		switch {
		case from == "km" && to == "m":
			result = input.Value * 1000
			formula = "× 1000"
		case from == "m" && to == "km":
			result = input.Value / 1000
			formula = "÷ 1000"
		case from == "mile" && to == "km":
			result = input.Value * 1.60934
			formula = "× 1.60934"
		case from == "km" && to == "mile":
			result = input.Value / 1.60934
			formula = "÷ 1.60934"
		case from == "kg" && to == "lb":
			result = input.Value * 2.20462
			formula = "× 2.20462"
		case from == "lb" && to == "kg":
			result = input.Value / 2.20462
			formula = "÷ 2.20462"
		case from == "celsius" && to == "fahrenheit":
			result = input.Value*9/5 + 32
			formula = "× 9/5 + 32"
		case from == "fahrenheit" && to == "celsius":
			result = (input.Value - 32) * 5 / 9
			formula = "(- 32) × 5/9"
		case from == "l" && to == "ml":
			result = input.Value * 1000
			formula = "× 1000"
		case from == "ml" && to == "l":
			result = input.Value / 1000
			formula = "÷ 1000"
		default:
			result = input.Value
			formula = "Không hỗ trợ"
		}

		return Output{
			Original:  fmt.Sprintf("%.4g %s", input.Value, input.FromUnit),
			Converted: fmt.Sprintf("📐 %.4g %s", result, input.ToUnit),
			Formula:   formula,
		}, nil
	}

	return functiontool.New(functiontool.Config{
		Name: "convert_unit",
		Description: `Chuyển đổi đơn vị đo lường.
Hỗ trợ: km↔m, mile↔km, kg↔lb, celsius↔fahrenheit, l↔ml
Ví dụ: "100km bằng bao nhiêu mile?", "30°C là bao nhiêu °F?"`,
	}, handler)
}

// ============================================================================
// ASSISTANT AGENT
// ============================================================================

func createAssistantAgent(ctx context.Context, m model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "personal_assistant",
		Model:       m,
		Description: "Trợ lý cá nhân đa năng",
		Instruction: `Bạn là trợ lý cá nhân thông minh với các công cụ:

🌤️ **get_weather** - Tra cứu thời tiết thành phố
🧮 **calculator** - Máy tính (cộng, trừ, nhân, chia, căn, lũy thừa)
🕐 **get_current_time** - Xem thời gian hiện tại
📐 **convert_unit** - Chuyển đổi đơn vị (km, mile, kg, lb, °C, °F)

**CÁCH SỬ DỤNG:**
1. Phân tích câu hỏi → chọn tool phù hợp
2. Trích xuất parameters từ câu hỏi
3. Gọi tool và trả kết quả thân thiện

**VÍ DỤ:**
- "Thời tiết Hà Nội?" → get_weather(city="Hanoi")
- "123 * 456 = ?" → calculator(expression="123*456")
- "Mấy giờ ở Tokyo?" → get_current_time(timezone="Asia/Tokyo")
- "10 mile = ? km" → convert_unit(value=10, from="mile", to="km")

**KHI CHÀO HỎI:**
"Xin chào! Tôi là trợ lý cá nhân của bạn.
Tôi có thể giúp bạn:
• 🌤️ Tra cứu thời tiết
• 🧮 Tính toán
• 🕐 Xem giờ
• 📐 Đổi đơn vị

Bạn cần giúp gì?"`,
		Tools: tools,
	})
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("Vui lòng set GOOGLE_API_KEY environment variable")
	}

	geminiModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Không thể tạo model: %v", err)
	}

	weatherTool, _ := createWeatherTool()
	calculatorTool, _ := createCalculatorTool()
	timeTool, _ := createTimeTool()
	converterTool, _ := createUnitConverterTool()

	assistant, err := createAssistantAgent(ctx, geminiModel, []tool.Tool{
		weatherTool,
		calculatorTool,
		timeTool,
		converterTool,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(assistant),
	}

	lch := full.NewLauncher()
	fmt.Println("=== Personal Assistant - Tool Use Pattern Demo ===")
	fmt.Println("Tools: Weather, Calculator, Time, Unit Converter")
	fmt.Println("Khởi động server...")

	err = lch.Execute(ctx, config, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
```

---

## Phần 11: Chạy thử và kiểm tra

### 11.1 Chuẩn bị

```bash
export GOOGLE_API_KEY="your-api-key"
go run cmd/tool_use/main.go
```

### 11.2 Test cases

```
User: Xin chào
→ Agent giới thiệu các tools có sẵn

User: Thời tiết Hà Nội thế nào?
→ Gọi get_weather, trả về thông tin thời tiết

User: Tính 15 nhân 27 cộng 100
→ Gọi calculator, trả về kết quả

User: Bây giờ là mấy giờ ở Tokyo?
→ Gọi get_current_time với timezone Asia/Tokyo

User: 100 độ F là bao nhiêu độ C?
→ Gọi convert_unit, trả về kết quả chuyển đổi
```

---

## Tài liệu tham khảo

1. [ADK-Go Tools Documentation](https://google.github.io/adk-docs/tools/)
2. [OpenAI Function Calling](https://platform.openai.com/docs/guides/function-calling)
3. [Chapter 5: Tool Use - Agentic Design Patterns](../doc_vi/05_Chapter_5_Tool_Use.md)
