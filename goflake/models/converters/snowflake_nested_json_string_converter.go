package converters

import (
	"database/sql"
	"encoding/json"
)

var (
	_ sql.Scanner = &NestedJsonStringConverter[any]{}
)

type NestedJsonStringConverter[T any] []T

func (t *NestedJsonStringConverter[T]) Scan(data interface{}) error {
	var converted []T
	if data != nil {
		err := json.Unmarshal([]byte(data.(string)), &converted)
		if err != nil {
			return err
		}
	}
	*t = converted
	return nil
}
