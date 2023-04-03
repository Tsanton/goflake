package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeEntity = Grant{}
)

type Grant struct {
	GranteeIdentifier  string                       `db:"grantee_name" json:"grantee_name"`
	PrincipalType      string                       `db:"granted_to" json:"granted_to"`
	GrantedOn          enum.SnowflakeObject         `db:"granted_on" json:"granted_on"`
	GrantedIdentifier  string                       `db:"name" json:"name"`
	Privilege          enum.Privilege               `db:"privilege" json:"privilege"`
	GrantOption        c.SnowflakeBoolConverter     `db:"grant_option" json:"grant_option"`
	GrantedBy          string                       `db:"granted_by" json:"granted_by"`
	Created            c.SnowflakeDatetimeConverter `db:"created_on" json:"created_on"`
	DistanceFromSource int                          `db:"distance_from_source" json:"distance_from_source"`
}

func (r Grant) GetIdentity() string {
	return "implements ISnowflakeEntity interface"
}
