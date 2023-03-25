package converters

import (
	"encoding/json"
	"strconv"
)

type SnowflakeintConverter int

func (num *SnowflakeintConverter) UnmarshalJSON(data []byte) error {
	var asString string
	err := json.Unmarshal(data, &asString)
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(asString)
	if err != nil {
		return err
	}
	*num = SnowflakeintConverter(i)

	return nil
}
