package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

var (
	_ i.ISnowflakeAsset     = &Role{}
	_ i.ISnowflakePrincipal = &Role{}
)

type Role struct {
	Name    string
	Owner   i.ISnowflakePrincipal
	Comment string
}

func (r *Role) GetCreateStatement() (string, int) {
	var principalType string
	switch r.Owner.(type) {
	case *Role:
		principalType = "ROLE"
	case *DatabaseRole:
		principalType = "DATABASE ROLE"
	default:
		panic("Ownership for this principal type is not implemented")
	}
	return fmt.Sprintf(`
	CREATE OR REPLACE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON ROLE %[1]s TO %[3]s %[4]s REVOKE CURRENT GRANTS;`,
		r.Name, r.Comment, principalType, r.Owner.GetIdentifier(),
	), 2
}

func (r *Role) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP ROLE IF EXISTS %[1]s;", r.Name), 1
}

func (r *Role) GetIdentifier() string {
	return r.Name
}
