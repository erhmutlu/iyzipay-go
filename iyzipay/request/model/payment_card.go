package model

import (
	. "github.com/erhmutlu/iyzipay-go/iyzipay/security"
	"github.com/erhmutlu/iyzipay-go/iyzipay/util"
)

type PaymentCard struct {
	CardHolderName string   `json:"cardHolderName,omitempty"`
	CardNumber     string   `json:"cardNumber,omitempty"`
	ExpireYear     string   `json:"expireYear,omitempty"`
	ExpireMonth    string   `json:"expireMonth,omitempty"`
	Cvc            string   `json:"cvc,omitempty"`
	RegisterCard   *int     `json:"registerCard,omitempty"`
	CardAlias      string   `json:"cardAlias,omitempty"`
	CardToken      string   `json:"cardToken,omitempty"`
	CardUserKey    string   `json:"cardUserKey,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}

func (paymentCard PaymentCard) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("cardHolderName", paymentCard.CardHolderName).
		Append("cardNumber", paymentCard.CardNumber).
		Append("expireYear", paymentCard.ExpireYear).
		Append("expireMonth", paymentCard.ExpireMonth).
		Append("cvc", paymentCard.Cvc).
		Append("registerCard", util.FormatInt(paymentCard.RegisterCard)).
		Append("cardAlias", paymentCard.CardAlias).
		Append("cardToken", paymentCard.CardToken).
		Append("cardUserKey", paymentCard.CardUserKey).
		Append("metadata", paymentCard.Metadata.ToPKIRequest()).
		ToString()

	return pki
}

type Metadata map[string]string

func (metadata Metadata) ToPKIRequest() string {
	if len(metadata) > 0 {
		pki := PKIRequest{}
		for key, val := range metadata {
			pki = pki.Append(key, val)
		}

		return pki.ToString()
	} else {
		return ""
	}
}
