package response

type CheckoutFormInitializeResponse struct {
	*Meta
	*CheckoutFormInfo
}

type CheckoutFormInfo struct {
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkoutFormContent"`
	TokenExpireTime     string `json:"tokenExpireTime"`
	PaymentPageUrl      string `json:"paymentPageUrl"`
}
