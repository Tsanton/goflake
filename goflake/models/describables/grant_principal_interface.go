package describables

type ISnowflakeGrantPrincipal interface {
	GetPrincipalType() string
	GetPrincipalIdentifier() string
}
