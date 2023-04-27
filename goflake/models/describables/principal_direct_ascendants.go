package describables

import (
	"fmt"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable = &PrincipalAscendants{}
)

// Get all direct parent (roles and database roles inheriting the RolePrincipal) for a given Snowflake role/database role
type PrincipalDirectAscendants struct {
	Principal ISnowflakeGrantPrincipal
}

func (p *PrincipalDirectAscendants) GetDescribeStatement() string {
	switch p.Principal.GetPrincipalType() {
	case enums.SnowflakePrincipalRole, enums.SnowflakePrincipalDatabaseRole:
		break
	default:
		panic("Show grants is not implementer for this principal type")
	}

	// return fmt.Sprintf("SHOW GRANTS OF %s %s", p.Principal.GetPrincipalType().GrantType(), p.Principal.GetPrincipalIdentifier())

	return fmt.Sprintf(`
with show_direct_ascendants_from_principal as procedure(principal_type varchar, principal_identifier varchar)
    returns variant not null
    language python
    runtime_version = '3.8'
    packages = ('snowflake-snowpark-python')
    handler = 'main_py'
as $$
def main_py(snowpark_session, principal_type_py:str, principal_identifier_py:str):
	res = []
	try:
		for row in snowpark_session.sql(f'SHOW GRANTS OF {principal_type_py} {principal_identifier_py}').to_local_iterator():
			res.append({
				**row.as_dict(),
				**{'distance_from_source': 0,'granted_on' : principal_type_py if principal_type_py != 'DATABASE ROLE' else 'DATABASE_ROLE'}
			})
	except:
		return res
	return res
$$
call show_direct_ascendants_from_principal('%[1]s', '%[2]s');
	`, p.Principal.GetPrincipalType().GrantType(), p.Principal.GetPrincipalIdentifier())
}

func (*PrincipalDirectAscendants) IsProcedure() bool {
	return true
}
