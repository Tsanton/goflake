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

func Test_create_database(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
}

func Test_describe_database(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))

	/* Act */
	ddb, err := g.DescribeOne[e.Database](cli, &d.Database{Name: db.Name})

	/* Assert */
	assert.Nil(t, err)
	assert.Equal(t, ddb.Name, db.Name)
}

// TODO: implement
func Test_describe_non_existing_database(t *testing.T) {}
