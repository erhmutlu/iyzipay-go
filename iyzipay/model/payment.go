package model

import (
	"strconv"
)

type PaymentRetrieveRequest struct {
	PaymentId             int64  `json:"paymentId"`
	PaymentConversationId string `json:"paymentConversationId"`
	Locale                string `json:"locale"`
	ConversationId        string `json:"conversationId"`
}

func (paymentRetrieveRequest PaymentRetrieveRequest) ToPKIRequest() string {
	//TODO: atanmamıs degerler ???
	//TODO: farklı bir sekilde concat ???
	return "[conversationId=" + paymentRetrieveRequest.ConversationId +
		//",locale=" + paymentRetrieveRequest.Locale +
		",paymentId=" + strconv.FormatInt(paymentRetrieveRequest.PaymentId, 10) +
		",paymentConversationId=" + paymentRetrieveRequest.PaymentConversationId + "]"
}

type PaymentRetrieveResponse struct {
	Status     string `json:"status"`
	Locale     string `json:"locale"`
	SystemTime int64  `json:"systemTime"`
}
