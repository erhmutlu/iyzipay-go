package util

import (
	"math/big"
	"strconv"
)

func FormatPrimitiveFloat(flt float64) string {
	if flt == 0 {
		return ""
	}
	float := big.NewFloat(flt)
	f, _ := big.NewFloat(flt).Float64()

	if float.IsInt() {
		return strconv.FormatFloat(f, 'f', 1, 64)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FormatInt(int *int) string {
	if int == nil {
		return ""
	}
	return strconv.Itoa(*int)
}

func FormatPrimitiveInt(int int) string {
	return strconv.Itoa(int)
}