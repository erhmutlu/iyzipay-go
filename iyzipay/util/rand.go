package util

import (
	"bytes"
	"math/rand"
)

var charset = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomAlphanumeric(length int) string {
	var bb bytes.Buffer
	bb.Grow(length)
	for i := 0; i < len(charset); i++ {
		bb.WriteRune(charset[rand.Intn(len(charset))])
	}
	return bb.String()
}
