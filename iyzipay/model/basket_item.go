package model

import (
	. "iyzipay-go/iyzipay/security"
	"iyzipay-go/iyzipay/util"
	"math/big"
)

type BasketItem struct {
	Id               string   `json:"id,omitempty"`
	Price            *big.Float  `json:"price,omitempty"`
	Name             string   `json:"name,omitempty"`
	Category1        string   `json:"category1,omitempty"`
	Category2        string   `json:"category2,omitempty"`
	ItemType         string   `json:"itemType,omitempty"`
	SubMerchantKey   string   `json:"subMerchantKey,omitempty"`
	SubMerchantPrice *big.Float `json:"subMerchantPrice,omitempty"`
}

func (basketItem BasketItem) ToPKIRequest() string {
	pki := PKIRequest{}.
		Append("id", basketItem.Id).
		Append("price", util.FormatPrice(basketItem.Price)).
		Append("name", basketItem.Name).
		Append("category1", basketItem.Category1).
		Append("category2", basketItem.Category2).
		Append("itemType", basketItem.ItemType).
		Append("subMerchantKey", basketItem.SubMerchantKey).
		Append("subMerchantPrice", util.FormatPrice(basketItem.SubMerchantPrice)).
		ToString()

	return pki
}