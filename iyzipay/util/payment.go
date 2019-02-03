package util

import (
	"math/big"
	"strconv"
)

func FormatPrice(price *big.Float) string {
	if price == nil {
		return ""
	}
	f, _ := price.Float64()

	if price.IsInt() {
		println("int")
		return strconv.FormatFloat(f, 'f', 1, 64)

	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}