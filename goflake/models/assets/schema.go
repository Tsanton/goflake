package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

var (
	_ i.ISnowflakeAsset = &Schema{}
)

type Schema struct {
	Database Database
	Name     string
	Comment  string
	Owner    i.ISnowflakePrincipal
}

func (r *Schema) GetCreateStatement() (string, int) {
	var principalType string
	switch r.Owner.(type) {
	case *Role:
		principalType = "ROLE"
	case *DatabaseRole:
		principalType = "DATABASE ROLE"
	default:
		panic("Ownership for this principal type is not implemented")
	}
	return fmt.Sprintf(`
CREATE OR REPLACE SCHEMA %[1]s.%[2]s WITH MANAGED ACCESS COMMENT = '%[3]s';
GRANT OWNERSHIP ON SCHEMA %[1]s.%[2]s TO %[4]s %[5]s REVOKE CURRENT GRANTS;
`,
		r.Database.Name, r.Name, r.Comment, principalType, r.Owner.GetIdentifier(),
	), 2
}

func (r *Schema) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("DROP SCHEMA IF EXISTS %[1]s.%[2]s CASCADE;", r.Database.Name, r.Name), 1
}
