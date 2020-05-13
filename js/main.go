package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Run(`
		abc = 2 + 2;
		console.log("The value of abc is "+ abc);
	`)
	vm.Set("i", 10)
	vm.Run(
		`
		if i < 10 {
			str = "i < 10"
		}
	`)
	object, _ := vm.Object(`{
		"type": "s",
		"name": "badboyd",
		"phone": "01655528882",
		"email": "test@chotot.vn",
		"region": 13,
		"city": 0,
		"category": "1010",
		"company_ad": true,
		"subject": "afsafas",
		"body": "🏡ĐẦU TƯ ĐÚNG CHỖ - KHÔNG LO BỊ LỖ🏡 🍀 Mở bán dự án mặt tiền Đinh Đức Thiện - với giá siêu hấp dẫn cho khách hàng muốn An Cư và nhà đầu tư. ➡Giá chỉ từ 178tr, SHR sang tên chính chủ ➡ Pháp lý: đã có sổ đỏ phân lô riêng từng nền, (thanh toán 95% nhận ngay sổ), thổ cư 70%, xây dựng tự do... ➡Vị trí: Gần KCN, nhà trọ công nhân, bệnh viện, trường học, nằm gần ngay chợ Bình Chánh. ➡ Hỗ trợ thanh toán trước 50% và trả chậm không lãi suất hoặc vay ngân hàng saccombank. ➡Cơ hộ''sleep();i nhận ngay 1 cây vàng SJC khi đặt mua dự án sớm... ➡Cam kết rẻ hơn thị trường hiện tại đang bán ra 100%. ☎️☎️A/C có thiện chí vui lòng liên hệ trực tiếp chủ đầu tư qua sdt này. 🚗Có xe ô tô đưa đón khách miễn phí khi tham quan dự án vào ngày mở bán ➡Theo dõi quy mô và vị trí dự án ĐINH ĐỨC THIỆN-BÌNH CHÁNH qua hình ảnh THẬT",
		"price": 3500000000,
		"orig_price": "35000000",
		"account_id": 4443367,
		"lang": "vi",
		"source": "desktop_site_flashad",
		"needs_payment": "0",
		"verified_account": "1",
		"action": ""
	}`)
	_, err := vm.Call(`
		if object.phone == "" {
			object.action = "refused"
		}
		`, nil, object)

	fmt.Println(err)
}
