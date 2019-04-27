package request

import (
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type PaymentPostAuthRequest struct {
	BaseRequest
	PaymentId string  `json:"paymentId,omitempty"`
	PaidPrice float64 `json:"paidPrice,omitempty"`
	Ip        string  `json:"ip,omitempty"`
	Currency  string  `json:"currency,omitempty"`
}

func (paymentPostAuthRequest PaymentPostAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentPostAuthRequest.Locale). //TODO: AppendBaseRequest gibi bir şey eklenebilir.
		Append("conversationId", paymentPostAuthRequest.ConversationId).
		Append("paymentId", paymentPostAuthRequest.PaymentId).
		Append("ip", paymentPostAuthRequest.Ip).
		Append("paidPrice", util.FormatPrimitiveFloat(paymentPostAuthRequest.PaidPrice)).
		Append("currency", paymentPostAuthRequest.Currency).
		ToString()

	return pki
}
