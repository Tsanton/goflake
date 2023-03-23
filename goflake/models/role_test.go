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

func Test_create_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
}

func Test_describe_role(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	r := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}

	i.ErrorFailNow(t, g.RegisterAsset(cli, &r, &stack))

	/* Act */
	dr, err := g.DescribeOne[e.Role](cli, &d.Role{Name: r.Name})
	i.ErrorFailNow(t, err)

	/* Assert */
	assert.Equal(t, r.Name, dr.Name)
	assert.Equal(t, r.Owner.GetIdentifier(), dr.Owner)
}
