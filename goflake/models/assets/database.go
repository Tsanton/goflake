package assets

import (
	"fmt"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	e "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset = &Database{}
)

type Database struct {
	Name    string
	Comment string
	Owner   i.ISnowflakePrincipal
}

// GetCreateStatement implements ISnowflakeAsset
func (r *Database) GetCreateStatement() (string, int) {
	var principal e.SnowflakePrincipal
	switch x := r.Owner.GetPrincipalType(); x {
	case e.SnowflakePrincipalRole, e.SnowflakePrincipalDatabaseRole:
		principal = x
	default:
		panic("Ownership for this principal type is not implemented")
	}
	return fmt.Sprintf(`
CREATE OR REPLACE DATABASE %[1]s COMMENT = '%[2]s';
GRANT OWNERSHIP ON DATABASE %[1]s TO %[3]s %[4]s;`,
		r.Name, r.Comment, principal.GrantType(), r.Owner.GetIdentifier(),
	), 2
}

// GetDeleteStatement implements ISnowflakeAsset
func (r *Database) GetDeleteStatement() (string, int) {
	return fmt.Sprintf(`DROP DATABASE IF EXISTS %[1]s CASCADE;`, r.Name), 1
}
