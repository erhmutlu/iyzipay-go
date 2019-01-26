package client

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"github.com/go-resty/resty"
	"iyzipay-go/iyzipay/model"
	. "iyzipay-go/iyzipay/util"
)

func PaymentRetrieve(paymentRetrieveRequest model.PaymentRetrieveRequest, options model.Options) model.PaymentRetrieveResponse {
	successResponse := model.PaymentRetrieveResponse{}

	resty.
		SetDebug(options.DebugMode).
		R().
		SetHeaders(httpHeaders(paymentRetrieveRequest, options)).
		SetResult(&successResponse).
		SetBody(paymentRetrieveRequest).
		Post(options.BaseUrl + "/payment/detail")
	return successResponse
}

func httpHeaders(paymentRetrieveRequest model.PaymentRetrieveRequest, options model.Options) map[string]string {
	random := RandomAlphanumeric(8)
	headers := make(map[string]string)
	headers["x-iyzi-client-version"] = "iyzipay-go-1.0.0"
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["x-iyzi-rnd"] = random
	headers["Authorization"] = prepareAuthorizationString(paymentRetrieveRequest, random, options)
	return headers
}

func prepareAuthorizationString(paymentRetrieveRequest model.PaymentRetrieveRequest, random string, options model.Options) string {
	hash := prepareHash(paymentRetrieveRequest, random, options)
	return "IYZWS " + options.ApiKey + ":" + hash
}

func prepareHash(paymentRetrieveRequest model.PaymentRetrieveRequest, random string, options model.Options) string{
	cipher := options.ApiKey + random + options.SecretKey + paymentRetrieveRequest.ToPKIRequest()
	println(cipher)
	crypt := sha1.New()
	crypt.Write([]byte(cipher))
	return b64.StdEncoding.EncodeToString(crypt.Sum(nil))
}
