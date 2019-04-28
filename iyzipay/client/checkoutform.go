package client

import (
	"github.com/go-resty/resty"
	. "iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/request/model"
	. "iyzipay-go/iyzipay/response"
)

func CheckoutFormInitialize(checkoutFormInitializeRequest CheckoutFormInitializeRequest, options Options) CheckoutFormInitializeResponse {
	successResponse := CheckoutFormInitializeResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(checkoutFormInitializeRequest, options)).
		SetResult(&successResponse).
		SetBody(checkoutFormInitializeRequest).
		Post(options.BaseUrl + "/payment/iyzipos/checkoutform/initialize/auth/ecom")
	return successResponse
}
