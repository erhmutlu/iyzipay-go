package model

import (
	"encoding/json"
	"strings"
)

type NullString struct {
	Data string
	Null bool
}

func NewNullString(str string) *NullString {
	return &NullString{Data: str, Null: str == ""}
}

func (ns *NullString) UnmarshalJSON(d []byte) error {
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

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Null == true {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.Data)
}
