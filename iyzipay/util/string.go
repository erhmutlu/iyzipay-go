package util

import (
	"iyzipay-go/iyzipay/nullable"
	"math/big"
	"strconv"
)

func FormatFloat(flt float64) string {
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

func FormatNullableFloat(float *nullable.Float) *nullable.String {
	if float == nil {
		return &nullable.String{Data: "", Null: true}
	}
	data := float.Data
	f, _ := data.Float64()

	if data.IsInt() {
		println("i")
		return &nullable.String{Data: strconv.FormatFloat(f, 'f', 1, 64), Null: false}
	}
	return &nullable.String{Data: strconv.FormatFloat(f, 'f', -1, 64), Null: false}
}

func FormatInt(int *int) string {
	if int == nil {
		return ""
	}
	return strconv.Itoa(*int)
}
