package nullable

import (
	"encoding/json"
	"strings"
)

type String struct {
	Data string
	Null bool
}

func NewNullableString(str string) *String {
	return &String{Data: str, Null: str == ""}
}

func (ns *String) UnmarshalJSON(d []byte) error {
	err := json.Unmarshal(d, &ns.Data)
	if err != nil {
		return err
	} else {
		if ns.Data == "null" || len(strings.TrimSpace(ns.Data)) == 0 {
			ns.Null = true
		}
		return nil
	}
}

func (ns *String) MarshalJSON() ([]byte, error) {
	if ns.Null == true {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.Data)
}
