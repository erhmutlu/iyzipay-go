package response

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	"github.com/erhmutlu/iyzipay-go/iyzipay/response/nullable"
)

type RefundResponse struct {
	*Meta
	*RefundInfo
}

type RefundInfo struct {
	Price                float64          `json:"price"`
	PaymentId            string           `json:"paymentId"`
	PaymentTransactionId string           `json:"paymentTransactionId"`
	Currency             Currency         `json:"currency"`
	ConnectorName        *nullable.String `json:"connectorName"`
	AuthCode             string           `json:"authCode"`
	HostReference        string           `json:"hostReference"`
}
