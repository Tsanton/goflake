package models_test

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_grant_database_role_future_schema_privilege(t *testing.T) {
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
	schema := a.Schema{
		Database: db,
		Name:     "IGT_GRANT",
		Comment:  "integration test goflake",
		Owner:    &a.Role{Name: "SYSADMIN"},
	}
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Principal:  &databaseRole,
		Target:     &a.GrantActionFutureSchemaGrant{DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectTable},
		Privileges: []enums.Privilege{enums.PrivilegeSelect},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[*e.FutureGrant](cli, &d.FutureGrant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 1)
	schemaFutureSelect, ok := lo.Find(grants, func(i *e.FutureGrant) bool { return i.Privilege == enums.Privilege(enums.PrivilegeSelect.String()) })
	assert.True(t, ok)
	assert.Equal(t, schemaFutureSelect.GrantedOn, enums.SnowflakeObjectTable)
	assert.Equal(t, fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectTable.ToSingular()), schemaFutureSelect.GrantedIdentifier)
}

func Test_grant_database_role_future_schema_privileges(t *testing.T) {
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
	schema := a.Schema{
		Database: db,
		Name:     "IGT_GRANT",
		Comment:  "integration test goflake",
		Owner:    &a.Role{Name: "SYSADMIN"},
	}
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege1 := a.GrantAction{
		Principal:  &databaseRole,
		Target:     &a.GrantActionFutureSchemaGrant{DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectTable},
		Privileges: []enums.Privilege{enums.PrivilegeSelect, enums.PrivilegeUpdate},
	}

	privilege2 := a.GrantAction{
		Principal:  &databaseRole,
		Target:     &a.GrantActionFutureSchemaGrant{DatabaseName: db.Name, SchemaName: schema.Name, ObjectType: enums.SnowflakeObjectView},
		Privileges: []enums.Privilege{enums.PrivilegeSelect, enums.PrivilegeReferences},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege1, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege2, &stack))

	grants, err := g.DescribeMany[*e.FutureGrant](cli, &d.FutureGrant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 4)

	tableSchemaScope := fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectTable.ToSingular())
	_, ok := lo.Find(grants, func(i *e.FutureGrant) bool {
		return i.Privilege == enums.PrivilegeSelect && i.GrantedIdentifier == tableSchemaScope
	})
	assert.True(t, ok)

	_, ok = lo.Find(grants, func(i *e.FutureGrant) bool {
		return i.Privilege == enums.PrivilegeUpdate && i.GrantedIdentifier == tableSchemaScope
	})
	assert.True(t, ok)

	viewSchemaScope := fmt.Sprintf("%[1]s.%[2]s.<%[3]s>", db.Name, schema.Name, enums.SnowflakeObjectView.ToSingular())
	_, ok = lo.Find(grants, func(i *e.FutureGrant) bool {
		return i.Privilege == enums.PrivilegeSelect && i.GrantedIdentifier == viewSchemaScope
	})
	assert.True(t, ok)

	_, ok = lo.Find(grants, func(i *e.FutureGrant) bool {
		return i.Privilege == enums.PrivilegeReferences && i.GrantedIdentifier == viewSchemaScope
	})
	assert.True(t, ok)
}
