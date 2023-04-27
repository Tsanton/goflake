package describables

import (
	"fmt"

	"github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeDescribable = &Grant{}
)

type Grant struct {
	Principal ISnowflakeGrantPrincipal
}

func (r *Grant) GetDescribeStatement() string {
	switch r.Principal.GetPrincipalType() {
	case enums.SnowflakePrincipalRole, enums.SnowflakePrincipalDatabaseRole:
		break
	default:
		panic("Show grants is not implementer for this principal type")
	}
	return fmt.Sprintf("SHOW GRANTS TO %s %s", r.Principal.GetPrincipalType().GrantType(), r.Principal.GetPrincipalIdentifier())
}

func (*Grant) IsProcedure() bool {
	return false
}
