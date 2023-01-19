package assets

import "fmt"

var (
	_ ISnowflakeAsset = &RoleRelationship{}
)

type RoleRelationship struct {
	ChildRoleName  string
	ParentRoleName string
}

func (r *RoleRelationship) GetCreateStatement() (string, int) {
	return fmt.Sprintf("GRANT ROLE %[1]s TO ROLE %[2]s;", r.ChildRoleName, r.ParentRoleName), 1
}

func (r *RoleRelationship) GetDeleteStatement() (string, int) {
	return fmt.Sprintf("REVOKE ROLE %[1]s FROM ROLE %[2]s;", r.ChildRoleName, r.ParentRoleName), 1
}
