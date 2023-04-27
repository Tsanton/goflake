package assets

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeGrantAsset = &GrantActionAccountGrant{}
)

type GrantActionAccountGrant struct{}

// GetGrantStatement implements ISnowflakeGrantAsset
func (g *GrantActionAccountGrant) GetGrantStatement(p i.ISnowflakePrincipal, privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch p.GetPrincipalType() {
	case enum.SnowflakePrincipalRole:
		return fmt.Sprintf("GRANT %[1]s ON ACCOUNT TO ROLE %[2]s;", privs, p.GetIdentifier()), 1
	case enum.SnowflakePrincipalDatabaseRole:
		panic("you can't grant account level privileges to database roles")
	default:
		panic("GetGrantStatement is not implemented for this interface type")
	}
}

// GetRevokeStatement implements ISnowflakeGrantAsset
func (g *GrantActionAccountGrant) GetRevokeStatement(p i.ISnowflakePrincipal, privileges []enum.Privilege) (string, int) {
	stringPrivileges := lo.Map(privileges, func(x enum.Privilege, index int) string { return x.String() })
	privs := strings.Join(stringPrivileges, ", ")
	switch p.GetPrincipalType() {
	case enum.SnowflakePrincipalRole:
		return fmt.Sprintf("REVOKE %[1]s ON ACCOUNT FROM ROLE %[2]s CASCADE;", privs, p.GetIdentifier()), 1
	case enum.SnowflakePrincipalDatabaseRole:
		panic("Account privileges cannot be neither granted to nor revoked from database roles")
	default:
		panic("GetRevokeStatement is not implemented for this interface type")
	}
}

// ValidatePrivileges implements ISnowflakeGrantAsset
func (*GrantActionAccountGrant) ValidatePrivileges(privileges []enum.Privilege) bool {
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
