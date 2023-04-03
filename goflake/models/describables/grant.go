package describables

import (
	"fmt"
)

var (
	_ ISnowflakeDescribable = &Grant{}
)

type Grant struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *Grant) GetDescribeStatement() string {
	var principalType string
	var principalIdentifier string
	switch any(r.Principal).(type) {
	case *Role:
		principalType = r.Principal.GetPrincipalType()
		principalIdentifier = r.Principal.GetPrincipalIdentifier()
	case *DatabaseRole:
		principalType = r.Principal.GetPrincipalType()
		principalIdentifier = r.Principal.GetPrincipalIdentifier()
	default:
		panic("Show grants is not implementer for this principal type")
	}
	return fmt.Sprintf("SHOW GRANTS TO %s %s", principalType, principalIdentifier)
}

func (*Grant) IsProcedure() bool {
	return false
}
