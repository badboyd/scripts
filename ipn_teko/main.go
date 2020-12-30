package main

const (
	ipnURL = "https://orders.tekoapis.com/command/payments/ipn"
)

type (
	// Order ...
	Order struct {
		Code     string `json:"code"`
		Payments []struct {
			MethodCode             string `json:"methodCode"`
			OrderID                string `json:"orderId"`
			OrderCode              string `json:"orderCode"`
			Amount                 int    `json:"amount"`
			BankCode               string `json:"bankCode"`
			BankName               string `json:"bankName"`
			CurrCode               string `json:"currCode"`
			BuyerName              string `json:"buyerName"`
			UpdatedAt              string `json:"updatedAt"`
			BuyerPhone             string `json:"buyerPhone"`
			ClientCode             string `json:"clientCode"`
			PartnerCode            string `json:"partnerCode"`
			PaymentType            string `json:"paymentType"`
			ServiceCode            string `json:"serviceCode"`
			CrmPartnerID           string `json:"crmPartnerId"`
			MerchantCode           string `json:"merchantCode"`
			ResponseCode           string `json:"responseCode"`
			TerminalCode           string `json:"terminalCode"`
			AsiaPartnerID          string `json:"asiaPartnerId"`
			CustomerPhone          string `json:"customerPhone"`
			PaymentMethod          string `json:"paymentMethod"`
			TransactionID          string `json:"transactionId"`
			PsTransactionCode      string `json:"psTransactionCode"`
			PsResponseTime         string `json:"psResponseTime"`
			ResponseMessage        string `json:"responseMessage"`
			ClientRequestTime      string `json:"clientRequestTime"`
			PartnerTransactionCode string `json:"partnerTransactionCode"`
		} `json:"payments"`
		ChildOrders []struct {
			ID         string      `json:"id"`
			Code       string      `json:"code"`
			State      interface{} `json:"state"`
			GrandTotal int         `json:"grandTotal"`
		} `json:"child_orders"`
	}
)

func checksum(data map[string]interface{}) string {
	return ""
}

func main() {

}
