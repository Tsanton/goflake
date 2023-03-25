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
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func bootstrapTagAssets(cli *g.GoflakeClient, stack *u.Stack[ai.ISnowflakeAsset]) (a.Database, a.Schema) {
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
	_ = g.RegisterAsset(cli, &db, stack)

	_ = g.RegisterAsset(cli, &sch, stack)

	return db, sch
}

func Test_describe_non_existing_tag(t *testing.T) {
	cli := i.Goflake()
	defer cli.Close()

	dbTag, err := g.DescribeOne[e.Tag](cli, &d.Tag{
		DatabaseName: "SNOWFLAKE",
		SchemaName:   "ACCOUNT_USAGE",
		TagName:      "I_DONT_EXIST_TAG",
	})
	assert.Nil(t, err)
	assert.Equal(t, e.Tag{}, dbTag)

}

func Test_create_tag_without_values(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTagAssets(cli, &stack)
	tag := a.Tag{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TagName:      "TEST_TAG",
		TagValues:    []string{},
		Comment:      "Goflake client test tag",
		Owner:        &a.Role{Name: "SYSADMIN"},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &tag, &stack))
	dbTag, err := g.DescribeOne[e.Tag](cli, &d.Tag{
		DatabaseName: tag.DatabaseName,
		SchemaName:   tag.SchemaName,
		TagName:      tag.TagName,
	})

	i.ErrorFailNow(t, err)
	assert.Equal(t, tag.DatabaseName, dbTag.DatabaseName)
	assert.Equal(t, tag.SchemaName, dbTag.SchemaName)
	assert.Equal(t, tag.TagName, dbTag.Name)
	assert.Equal(t, tag.Owner.GetIdentifier(), dbTag.Owner)
	assert.Empty(t, tag.TagValues)
}

func Test_create_tag_with_values(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTagAssets(cli, &stack)
	tag := a.Tag{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TagName:      "TEST_TAG",
		TagValues:    []string{"FOO", "BAR"},
		Comment:      "Goflake client test tag",
		Owner:        &a.Role{Name: "SYSADMIN"},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &tag, &stack))
	dbTag, err := g.DescribeOne[e.Tag](cli, &d.Tag{
		DatabaseName: tag.DatabaseName,
		SchemaName:   tag.SchemaName,
		TagName:      tag.TagName,
	})

	i.ErrorFailNow(t, err)
	assert.Equal(t, tag.DatabaseName, dbTag.DatabaseName)
	assert.Equal(t, tag.SchemaName, dbTag.SchemaName)
	assert.Equal(t, tag.TagName, dbTag.Name)
	assert.Equal(t, tag.Owner.GetIdentifier(), dbTag.Owner)
	assert.True(t, lo.Every(tag.TagValues, dbTag.AllowedValues))
}
