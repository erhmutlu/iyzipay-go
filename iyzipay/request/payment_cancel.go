package request

import (
	. "iyzipay-go/iyzipay/security"
)

type PaymentCancelRequest struct {
	PaymentId      string `json:"paymentId,omitempty"`
	Ip             string `json:"ip,omitempty"`
	RefundReason   string `json:"refundReason,omitempty"` //TODO: RefundReason olarak gönderilebilecek değerler belli. burası ona göre güncellenmeli
	Description    string `json:"description,omitempty"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

func (paymentCancelRequest PaymentCancelRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentCancelRequest.Locale).
		Append("conversationId", paymentCancelRequest.ConversationId).
		Append("paymentId", paymentCancelRequest.PaymentId).
		Append("ip", paymentCancelRequest.Ip).
		Append("reason", paymentCancelRequest.RefundReason).
		Append("description", paymentCancelRequest.Description).
		ToString()

	return pki
}
