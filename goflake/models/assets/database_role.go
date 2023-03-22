package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
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

// GetIdentifier implements ISnowflakeRole
func (r *DatabaseRole) GetIdentifier() string {
	return fmt.Sprintf("%[1]s.%[2]s", r.DatabaseName, r.Name)
}

func (r *DatabaseRole) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
	CREATE OR REPLACE DATABASE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON DATABASE ROLE %[1]s TO %[3]s REVOKE CURRENT GRANTS;`,
		r.GetIdentifier(), r.Comment, r.Owner,
	), 2
}

func (r *DatabaseRole) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP DATABASE ROLE %[1]s;", r.GetIdentifier()), 1
}
