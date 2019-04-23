package request

import (
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type PaymentRefundRequest struct {
	PaymentTransactionId string  `json:"paymentTransactionId,omitempty"`
	Price                float64 `json:"price,omitempty"`
	Ip                   string  `json:"ip,omitempty"`
	Currency             string  `json:"currency,omitempty"`
	RefundReason         string  `json:"refundReason,omitempty"`
	Description          string  `json:"description,omitempty"`
	Locale               string  `json:"locale,omitempty"`
	ConversationId       string  `json:"conversationId,omitempty"`
}

func (paymentRefundRequest PaymentRefundRequest) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("locale", paymentRefundRequest.Locale).
		Append("conversationId", paymentRefundRequest.ConversationId).
		Append("paymentTransactionId", paymentRefundRequest.PaymentTransactionId).
		Append("price", util.FormatFloat(paymentRefundRequest.Price)).
		Append("ip", paymentRefundRequest.Ip).
		Append("currency", paymentRefundRequest.Currency).
		Append("reason", paymentRefundRequest.RefundReason).
		Append("description", paymentRefundRequest.Description).
		ToString()

	return pki
}
