# Trong trường hợp tạo cột đếm thì làm sao để update cột đó? 
- Tại tầng storage ta dùng thư viện GORM vời câu lệnh:" db.<Tên model>(&tên cột).Updates(map[string]interface{}{"tên cột": gorm.Expr("tên cột+ ?", 1)}) " để update giá trị cho cột
# Làm sao để API chính không bị block vì phải update số đếm?
- Bắt lỗi tại tầng transport, cho nó panic khi gặp lỗi, trả về lỗi, dùng middleware Recovery để tránh crash
