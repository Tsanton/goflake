package models_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_role_ascendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	rr := a.Role{
		Name:    "SOME_SCHEMA_R",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rrw := a.Role{
		Name:    "SOME_SCHEMA_RW",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rrwc := a.Role{
		Name:    "SOME_SCHEMA_RWC",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rel1 := a.RoleRelationship{
		ChildRoleName:  rr.Name,
		ParentRoleName: rrw.Name,
	}
	rel2 := a.RoleRelationship{
		ChildRoleName:  rrw.Name,
		ParentRoleName: rrwc.Name,
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rr, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rrw, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rrwc, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel1, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel2, &stack))

	/* Act */
	hier, err := g.DescribeOne[e.PrincipalAscendants](cli, &d.PrincipalAscendants{Principal: &d.Role{Name: rr.Name}})

	/* Assert */
	assert.Nil(t, err)
	assert.Equal(t, 2, len(hier.Ascendants))

	qrrw, ok := lo.Find(hier.Ascendants, func(i e.PrincipalAscendant) bool { return i.GranteeIdentifier == rrw.Name })
	assert.True(t, ok)
	assert.Equal(t, 0, qrrw.DistanceFromSource)

	qrrwc, ok := lo.Find(hier.Ascendants, func(i e.PrincipalAscendant) bool { return i.GranteeIdentifier == rrwc.Name })
	assert.True(t, ok)
	assert.Equal(t, 1, qrrwc.DistanceFromSource)
}

func Test_user_admin_ascendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	/* Act */
	hier, err := g.DescribeOne[e.PrincipalAscendants](cli, &d.PrincipalAscendants{Principal: &d.Role{Name: "USERADMIN"}})

	/* Assert */
	assert.Nil(t, err)

	sa, ok := lo.Find(hier.Ascendants, func(i e.PrincipalAscendant) bool { return i.GranteeIdentifier == "SECURITYADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, sa.DistanceFromSource)

	aa, ok := lo.Find(hier.Ascendants, func(i e.PrincipalAscendant) bool { return i.GranteeIdentifier == "ACCOUNTADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 1, aa.DistanceFromSource)
}
