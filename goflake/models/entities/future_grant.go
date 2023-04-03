package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeEntity = FutureGrant{}
)

type FutureGrant struct {
	GranteeIdentifier string                       `db:"grantee_name" json:"grantee_name"`
	PrincipalType     string                       `db:"grant_to" json:"grant_to"`
	GrantedOn         enum.SnowflakeObject         `db:"grant_on" json:"grant_on"`
	GrantedIdentifier string                       `db:"name" json:"name"`
	Privilege         enum.Privilege               `db:"privilege" json:"privilege"`
	GrantOption       c.SnowflakeBoolConverter     `db:"grant_option" json:"grant_option"`
	Created           c.SnowflakeDatetimeConverter `db:"created_on" json:"created_on"`
}

func (r FutureGrant) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}
