# Chương 9: Learning and Adaptation (Học Hỏi và Thích Nghi)

Học hỏi và thích nghi là yếu tố then chốt để nâng cao khả năng của các tác nhân trí tuệ nhân tạo. Các quy trình này cho phép các tác nhân phát triển vượt ra ngoài các tham số được xác định trước, cho phép chúng tự chủ cải thiện thông qua kinh nghiệm và tương tác với môi trường. Bằng cách học hỏi và thích nghi, các tác nhân có thể quản lý hiệu quả các tình huống mới lạ và tối ưu hóa hiệu suất của chúng mà không cần can thiệp thủ công liên tục. Chương này khám phá các nguyên tắc và cơ chế củng cố việc học hỏi và thích nghi của tác nhân một cách chi tiết.

## Bức tranh toàn cảnh

Các tác nhân học hỏi và thích nghi bằng cách thay đổi suy nghĩ, hành động hoặc kiến thức của chúng dựa trên những trải nghiệm và dữ liệu mới. Điều này cho phép các tác nhân phát triển từ việc chỉ đơn giản làm theo các hướng dẫn sang trở nên thông minh hơn theo thời gian.

*   **Học Tăng cường (Reinforcement Learning - RL):** Các tác nhân thử các hành động và nhận phần thưởng cho các kết quả tích cực và hình phạt cho các kết quả tiêu cực, học các hành vi tối ưu trong các tình huống thay đổi. Hữu ích cho các tác nhân điều khiển robot hoặc chơi trò chơi.
*   **Học có Giám sát (Supervised Learning):** Các tác nhân học từ các ví dụ được dán nhãn, kết nối đầu vào với đầu ra mong muốn, cho phép các tác vụ như ra quyết định và nhận dạng mẫu. Lý tưởng cho các tác nhân phân loại email hoặc dự đoán xu hướng.
*   **Học không Giám sát (Unsupervised Learning):** Các tác nhân khám phá các kết nối và mẫu ẩn trong dữ liệu không được dán nhãn, hỗ trợ việc tìm hiểu sâu, tổ chức và tạo bản đồ tư duy về môi trường của chúng. Hữu ích cho các tác nhân khám phá dữ liệu mà không cần hướng dẫn cụ thể.
*   **Học Few-Shot/Zero-Shot với Tác nhân Dựa trên LLM:** Các tác nhân tận dụng LLM có thể nhanh chóng thích nghi với các tác vụ mới với các ví dụ tối thiểu hoặc hướng dẫn rõ ràng, cho phép phản ứng nhanh với các lệnh hoặc tình huống mới.
*   **Học Trực tuyến (Online Learning):** Các tác nhân liên tục cập nhật kiến thức với dữ liệu mới, điều cần thiết cho các phản ứng thời gian thực và thích nghi liên tục trong các môi trường động. Quan trọng cho các tác nhân xử lý các luồng dữ liệu liên tục.
*   **Học Dựa trên Bộ nhớ (Memory-Based Learning):** Các tác nhân nhớ lại các trải nghiệm trong quá khứ để điều chỉnh các hành động hiện tại trong các tình huống tương tự, nâng cao nhận thức ngữ cảnh và ra quyết định. Hiệu quả cho các tác nhân có khả năng nhớ lại bộ nhớ.

Các tác nhân thích nghi bằng cách thay đổi chiến lược, sự hiểu biết hoặc mục tiêu dựa trên việc học. Điều này rất quan trọng đối với các tác nhân trong các môi trường không thể đoán trước, thay đổi hoặc mới.

## Các Kỹ thuật Cốt lõi

### Proximal Policy Optimization (PPO)
Proximal Policy Optimization (PPO) là một thuật toán học tăng cường được sử dụng để huấn luyện các tác nhân trong các môi trường với phạm vi hành động liên tục, như điều khiển các khớp của robot hoặc một nhân vật trong trò chơi. Mục tiêu chính của nó là cải thiện một cách đáng tin cậy và ổn định chiến lược ra quyết định của tác nhân, được gọi là chính sách (policy).

Ý tưởng cốt lõi đằng sau PPO là thực hiện các cập nhật nhỏ, cẩn thận cho chính sách của tác nhân. Nó tránh những thay đổi mạnh mẽ có thể gây ra sự sụp đổ về hiệu suất. Đây là cách nó hoạt động:
1.  **Thu thập Dữ liệu:** Tác nhân tương tác với môi trường của nó (ví dụ: chơi một trò chơi) bằng cách sử dụng chính sách hiện tại của nó và thu thập một loạt các trải nghiệm (trạng thái, hành động, phần thưởng).
2.  **Đánh giá Mục tiêu "Thay thế":** PPO tính toán xem một bản cập nhật chính sách tiềm năng sẽ thay đổi phần thưởng mong đợi như thế nào. Tuy nhiên, thay vì chỉ tối đa hóa phần thưởng này, nó sử dụng một hàm mục tiêu "được cắt" (clipped) đặc biệt.
3.  **Cơ chế "Cắt" (Clipping):** Đây là chìa khóa cho sự ổn định của PPO. Nó tạo ra một "vùng tin cậy" hoặc một vùng an toàn xung quanh chính sách hiện tại. Thuật toán bị ngăn cản thực hiện một bản cập nhật quá khác biệt so với chiến lược hiện tại. Việc cắt này hoạt động giống như một phanh an toàn, đảm bảo tác nhân không thực hiện một bước đi lớn, rủi ro làm hỏng quá trình học tập của nó.

Tóm lại, PPO cân bằng việc cải thiện hiệu suất với việc giữ gần một chiến lược đã biết, hoạt động tốt, ngăn ngừa các thất bại thảm khốc trong quá trình huấn luyện và dẫn đến việc học tập ổn định hơn.

### Direct Preference Optimization (DPO)
Direct Preference Optimization (DPO) là một phương pháp gần đây hơn được thiết kế đặc biệt để căn chỉnh các Mô hình Ngôn ngữ Lớn (LLM) với sở thích của con người. Nó cung cấp một giải pháp thay thế đơn giản hơn, trực tiếp hơn cho việc sử dụng PPO cho nhiệm vụ này.

Để hiểu DPO, trước tiên cần hiểu phương pháp căn chỉnh dựa trên PPO truyền thống:
*   **Cách tiếp cận PPO (Quy trình Hai bước):**
    1.  **Huấn luyện Mô hình Phần thưởng:** Đầu tiên, bạn thu thập dữ liệu phản hồi của con người, nơi mọi người đánh giá hoặc so sánh các phản hồi LLM khác nhau (ví dụ: "Phản hồi A tốt hơn Phản hồi B"). Dữ liệu này được sử dụng để huấn luyện một mô hình AI riêng biệt, được gọi là mô hình phần thưởng, có nhiệm vụ dự đoán điểm số mà con người sẽ đưa ra cho bất kỳ phản hồi mới nào.
    2.  **Tinh chỉnh với PPO:** Tiếp theo, LLM được tinh chỉnh bằng cách sử dụng PPO. Mục tiêu của LLM là tạo ra các phản hồi nhận được điểm số cao nhất có thể từ mô hình phần thưởng. Mô hình phần thưởng đóng vai trò là "giám khảo" trong trò chơi huấn luyện.

Quy trình hai bước này có thể phức tạp và không ổn định. Ví dụ, LLM có thể tìm ra một lỗ hổng và học cách "hack" mô hình phần thưởng để đạt điểm cao cho các phản hồi tồi.

*   **Cách tiếp cận DPO (Quy trình Trực tiếp):** DPO bỏ qua hoàn toàn mô hình phần thưởng. Thay vì dịch sở thích của con người thành điểm thưởng và sau đó tối ưu hóa cho điểm đó, DPO sử dụng dữ liệu sở thích trực tiếp để cập nhật chính sách của LLM.
*   Nó hoạt động bằng cách sử dụng một mối quan hệ toán học liên kết trực tiếp dữ liệu sở thích với chính sách tối ưu. Về cơ bản, nó dạy mô hình: "Tăng xác suất tạo ra các phản hồi giống như phản hồi được ưa thích và giảm xác suất tạo ra các phản hồi giống như phản hồi bị loại bỏ."

Về bản chất, DPO đơn giản hóa việc căn chỉnh bằng cách tối ưu hóa trực tiếp mô hình ngôn ngữ dựa trên dữ liệu sở thích của con người. Điều này tránh được sự phức tạp và bất ổn tiềm tàng của việc huấn luyện và sử dụng một mô hình phần thưởng riêng biệt, làm cho quá trình căn chỉnh hiệu quả và mạnh mẽ hơn.

## Các Ứng dụng Thực tế & Trường hợp Sử dụng

Các tác nhân thích nghi thể hiện hiệu suất nâng cao trong các môi trường biến đổi thông qua các cập nhật lặp lại được thúc đẩy bởi dữ liệu kinh nghiệm.

*   **Tác nhân trợ lý cá nhân hóa:** Tinh chỉnh các giao thức tương tác thông qua phân tích dọc hành vi của từng người dùng, đảm bảo tạo phản hồi được tối ưu hóa cao.
*   **Tác nhân bot giao dịch:** Tối ưu hóa các thuật toán ra quyết định bằng cách điều chỉnh động các tham số mô hình dựa trên dữ liệu thị trường thời gian thực, độ phân giải cao, do đó tối đa hóa lợi nhuận tài chính và giảm thiểu các yếu tố rủi ro.
*   **Tác nhân ứng dụng:** Tối ưu hóa giao diện người dùng và chức năng thông qua sửa đổi động dựa trên hành vi người dùng quan sát được, dẫn đến tăng sự tham gia của người dùng và tính trực quan của hệ thống.
*   **Tác nhân xe tự lái và robot:** Nâng cao khả năng điều hướng và phản ứng bằng cách tích hợp dữ liệu cảm biến và phân tích hành động lịch sử, cho phép hoạt động an toàn và hiệu quả trong các điều kiện môi trường đa dạng.
*   **Tác nhân phát hiện gian lận:** Cải thiện việc phát hiện bất thường bằng cách tinh chỉnh các mô hình dự đoán với các mẫu gian lận mới được xác định, tăng cường bảo mật hệ thống và giảm thiểu tổn thất tài chính.
*   **Tác nhân đề xuất:** Cải thiện độ chính xác lựa chọn nội dung bằng cách sử dụng các thuật toán học sở thích người dùng, cung cấp các đề xuất phù hợp theo ngữ cảnh và được cá nhân hóa cao.
*   **Tác nhân Game AI:** Tăng cường sự tham gia của người chơi bằng cách thích nghi động các thuật toán chiến lược, do đó làm tăng độ phức tạp và thách thức của trò chơi.
*   **Tác nhân Học Cơ sở Tri thức:** Các tác nhân có thể tận dụng Retrieval Augmented Generation (RAG) để duy trì một cơ sở tri thức động về các mô tả vấn đề và giải pháp đã được chứng minh (xem Chương 14). Bằng cách lưu trữ các chiến lược thành công và những thách thức gặp phải, tác nhân có thể tham khảo dữ liệu này trong quá trình ra quyết định, cho phép nó thích nghi với các tình huống mới hiệu quả hơn bằng cách áp dụng các mẫu thành công trước đó hoặc tránh các cạm bẫy đã biết.

## Nghiên cứu Tình huống: Tác nhân Lập trình Tự Cải thiện (SICA)

Tác nhân Lập trình Tự Cải thiện (SICA - Self-Improving Coding Agent), được phát triển bởi Maxime Robeyns, Laurence Aitchison và Martin Szummer, đại diện cho một sự tiến bộ trong học tập dựa trên tác nhân, chứng minh khả năng của một tác nhân trong việc sửa đổi mã nguồn của chính nó. Điều này trái ngược với các phương pháp truyền thống nơi một tác nhân có thể huấn luyện một tác nhân khác; SICA đóng vai trò là cả thực thể sửa đổi và thực thể được sửa đổi, liên tục tinh chỉnh cơ sở mã của nó để cải thiện hiệu suất qua các thách thức lập trình khác nhau.

Sự tự cải thiện của SICA hoạt động thông qua một chu trình lặp lại. Ban đầu, SICA xem xét một kho lưu trữ các phiên bản trước đây của nó và hiệu suất của chúng trên các bài kiểm tra điểm chuẩn. Nó chọn phiên bản có điểm hiệu suất cao nhất, được tính toán dựa trên công thức có trọng số xem xét thành công, thời gian và chi phí tính toán. Phiên bản được chọn này sau đó thực hiện vòng tự sửa đổi tiếp theo. Nó phân tích kho lưu trữ để xác định các cải tiến tiềm năng và sau đó trực tiếp thay đổi cơ sở mã của nó. Tác nhân đã sửa đổi sau đó được kiểm tra lại so với các điểm chuẩn, với kết quả được ghi lại trong kho lưu trữ. Quá trình này lặp lại, tạo điều kiện cho việc học trực tiếp từ hiệu suất trong quá khứ. Cơ chế tự cải thiện này cho phép SICA phát triển khả năng của mình mà không cần các mô hình đào tạo truyền thống.

SICA đã trải qua quá trình tự cải thiện đáng kể, dẫn đến những tiến bộ trong việc chỉnh sửa mã và điều hướng. Ban đầu, SICA sử dụng một phương pháp ghi đè tệp cơ bản cho các thay đổi mã. Sau đó, nó đã phát triển một "Trình chỉnh sửa Thông minh" (Smart Editor) có khả năng chỉnh sửa thông minh và theo ngữ cảnh hơn. Điều này đã phát triển thành một "Trình chỉnh sửa Thông minh Nâng cao bằng Diff" (Diff-Enhanced Smart Editor), kết hợp các diffs cho các sửa đổi được nhắm mục tiêu và chỉnh sửa dựa trên mẫu, và một "Công cụ Ghi đè Nhanh" (Quick Overwrite Tool) để giảm nhu cầu xử lý.

SICA tiếp tục triển khai "Tối ưu hóa Đầu ra Diff Tối thiểu" và "Tối thiểu hóa Diff Nhạy cảm Ngữ cảnh", sử dụng phân tích Cây Cú pháp Trừu tượng (AST) để đạt hiệu quả. Ngoài ra, một "Trình chuẩn hóa Đầu vào SmartEditor" đã được thêm vào. Về điều hướng, SICA đã độc lập tạo ra một "Bộ định vị Ký hiệu AST", sử dụng bản đồ cấu trúc của mã (AST) để xác định các định nghĩa trong cơ sở mã. Sau đó, một "Bộ định vị Ký hiệu Lai" đã được phát triển, kết hợp tìm kiếm nhanh với kiểm tra AST. Điều này đã được tối ưu hóa hơn nữa thông qua "Phân tích AST Tối ưu hóa trong Bộ định vị Ký hiệu Lai" để tập trung vào các phần mã liên quan, cải thiện tốc độ tìm kiếm.

Kiến trúc của SICA bao gồm một bộ công cụ nền tảng cho các hoạt động tệp cơ bản, thực thi lệnh và tính toán số học. Nó bao gồm các cơ chế để gửi kết quả và gọi các tác nhân phụ chuyên biệt (lập trình, giải quyết vấn đề và suy luận). Các tác nhân phụ này phân rã các tác vụ phức tạp và quản lý độ dài ngữ cảnh của LLM, đặc biệt là trong các chu kỳ cải thiện mở rộng.

Một người giám sát bất đồng bộ (asynchronous overseer), một LLM khác, giám sát hành vi của SICA, xác định các vấn đề tiềm ẩn như vòng lặp hoặc đình trệ. Nó giao tiếp với SICA và có thể can thiệp để tạm dừng thực thi nếu cần thiết. Người giám sát nhận được một báo cáo chi tiết về các hành động của SICA, bao gồm một đồ thị cuộc gọi (callgraph) và nhật ký các tin nhắn và hành động công cụ, để xác định các mẫu và sự kém hiệu quả.

LLM của SICA tổ chức thông tin trong cửa sổ ngữ cảnh của nó, bộ nhớ ngắn hạn của nó, theo một cách có cấu trúc quan trọng cho hoạt động của nó. Cấu trúc này bao gồm một Prompt Hệ thống xác định mục tiêu của tác nhân, tài liệu về công cụ và tác nhân phụ, và hướng dẫn hệ thống. Một Prompt Cốt lõi chứa tuyên bố vấn đề hoặc hướng dẫn, nội dung của các tệp đang mở và bản đồ thư mục. Các Tin nhắn Trợ lý ghi lại lý luận từng bước của tác nhân, hồ sơ cuộc gọi và kết quả của công cụ và tác nhân phụ, và thông tin liên lạc của người giám sát. Tổ chức này tạo điều kiện cho luồng thông tin hiệu quả, nâng cao hoạt động của LLM và giảm thời gian và chi phí xử lý. Ban đầu, các thay đổi tệp được ghi lại dưới dạng diffs, chỉ hiển thị các sửa đổi và được hợp nhất định kỳ.

Dự án đang được phát triển tích cực và nhằm mục đích cung cấp một khuôn khổ mạnh mẽ cho những người quan tâm đến việc đào tạo sau (post-training) các LLM về sử dụng công cụ và các tác vụ agentic khác, với mã nguồn đầy đủ có sẵn để khám phá và đóng góp thêm tại kho lưu trữ GitHub: https://github.com/MaximeRobeyns/self_improving_coding_agent/

Về bảo mật, dự án nhấn mạnh mạnh mẽ vào việc container hóa Docker, nghĩa là tác nhân chạy trong một container Docker chuyên dụng. Đây là một biện pháp quan trọng, vì nó cung cấp sự cách ly khỏi máy chủ, giảm thiểu rủi ro như thao tác hệ thống tệp vô tình do khả năng thực thi các lệnh shell của tác nhân.

Để đảm bảo tính minh bạch và kiểm soát, hệ thống có tính năng quan sát mạnh mẽ thông qua một trang web tương tác trực quan hóa các sự kiện trên bus sự kiện và đồ thị cuộc gọi của tác nhân. Điều này cung cấp thông tin chi tiết toàn diện về các hành động của tác nhân, cho phép người dùng kiểm tra các sự kiện riêng lẻ, đọc tin nhắn của người giám sát và thu gọn các dấu vết tác nhân phụ để hiểu rõ hơn.

Về trí tuệ cốt lõi của nó, khuôn khổ tác nhân hỗ trợ tích hợp LLM từ nhiều nhà cung cấp khác nhau, cho phép thử nghiệm với các mô hình khác nhau để tìm ra sự phù hợp nhất cho các tác vụ cụ thể. Cuối cùng, một thành phần quan trọng là người giám sát bất đồng bộ, một LLM chạy đồng thời với tác nhân chính. Người giám sát này định kỳ đánh giá hành vi của tác nhân để tìm các sai lệch bệnh lý hoặc sự đình trệ và có thể can thiệp bằng cách gửi thông báo hoặc thậm chí hủy bỏ việc thực thi của tác nhân nếu cần thiết. Nó nhận được một biểu diễn văn bản chi tiết về trạng thái của hệ thống, bao gồm một đồ thị cuộc gọi và một luồng sự kiện của các tin nhắn LLM, các cuộc gọi công cụ và phản hồi, cho phép nó phát hiện các mẫu không hiệu quả hoặc công việc lặp đi lặp lại.

Một thách thức đáng chú ý trong việc triển khai SICA ban đầu là thúc đẩy tác nhân dựa trên LLM độc lập đề xuất các sửa đổi mới lạ, sáng tạo, khả thi và hấp dẫn trong mỗi lần lặp lại siêu cải thiện (meta-improvement). Hạn chế này, đặc biệt là trong việc thúc đẩy việc học mở và sự sáng tạo đích thực trong các tác nhân LLM, vẫn là một lĩnh vực điều tra chính trong nghiên cứu hiện tại.

## AlphaEvolve và OpenEvolve

**AlphaEvolve** là một tác nhân AI được phát triển bởi Google được thiết kế để khám phá và tối ưu hóa các thuật toán. Nó sử dụng sự kết hợp của các LLM, cụ thể là các mô hình Gemini (Flash và Pro), các hệ thống đánh giá tự động và một khuôn khổ thuật toán tiến hóa. Hệ thống này nhằm mục đích thúc đẩy cả toán học lý thuyết và các ứng dụng tính toán thực tế.

AlphaEvolve sử dụng một tập hợp các mô hình Gemini. Flash được sử dụng để tạo ra một loạt các đề xuất thuật toán ban đầu, trong khi Pro cung cấp phân tích và tinh chỉnh chuyên sâu hơn. Các thuật toán được đề xuất sau đó được đánh giá và chấm điểm tự động dựa trên các tiêu chí được xác định trước. Đánh giá này cung cấp phản hồi được sử dụng để cải thiện lặp lại các giải pháp, dẫn đến các thuật toán mới lạ và được tối ưu hóa.

Trong tính toán thực tế, AlphaEvolve đã được triển khai trong cơ sở hạ tầng của Google. Nó đã chứng minh những cải tiến trong lập lịch trung tâm dữ liệu, dẫn đến giảm 0,7% mức sử dụng tài nguyên tính toán toàn cầu. Nó cũng đã đóng góp vào thiết kế phần cứng bằng cách đề xuất tối ưu hóa cho mã Verilog trong các Đơn vị Xử lý Tensor (TPU) sắp tới. Hơn nữa, AlphaEvolve đã tăng tốc hiệu suất AI, bao gồm cải thiện tốc độ 23% trong nhân cốt lõi của kiến trúc Gemini và tối ưu hóa lên đến 32,5% các lệnh GPU cấp thấp cho FlashAttention.

Trong lĩnh vực nghiên cứu cơ bản, AlphaEvolve đã góp phần khám phá các thuật toán mới cho phép nhân ma trận, bao gồm một phương pháp cho các ma trận giá trị phức 4x4 sử dụng 48 phép nhân vô hướng, vượt qua các giải pháp đã biết trước đó. Trong nghiên cứu toán học rộng hơn, nó đã khám phá lại các giải pháp hiện đại hiện có cho hơn 50 bài toán mở trong 75% trường hợp và cải thiện các giải pháp hiện có trong 20% trường hợp, với các ví dụ bao gồm những tiến bộ trong bài toán số hôn (kissing number problem).

**OpenEvolve** là một tác nhân mã hóa tiến hóa tận dụng các LLM để tối ưu hóa mã lặp đi lặp lại. Nó điều phối một đường ống tạo mã, đánh giá và lựa chọn do LLM điều khiển để liên tục nâng cao các chương trình cho một loạt các tác vụ. Một khía cạnh chính của OpenEvolve là khả năng phát triển toàn bộ các tệp mã, thay vì bị giới hạn ở các hàm đơn lẻ. Tác nhân được thiết kế để linh hoạt, cung cấp hỗ trợ cho nhiều ngôn ngữ lập trình và khả năng tương thích với các API tương thích OpenAI cho bất kỳ LLM nào. Hơn nữa, nó kết hợp tối ưu hóa đa mục tiêu, cho phép kỹ thuật prompt linh hoạt và có khả năng đánh giá phân tán để xử lý hiệu quả các thách thức mã hóa phức tạp.

Kiến trúc nội bộ của OpenEvolve được quản lý bởi một bộ điều khiển (controller). Bộ điều khiển này điều phối một số thành phần chính: trình lấy mẫu chương trình (program sampler), Cơ sở dữ liệu Chương trình, Nhóm Người đánh giá (Evaluator Pool) và Tập hợp LLM. Chức năng chính của nó là tạo điều kiện thuận lợi cho các quá trình học tập và thích nghi của chúng để nâng cao chất lượng mã.

Đoạn mã này sử dụng thư viện OpenEvolve để thực hiện tối ưu hóa tiến hóa trên một chương trình. Nó khởi tạo hệ thống OpenEvolve với đường dẫn đến chương trình ban đầu, tệp đánh giá và tệp cấu hình. Dòng `evolve.run(iterations=1000)` bắt đầu quá trình tiến hóa, chạy trong 1000 lần lặp để tìm phiên bản cải tiến của chương trình. Cuối cùng, nó in các số liệu của chương trình tốt nhất được tìm thấy trong quá trình tiến hóa, được định dạng đến bốn chữ số thập phân.

```python
from openevolve import OpenEvolve

# Khởi tạo hệ thống
evolve = OpenEvolve(
    initial_program_path="path/to/initial_program.py",
    evaluation_file="path/to/evaluator.py",
    config_path="path/to/config.yaml"
)

# Chạy quá trình tiến hóa
best_program = await evolve.run(iterations=1000)

print(f"Best program metrics:")
for name, value in best_program.metrics.items():
    print(f"  {name}: {value:.4f}")
```

## Tóm tắt Nhanh (At a Glance)

*   **Cái gì:** Các tác nhân AI thường hoạt động trong các môi trường năng động và không thể đoán trước, nơi logic được lập trình trước là không đủ. Hiệu suất của chúng có thể suy giảm khi đối mặt với các tình huống mới lạ không được dự đoán trong thiết kế ban đầu của chúng. Nếu không có khả năng học hỏi từ kinh nghiệm, các tác nhân không thể tối ưu hóa chiến lược hoặc cá nhân hóa các tương tác của chúng theo thời gian. Sự cứng nhắc này hạn chế hiệu quả của chúng và ngăn chúng đạt được quyền tự chủ thực sự trong các kịch bản thế giới thực phức tạp.
*   **Tại sao:** Giải pháp tiêu chuẩn hóa là tích hợp các cơ chế học hỏi và thích nghi, biến các tác nhân tĩnh thành các hệ thống năng động, phát triển. Điều này cho phép một tác nhân tự chủ tinh chỉnh kiến thức và hành vi của mình dựa trên dữ liệu và tương tác mới. Các hệ thống Agentic có thể sử dụng nhiều phương pháp khác nhau, từ học tăng cường đến các kỹ thuật tiên tiến hơn như tự sửa đổi, như được thấy trong Tác nhân Lập trình Tự Cải thiện (SICA). Các hệ thống tiên tiến như AlphaEvolve của Google tận dụng LLM và các thuật toán tiến hóa để khám phá các giải pháp hoàn toàn mới và hiệu quả hơn cho các vấn đề phức tạp. Bằng cách liên tục học hỏi, các tác nhân có thể làm chủ các nhiệm vụ mới, nâng cao hiệu suất và thích nghi với các điều kiện thay đổi mà không cần lập trình lại thủ công liên tục.
*   **Quy tắc ngón tay cái:** Sử dụng mẫu này khi xây dựng các tác nhân phải hoạt động trong các môi trường năng động, không chắc chắn hoặc đang phát triển. Nó rất cần thiết cho các ứng dụng yêu cầu cá nhân hóa, cải thiện hiệu suất liên tục và khả năng xử lý các tình huống mới lạ một cách tự chủ.

## Những Điểm Chính (Key Takeaways)

*   Học hỏi và Thích nghi là về việc các tác nhân trở nên tốt hơn trong những gì chúng làm và xử lý các tình huống mới bằng cách sử dụng kinh nghiệm của chúng.
*   "Thích nghi" là sự thay đổi rõ ràng trong hành vi hoặc kiến thức của tác nhân đến từ việc học hỏi.
*   SICA, Tác nhân Lập trình Tự Cải thiện, tự cải thiện bằng cách sửa đổi mã của nó dựa trên hiệu suất trong quá khứ. Điều này dẫn đến các công cụ như Smart Editor và AST Symbol Locator.
*   Có các "tác nhân phụ" chuyên biệt và một "người giám sát" giúp các hệ thống tự cải thiện này quản lý các nhiệm vụ lớn và đi đúng hướng.
*   Cách thiết lập "cửa sổ ngữ cảnh" của một LLM (với prompt hệ thống, prompt cốt lõi và tin nhắn trợ lý) là cực kỳ quan trọng đối với mức độ hiệu quả của các tác nhân làm việc.
*   Mẫu này rất quan trọng đối với các tác nhân cần hoạt động trong các môi trường luôn thay đổi, không chắc chắn hoặc yêu cầu sự tiếp xúc cá nhân.
*   Xây dựng các tác nhân học hỏi thường có nghĩa là kết nối chúng với các công cụ học máy và quản lý cách dữ liệu chảy.
*   Một hệ thống tác nhân, được trang bị các công cụ mã hóa cơ bản, có thể tự chủ chỉnh sửa chính nó và do đó cải thiện hiệu suất của nó trên các tác vụ điểm chuẩn.
*   AlphaEvolve là tác nhân AI của Google tận dụng các LLM và một khuôn khổ tiến hóa để tự chủ khám phá và tối ưu hóa các thuật toán, nâng cao đáng kể cả nghiên cứu cơ bản và các ứng dụng tính toán thực tế.

## Kết luận

Chương này kiểm tra các vai trò quan trọng của học hỏi và thích nghi trong Trí tuệ Nhân tạo. Các tác nhân AI nâng cao hiệu suất của chúng thông qua việc thu thập dữ liệu liên tục và kinh nghiệm. Tác nhân Lập trình Tự Cải thiện (SICA) minh họa điều này bằng cách tự chủ cải thiện khả năng của nó thông qua các sửa đổi mã.

Chúng ta đã xem xét các thành phần cơ bản của AI agentic, bao gồm kiến trúc, ứng dụng, lập kế hoạch, hợp tác đa tác nhân, quản lý bộ nhớ, và học hỏi và thích nghi. Các nguyên tắc học hỏi đặc biệt quan trọng để cải thiện phối hợp trong các hệ thống đa tác nhân. Để đạt được điều này, dữ liệu điều chỉnh (tuning data) phải phản ánh chính xác quỹ đạo tương tác hoàn chỉnh, nắm bắt các đầu vào và đầu ra riêng lẻ của từng tác nhân tham gia.

Các yếu tố này đóng góp vào những tiến bộ đáng kể, chẳng hạn như AlphaEvolve của Google. Hệ thống AI này độc lập khám phá và tinh chỉnh các thuật toán bằng LLM, đánh giá tự động và phương pháp tiếp cận tiến hóa, thúc đẩy tiến bộ trong nghiên cứu khoa học và các kỹ thuật tính toán. Các mẫu như vậy có thể được kết hợp để xây dựng các hệ thống AI tinh vi. Những phát triển như AlphaEvolve chứng minh rằng việc khám phá và tối ưu hóa thuật toán tự chủ bởi các tác nhân AI là có thể đạt được.

## Tài liệu tham khảo
1.  Sutton, R. S., & Barto, A. G. (2018). Reinforcement Learning: An Introduction. MIT Press.
2.  Goodfellow, I., Bengio, Y., & Courville, A. (2016). Deep Learning. MIT Press.
3.  Mitchell, T. M. (1997). Machine Learning. McGraw-Hill.
4.  Proximal Policy Optimization Algorithms by John Schulman, Filip Wolski, Prafulla Dhariwal, Alec Radford, and Oleg Klimov. You can find it on arXiv: https://arxiv.org/abs/1707.06347
5.  Robeyns, M., Aitchison, L., & Szummer, M. (2025). A Self-Improving Coding Agent. arXiv:2504.15228v2. https://arxiv.org/pdf/2504.15228 https://github.com/MaximeRobeyns/self_improving_coding_agent
6.  AlphaEvolve blog, https://deepmind.google/discover/blog/alphaevolve-a-gemini-powered-coding-agent-for-designing-advanced-algorithms/
7.  OpenEvolve, https://github.com/codelion/openevolve