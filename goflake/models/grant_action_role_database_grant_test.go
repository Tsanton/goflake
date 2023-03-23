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

func Test_grant_role_database_privilege(t *testing.T) {
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
	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.Role]{Principal: &role, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: "IGT_DEMO_ROLE"})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 1)
	dbUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbUsage.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbUsage.GrantedOn)
}

func Test_grant_role_database_privileges(t *testing.T) {
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
	role := a.Role{
		Name:    "IGT_DEMO_ROLE",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.Role]{Principal: &role, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage, enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: "IGT_DEMO_ROLE"})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Equal(t, role.Name, res.RoleName)
	assert.Len(t, res.Grants, 2)

	dbUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbUsage.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbUsage.GrantedOn)

	dbMonitor, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbMonitor.GrantedOn)
}
