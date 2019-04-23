package client

import (
	"github.com/go-resty/resty"
	. "iyzipay-go/iyzipay/response"
	"iyzipay-go/iyzipay/util"
	"log"
)

func ApiTest(options util.Options) ApiTestResponse {
	successResponse := ApiTestResponse{}

	_, err := resty.R().SetResult(&successResponse).Get(options.BaseUrl + "/payment/test")
	if err != nil {
		log.Fatal("error") //TODO: fix here
	}
	return successResponse
}
