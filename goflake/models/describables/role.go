package describables

import "fmt"

var (
	_ ISnowflakeDescribable = &Role{}
)

type Role struct {
	Name string
}

func (r *Role) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW ROLES LIKE '%[1]s';", r.Name)
}

func (r *Role) IsProcedure() bool {
	return false
}

func (r *Role) GetPrincipalType() string {
	return "ROLE"
}
func (r *Role) GetPrincipalIdentifier() string {
	return r.Name
}
