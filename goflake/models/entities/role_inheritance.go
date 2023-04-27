package entities

import (
	"fmt"
	"strings"

	c "github.com/tsanton/goflake-client/goflake/models/converters"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

var (
	_ ISnowflakeEntity = RoleInheritance{}
)

type RoleInheritance struct {
	ChildRoleName  string                   `json:"child_role_name"`
	ParentRoleName string                   `json:"parent_role_name"`
	GrantOption    c.SnowflakeBoolConverter `json:"grant_option"`
	GrantedBy      string                   `json:"granted_by"`
	GrantedOn      u.SnowTime               `json:"created_on"`
}

func (r RoleInheritance) GetIdentity() string {
	return strings.ToUpper(fmt.Sprintf("%[1]s.%[2]s", r.ChildRoleName, r.ParentRoleName))
}
