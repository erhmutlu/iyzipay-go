package request

import (
	. "iyzipay-go/iyzipay/security"
)

type PaymentCancelRequest struct {
	PaymentId      string `json:"paymentId,omitempty"`
	Ip             string `json:"ip,omitempty"`
	RefundReason   string `json:"refundReason,omitempty"`
	Description    string `json:"description,omitempty"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

func (PaymentCancelRequest PaymentCancelRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", PaymentCancelRequest.Locale).
		Append("conversationId", PaymentCancelRequest.ConversationId).
		Append("paymentId", PaymentCancelRequest.PaymentId).
		Append("ip", PaymentCancelRequest.Ip).
		Append("reason", PaymentCancelRequest.RefundReason).
		Append("description", PaymentCancelRequest.Description).
		ToString()

	return pki
}
