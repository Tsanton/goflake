package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	e "github.com/tsanton/goflake-client/goflake/models/enums"
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

// GetCreateStatement implements ISnowflakeAsset
func (r *Role) GetCreateStatement() (string, int) {
	var principal e.SnowflakePrincipal
	switch x := r.Owner.GetPrincipalType(); x {
	case e.SnowflakePrincipalRole, e.SnowflakePrincipalDatabaseRole:
		principal = x
	default:
		panic("Ownership for this principal type is not implemented")
	}
	return fmt.Sprintf(`
	CREATE OR REPLACE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON ROLE %[1]s TO %[3]s %[4]s REVOKE CURRENT GRANTS;`,
		r.Name, r.Comment, principal.GrantType(), r.Owner.GetIdentifier(),
	), 2
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *Role) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP ROLE IF EXISTS %[1]s;", r.Name), 1
}

// GetIdentifier implements ISnowflakePrincipal
func (r *Role) GetIdentifier() string {
	return r.Name
}

// GetPrincipalType implements ISnowflakePrincipal
func (r *Role) GetPrincipalType() e.SnowflakePrincipal {
	return e.SnowflakePrincipalRole
}
