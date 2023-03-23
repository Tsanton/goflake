package describables

import "fmt"

var (
	_ ISnowflakeDescribable = &DatabaseRole{}
)

// Beware that you cannot grant account level privleges to database roles
type DatabaseRole struct {
	Name         string
	DatabaseName string
}

func (r *DatabaseRole) GetDescribeStatement() string {
	return fmt.Sprintf("SHOW DATABASE ROLES LIKE '%[1]s' IN DATABASE %[2]s;", r.Name, r.DatabaseName)
}

func (r *DatabaseRole) IsProcedure() bool {
	return false
}

func (r *DatabaseRole) GetPrincipalType() string {
	return "DATABASE ROLE"
}
func (r *DatabaseRole) GetPrincipalIdentifier() string {
	return fmt.Sprintf("%[1]s.%[2]s", r.DatabaseName, r.Name)
}
