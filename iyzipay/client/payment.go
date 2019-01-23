package client

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"github.com/go-resty/resty"
	"iyzipay-go/iyzipay/model"
)

func PaymentRetrieve(paymentRetrieveRequest model.PaymentRetrieveRequest, options model.Options) model.PaymentRetrieveResponse {
	successResponse := model.PaymentRetrieveResponse{}

	headers := httpHeaders(paymentRetrieveRequest, options)
	println(headers["Authorization"])

	resp, _ := resty.R().SetHeaders(headers).SetResult(&successResponse).SetBody(paymentRetrieveRequest).Post(options.BaseUrl + "/payment/detail")

	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
	//fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	return successResponse
}

func httpHeaders(paymentRetrieveRequest model.PaymentRetrieveRequest, options model.Options) map[string]string {
	random := "123456"
	headers := make(map[string]string)
	//headers["Accept"] = "application/json"
	headers["x-iyzi-rnd"] = random
	//headers["x-iyzi-client-version"] = "iyzipay-go-1.0.0"
	headers["Authorization"] = prepareAuthorizationString(paymentRetrieveRequest, random, options)
	return headers
}

func prepareAuthorizationString(paymentRetrieveRequest model.PaymentRetrieveRequest, random string, options model.Options) string {
	hash := prepareHash(paymentRetrieveRequest, random, options)
	//println(hash)
	return "IYZWS " + options.ApiKey + ":" + hash
}

func prepareHash(paymentRetrieveRequest model.PaymentRetrieveRequest, random string, options model.Options) string{
	cipher := options.ApiKey + random + options.SecretKey + paymentRetrieveRequest.ToPKIRequest()
	println(cipher)
	crypt := sha1.New()
	crypt.Write([]byte(cipher))
	return b64.StdEncoding.EncodeToString(crypt.Sum(nil))
}
