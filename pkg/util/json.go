package util

import (
	"database/sql/driver"
	"encoding/json"
)

type Json map[string]any

// Value implements the driver.Valuer interface
func (f Json) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *Json) Scan(value any) error {
	var data = []byte(value.([]uint8))
	return json.Unmarshal(data, &f)
}

func JsonParse(d any) []byte {
	result, _ := json.Marshal(d)
	return result
}

func JsonStringify(b []byte, d any) any {
	json.Unmarshal(b, d)

	return d
}
