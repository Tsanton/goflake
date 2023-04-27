package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset = &RoleInheritance{}
)

type RoleInheritance struct {
	ChildPrincipal  i.ISnowflakePrincipal
	ParentPrincipal i.ISnowflakePrincipal
}

// GetCreateStatement implements ISnowflakeAsset
func (r *RoleInheritance) GetCreateStatement() (string, int) {
	switch r.ChildPrincipal.GetPrincipalType() {
	case enums.SnowflakePrincipalUser:
		panic("users cannot be granted to other roles")
	}
	return fmt.Sprintf("GRANT %[1]s %[2]s TO %[3]s %[4]s;",
		r.ChildPrincipal.GetPrincipalType().GrantType(),
		r.ChildPrincipal.GetIdentifier(),
		r.ParentPrincipal.GetPrincipalType().GrantType(),
		r.ParentPrincipal.GetIdentifier(),
	), 1
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *RoleInheritance) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("REVOKE %[1]s %[2]s FROM %[3]s %[4]s;",
		r.ChildPrincipal.GetPrincipalType().GrantType(),
		r.ChildPrincipal.GetIdentifier(),
		r.ParentPrincipal.GetPrincipalType().GrantType(),
		r.ParentPrincipal.GetIdentifier(),
	), 1
}
