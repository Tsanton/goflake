package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeEntity = &Grant{}
)

type Grant struct {
	GranteeIdentifier  string                       `json:"grantee_name"`
	PrincipalType      string                       `json:"granted_to"`
	GrantedOn          enum.SnowflakeObject         `json:"granted_on"`
	GrantedIdentifier  string                       `json:"name"`
	Privilege          enum.Privilege               `json:"privilege"`
	GrantOption        c.SnowflakeBoolConverter     `json:"grant_option"`
	GrantedBy          string                       `json:"granted_by"`
	Created            c.SnowflakeDatetimeConverter `json:"created_on"`
	DistanceFromSource int                          `json:"distance_from_source"`
}

func (r *Grant) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}
