package util

import (
	"math/big"
	"strconv"
)

func FormatFloat(float *big.Float) string {
	if float == nil {
		return ""
	}
	f, _ := float.Float64()

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