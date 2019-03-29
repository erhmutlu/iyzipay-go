package nullable

import (
	"encoding/json"
	"math/big"
)

type Float struct {
	Data *big.Float
	Null bool
}

func NewNullableFloat(flt float64) *Float {
	println(flt)
	return &Float{Data: big.NewFloat(flt), Null: flt == 0}
}

func (ns *Float) UnmarshalJSON(d []byte) error {
	err := json.Unmarshal(d, &ns.Data)
	if err != nil {
		return err
	} else {
		if ns.Data == nil {
			ns.Null = true
		}
		return nil
	}
}

func (ns *Float) MarshalJSON() ([]byte, error) {
	if ns.Null == true {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.Data)
}
