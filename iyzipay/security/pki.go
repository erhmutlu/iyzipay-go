package security

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

func (pkiRequest PKIRequest) ToString() string {
	return "[" + pkiRequest.value + "]"
}
