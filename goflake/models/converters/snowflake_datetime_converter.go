package converters

import (
	"encoding/json"
	"time"
)

type SnowflakeDatetimeConverter time.Time

// 2023-03-23 14:37:56.977000+01:00
// https://yourbasic.org/golang/format-parse-string-time-date-example/
func (t *SnowflakeDatetimeConverter) UnmarshalJSON(data []byte) error {
	var asString string
	err := json.Unmarshal(data, &asString)
	if err != nil {
		return err
	}

	format := "2006-01-02 15:04:05.000000-07:00"
	parsedTime, err := time.Parse(format, asString)
	if err != nil {
		return err
	}
	*t = SnowflakeDatetimeConverter(parsedTime)

	return nil
}
