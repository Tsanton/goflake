package models_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	dg "github.com/tsanton/goflake-client/goflake/models/describables/grants"
	eg "github.com/tsanton/goflake-client/goflake/models/entities/grants"
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_grant_database_role_database_privilege(t *testing.T) {
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
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.DatabaseRole]{Principal: &databaseRole, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateDatabaseRole},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.DatabaseRoleGrant{RoleName: "IGT_DEMO_ROLE", DatabaseName: db.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, databaseRole.Name, res.RoleName)
	assert.Len(t, res.Grants, 2) //Database roles are created with usage on database
	dbCreateRole, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeCreateDatabaseRole })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbCreateRole.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbCreateRole.GrantedOn)
}

func Test_grant_database_role_database_privileges(t *testing.T) {
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
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.DatabaseRole]{Principal: &databaseRole, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateDatabaseRole, enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.DatabaseRoleGrant{RoleName: "IGT_DEMO_ROLE", DatabaseName: db.Name})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, databaseRole.Name, res.RoleName)
	assert.Len(t, res.Grants, 3) //Database roles are created with usage on database

	dbCreateRole, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeCreateDatabaseRole })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbCreateRole.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbCreateRole.GrantedOn)

	dbMonitor, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbMonitor.GrantedOn)
}
