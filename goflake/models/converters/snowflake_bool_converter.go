package converters

import (
	"encoding/json"
	"fmt"
)

type SnowflakeBoolConverter bool

func (bit *SnowflakeBoolConverter) UnmarshalJSON(data []byte) error {
	var asString string
	err := json.Unmarshal(data, &asString)
	if err != nil {
		return err
	}
	switch asString {
	case "1", "true":
		*bit = true
	case "0", "false":
		*bit = false
	default:
		return fmt.Errorf("boolean unmarshal error: invalid input %s", asString)
	}
	return nil
}
