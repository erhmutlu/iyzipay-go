package response

import (
	. "iyzipay-go/iyzipay/request/model"
	"iyzipay-go/iyzipay/response/nullable"
)

type CancelResponse struct {
	*Meta
	*CancelInfo
}

type CancelInfo struct {
	Price         float64          `json:"price"`
	PaymentId     string           `json:"paymentId"`
	Currency      Currency         `json:"currency"`
	ConnectorName *nullable.String `json:"connectorName"`
	AuthCode      string           `json:"authCode"`
	HostReference string           `json:"hostReference"`
}
