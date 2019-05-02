package security

import (
	"bytes"
	"github.com/erhmutlu/iyzipay-go/iyzipay/util"
	"reflect"
)

type PKIRequest struct {
	value string
}

//TODO: value değeri herşeyi alabilir interface{} ???
//TODO: farklı bir sekilde concat ???
func (pkiRequest PKIRequest) Append(key string, value string) PKIRequest {
	if len(value) > 0 {
		tmp := key + "=" + value
		if len(pkiRequest.value) > 0 {
			pkiRequest.value += "," + tmp
		} else {
			pkiRequest.value += tmp
		}
	}
	return pkiRequest
}

func (pkiRequest PKIRequest) AppendSlice(key string, pkiRequestItems interface{}) PKIRequest {
	if reflect.TypeOf(pkiRequestItems).Kind() == reflect.Slice {
		slice := reflect.ValueOf(pkiRequestItems)

		var buffer bytes.Buffer
		for i := 0; i < slice.Len(); i++ {
			item := slice.Index(i)
			//TODO: ToPKIRequestReflection ile kod tekrarı var
			pki := callToPKIRequest(item)

			if i > 0 {
				buffer.WriteString(", ")
				buffer.WriteString(pki)
			} else {
				buffer.WriteString(pki)
			}
		}
		s := "[" + buffer.String() + "]"
		return pkiRequest.Append(key, s)
	}

	return pkiRequest
}

func (pkiRequest PKIRequest) AppendIntSlice(key string, pkiRequestItems interface{}) PKIRequest {
	if reflect.TypeOf(pkiRequestItems).Kind() == reflect.Slice {
		slice := reflect.ValueOf(pkiRequestItems)

		var buffer bytes.Buffer
		for i := 0; i < slice.Len(); i++ {
			item := slice.Index(i)
			str := util.FormatPrimitiveInt(item.Interface().(int))

			if i > 0 {
				buffer.WriteString(", ")
				buffer.WriteString(str)
			} else {
				buffer.WriteString(str)
			}
		}
		s := "[" + buffer.String() + "]"
		return pkiRequest.Append(key, s)
	}

	return pkiRequest
}

func (pkiRequest PKIRequest) ToString() string {
	return "[" + pkiRequest.value + "]"
}

func ToPKIRequestReflection(i interface{}) string {
	if reflect.TypeOf(i).Kind() == reflect.Struct {
		item := reflect.ValueOf(i)
		return callToPKIRequest(item)
	}

	return ""
}

func callToPKIRequest(item reflect.Value) string {
	methodVal := item.MethodByName("ToPKIRequest")
	methodIface := methodVal.Interface()
	method := methodIface.(func() string)
	return method()
}
