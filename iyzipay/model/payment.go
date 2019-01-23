package model

type PaymentAuthRequest struct {
	Price          float64 `json:"price"`
	ConversationId string  `json:"conversationId"`
	ApiKey         string  `json:"apiKey"`
	SecretKey      string  `json:"secretKey"`
}

type PaymentRetrieveRequest struct {
	PaymentId             int64  `json:"paymentId"`
	PaymentConversationId string `json:"paymentConversationId"`
}

type PaymentRetrieveResponse struct {
	Status     string `json:"status"`
	Locale     string `json:"locale"`
	SystemTime int64  `json:"systemTime"`
}
