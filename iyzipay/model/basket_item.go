package model

import (
	. "iyzipay-go/iyzipay/security"
	"strconv"
)

type BasketItem struct {
	Id               string  `json:"id,omitempty"`
	Price            float64 `json:"price,omitempty"`
	Name             string  `json:"name,omitempty"`
	Category1        string  `json:"category1,omitempty"`
	Category2        string  `json:"category2,omitempty"`
	ItemType         string  `json:"itemType,omitempty"`
	SubMerchantKey   string  `json:"subMerchantKey,omitempty"`
	SubMerchantPrice float64 `json:"subMerchantPrice,omitempty"`
}

func (basketItem BasketItem) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("id", basketItem.Id).
		Append("price", strconv.FormatFloat(basketItem.Price, 'f', -1, 64)).
		Append("name", basketItem.Name).
		Append("category1", basketItem.Category1).
		Append("category2", basketItem.Category2).
		Append("itemType", basketItem.ItemType).
		Append("subMerchantKey", basketItem.SubMerchantKey).
		Append("subMerchantPrice", strconv.FormatFloat(basketItem.SubMerchantPrice, 'f', -1, 64)).
		ToString()

	return pki
}