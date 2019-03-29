package model

import (
	"iyzipay-go/iyzipay/nullable"
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
)

type BasketItem struct {
	Id               string           `json:"id,omitempty"`
	Price            float64          `json:"price,omitempty"`
	Name             string           `json:"name,omitempty"`
	Category1        string           `json:"category1,omitempty"`
	Category2        string           `json:"category2,omitempty"`
	ItemType         string           `json:"itemType,omitempty"`
	SubMerchantKey   *nullable.String `json:"subMerchantKey,omitempty"` //TODO: bu pointer olmak zorunda mı ?
	SubMerchantPrice *nullable.Float  `json:"subMerchantPrice,omitempty"`
}

func (basketItem BasketItem) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("id", basketItem.Id).
		Append("price", util.FormatFloat(basketItem.Price)).
		Append("name", basketItem.Name).
		Append("category1", basketItem.Category1).
		Append("category2", basketItem.Category2).
		Append("itemType", basketItem.ItemType).
		AppendNullableString("subMerchantKey", basketItem.SubMerchantKey).
		AppendNullableString("subMerchantPrice", util.FormatNullableFloat(basketItem.SubMerchantPrice)).
		ToString()

	return pki
}
