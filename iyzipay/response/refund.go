package response

type RefundResponse struct {
	*Meta
	*RefundInfo
}

type RefundInfo struct {
	Price                float64 `json:"price"`
	PaymentId            string  `json:"paymentId"`
	PaymentTransactionId string  `json:"paymentTransactionId"`
	Currency             string  `json:"currency"`
	ConnectorName        *string `json:"connectorName"`
	AuthCode             string  `json:"authCode"`
	HostReference        string  `json:"hostReference"`
}