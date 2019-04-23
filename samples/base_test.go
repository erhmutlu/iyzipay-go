package samples_test

import . "iyzipay-go/iyzipay/util"

func RetrieveOptions() Options {
	return Options{
		BaseUrl:   "https://sandbox-api.iyzipay.com",
		ApiKey:    "sandbox-9I7aOULJ0qcEqlMWeEO6Q8Rt6i6wHNVx",
		SecretKey: "sandbox-FANniqFxZ4IdufG5oNcv4SULKT1BDi7h",
		DebugMode: true}
}
