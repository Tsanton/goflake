package enums

type Privilege string

func (p Privilege) String() string {
	return string(p)
}

// https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/e1227159be50bf26841acead8730dad516a96ebc/pkg/resources/privileges.go#L73
const (
	PrivilegeSelect                    Privilege = "SELECT"
	PrivilegeInsert                    Privilege = "INSERT"
	PrivilegeUpdate                    Privilege = "UPDATE"
	PrivilegeDelete                    Privilege = "DELETE"
	PrivilegeTruncate                  Privilege = "TRUNCATE"
	PrivilegeReferences                Privilege = "REFERENCES"
	PrivilegeRebuild                   Privilege = "REBUILD"
	PrivilegeCreateSchema              Privilege = "CREATE SCHEMA"
	PrivilegeCreateDatabaseRole        Privilege = "CREATE DATABASE ROLE"
	PrivilegeImportedPrivileges        Privilege = "IMPORTED PRIVILEGES"
	PrivilegeModify                    Privilege = "MODIFY"
	PrivilegeOperate                   Privilege = "OPERATE"
	PrivilegeMonitor                   Privilege = "MONITOR"
	PrivilegeOwnership                 Privilege = "OWNERSHIP"
	PrivilegeRead                      Privilege = "READ"
	PrivilegeReferenceUsage            Privilege = "REFERENCE_USAGE"
	PrivilegeUsage                     Privilege = "USAGE"
	PrivilegeWrite                     Privilege = "WRITE"
	PrivilegeCreateTable               Privilege = "CREATE TABLE"
	PrivilegeCreateTag                 Privilege = "CREATE TAG"
	PrivilegeCreateView                Privilege = "CREATE VIEW"
	PrivilegeCreateFileFormat          Privilege = "CREATE FILE FORMAT"
	PrivilegeCreateStage               Privilege = "CREATE STAGE"
	PrivilegeCreatePipe                Privilege = "CREATE PIPE"
	PrivilegeCreateStream              Privilege = "CREATE STREAM"
	PrivilegeCreateTask                Privilege = "CREATE TASK"
	PrivilegeCreateSequence            Privilege = "CREATE SEQUENCE"
	PrivilegeCreateFunction            Privilege = "CREATE FUNCTION"
	PrivilegeCreateProcedure           Privilege = "CREATE PROCEDURE"
	PrivilegeCreateExternalTable       Privilege = "CREATE EXTERNAL TABLE"
	PrivilegeCreateMaterializedView    Privilege = "CREATE MATERIALIZED VIEW"
	PrivilegeCreateRowAccessPolicy     Privilege = "CREATE ROW ACCESS POLICY"
	PrivilegeCreateTemporaryTable      Privilege = "CREATE TEMPORARY TABLE"
	PrivilegeCreateMaskingPolicy       Privilege = "CREATE MASKING POLICY"
	PrivilegeCreateNetworkPolicy       Privilege = "CREATE NETWORK POLICY"
	PrivilegeCreateDataExchangeListing Privilege = "CREATE DATA EXCHANGE LISTING"
	PrivilegeCreateAccount             Privilege = "CREATE ACCOUNT"
	PrivilegeCreateShare               Privilege = "CREATE SHARE"
	PrivilegeCreateSecret              Privilege = "CREATE SECRET"
	PrivilegeCreateSessionPolicy       Privilege = "CREATE SESSION POLICY"
	PrivilegeCreatePasswordPolicy      Privilege = "CREATE PASSWORD POLICY"
	PrivilegeCreateAlert               Privilege = "CREATE ALERT"
	PrivilegeImportShare               Privilege = "IMPORT SHARE"
	PrivilegeOverrideShareRestrictions Privilege = "OVERRIDE SHARE RESTRICTIONS"
	PrivilegeAddSearchOptimization     Privilege = "ADD SEARCH OPTIMIZATION"
	PrivilegeApplyMaskingPolicy        Privilege = "APPLY MASKING POLICY"
	PrivilegeApplyRowAccessPolicy      Privilege = "APPLY ROW ACCESS POLICY"
	PrivilegeApplyTag                  Privilege = "APPLY TAG"
	PrivilegeApply                     Privilege = "APPLY"
	PrivilegeAttachPolicy              Privilege = "ATTACH POLICY"

	PrivilegeCreateRole               Privilege = "CREATE ROLE"
	PrivilegeCreateUser               Privilege = "CREATE USER"
	PrivilegeCreateWarehouse          Privilege = "CREATE WAREHOUSE"
	PrivilegeCreateDatabase           Privilege = "CREATE DATABASE"
	PrivilegeCreateIntegration        Privilege = "CREATE INTEGRATION"
	PrivilegeManageGrants             Privilege = "MANAGE GRANTS"
	PrivilegeMonitorUsage             Privilege = "MONITOR USAGE"
	PrivilegeMonitorExecution         Privilege = "MONITOR EXECUTION"
	PrivilegeExecuteTask              Privilege = "EXECUTE TASK"
	PrivilegeExecuteManagedTask       Privilege = "EXECUTE MANAGED TASK"
	PrivilegeOrganizationSupportCases Privilege = "MANAGE ORGANIZATION SUPPORT CASES"
	PrivilegeAccountSupportCases      Privilege = "MANAGE ACCOUNT SUPPORT CASES"
	PrivilegeUserSupportCases         Privilege = "MANAGE USER SUPPORT CASES"
)

type PrivilegeSet map[Privilege]struct{}

func NewPrivilegeSet(privileges ...Privilege) PrivilegeSet {
	ps := PrivilegeSet{}
	for _, priv := range privileges {
		ps[priv] = struct{}{}
	}
	return ps
}

func (ps PrivilegeSet) ToList() []string {
	privs := []string{}
	for p := range ps {
		privs = append(privs, string(p))
	}
	return privs
}
