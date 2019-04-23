package response

type CancelResponse struct {
	*Meta
	*CancelInfo
}

type CancelInfo struct {
	Price         float64 `json:"price"`
	PaymentId     string  `json:"paymentId"`
	Currency      string  `json:"currency"`
	ConnectorName *string `json:"connectorName"`
	AuthCode      string  `json:"authCode"`
	HostReference string  `json:"hostReference"`
}
