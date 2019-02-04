package client

import (
	"github.com/go-resty/resty"
	. "iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/response"
	. "iyzipay-go/iyzipay/util"
)

func PaymentRetrieve(paymentRetrieveRequest PaymentRetrieveRequest, options Options) PaymentRetrieveResponse {
	successResponse := PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentRetrieveRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentRetrieveRequest).
		Post(options.BaseUrl + "/payment/detail")
	return successResponse
}

func PaymentAuth(paymentAuthRequest PaymentAuthRequest, options Options) PaymentRetrieveResponse {
	successResponse := PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentAuthRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentAuthRequest).
		Post(options.BaseUrl + "/payment/auth")
	return successResponse
}