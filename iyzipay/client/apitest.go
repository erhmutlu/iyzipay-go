package client

import (
	"github.com/go-resty/resty"
	"iyzipay-go/iyzipay/model"
	"iyzipay-go/iyzipay/util"
	"log"
)

func ApiTest(options util.Options) model.ApiTestResponse {
	successResponse := model.ApiTestResponse{}

	_, err := resty.R().SetResult(&successResponse).Get(options.BaseUrl + "/payment/test")
	if err != nil {
		log.Fatal("lalala")
	}
	return successResponse
}
