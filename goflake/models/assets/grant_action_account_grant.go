package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionAccountGrant[i.ISnowflakePrincipal]{}
)

type GrantActionAccountGrant[T i.ISnowflakePrincipal] struct {
	Principal T
}

func (r *GrantActionAccountGrant[T]) GetGrantStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("GRANT %[1]s ON ACCOUNT TO ROLE %[2]s;", privs, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		panic("you can't grant account level privileges to database roles")
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

func (r *GrantActionAccountGrant[T]) GetRevokeStatement(privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf("REVOKE %[1]s ON ACCOUNT FROM ROLE %[2]s CASCADE;", privs, r.Principal.GetIdentifier()), 1
	case *DatabaseRole:
		panic("Account privileges cannot be neither granted to nor revoked from database roles")
	default:
		panic("GetRevokeStatement is not implemented for this interface type")
	}
}

func (*GrantActionAccountGrant[T]) ValidatePrivileges(privileges []enum.Privilege) bool {
	allowedPrivileges := []enum.Privilege{
		enum.PrivilegeCreateAccount,
		enum.PrivilegeCreateDataExchangeListing,
		enum.PrivilegeCreateDatabase,
		enum.PrivilegeCreateIntegration,
		enum.PrivilegeCreateNetworkPolicy,
		enum.PrivilegeCreateRole,
		enum.PrivilegeCreateShare,
		enum.PrivilegeCreateUser,
		enum.PrivilegeCreateWarehouse,

		enum.PrivilegeApplyMaskingPolicy,
		// enum.PrivilegeApplyPasswordPolicy, //Missing enum
		enum.PrivilegeApplyRowAccessPolicy,
		// enum.PrivilegeApplySessionPolicy, //Missing enum
		enum.PrivilegeApplyTag,
		enum.PrivilegeAttachPolicy,
		enum.PrivilegeExecuteTask,
		enum.PrivilegeImportShare,
		enum.PrivilegeManageGrants,
		enum.PrivilegeMonitorExecution,
		enum.PrivilegeMonitorUsage,
		enum.PrivilegeOverrideShareRestrictions,
		enum.PrivilegeExecuteManagedTask,
		enum.PrivilegeOrganizationSupportCases,
		enum.PrivilegeAccountSupportCases,
		enum.PrivilegeUserSupportCases,
	}
	return lo.Every(allowedPrivileges, privileges)
}
