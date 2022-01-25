## Khoá chính (Primary Key): ## 
    - Khóa chính (hay ràng buộc khóa chính) được sử dụng để định danh duy nhất mỗi record tại table của cơ sở dữ liệu.  
    - Ngoài ra, nó còn sử dụng để thiết lập quan hệ 1-n (hay ràng buộc tham chiếu) giữa hai table tại cơ sở dữ liệu.  
    - Dữ liệu (value) của field khóa chính nên có tính độc nhất. và không có các giá trị Null.  
    - Mỗi table nên chỉ có một khóa chủ đạo, khóa chủ đạo có khả năng sản sinh ra từ nhiều field của table.
## Giá trị của ID tự tăng vì nó đảm bảo tính uniqe của khoá chính, không có ID nào được giống nhau ##  
## Dùng nhiều khoá chính (Primary Key) khi: ##  
    - Cần đảm bảo tính unique của dữ liệu
    - Có thể loại bỏ sự cần thiết phải tham gia trong một số bảng con
    - Có nhiều định danh hơn cho dữ liệu
