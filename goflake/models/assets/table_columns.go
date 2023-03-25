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

func (i Identity) String() string {
	return fmt.Sprintf("IDENTITY START %[1]d INCREMENT %[2]d", i.StartNumber, i.IncrementNumber)
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
	DefaultValue *string
	Nullable     bool
	Unique       bool
	ColumnFields
}

func (s *Varchar) GetColumn() *ColumnFields {
	return &s.ColumnFields
}

func (s *Varchar) GetColumnDefinition() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\t%[1]s VARCHAR(%[2]d)", s.Name, s.Length))
	if !s.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if s.Unique {
		sb.WriteString(" UNIQUE")
	}
	if s.DefaultValue != nil {
		sb.WriteString(fmt.Sprintf(" DEFAULT '%[1]s'", *s.DefaultValue))
	}
	if s.Collation != "" {
		sb.WriteString(fmt.Sprintf(" COLLATE '%[1]s'", *s.DefaultValue))
	}
	if (s.ForeignKey != ForeignKey{}) {
		panic("foreign keys are not yet implemented")
	}
	return sb.String()
}

/*###################
### Number column ###
###################*/

var _ ISnowflakeColumn = &Number{}

type Number struct {
	//Precision refers to `xxx` of xxx.yyy -> max 38
	Precision int
	//Scale refers to `yyy` of xxx.yyy -> max 37
	Scale        int
	DefaultValue *float64
	Identity     Identity
	Sequence     SequenceAssociation
	Nullable     bool
	Unique       bool
	ColumnFields
}

func (s *Number) GetColumn() *ColumnFields {
	return &s.ColumnFields
}

func (s *Number) GetColumnDefinition() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\t%[1]s NUMBER(%[2]d,%[3]d)", s.Name, s.Precision, s.Scale))
	if !s.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if s.Unique {
		sb.WriteString(" UNIQUE")
	}
	if s.DefaultValue != nil {
		sb.WriteString(fmt.Sprintf(" DEFAULT %[1]d", s.DefaultValue))
	}
	if (s.Identity != Identity{}) {
		sb.WriteString(fmt.Sprintf(" IDENTITY(%[1]d, %[2]d)", s.Identity.StartNumber, s.Identity.IncrementNumber))
	}
	if (s.Sequence != SequenceAssociation{}) {
		panic("sequence associations are not yet implemented")
	}
	if (s.ForeignKey != ForeignKey{}) {
		panic("foreign keys are not yet implemented")
	}
	return sb.String()
}

/*#################
### Bool column ###
#################*/

var _ ISnowflakeColumn = &Boolean{}

type Boolean struct {
	DefaultValue *bool
	Nullable     bool
	Unique       bool
	ColumnFields
}

func (s *Boolean) GetColumn() *ColumnFields {
	return &s.ColumnFields
}

func (s *Boolean) GetColumnDefinition() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\t%[1]s BOOLEAN", s.Name))
	if !s.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if s.Unique {
		sb.WriteString(" UNIQUE")
	}
	if s.DefaultValue != nil {
		sb.WriteString(fmt.Sprintf(" DEFAULT %[1]t", *s.DefaultValue))
	}
	if (s.ForeignKey != ForeignKey{}) {
		panic("foreign keys are not yet implemented")
	}
	return sb.String()
}

/*
#################
### Date column ###
#################
*/
type Date struct {
	DefaultValue *string
	Nullable     bool
	Unique       bool
	ColumnFields
}

func (s *Date) GetColumn() *ColumnFields {
	return &s.ColumnFields
}

func (s *Date) GetColumnDefinition() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\t%[1]s DATE", s.Name))
	if !s.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if s.Unique {
		sb.WriteString(" UNIQUE")
	}
	if s.DefaultValue != nil {
		sb.WriteString(fmt.Sprintf(" DEFAULT '%[1]s'", *s.DefaultValue))
	}
	if (s.ForeignKey != ForeignKey{}) {
		panic("foreign keys are not yet implemented")
	}
	return sb.String()
}

/*#################
### Time column ###
#################*/

/*######################
### Timestamp column ###
######################*/

/*####################
### Variant column ###
####################*/
