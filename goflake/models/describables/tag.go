package describables

import "fmt"

var (
	_ ISnowflakeDescribable = &Tag{}
)

type Tag struct {
	DatabaseName string
	SchemaName   string
	TagName      string
}

// GetDescribeStatement implements ISnowflakeDescribable
func (r *Tag) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW TAGS LIKE '%[1]s' IN SCHEMA %[2]s.%[3]s;", r.TagName, r.DatabaseName, r.SchemaName)
}

// IsProcedure implements ISnowflakeDescribable
func (*Tag) IsProcedure() bool {
	return false
}
