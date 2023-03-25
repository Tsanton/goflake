package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
)

type ColumnType struct {
	Type      string
	Nullable  bool
	Precision int
	Scale     int
	Length    int
	Fixed     bool
}

type Column struct {
	Name          string
	ColumnType    ColumnType `json:"data_type"`
	Default       *string
	Check         *string
	Expression    *string
	PrimaryKey    c.SnowflakeBoolConverter `json:"primary key"`
	UniqueKey     c.SnowflakeBoolConverter `json:"unique key"`
	PolicyName    *string
	AutoIncrement *string `json:"auto_increment"`
	Tags          []ClassificationTag
	Comment       string
}
