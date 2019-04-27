package nullable

import (
	"encoding/json"
)

type String struct {
	Data string
}

func (s *String) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.Data = str
	return nil
}

type Float struct {
	Data float64
}

func (f *Float) UnmarshalJSON(data []byte) error {
	var flt float64
	if err := json.Unmarshal(data, &flt); err != nil {
		return err
	}
	f.Data = flt
	return nil
}

type Int64 struct {
	Data int64
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	var i64 int64
	if err := json.Unmarshal(data, &i64); err != nil {
		return err
	}
	i.Data = i64
	return nil
}