package inter

import e "github.com/tsanton/goflake-client/goflake/models/enums"

type ISnowflakePrincipal interface {
	GetIdentifier() string
	GetPrincipalType() e.SnowflakePrincipal
}
