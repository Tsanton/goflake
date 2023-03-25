package models_test

import (
	g "github.com/tsanton/goflake-client/goflake"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func bootstrapTableAssets(cli *g.GoflakeClient, stack u.Stack[ai.ISnowflakeAsset]) (a.Database, a.Schema) {
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
	_ = g.RegisterAsset(cli, &db, &stack)

	_ = g.RegisterAsset(cli, &sch, &stack)

	return db, sch
}
