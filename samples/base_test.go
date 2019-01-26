package samples

import "iyzipay-go/iyzipay/model"

func RetrieveOptions() model.Options {
	return model.Options{
		BaseUrl:   "https://sandbox-api.iyzipay.com",
		ApiKey:    "mrI3mIMuNwGiIxanQslyJBRYa8nYrCU5",
		SecretKey: "9lkVluNHBABPw0LIvyn50oYZcrSJ8oNo",
		DebugMode: true}
}
