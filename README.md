# Nếu chúng ta có nhiều hơn 1 module, làm sao để giao tiếp với nhau.
- khi có nhiều module, muốn các module liên kết ta store dữ liệu của 1 module trong tầng storage của module này sau đó gọi và lấy dữ liệu rồi thao tác trong business của 1 module khác 
# Giả sử module "Restaurant" cần data số lượt like từ module "Like Restaurant" thì sẽ truy xuất như thế nào?
- Khi cần lấy lượt like từ module Restaurantlike sang module Restaurant ta store dữ liệu ở tầng storage ở đây là dùng cout để tính số lần nhà hàng được like, sau đó bên business của Restaurant ta gọi lại store này lấy dữ liệu
