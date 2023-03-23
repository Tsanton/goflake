package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionFutureSchemaGrant[i.ISnowflakePrincipal]{}
)

type GrantActionFutureSchemaGrant[T i.ISnowflakePrincipal] struct {
	Principal    T
	DatabaseName string
	SchemaName   string
	ObjectType   enum.SnowflakeObject
}

func (r *GrantActionFutureSchemaGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")

	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("GRANT %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s TO ROLE %[5]s;", privs, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("GRANT %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s TO DATABASE ROLE %[5]s;", privs, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (r *GrantActionFutureSchemaGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")

	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("REVOKE %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s FROM ROLE %[5]s CASCADE;", privs, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		return fmt.Sprintf("REVOKE %[1]s ON FUTURE %[2]s IN SCHEMA %[3]s.%[4]s FROM DATABASE ROLE %[5]s CASCADE;", privs, r.ObjectType.ToPlural(), r.DatabaseName, r.SchemaName, r.Principal.GetIdentifier()), 1
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (r *GrantActionFutureSchemaGrant[T]) ValidatePrivileges(privileges []enum.Privilege) bool {
	var allowedPrivileges []enum.Privilege

	switch r.ObjectType {
	case enum.SnowflakeObjectTable:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeInsert,
			enum.PrivilegeUpdate,
			enum.PrivilegeDelete,
			enum.PrivilegeTruncate,
			enum.PrivilegeReferences,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectView, enum.SnowflakeObjectMatView:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeReferences,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectSequence, enum.SnowflakeObjectFunction, enum.SnowflakeObjectProcedure, enum.SnowflakeObjectFileFormat:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeUsage,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectInternalStage:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeRead,
			enum.PrivilegeWrite,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectExternalStage:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeUsage,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectPipe:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeMonitor,
			enum.PrivilegeOperate,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectStream:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeSelect,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectTask:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeMonitor,
			enum.PrivilegeOperate,
			enum.PrivilegeOwnership,
		}
	case enum.SnowflakeObjectMaskingPolicy, enum.SnowflakeObjectPasswordPolicy, enum.SnowflakeObjectRowAccessPolicy, enum.SnowflakeObjectTag:
		allowedPrivileges = []enum.Privilege{
			enum.PrivilegeApply,
			enum.PrivilegeOwnership,
		}
	}
	return lo.Every(allowedPrivileges, privileges)
}
