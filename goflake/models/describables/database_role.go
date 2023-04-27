package describables

import (
	"fmt"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable    = &DatabaseRole{}
	_ ISnowflakeGrantPrincipal = &DatabaseRole{}
)

// Beware that you cannot grant account level privleges to database roles
type DatabaseRole struct {
	Name         string
	DatabaseName string
}

func (r *DatabaseRole) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW DATABASE ROLES LIKE '%[1]s' IN DATABASE %[2]s;", r.Name, r.DatabaseName)
}

func (r *DatabaseRole) IsProcedure() bool {
	return false
}

// GetPrincipalType implements ISnowflakeGrantPrincipal
func (r *DatabaseRole) GetPrincipalType() enums.SnowflakePrincipal {
	return enums.SnowflakePrincipalDatabaseRole
}

// GetPrincipalIdentifier implements ISnowflakeGrantPrincipal
func (r *DatabaseRole) GetPrincipalIdentifier() string {
	return fmt.Sprintf("%[1]s.%[2]s", r.DatabaseName, r.Name)
}
