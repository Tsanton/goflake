package entities

import (
	"fmt"

	c "github.com/tsanton/goflake-client/goflake/models/converters"
	// c "github.com/tsanton/goflake-client/goflake/models/converters"
)

var (
	_ ISnowflakeEntity = Tag{}
)

type Tag struct {
	DatabaseName string `db:"database_name"`
	SchemaName   string `db:"schema_name"`
	Name         string
	Owner        string
	Comment      string
	// AllowedValues []string                     `db:"allowed_values"`
	AllowedValues c.NestedJsonStringConverter[string] `db:"allowed_values" json:"allowed_values"`
	CreatedOn     c.SnowflakeDatetimeConverter        `db:"created_on"`
}

// GetIdentity implements ISnowflakeEntity
func (r Tag) GetIdentity() string {
	return fmt.Sprintf("%[1]s.%[2]s.%[3]s", r.DatabaseName, r.SchemaName, r.Name)
}
