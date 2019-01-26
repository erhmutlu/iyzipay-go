package model

import (
	. "iyzipay-go/iyzipay/security"
)

type PaymentRetrieveRequest struct {
	PaymentId             string `json:"paymentId,omitempty"`
	PaymentConversationId string `json:"paymentConversationId,omitempty"`
	Locale                string `json:"locale,omitempty"`
	ConversationId        string `json:"conversationId,omitempty"`
}

func (paymentRetrieveRequest PaymentRetrieveRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentRetrieveRequest.Locale).
		Append("conversationId", paymentRetrieveRequest.ConversationId).
		Append("paymentId", paymentRetrieveRequest.PaymentId).
		Append("paymentConversationId", paymentRetrieveRequest.PaymentConversationId).
		ToString()

	return pki
}

type PaymentRetrieveResponse struct {
	Status      string `json:"status"`
	Locale      string `json:"locale"`
	SystemTime  int64  `json:"systemTime"`
	PaymentInfo PaymentInfo
}

type PaymentInfo struct {
}
