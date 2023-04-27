package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	e "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset     = &DatabaseRole{}
	_ i.ISnowflakePrincipal = &DatabaseRole{}
)

type DatabaseRole struct {
	Name         string
	DatabaseName string
	Owner        i.ISnowflakePrincipal
	Comment      string
}

// GetCreateStatement implements ISnowflakeAsset
func (r *DatabaseRole) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
	CREATE OR REPLACE DATABASE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON DATABASE ROLE %[1]s TO %[3]s REVOKE CURRENT GRANTS;`,
		r.GetIdentifier(), r.Comment, r.Owner.GetIdentifier(),
	), 2
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *DatabaseRole) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP DATABASE ROLE %[1]s;", r.GetIdentifier()), 1
}

// GetIdentifier implements ISnowflakePrincipal
func (r *DatabaseRole) GetIdentifier() string {
	return fmt.Sprintf("%[1]s.%[2]s", r.DatabaseName, r.Name)
}

// GetPrincipalType implements ISnowflakePrincipal
func (r *DatabaseRole) GetPrincipalType() e.SnowflakePrincipal {
	return e.SnowflakePrincipalDatabaseRole
}
