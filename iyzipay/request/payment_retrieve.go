package request

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
)

type PaymentRetrieveRequest struct {
	PaymentId             string `json:"paymentId,omitempty"`
	PaymentConversationId string `json:"paymentConversationId,omitempty"`
	BaseRequest
}

func (paymentRetrieveRequest PaymentRetrieveRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", string(paymentRetrieveRequest.Locale)).
		Append("conversationId", paymentRetrieveRequest.ConversationId).
		Append("paymentId", paymentRetrieveRequest.PaymentId).
		Append("paymentConversationId", paymentRetrieveRequest.PaymentConversationId).
		ToString()

	return pki
}
