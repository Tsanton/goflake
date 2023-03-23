package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_create_database_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DATABASE_ROLES",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}

	role := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
}

func Test_describe_database_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DATABASE_ROLES",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}

	role := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}

	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))

	/* Act */
	dr, err := g.Describe[e.Role](cli, &d.DatabaseRole{Name: role.Name, DatabaseName: db.Name})
	i.ErrorFailNow(t, err)

	/* Assert */
	assert.Equal(t, role.Name, dr.Name)
	assert.Equal(t, role.Owner.GetIdentifier(), dr.Owner)
	assert.Equal(t, role.DatabaseName, db.Name)
}
