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

func PaymentPreAuth(paymentPreAuthRequest PaymentPreAuthRequest, options Options) PaymentRetrieveResponse {
	successResponse := PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentPreAuthRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentPreAuthRequest).
		Post(options.BaseUrl + "/payment/preauth")
	return successResponse
}

func PaymentPostAuth(paymentPostAuthRequest PaymentPostAuthRequest, options Options) PaymentRetrieveResponse {
	successResponse := PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentPostAuthRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentPostAuthRequest).
		Post(options.BaseUrl + "/payment/postauth")
	return successResponse
}

func PaymentCancel(paymentCancelRequest PaymentCancelRequest, options Options) CancelResponse {
	successResponse := CancelResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentCancelRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentCancelRequest).
		Post(options.BaseUrl + "/payment/cancel")
	return successResponse
}

func PaymentRefund(paymentRefundRequest PaymentRefundRequest, options Options) RefundResponse {
	successResponse := RefundResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(HttpHeaders(paymentRefundRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentRefundRequest).
		Post(options.BaseUrl + "/payment/refund")
	return successResponse
}
