package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionSchemaGrant[i.ISnowflakePrincipal]{}
)

type GrantActionSchemaGrant[T i.ISnowflakePrincipal] struct {
	Principal    T
	DatabaseName string
	SchemaName   string
}

func (r *GrantActionSchemaGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("GRANT %[1]s ON SCHEMA %[2]s.%[3]s TO ROLE %[4]s;", privs, r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("GRANT %[1]s ON SCHEMA %[2]s.%[3]s TO DATABASE ROLE %[4]s;", privs, r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (r *GrantActionSchemaGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("REVOKE %[1]s ON SCHEMA %[2]s.%[3]s FROM ROLE %[4]s CASCADE;", privs, r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("REVOKE %[1]s ON SCHEMA %[2]s.%[3]s FROM DATABASE ROLE %[4]s CASCADE;", privs, r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (*GrantActionSchemaGrant[T]) ValidatePrivileges(privileges []enum.Privilege) bool {
	allowedPrivileges := []enum.Privilege{
		enum.PrivilegeModify,
		enum.PrivilegeMonitor,
		enum.PrivilegeMonitorUsage,
		enum.PrivilegeCreateTable,
		enum.PrivilegeCreateExternalTable,
		enum.PrivilegeCreateView,
		enum.PrivilegeCreateMaskingPolicy,
		enum.PrivilegeCreateMaterializedView,
		enum.PrivilegeCreateRowAccessPolicy,
		enum.PrivilegeCreateSecret,
		enum.PrivilegeCreateSessionPolicy,
		enum.PrivilegeCreateStage,
		enum.PrivilegeCreateFileFormat,
		enum.PrivilegeCreateSequence,
		enum.PrivilegeCreateFunction,
		enum.PrivilegeCreatePasswordPolicy,
		enum.PrivilegeCreatePipe,
		enum.PrivilegeCreateStream,
		enum.PrivilegeCreateTag,
		enum.PrivilegeCreateTask,
		enum.PrivilegeCreateProcedure,
		enum.PrivilegeCreateAlert,
		enum.PrivilegeAddSearchOptimization,
		enum.PrivilegeOwnership,
	}
	return lo.Every(allowedPrivileges, privileges)
}
