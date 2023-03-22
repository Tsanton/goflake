package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

var (
	_ i.ISnowflakeAsset     = &Role{}
	_ i.ISnowflakePrincipal = &Role{}
)

type Role struct {
	Name    string
	Owner   string
	Comment string
	//TODO: DatabaseName string: if != nil then https://docs.snowflake.com/en/sql-reference/sql/create-database-role.html
}

func (r *Role) GetCreateStatement() (string, int) {
	return fmt.Sprintf(`
	CREATE OR REPLACE ROLE %[1]s COMMENT = '%[2]s';
	GRANT OWNERSHIP ON ROLE %[1]s TO %[3]s REVOKE CURRENT GRANTS;`,
		r.Name, r.Comment, r.Owner,
	), 2
}

func (r *Role) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP ROLE IF EXISTS %[1]s;", r.Name), 1
}

func (r *Role) GetIdentifier() string {
	return r.Name
}
