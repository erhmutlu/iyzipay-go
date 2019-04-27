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