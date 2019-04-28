package response

type CheckoutFormInitializeResponse struct {
	*Meta
	*CheckoutFormInitiliazeInfo
}

type CheckoutFormInitiliazeInfo struct {
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkoutFormContent"`
	TokenExpireTime     string `json:"tokenExpireTime"`
	PaymentPageUrl      string `json:"paymentPageUrl"`
}

type CheckoutFormResultResponse struct {
	*Meta
	*PaymentInfo
	*CheckoutFormResultInfo
}

type CheckoutFormResultInfo struct {
	Token       string `json:"token"`
	CallbackUrl string `json:"callbackUrl"`
}