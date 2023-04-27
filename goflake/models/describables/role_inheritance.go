package describables

import (
	"fmt"
)

var (
	_ ISnowflakeDescribable = &RoleInheritance{}
)

type RoleInheritance struct {
	ChildRoleName  string
	ParentRoleName string
}

// GetDescribeStatement implements ISnowflakeDescribable
func (r *RoleInheritance) GetDescribeStatement() string {
	return fmt.Sprintf(`
with show_role_inheritance_relationship as procedure(child_role varchar, parent_role varchar)
    returns variant not null
    language python
    runtime_version = '3.8'
    packages = ('snowflake-snowpark-python')
    handler = 'show_role_inheritance_relationship_py'
as '
def show_role_inheritance_relationship_py(snowpark_session, child_role_py: str, parent_role_py:str):
    res = {}
    for row in snowpark_session.sql(f"SHOW GRANTS ON ROLE {child_role_py.upper()}").to_local_iterator():
        if row["granted_on"] == "ROLE" and row["privilege"] == "USAGE" and row["granted_to"] == "ROLE" and row["grantee_name"] == parent_role_py.upper():
                return {**row.as_dict(), **{"child_role_name": child_role_py.upper(), "parent_role_name": parent_role_py.upper()}}
    return res
'
call show_role_inheritance_relationship('%[1]s', '%[2]s');
`,
		r.ChildRoleName, r.ParentRoleName,
	)
}

func (*RoleInheritance) IsProcedure() bool {
	return true
}
