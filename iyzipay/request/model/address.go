package model

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
)

type Address struct {
	Address     string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}

func (address Address) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("address", address.Address).
		Append("zipCode", address.ZipCode).
		Append("contactName", address.ContactName).
		Append("city", address.City).
		Append("country", address.Country).
		ToString()

	return pki
}
