package assets

import (
	"fmt"
	"strings"
)

type ForeignKey struct {
	DatabaseName string
	SchemaName   string
	TableName    string
	ColumnName   string
}

type Identity struct {
	StartNumber     int
	IncrementNumber int
}

type MaskingPolicyAssociation struct {
	DatabaseName string
	SchemaName   string
	PolicyName   string
}

type ClassificationTag struct {
	DatabaseName string
	SchemaName   string
	TagName      string
	TagValue     string
}

func (t *ClassificationTag) GetIdentifier() string {
	return fmt.Sprintf("%[1]s.%[2]s.%[3]s", t.DatabaseName, t.SchemaName, t.TagName)
}

type SequenceAssociation struct {
	DatabaseName string
	SchemaName   string
	SequenceName string
}

type ColumnFields struct {
	Name       string
	PrimaryKey bool
	ForeignKey ForeignKey
	Tags       []ClassificationTag
}

type ISnowflakeColumn interface {
	GetColumnDefinition() string
	GetColumn() *ColumnFields
}

/*####################
### Varchar column ###
####################*/

var _ ISnowflakeColumn = &Varchar{}

type Varchar struct {
	//Constraint: 0 <= x <= 16777216
	Length       int
	Collation    string
	DefaultValue string
	Nullable     bool
	Unique       bool
	ColumnFields
}

func (s *Varchar) GetColumn() *ColumnFields {
	return &s.ColumnFields
}

func (s *Varchar) GetColumnDefinition() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\t%[1]s VARCHAR(%[2]d)", s.ColumnFields.Name, s.Length))
	if !s.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if s.Unique {
		sb.WriteString(" UNIQUE")
	}
	if s.DefaultValue != "" {
		sb.WriteString(fmt.Sprintf(" DEFAULT '%[1]s'", s.DefaultValue))
	}
	if s.Collation != "" {
		sb.WriteString(fmt.Sprintf(" COLLATE '%[1]s'", s.DefaultValue))
	}
	if (s.ForeignKey != ForeignKey{}) {
		panic("foreign keys are not yet implemented")
	}
	return sb.String()
}
