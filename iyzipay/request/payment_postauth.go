package request

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
	"github.com/erhmutlu/iyzipay-go/iyzipay/util"
)

type PaymentPostAuthRequest struct {
	BaseRequest
	PaymentId string  `json:"paymentId,omitempty"`
	PaidPrice float64 `json:"paidPrice,omitempty"`
	Ip        string  `json:"ip,omitempty"`
	Currency  Currency  `json:"currency,omitempty"`
}

func (paymentPostAuthRequest PaymentPostAuthRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(paymentPostAuthRequest.Locale)).
		Append("conversationId", paymentPostAuthRequest.ConversationId).
		Append("paymentId", paymentPostAuthRequest.PaymentId).
		Append("ip", paymentPostAuthRequest.Ip).
		Append("paidPrice", util.FormatPrimitiveFloat(paymentPostAuthRequest.PaidPrice)).
		Append("currency", string(paymentPostAuthRequest.Currency)).
		ToString()

	return pki
}
