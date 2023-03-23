package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionDatabaseGrant[i.ISnowflakePrincipal]{}
)

type GrantActionDatabaseGrant[T i.ISnowflakePrincipal] struct {
	Principal    T
	DatabaseName string
}

func (r *GrantActionDatabaseGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO ROLE %[3]s;", privs, r.DatabaseName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("GRANT %[1]s ON DATABASE %[2]s TO DATABASE ROLE %[3]s;", privs, r.DatabaseName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (r *GrantActionDatabaseGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM ROLE %[3]s CASCADE;", privs, r.DatabaseName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("REVOKE %[1]s ON DATABASE %[2]s FROM DATABASE ROLE %[3]s CASCADE;", privs, r.DatabaseName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (*GrantActionDatabaseGrant[T]) ValidatePrivileges(privileges []enum.Privilege) bool {
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
