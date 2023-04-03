package assets

import (
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset = &GrantAction{}
)

type GrantAction struct {
	Principal  i.ISnowflakePrincipal
	Target     i.ISnowflakeGrantAsset
	Privileges []enum.Privilege
}

func (g *GrantAction) GetCreateStatement() (string, int) {
	return g.Target.GetGrantStatement(g.Principal, g.Privileges)
}

func (g *GrantAction) GetDeleteStatement() (string, int) {
	return g.Target.GetRevokeStatement(g.Principal, g.Privileges)
}
