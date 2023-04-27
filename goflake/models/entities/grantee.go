package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
	enum "github.com/tsanton/goflake-client/goflake/models/enums"
)

var (
	_ ISnowflakeEntity = Grantee{}
)

type Grantee struct {
	GranteeIdentifier  string                       `db:"grantee_name" json:"grantee_name"`
	PrincipalType      string                       `db:"granted_to" json:"granted_to"`
	GrantedOn          enum.SnowflakeObject         `db:"granted_on" json:"granted_on"`
	GrantedIdentifier  string                       `db:"role" json:"role"`
	GrantedBy          string                       `db:"granted_by" json:"granted_by"`
	Created            c.SnowflakeDatetimeConverter `db:"created_on" json:"created_on"`
	DistanceFromSource int                          `db:"distance_from_source" json:"distance_from_source"`
}

// GetIdentity implements ISnowflakeEntity
func (Grantee) GetIdentity() string {
	panic("unimplemented")
}
