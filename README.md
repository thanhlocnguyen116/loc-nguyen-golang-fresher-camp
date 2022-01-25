## Index trong database : ##  
- Index là một cấu trúc dữ liệu được dùng để định vị và truy cập nhanh nhất vào dữ liệu trong các bảng database
- Index là một cách tối ưu hiệu suất truy vấn database bằng việc giảm lượng truy cập vào bộ nhớ khi thực hiện truy vấn
- Index có thể được tạo cho một hoặc nhiều cột trong database. Index thường được tạo mặc định cho primary key, foreign key. Ngoài ra, ta cũng có thể tạo thêm index cho các cột nếu cần.
- Index là một cấu trúc dữ liệu gồm:
    + Cột Search Key: chứa bản sao các giá trị của cột được tạo Index
    + Cột Data Reference: chứa con trỏ trỏ đến địa chỉ của bản ghi có giá trị cột index tương ứng

## Khác biệt khi có và không có index: ##  

- Khi không có Index cho cột Name, truy vấn sẽ phải chạy qua tất cả các Row của bảng User để so sánh và lấy ra những Row thỏa mãn. Vì vậy, khi số lượng bản ghi lớn. Index được sinh ra để giải quyết vấn đề này. Nói đơn giản, index trỏ tới địa chỉ dữ liệu trong một bảng, giống như Mục lục của một cuốn sách (Gồm tên đề mục và số trang), nó giúp truy vấn trở nên nhanh chóng như việc bạn xem mục lục và tìm đúng trang cần đọc.
- Dù Index đóng vai trò quan trọng trong việc tối ưu truy vấn và tăng tốc độ tìm kiếm trong Database nhưng nhược điểm của nó là tốn thêm bộ nhớ để lưu trữ. Do vậy, việc Index cho các cột phải được tính toán, tránh lạm dụng.
