package common_model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type MultiLangFields map[string]string

// Value Marshal
func (a MultiLangFields) Value() (driver.Value, error) {
	if a == nil {
		return json.Marshal(map[string]string{"en": "", "id": ""})
	}
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *MultiLangFields) Scan(value interface{}) error {
	if value == nil {
		return json.Unmarshal([]byte(`{"en":"","id":""}`), &a)
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	if string(b) == "null" {
		b = []byte(`{"en":"","id":""}`)
	}
	return json.Unmarshal(b, &a)
}
