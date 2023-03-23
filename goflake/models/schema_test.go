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

func Test_create_schema(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	db := a.Database{
		Name:    "IGT_TEST_DB",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}
	sch := a.Schema{
		Database: db,
		Name:     "IGT_TEST_SCHEMA",
		Comment:  "integration test goflake",
		Owner:    &a.Role{Name: "SYSADMIN"},
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &sch, &stack))
}

func Test_describe_schema(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	/* Arrange */
	db := a.Database{
		Name:    "IGT_TEST_DB",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}
	sch := a.Schema{
		Database: db,
		Name:     "IGT_TEST_SCHEMA",
		Comment:  "integration test goflake",
		Owner:    &a.Role{Name: "SYSADMIN"},
	}
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &sch, &stack))

	/* Act */
	dsch, err := g.DescribeOne[e.Schema](cli, &d.Schema{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
	})
	i.ErrorFailNow(t, err)

	if dsch.Name != sch.Name || dsch.DatabaseName != sch.Database.Name {
		t.Fail()
	}
}
