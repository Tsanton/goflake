package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

var (
	_ i.ISnowflakeAsset = &Database{}
)

type Database struct {
	Name    string
	Comment string
	Owner   i.ISnowflakePrincipal
}

func (r *Database) GetCreateStatement() (string, int) {
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
CREATE OR REPLACE DATABASE %[1]s COMMENT = '%[2]s';
GRANT OWNERSHIP ON DATABASE %[1]s TO %[3]s %[4]s;`,
		r.Name, r.Comment, principalType, r.Owner.GetIdentifier(),
	), 2
}

func (r *Database) GetDeleteStatement() (string, int) {
	return fmt.Sprintf(`DROP DATABASE IF EXISTS %[1]s CASCADE;`, r.Name), 1
}
