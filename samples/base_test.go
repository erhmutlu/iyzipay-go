package samples

import "iyzipay-go/iyzipay/model"

func RetrieveOptions() model.Options{
	return model.Options{BaseUrl: "https://sandbox-api.iyzipay.com", ApiKey: "apiKey", SecretKey: "secretKey"}
}