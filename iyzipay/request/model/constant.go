package model

type Currency string

const (
	TRY Currency = "TRY"
	EUR Currency = "EUR"
	USD Currency = "USD"
	GBP Currency = "GBP"
	IRR Currency = "IRR"
	NOK Currency = "NOK"
	RUB Currency = "RUB"
	CHF Currency = "CHF"
)

type Locale string

const (
	TR Locale = "tr"
	EN Locale = "en"
)

type PaymentChannel string

const (
	MOBILE         PaymentChannel = "MOBILE"
	WEB            PaymentChannel = "WEB"
	MOBILE_WEB     PaymentChannel = "MOBILE_WEB"
	MOBILE_IOS     PaymentChannel = "MOBILE_IOS"
	MOBILE_ANDROID PaymentChannel = "MOBILE_ANDROID"
	MOBILE_WINDOWS PaymentChannel = "MOBILE_WINDOWS"
	MOBILE_TABLET  PaymentChannel = "MOBILE_TABLET"
	MOBILE_PHONE   PaymentChannel = "MOBILE_PHONE"
)

type PaymentGroup string

const (
	PRODUCT      PaymentGroup = "PRODUCT"
	LISTING      PaymentGroup = "LISTING"
	SUBSCRIPTION PaymentGroup = "SUBSCRIPTION"
)
