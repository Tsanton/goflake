package converters

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SnowflakeBoolConverter bool

func (bit *SnowflakeBoolConverter) UnmarshalJSON(data []byte) error {
	var asString string
	err := json.Unmarshal(data, &asString)
	if err != nil {
		return err
	}
	switch strings.ToLower(asString) {
	case "1", "true", "t", "y":
		*bit = true
	case "0", "false", "f", "n":
		*bit = false
	default:
		return fmt.Errorf("boolean unmarshal error: invalid input %s", asString)
	}
	return nil
}
