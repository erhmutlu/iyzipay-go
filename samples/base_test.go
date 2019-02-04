package samples

import . "iyzipay-go/iyzipay/util"

func RetrieveOptions() Options {
	return Options{
		BaseUrl:   "https://sandbox-api.iyzipay.com",
		ApiKey:    "apiKey",
		SecretKey: "secretKey",
		DebugMode: true}
}
