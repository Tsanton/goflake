package enums

import "fmt"

type SnowflakeObject string

func (p SnowflakeObject) String() string {
	return string(p)
}

func (p SnowflakeObject) ToSingular() string {
	return string(p)
}
func (p SnowflakeObject) ToPlural() string {
	switch p {
	case SnowflakeObjectMaskingPolicy:
		return "MASKING POLICIES"
	case SnowflakeObjectPasswordPolicy:
		return "PASSWORD POLICIES"
	case SnowflakeObjectRowAccessPolicy:
		return "ROW ACCESS POLICIES"
	default:
		return fmt.Sprintf("%[1]sS", p)
	}
}

const (
	SnowflakeObjectTable           SnowflakeObject = "TABLE"
	SnowflakeObjectView            SnowflakeObject = "VIEW"
	SnowflakeObjectMatView         SnowflakeObject = "MATERIALIZED VIEW"
	SnowflakeObjectAccount         SnowflakeObject = "ACCOUNT"
	SnowflakeObjectDatabase        SnowflakeObject = "DATABASE"
	SnowflakeObjectDatabaseRole    SnowflakeObject = "DATABASE_ROLE"
	SnowflakeObjectFunction        SnowflakeObject = "FUNCTION"
	SnowflakeObjectRole            SnowflakeObject = "ROLE"
	SnowflakeObjectSchema          SnowflakeObject = "SCHEMA"
	SnowflakeObjectTag             SnowflakeObject = "TAG"
	SnowflakeObjectUser            SnowflakeObject = "USER"
	SnowflakeObjectSequence        SnowflakeObject = "SEQUENCE"
	SnowflakeObjectProcedure       SnowflakeObject = "PROCEDURE"
	SnowflakeObjectFileFormat      SnowflakeObject = "FILE FORMAT"
	SnowflakeObjectInternalStage   SnowflakeObject = "INTERNAL STAGE"
	SnowflakeObjectExternalStage   SnowflakeObject = "EXTERNAL STAGE"
	SnowflakeObjectPipe            SnowflakeObject = "PIPE"
	SnowflakeObjectStream          SnowflakeObject = "STREAM"
	SnowflakeObjectTask            SnowflakeObject = "TASK"
	SnowflakeObjectMaskingPolicy   SnowflakeObject = "MASKING POLICY"
	SnowflakeObjectPasswordPolicy  SnowflakeObject = "PASSWORD POLICY"
	SnowflakeObjectRowAccessPolicy SnowflakeObject = "ROW ACCESS POLICY"
	SnowflakeObjectWarehouse       SnowflakeObject = "WAREHOUSE"
)
