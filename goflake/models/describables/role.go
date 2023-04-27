package describables

import (
	"fmt"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable    = &Role{}
	_ ISnowflakeGrantPrincipal = &Role{}
)

type Role struct {
	Name string
}

func (r *Role) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW ROLES LIKE '%[1]s';", r.Name)
}

func (r *Role) IsProcedure() bool {
	return false
}

func (r *Role) GetPrincipalIdentifier() string {
	return r.Name
}

// GetPrincipalType implements ISnowflakeGrantPrincipal
func (r *Role) GetPrincipalType() enums.SnowflakePrincipal {
	return enums.SnowflakePrincipalRole
}
