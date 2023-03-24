package entities

import (
	c "github.com/tsanton/goflake-client/goflake/models/converters"
)

type PrincipalAscendant struct {
	GranteeIdentifier  string                       `json:"grantee_name"`
	PrincipalType      string                       `json:"granted_to"`
	GrantedIdentifier  string                       `json:"role"`
	GrantedOn          string                       `json:"granted_on"`
	GrantedBy          string                       `json:"granted_by"`
	CreatedOn          c.SnowflakeDatetimeConverter `json:"created_on"`
	DistanceFromSource int                          `json:"distance_from_source"`
}

var (
	_ ISnowflakeEntity = PrincipalAscendants{}
)

type PrincipalAscendants struct {
	PrincipalIdentifier string               `json:"principal_identifier"`
	PrincipalType       string               `json:"principal_type"`
	Ascendants          []PrincipalAscendant `json:"ascendants"`
}

// GetIdentity implements ISnowflakeEntity
func (r PrincipalAscendants) GetIdentity() string {
	return r.PrincipalIdentifier
}
