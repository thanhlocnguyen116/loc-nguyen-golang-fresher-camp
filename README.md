## Khoá chính (Primary Key): ## 
- Khóa chính (hay ràng buộc khóa chính) được sử dụng để định danh duy nhất mỗi record tại table của cơ sở dữ liệu.  
- Ngoài ra, nó còn sử dụng để thiết lập quan hệ 1-n (hay ràng buộc tham chiếu) giữa hai table tại cơ sở dữ liệu.  
- Dữ liệu (value) của field khóa chính nên có tính độc nhất. và không có các giá trị Null.  
- Mỗi table nên chỉ có một khóa chủ đạo, khóa chủ đạo có khả năng sản sinh ra từ nhiều field của table.

## Giá trị của ID tự tăng vì : ##  
- AUTO_INCREMENT chỉ thiết lập được cho kiểu INT và mỗi bảng chỉ có một field duy nhất.
- Khi bạn thêm dữ liệu nếu bạn có truyền data thì nó sẽ lấy data đó thay vì tăng tự động, ngược lại nó sẽ lấy giá trị lớn nhất hiện tại và tăng lên 1.
- Khi bạn xóa một record thì sẽ bị khuyết mất một giá trị, lúc này nếu bạn thêm thì nó sẽ không lấp vào vị trí này mà nó tuân theo quy luật trên.
- Giả sử giá trị 120 là lớn nhất, bạn xóa đi 120 thì lúc này lớn nhất là 119. Lúc này nếu ban thêm mới thì nó sẽ lấy 121 chứ không phải là 120 vì giá trị lớn nhất nó lưu trong config của table.

## Dùng nhiều khoá chính (Primary Key) khi: ##  
- Cần đảm bảo tính unique của dữ liệu
- Có thể loại bỏ sự cần thiết phải tham gia trong một số bảng con
- Có nhiều định danh hơn cho dữ liệu
