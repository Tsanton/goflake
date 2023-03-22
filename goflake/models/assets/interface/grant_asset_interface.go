package inter

import (
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

type ISnowflakeGrantAsset interface {
	GetGrantStatement(privileges []enum.Privilege) (string, int)
	GetRevokeStatement(privileges []enum.Privilege) (string, int)
	ValidatePrivileges(privileges []enum.Privilege) bool
}
