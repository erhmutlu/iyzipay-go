package client

import (
	"github.com/go-resty/resty"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/response"
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

func CheckoutFormInitializePreAuth(checkoutFormInitializeRequest CheckoutFormInitializeRequest, options Options) CheckoutFormInitializeResponse {
	successResponse := CheckoutFormInitializeResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(checkoutFormInitializeRequest, options)).
		SetResult(&successResponse).
		SetBody(checkoutFormInitializeRequest).
		Post(options.BaseUrl + "/payment/iyzipos/checkoutform/initialize/preauth/ecom")
	return successResponse
}

func CheckoutFormRetrieve(checkoutFormRequest CheckoutFormRequest, options Options) CheckoutFormResultResponse {
	successResponse := CheckoutFormResultResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(checkoutFormRequest, options)).
		SetResult(&successResponse).
		SetBody(checkoutFormRequest).
		Post(options.BaseUrl + "/payment/iyzipos/checkoutform/auth/ecom/detail")
	return successResponse
}
