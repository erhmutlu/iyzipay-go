package client

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/request/model"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/util"
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
)

func HttpHeaders(request interface{}, options Options) map[string]string {
	random := RandomAlphanumeric(8)
	headers := make(map[string]string)
	headers["x-iyzi-client-version"] = "iyzipay-go-1.0.0"
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["x-iyzi-rnd"] = random
	headers["Authorization"] = prepareAuthorizationString(request, random, options)
	return headers
}

func prepareAuthorizationString(request interface{}, random string, options Options) string {
	hash := PrepareHashV1(request, random, options.ApiKey, options.SecretKey)
	return "IYZWS " + options.ApiKey + ":" + hash
}
