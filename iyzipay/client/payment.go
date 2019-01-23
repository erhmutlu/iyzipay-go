package client

import (
	"fmt"
	"github.com/go-resty/resty"
	"iyzipay-go/iyzipay/model"
	"log"
)

func PaymentRetrieve(paymentRetrieveRequest model.PaymentRetrieveRequest, options model.Options) model.PaymentRetrieveResponse {
	successResponse := model.PaymentRetrieveResponse{}

	resp, err := resty.R().SetResult(&successResponse).SetBody(paymentRetrieveRequest).Post(options.BaseUrl + "/payment/detail")
	if err != nil {
		log.Fatal("lalala")
	}
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	return successResponse
}
