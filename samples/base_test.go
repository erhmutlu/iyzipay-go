package samples

import . "iyzipay-go/iyzipay/request/model"

func RetrieveOptions() Options {
	return Options{
		BaseUrl:   "https://sandbox-api.iyzipay.com",
		ApiKey:    "apiKey",
		SecretKey: "secretKey",
		DebugMode: true}
}
