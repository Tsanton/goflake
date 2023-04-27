package describables

import (
	"fmt"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable = &PrincipalAscendants{}
)

// Get all parent (roles and database roles inheriting the RolePrincipal) for a given Snowflake role/database role
type PrincipalAscendants struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *PrincipalAscendants) GetDescribeStatement() string {
	switch r.Principal.GetPrincipalType() {
	case enums.SnowflakePrincipalRole, enums.SnowflakePrincipalDatabaseRole:
		break
	default:
		panic("Show grants is not implementer for this principal type")
	}

	return fmt.Sprintf(`
with show_all_roles_that_inherit_source as procedure(principal_type varchar, principal_identifier varchar)
    returns variant not null
    language python
    runtime_version = '3.8'
    packages = ('snowflake-snowpark-python')
    handler = 'main_py'
as $$
def show_grants_on_py(snowpark_session, principal_type_py:str, principal_identifier_py:str, links_removed:int):
    res = []
    try:
        for row in snowpark_session.sql(f'SHOW GRANTS OF {principal_type_py} {principal_identifier_py}').to_local_iterator():
            res.append({
                **row.as_dict(),
                **{ 'distance_from_source': links_removed, 'granted_on' : principal_type_py if principal_type_py != 'DATABASE ROLE' else 'DATABASE_ROLE' }
            })
    except:
        return res
    return res

def show_all_roles_that_inherit_source_py(snowpark_session, principal_type: str, principal_identifier: str, links_removed:int, result:list, roles_shown:set = set()):
    roles = show_grants_on_py(snowpark_session, principal_type, principal_identifier, links_removed)
    show_inheritance = []
    for role in roles:
        if not role['grantee_name'] in roles_shown:
            result.append(role)
            roles_shown.add(role['grantee_name'].upper())
            show_inheritance.append(role)
    for role in show_inheritance:
        principal_type_iter:str = role['granted_to'] if role['granted_to'] != 'DATABASE_ROLE' else 'DATABASE ROLE'
        show_all_roles_that_inherit_source_py(snowpark_session, principal_type_iter, role['grantee_name'], links_removed +1, result, roles_shown)

def main_py(snowpark_session, principal_type_py:str, principal_identifier_py:str):
    res = []
    show_all_roles_that_inherit_source_py(snowpark_session, principal_type_py, principal_identifier_py, 0, res)
    return {
        'principal_identifier': principal_identifier_py,
        'principal_type': principal_type_py if principal_type_py != 'DATABASE ROLE' else 'DATABASE_ROLE',
        'ascendants': res
    }
$$
call show_all_roles_that_inherit_source('%[1]s', '%[2]s')
`, r.Principal.GetPrincipalType().GrantType(), r.Principal.GetPrincipalIdentifier())
}

func (*PrincipalAscendants) IsProcedure() bool {
	return true
}
