package models_test

import (
	"testing"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_create_role_relationship(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	rc := a.Role{
		Name:    "IGT_CHILD_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rp := a.Role{
		Name:    "IGT_PARENT_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rel := a.RoleRelationship{
		ChildRoleName:  rc.Name,
		ParentRoleName: rp.Name,
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rc, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rp, &stack))

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel, &stack))
	child, cerr := g.DescribeOne[e.Role](cli, &d.Role{Name: rc.Name})
	parent, perr := g.DescribeOne[e.Role](cli, &d.Role{Name: rp.Name})

	/* Assert */
	if cerr != nil || rc.Name != child.Name || child.GrantedToRoles != 1 {
		t.FailNow()
	}

	if perr != nil || rp.Name != parent.Name || parent.GrantedRoles != 1 {
		t.FailNow()
	}
}

func Test_describe_role_relationship(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	rc := a.Role{
		Name:    "IGT_CHILD_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rp := a.Role{
		Name:    "IGT_PARENT_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rel := a.RoleRelationship{
		ChildRoleName:  rc.Name,
		ParentRoleName: rp.Name,
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rc, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rp, &stack))

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel, &stack))
	dr, err := g.DescribeOne[e.RoleRelationship](cli, &d.RoleRelationship{ChildRoleName: rc.Name, ParentRoleName: rp.Name})
	i.ErrorFailNow(t, err)

	/* Assert */
	if dr.ParentRoleName != rp.Name || dr.ChildRoleName != rc.Name {
		t.Fail()
	}
}
