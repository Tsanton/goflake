package entities

import (
	"fmt"

	c "github.com/tsanton/goflake-client/goflake/models/converters"
)

var (
	_ ISnowflakeEntity = Table{}
)

type Table struct {
	DatabaseName   string `json:"database_name"`
	SchemaName     string `json:"schema_name"`
	Name           string
	Kind           string
	Comment        string
	ChangeTracking string `json:"change_tracking"`
	AutoClustering string `json:"automatic_clustering"`
	Rows           int
	Owner          string
	RetentionTime  c.SnowflakeintConverter `json:"retention_time"`
	Tags           []ClassificationTag
	Columns        []Column
	CreatedOn      c.SnowflakeDatetimeConverter
}

// GetIdentity implements ISnowflakeEntity
func (r Table) GetIdentity() string {
	return fmt.Sprintf("%[1]s.%[2]s.%[3]s", r.DatabaseName, r.SchemaName, r.Name)
}
