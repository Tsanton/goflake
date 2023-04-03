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
		Principal:  &role,
		Target:     &a.GrantActionDatabaseGrant{DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[e.Grant](cli, &d.Grant{Principal: &d.Role{Name: role.Name}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 1)
	dbUsage, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeUsage })
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
		Principal:  &role,
		Target:     &a.GrantActionDatabaseGrant{DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeUsage, enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &role, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[e.Grant](cli, &d.Grant{Principal: &d.Role{Name: role.Name}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 2)

	dbUsage, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeUsage })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbUsage.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbUsage.GrantedOn)

	dbMonitor, ok := lo.Find(grants, func(i e.Grant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbMonitor.GrantedOn)
}
