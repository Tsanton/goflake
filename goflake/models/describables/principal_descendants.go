package describables

import "fmt"

var (
	_ ISnowflakeDescribable = &PrincipalDescendants{}
)

// / Get direct children (inherited roles and database roles) for any snowflake principal. Only one level removed
type PrincipalDescendants struct {
	Principal ISnowflakeGrantPrincipal
}

// GetDescribeStatement implements ISnowflakeDescribable
func (p *PrincipalDescendants) GetDescribeStatement() string {
	var principalType string
	var principalIdentifier string
	switch any(p.Principal).(type) {
	case *Role:
		principalType = p.Principal.GetPrincipalType()
		principalIdentifier = p.Principal.GetPrincipalIdentifier()
	case *DatabaseRole:
		principalType = p.Principal.GetPrincipalType()
		principalIdentifier = p.Principal.GetPrincipalIdentifier()
	default:
		panic("GetDescribeStatement is not implementer for this principal type")
	}
	return fmt.Sprintf(`
with show_direct_descendants_from_principal as procedure(principal_type varchar, principal_identifier varchar)
    returns variant not null
    language python
    runtime_version = '3.8'
    packages = ('snowflake-snowpark-python')
    handler = 'show_direct_descendants_from_principal_py'
as $$
def show_direct_descendants_from_principal_py(snowpark_session, principal_type_py:str, principal_identifier_py:str):
    res = []
    for row in snowpark_session.sql(f'SHOW GRANTS TO {principal_type_py} {principal_identifier_py}').to_local_iterator():
        if row['privilege'] == 'USAGE' and row['granted_on'] in ['ROLE', 'DATABASE_ROLE']:
            res.append({
                **row.as_dict(),
                **{'distance_from_source': 0 }
            })
    return {
        'principal_identifier': principal_identifier_py,
        'principal_type': principal_type_py if principal_type_py != 'DATABASE ROLE' else 'DATABASE_ROLE',
        'descendants': res
    }
$$
call show_direct_descendants_from_principal('%[1]s', '%[2]s');
	`, principalType, principalIdentifier)
}

// IsProcedure implements ISnowflakeDescribable
func (*PrincipalDescendants) IsProcedure() bool {
	return true
}
