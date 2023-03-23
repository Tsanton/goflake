package describables

import (
	"fmt"
)

var (
	_ ISnowflakeDescribable = &FutureGrant{}
)

type FutureGrant struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *FutureGrant) GetDescribeStatement() string {
	switch any(r.Principal).(type) {
	case *Role:
		return fmt.Sprintf(`
with show_future_grants_to_role as procedure(role_name varchar)
	returns variant not null
	language python
	runtime_version = '3.8'
	packages = ('snowflake-snowpark-python')
	handler = 'show_grants_to_role_py'
as $$
def show_grants_to_role_py(snowpark_session, role_name_py:str):
	res = []
	for row in snowpark_session.sql(f'SHOW FUTURE GRANTS TO ROLE {role_name_py.upper()}').to_local_iterator():
		res.append(row.as_dict())
	return res
$$
call show_future_grants_to_role('%[1]s');
	`, r.Principal.GetPrincipalIdentifier())
	case *DatabaseRole:
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
		`, r.Principal.(*DatabaseRole).DatabaseName, r.Principal.(*DatabaseRole).Name)
	default:
		panic("Show future grants is not implementer for this principal type")
	}
}

func (*FutureGrant) IsProcedure() bool {
	return true
}
