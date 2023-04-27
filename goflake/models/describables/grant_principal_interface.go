package describables

import (
	e "github.com/tsanton/goflake-client/goflake/models/enums"
)

type ISnowflakeGrantPrincipal interface {
	GetPrincipalIdentifier() string
	GetPrincipalType() e.SnowflakePrincipal
}
