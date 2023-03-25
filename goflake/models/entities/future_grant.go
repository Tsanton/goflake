package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeEntity = &FutureGrant{}
)

type FutureGrant struct {
	GranteeIdentifier string                       `json:"grantee_name"`
	PrincipalType     string                       `json:"grant_to"`
	GrantedOn         enum.SnowflakeObject         `json:"grant_on"`
	GrantedIdentifier string                       `json:"name"`
	Privilege         enum.Privilege               `json:"privilege"`
	GrantOption       c.SnowflakeBoolConverter     `json:"grant_option"`
	Created           c.SnowflakeDatetimeConverter `json:"created_on"`
}

func (r *FutureGrant) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}
