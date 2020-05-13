package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/antonmedv/expr"
)

var (
	ad = `{
		"ad": {
			"type": "s",
			"name": "badboyd",
			"phone": "01655528881",
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
			"verified_account": "1"
		},
		"params": {
			"area": "111",
			"address": "130-132, Đường Hồng Hà, Phú Nhuận, Hồ Chí Minh",
			"apartment_type": "4",
			"size": "200",
			"projectimages": "14_overview_6,14_floor_plan_project_7,14_process_11",
			"property_status": "1",
			"rooms": "4"
		}
	}`
)

type mapStringInterface map[string]interface{}

func (m mapStringInterface) Get(name string) (interface{}, error) {
	return findInMap(name, map[string]interface{}(m))
}

func findInMap(name string, m map[string]interface{}) (interface{}, error) {
	tkns := strings.SplitN(strings.ToLower(name), ".", 2)

	next := m[tkns[0]]
	nextMap, ok := next.(map[string]interface{})

	if next == nil || len(tkns) == 1 || !ok {
		return next, nil
	}

	return findInMap(tkns[1], nextMap)
}

func main() {
	adMap := mapStringInterface{}
	err := json.Unmarshal([]byte(ad), &adMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(adMap.Get("ad.phone"))
	phone := "01655528881"
	adMap["ad"].(map[string]interface{})["phone"] = &phone
	expression, err := expr.Parse(`ad.phone == "01655528881" and ad.category == "1010"`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", expression)

	fmt.Println(adMap)
	result, err := expr.Run(expression, adMap)
	fmt.Println(result, err)

	regexEpx, err := expr.Parse(`("foo" matches "bar")`)
	if err != nil {
		panic(err)
	}
	res, err := expr.Run(regexEpx, nil)
	fmt.Println(res, err)

	containExp, err := expr.Parse(`ad.body matches ".*SJC.*" and ad.category == "1010"`, expr.With(adMap))
	if err != nil {
		panic(err)
	}
	res, err = expr.Run(containExp, adMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// notInEpx, err := expr.Parse(`Req.ActionParams.source=="ios" and Req.Ad.Category in ["1020","2010","2020","2030","2050","2060","2090","3060","3070","4010","4020","4040","5020","5040","5050","5060","8010","8030","11010","12010","13020"] and len(Req.Images) >= 3`)
}

type kneticExpr struct {
	ID *int
}

func (k *kneticExpr) Get() {

}

func tryKneticExpr() {
	id := 13
	tmp := map[string]interface{}{
		"ID": id,
	}
	expr, err := govaluate.NewEvaluableExpression("ID==13")
	if err != nil {
		panic(err)
	}
	result, err := expr.Evaluate(tmp)
	fmt.Println(result, err)
}
