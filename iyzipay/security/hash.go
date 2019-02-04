package security

import (
	"crypto/sha1"
	b64 "encoding/base64"
)

func PrepareHashV1(request interface{}, random string, apiKey string, secretKey string) string {
	cipher := apiKey + random + secretKey + ToPKIRequestReflection(request)
	println(cipher)
	crypt := sha1.New()
	crypt.Write([]byte(cipher))
	return b64.StdEncoding.EncodeToString(crypt.Sum(nil))
}
