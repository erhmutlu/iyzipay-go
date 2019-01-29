package security

import (
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
	//TODO: bu kısmı util'e çıkabiliriz
	switch reflect.TypeOf(pkiRequestItems).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(pkiRequestItems)

		for i := 0; i < s.Len(); i++ {
			item := s.Index(i)
			methodVal := item.MethodByName("ToPKIRequest")
			methodIface := methodVal.Interface()
			method := methodIface.(func() string)

			//TODO: not finished implementation
			method()
		}
	}

	return pkiRequest
}

func (pkiRequest PKIRequest) ToString() string {
	return "[" + pkiRequest.value + "]"
}
