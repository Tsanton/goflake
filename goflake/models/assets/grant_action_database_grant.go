package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionDatabaseGrant{}
)

type GrantActionDatabaseGrant struct {
	DatabaseName string
}

func (g *GrantActionDatabaseGrant) GetGrantStatement(p i.ISnowflakePrincipal, privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch p.GetPrincipalType() {
	case enum.SnowflakePrincipalRole:
		return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO ROLE %[3]s;", privs, g.DatabaseName, p.GetIdentifier()), 1
	case enum.SnowflakePrincipalDatabaseRole:
		return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO DATABASE ROLE %[3]s;", privs, g.DatabaseName, p.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (g *GrantActionDatabaseGrant) GetRevokeStatement(p i.ISnowflakePrincipal, privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch p.GetPrincipalType() {
	case enum.SnowflakePrincipalRole:
		return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM ROLE %[3]s CASCADE;", privs, g.DatabaseName, p.GetIdentifier()), 1
	case enum.SnowflakePrincipalDatabaseRole:
		return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM DATABASE ROLE %[3]s CASCADE;", privs, g.DatabaseName, p.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (*GrantActionDatabaseGrant) ValidatePrivileges(privileges []enum.Privilege) bool {
	allowedPrivileges := []enum.Privilege{
		// enum.PrivilegeCreateDatabaseRole, //Missing enum
		enum.PrivilegeCreateSchema,
		enum.PrivilegeImportedPrivileges,
		enum.PrivilegeModify,
		enum.PrivilegeMonitor,
		enum.PrivilegeUsage,
	}
	return lo.Every(allowedPrivileges, privileges)
}
