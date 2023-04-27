package entities

var (
	_ ISnowflakeEntity = PrincipalAscendants{}
)

type PrincipalAscendants struct {
	PrincipalIdentifier string    `json:"principal_identifier"`
	PrincipalType       string    `json:"principal_type"`
	Ascendants          []Grantee `json:"ascendants"`
}

// GetIdentity implements ISnowflakeEntity
func (r PrincipalAscendants) GetIdentity() string {
	return r.PrincipalIdentifier
}
