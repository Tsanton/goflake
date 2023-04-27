package enums

type SnowflakePrincipal string

func (p SnowflakePrincipal) SnowflakeType() string {
	return string(p)
}
func (p SnowflakePrincipal) GrantType() string {
	switch p {
	case SnowflakePrincipalDatabaseRole:
		return "DATABASE ROLE"
	default:
		return string(p)
	}
}

const (
	SnowflakePrincipalUser         SnowflakePrincipal = "USER"
	SnowflakePrincipalRole         SnowflakePrincipal = "ROLE"
	SnowflakePrincipalDatabaseRole SnowflakePrincipal = "DATABASE_ROLE"
)
