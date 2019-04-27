package response

import . "iyzipay-go/iyzipay/request/model"

type CancelResponse struct {
	*Meta
	*CancelInfo
}

type CancelInfo struct {
	Price         float64  `json:"price"`
	PaymentId     string   `json:"paymentId"`
	Currency      Currency `json:"currency"`
	ConnectorName *string  `json:"connectorName"`
	AuthCode      string   `json:"authCode"`
	HostReference string   `json:"hostReference"`
}
