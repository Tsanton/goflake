package entities

var (
	_ ISnowflakeEntity = PrincipalDescendants{}
)

type PrincipalDescendants struct {
	PrincipalIdentifier string  `json:"principal_identifier"`
	PrincipalType       string  `json:"principal_type"`
	Descendants         []Grant `json:"descendants"`
}

// GetIdentity implements ISnowflakeEntity
func (r PrincipalDescendants) GetIdentity() string {
	return r.PrincipalIdentifier
}
