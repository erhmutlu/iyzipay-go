package main

import "iyzipay-go/iyzipay/model"

func main() {
	model.PaymentRetrieve()
	model.PaymentAuth(100.2, "lalal", "apiKey", "secretKey")
}
