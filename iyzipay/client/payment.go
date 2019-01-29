package client

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"github.com/go-resty/resty"
	. "iyzipay-go/iyzipay/model"
	. "iyzipay-go/iyzipay/request"
	. "iyzipay-go/iyzipay/util"
)

func PaymentRetrieve(paymentRetrieveRequest PaymentRetrieveRequest, options Options) PaymentRetrieveResponse {
	successResponse := PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(httpHeaders(paymentRetrieveRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentRetrieveRequest).
		Post(options.BaseUrl + "/payment/detail")
	return successResponse
}

func httpHeaders(paymentRetrieveRequest PaymentRetrieveRequest, options Options) map[string]string {
	random := RandomAlphanumeric(8)
	headers := make(map[string]string)
	headers["x-iyzi-client-version"] = "iyzipay-go-1.0.0"
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["x-iyzi-rnd"] = random
	headers["Authorization"] = prepareAuthorizationString(paymentRetrieveRequest, random, options)
	return headers
}

func prepareAuthorizationString(paymentRetrieveRequest PaymentRetrieveRequest, random string, options Options) string {
	hash := prepareHash(paymentRetrieveRequest, random, options)
	return "IYZWS " + options.ApiKey + ":" + hash
}

func prepareHash(paymentRetrieveRequest PaymentRetrieveRequest, random string, options Options) string {
	cipher := options.ApiKey + random + options.SecretKey + paymentRetrieveRequest.ToPKIRequest()
	println(cipher)
	crypt := sha1.New()
	crypt.Write([]byte(cipher))
	return b64.StdEncoding.EncodeToString(crypt.Sum(nil))
}
