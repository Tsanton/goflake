package assets

import (
	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ i.ISnowflakeAsset = &GrantAction{}
)

type GrantAction struct {
	Target     i.ISnowflakeGrantAsset
	Privileges []enum.Privilege
}

func (r *GrantAction) GetCreateStatement() (string, int) {
	return r.Target.GetGrantStatement(r.Privileges)
}

func (r *GrantAction) GetDeleteStatement() (string, int) {
	return r.Target.GetRevokeStatement(r.Privileges)
}
