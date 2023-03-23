package describables

import (
	"fmt"
)

var (
	_ ISnowflakeDescribable = &Grant{}
)

type Grant struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *Grant) GetDescribeStatement() string {
	var principalType string
	var principalIdentifier string
	switch any(r.Principal).(type) {
	case *Role:
		principalType = r.Principal.GetPrincipalType()
		principalIdentifier = r.Principal.GetPrincipalIdentifier()
	case *DatabaseRole:
		principalType = r.Principal.GetPrincipalType()
		principalIdentifier = r.Principal.GetPrincipalIdentifier()
	default:
		panic("Show grants is not implementer for this principal type")
	}
	return fmt.Sprintf(`
with show_grants_to_principal as procedure(principal_type varchar, principal_identifier varchar)
    returns variant not null
    language python
    runtime_version = '3.8'
    packages = ('snowflake-snowpark-python')
    handler = 'show_grants_to_principal_py'
as $$
def show_grants_to_principal_py(snowpark_session, principal_type_py:str, principal_identifier_py:str):
    res = []
    for row in snowpark_session.sql(f'SHOW GRANTS TO {principal_type_py} {principal_identifier_py}').to_local_iterator():
        res.append(row.as_dict())
    return res
$$
call show_grants_to_principal('%[1]s', '%[2]s');
	`, principalType, principalIdentifier)
}

func (*Grant) IsProcedure() bool {
	return true
}
