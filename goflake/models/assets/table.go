package assets

import (
	"fmt"
	"strings"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

var (
	_ i.ISnowflakeAsset = &Table{}
)

type Table struct {
	DatabaseName string
	SchemaName   string
	TableName    string
	Columns      u.Queue[ISnowflakeColumn]
	Tags         []ClassificationTag
}

// GetCreateStatement implements ISnowflakeAsset
func (r *Table) GetCreateStatement() (string, int) {
	statements := 1
	var sb strings.Builder
	var columnDefinitions []string
	var primaryKeys []string
	var columnsWithTags []ISnowflakeColumn
	sb.WriteString(fmt.Sprintf("CREATE TABLE %[1]s.%[2]s.%[3]s (\n", r.DatabaseName, r.SchemaName, r.TableName))
	for !r.Columns.IsEmpty() {
		col := r.Columns.Get()
		if len(col.GetColumn().Tags) > 0 {
			columnsWithTags = append(columnsWithTags, col)
		}

		columnDefinitions = append(columnDefinitions, col.GetColumnDefinition())
		if col.GetColumn().PrimaryKey {
			primaryKeys = append(primaryKeys, col.GetColumn().Name)
		}
	}
	if len(primaryKeys) > 0 {
		sb.WriteString(fmt.Sprintf(" %s,\n", strings.Join(columnDefinitions, ",\n")))
		sb.WriteString(fmt.Sprintf("\tPRIMARY KEY(%s)\n", strings.Join(primaryKeys, ", ")))
	} else {
		sb.WriteString(fmt.Sprintf(" %s\n", strings.Join(columnDefinitions, ",\n")))
	}
	sb.WriteString(")\n")
	for _, tag := range r.Tags {
		sb.WriteString(fmt.Sprintf("ALTER TABLE %[1]s.%[2]s.%[3]s SET TAG %[4]s = '%[5]s';\n", r.DatabaseName, r.SchemaName, r.TableName, tag.GetIdentifier(), tag.TagValue))
		statements += 1
	}
	for _, ct := range columnsWithTags {
		for _, tag := range ct.GetColumn().Tags {
			sb.WriteString(fmt.Sprintf("ALTER TABLE %[1]s.%[2]s.%[3]s ALTER COLUMN %[4]s SET TAG %[5]s = '%[6]s';\n", r.DatabaseName, r.SchemaName, r.TableName, ct.GetColumn().Name, tag.GetIdentifier(), tag.TagValue))
			statements += 1
		}
	}
	return sb.String(), statements
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *Table) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP TABLE %[1]s.%[2]s.%[3]s;", r.DatabaseName, r.SchemaName, r.TableName), 1
}
