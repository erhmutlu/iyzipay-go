package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type PaymentRefundRequest struct {
	BaseRequest
	PaymentTransactionId string  `json:"paymentTransactionId,omitempty"`
	Price                float64 `json:"price,omitempty"`
	Ip                   string  `json:"ip,omitempty"`
	Currency             string  `json:"currency,omitempty"`
	RefundReason         string  `json:"refundReason,omitempty"`  //TODO: RefundReason olarak gönderilebilecek değerler belli. burası ona göre güncellenmeli
	Description          string  `json:"description,omitempty"`
}

func (paymentRefundRequest PaymentRefundRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(paymentRefundRequest.Locale)).
		Append("conversationId", paymentRefundRequest.ConversationId).
		Append("paymentTransactionId", paymentRefundRequest.PaymentTransactionId).
		Append("price", util.FormatPrimitiveFloat(paymentRefundRequest.Price)).
		Append("ip", paymentRefundRequest.Ip).
		Append("currency", paymentRefundRequest.Currency).
		Append("reason", paymentRefundRequest.RefundReason).
		Append("description", paymentRefundRequest.Description).
		ToString()

	return pki
}
