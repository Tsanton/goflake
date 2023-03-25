package enums

type SnowflakeTimestamp string

func (p SnowflakeTimestamp) String() string {
	return string(p)
}

const (
	SnowflakeTimestampLtz SnowflakeTimestamp = "TIMESTAMP_LTZ"
	SnowflakeTimestampNtz SnowflakeTimestamp = "TIMESTAMP_NTZ"
	SnowflakeTimestampTz  SnowflakeTimestamp = "TIMESTAMP_TZ"
)
