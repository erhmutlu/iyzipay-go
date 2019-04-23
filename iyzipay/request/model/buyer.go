package model

import (
	. "iyzipay-go/iyzipay/security"
)

type Buyer struct {
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Surname             string `json:"surname,omitempty"`
	IdentityNumber      string `json:"identityNumber,omitempty"`
	Email               string `json:"email,omitempty"`
	GsmNumber           string `json:"gsmNumber,omitempty"`
	RegistrationDate    string `json:"registrationDate,omitempty"`
	LastLoginDate       string `json:"lastLoginDate,omitempty"`
	RegistrationAddress string `json:"registrationAddress,omitempty"`
	City                string `json:"city,omitempty"`
	Country             string `json:"country,omitempty"`
	ZipCode             string `json:"zipCode,omitempty"`
	Ip                  string `json:"ip,omitempty"`
}

func (buyer Buyer) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("id", buyer.Id).
		Append("name", buyer.Name).
		Append("surname", buyer.Surname).
		Append("identityNumber", buyer.IdentityNumber).
		Append("email", buyer.Email).
		Append("gsmNumber", buyer.GsmNumber).
		Append("registrationDate", buyer.RegistrationDate).
		Append("lastLoginDate", buyer.LastLoginDate).
		Append("registrationAddress", buyer.RegistrationAddress).
		Append("city", buyer.City).
		Append("country", buyer.Country).
		Append("zipCode", buyer.ZipCode).
		Append("ip", buyer.Ip).
		ToString()

	return pki
}
