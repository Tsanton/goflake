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

func Test_create_table_time(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTableAssets(cli, &stack)
	tbl := a.Table{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TableName:    "TEST_TABLE",
		Columns:      u.Queue[a.ISnowflakeColumn]{},
		Tags:         []a.ClassificationTag{},
	}
	col1 := a.Time{
		Scale:        7,
		DefaultValue: nil,
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "TIME_COLUMN",
			PrimaryKey: false,
			ForeignKey: a.ForeignKey{},
			Tags:       []a.ClassificationTag{},
		},
	}
	tbl.Columns.Put(&col1)

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &tbl, &stack))
	dbTable, err := g.DescribeOne[e.Table](cli, &d.Table{
		DatabaseName: tbl.DatabaseName,
		SchemaName:   tbl.SchemaName,
		TableName:    tbl.TableName,
	})

	i.ErrorFailNow(t, err)
	assert.Equal(t, tbl.DatabaseName, dbTable.DatabaseName)
	assert.Equal(t, tbl.SchemaName, dbTable.SchemaName)
	assert.Equal(t, tbl.TableName, dbTable.Name)
	assert.Equal(t, 1, len(dbTable.Columns))
	assert.Equal(t, "TIME", dbTable.Columns[0].ColumnType.Type)
	assert.False(t, dbTable.Columns[0].ColumnType.Nullable)
	assert.Equal(t, col1.Scale, dbTable.Columns[0].ColumnType.Scale)
	assert.False(t, bool(dbTable.Columns[0].PrimaryKey))
	assert.False(t, bool(dbTable.Columns[0].UniqueKey))
	assert.Nil(t, dbTable.Columns[0].Default)
	assert.Nil(t, dbTable.Columns[0].Expression)
	assert.Nil(t, dbTable.Columns[0].Check)
	assert.Nil(t, dbTable.Columns[0].PolicyName)
}
