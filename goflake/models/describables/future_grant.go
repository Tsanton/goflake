package describables

import (
	"fmt"
	"strings"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable = &FutureGrant{}
)

type FutureGrant struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *FutureGrant) GetDescribeStatement() string {
	switch x := r.Principal.GetPrincipalType(); x {
	case enums.SnowflakePrincipalRole:
		return fmt.Sprintf("SHOW FUTURE GRANTS TO ROLE %s", r.Principal.GetPrincipalIdentifier())
	case enums.SnowflakePrincipalDatabaseRole:
		id := r.Principal.GetPrincipalIdentifier()
		databaseName, roleName := strings.Split(id, ".")[0], strings.Split(id, ".")[1]
		return fmt.Sprintf(`
with show_future_grants_to_database_role as procedure(database_name varchar, database_role_name varchar)
	returns variant not null
	language python
	runtime_version = '3.8'
	packages = ('snowflake-snowpark-python')
	handler = 'show_future_grants_to_database_role_py'
as $$
def show_future_grants_to_database_role_py(snowpark_session, database_name_py:str, database_role_name_py:str):
	res = []
	for row in snowpark_session.sql(f'SHOW FUTURE GRANTS IN DATABASE {database_name_py.upper()}').to_local_iterator():
		if row['grant_to'] == 'DATABASE_ROLE' and row['grant_to'] == 'DATABASE_ROLE':
			res.append(row.as_dict())
	for schema_object in snowpark_session.sql(f'SHOW SCHEMAS IN DATABASE {database_name_py.upper()}').to_local_iterator():
		schema_name:str = schema_object['name']
		if schema_name not in('INFORMATION_SCHEMA', 'PUBLIC'):
			query:str = f'SHOW FUTURE GRANTS IN SCHEMA {database_name_py}.{schema_name}'.upper()
			for row in snowpark_session.sql(query).to_local_iterator():
				if row['grant_to'] == 'DATABASE_ROLE' and row['grantee_name'] == database_role_name_py:
					res.append(row.as_dict())
	return res
$$
call show_future_grants_to_database_role('%[1]s','%[2]s');
		`, databaseName, roleName)
	default:
		panic("Show future grants is not implementer for this principal type")
	}
}

func (r *FutureGrant) IsProcedure() bool {
	switch r.Principal.GetPrincipalType() {
	case enums.SnowflakePrincipalRole:
		return false
	case enums.SnowflakePrincipalDatabaseRole:
		return true
	default:
		panic("Show future grants is not implementer for this principal type")
	}
}
