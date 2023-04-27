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

func Test_user_admin_direct_ascendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	/* Act */
	ascendants, err := g.DescribeMany[e.Grantee](cli, &d.PrincipalDirectAscendants{Principal: &d.Role{Name: "USERADMIN"}})

	/* Assert */
	assert.Nil(t, err)

	sa, ok := lo.Find(ascendants, func(i e.Grantee) bool { return i.GranteeIdentifier == "SECURITYADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, sa.DistanceFromSource)
}

func Test_security_admin_direct_ascendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	/* Act */
	ascendants, err := g.DescribeMany[e.Grantee](cli, &d.PrincipalDirectAscendants{Principal: &d.Role{Name: "SECURITYADMIN"}})
	//fmt.Printf(err.Error())

	/* Assert */
	assert.Nil(t, err)

	aa, ok := lo.Find(ascendants, func(i e.Grantee) bool { return i.GranteeIdentifier == "ACCOUNTADMIN" })
	assert.True(t, ok)
	assert.Equal(t, 0, aa.DistanceFromSource)
}

func Test_database_role_ascendants(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}

	/* Act */
	dbc := a.DatabaseRole{
		DatabaseName: db.Name,
		Name:         "DATABASE_CHILD",
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	dbp := a.DatabaseRole{
		DatabaseName: db.Name,
		Name:         "DATABASE_PARENT",
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	ap := a.Role{
		Name:    "ACCOUNT_PARENT",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	rel1 := a.RoleInheritance{
		ChildPrincipal:  &dbc,
		ParentPrincipal: &dbp,
	}
	rel2 := a.RoleInheritance{
		ChildPrincipal:  &dbc,
		ParentPrincipal: &ap,
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &dbc, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &dbp, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &ap, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel1, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &rel2, &stack))

	/* Act */
	ascendants, err := g.DescribeMany[e.Grantee](cli, &d.PrincipalDirectAscendants{Principal: &d.DatabaseRole{DatabaseName: dbc.DatabaseName, Name: dbc.Name}})

	/* Assert */
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ascendants))

	dpr, ok := lo.Find(ascendants, func(i e.Grantee) bool { return i.GranteeIdentifier == dbp.GetIdentifier() })
	assert.True(t, ok)
	assert.Equal(t, 0, dpr.DistanceFromSource)

	apr, ok := lo.Find(ascendants, func(i e.Grantee) bool { return i.GranteeIdentifier == ap.Name })
	assert.True(t, ok)
	assert.Equal(t, 0, apr.DistanceFromSource)
}
