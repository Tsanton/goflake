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
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_grant_database_role_schema_privilege(t *testing.T) {
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
		Target:     &a.GrantActionSchemaGrant{DatabaseName: db.Name, SchemaName: schema.Name},
		Privileges: []enums.Privilege{enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[e.Grant](cli, &d.Grant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 2)
	_, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)

	schemaMonitor, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectSchema, schemaMonitor.GrantedOn)

}

func Test_grant_database_role_schema_privileges(t *testing.T) {
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
		Target:     &a.GrantActionSchemaGrant{DatabaseName: db.Name, SchemaName: schema.Name},
		Privileges: []enums.Privilege{enums.PrivilegeMonitor, enums.PrivilegeUsage},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &schema, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[e.Grant](cli, &d.Grant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 3)
	_, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)

	schemaMonitor, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectSchema, schemaMonitor.GrantedOn)

	schemaUsage, ok := lo.Find(grants, func(i e.Grant) bool {
		return i.Privilege == enums.PrivilegeUsage && i.GrantedOn == enums.SnowflakeObjectSchema
	})
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", schemaUsage.GrantedBy)
}
