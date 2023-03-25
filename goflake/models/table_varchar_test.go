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

func Test_create_table_varchar(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTableAssets(cli, stack)
	tbl := a.Table{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TableName:    "TEST_TABLE",
		Columns:      u.Queue[a.ISnowflakeColumn]{},
		Tags:         []a.ClassificationTag{},
	}
	col1 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN",
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
	// assert.True(t, lo.FindOrElse(dbTable.Columns, e.Column{}, func(i e.Column) bool {return i.}))
	assert.Equal(t, "TEXT", dbTable.Columns[0].ColumnType.Type)
	assert.False(t, dbTable.Columns[0].ColumnType.Nullable)
	assert.Equal(t, col1.Length, dbTable.Columns[0].ColumnType.Length)
	assert.False(t, bool(dbTable.Columns[0].PrimaryKey))
	assert.False(t, bool(dbTable.Columns[0].UniqueKey))
	assert.Equal(t, "", dbTable.Columns[0].Default)
	assert.Equal(t, "", dbTable.Columns[0].Expression)
	assert.Equal(t, "", dbTable.Columns[0].Check)
	assert.Equal(t, "", dbTable.Columns[0].PolicyName)
}

func Test_create_table_varchar_primary_key(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTableAssets(cli, stack)
	tbl := a.Table{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TableName:    "TEST_TABLE",
		Columns:      u.Queue[a.ISnowflakeColumn]{},
		Tags:         []a.ClassificationTag{},
	}
	col1 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN",
			PrimaryKey: true,
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
	// assert.True(t, lo.FindOrElse(dbTable.Columns, e.Column{}, func(i e.Column) bool {return i.}))
	assert.Equal(t, "TEXT", dbTable.Columns[0].ColumnType.Type)
	assert.False(t, dbTable.Columns[0].ColumnType.Nullable)
	assert.Equal(t, col1.Length, dbTable.Columns[0].ColumnType.Length)
	assert.True(t, bool(dbTable.Columns[0].PrimaryKey))
	assert.False(t, bool(dbTable.Columns[0].UniqueKey))
	assert.Equal(t, "", dbTable.Columns[0].Default)
	assert.Equal(t, "", dbTable.Columns[0].Expression)
	assert.Equal(t, "", dbTable.Columns[0].Check)
	assert.Equal(t, "", dbTable.Columns[0].PolicyName)
}

func Test_create_table_varchar_multiple_columns(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTableAssets(cli, stack)
	tbl := a.Table{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TableName:    "TEST_TABLE",
		Columns:      u.Queue[a.ISnowflakeColumn]{},
		Tags:         []a.ClassificationTag{},
	}
	col1 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN_1",
			PrimaryKey: false,
			ForeignKey: a.ForeignKey{},
			Tags:       []a.ClassificationTag{},
		},
	}
	col2 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN_2",
			PrimaryKey: false,
			ForeignKey: a.ForeignKey{},
			Tags:       []a.ClassificationTag{},
		},
	}
	tbl.Columns.Put(&col1)
	tbl.Columns.Put(&col2)

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
	assert.Equal(t, 2, len(dbTable.Columns))
	dbCol1, ok := lo.Find(dbTable.Columns, func(i e.Column) bool { return i.Name == col1.Name })
	assert.True(t, ok)
	assert.Equal(t, "TEXT", dbCol1.ColumnType.Type)
	assert.False(t, dbCol1.ColumnType.Nullable)
	assert.Equal(t, col1.Length, dbCol1.ColumnType.Length)
	assert.False(t, bool(dbCol1.PrimaryKey))
	assert.False(t, bool(dbCol1.UniqueKey))
	assert.Equal(t, "", dbCol1.Default)
	assert.Equal(t, "", dbCol1.Expression)
	assert.Equal(t, "", dbCol1.Check)
	assert.Equal(t, "", dbCol1.PolicyName)
}

func Test_create_table_varchar_composite_primary_key(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)
	db, sch := bootstrapTableAssets(cli, stack)
	tbl := a.Table{
		DatabaseName: db.Name,
		SchemaName:   sch.Name,
		TableName:    "TEST_TABLE",
		Columns:      u.Queue[a.ISnowflakeColumn]{},
		Tags:         []a.ClassificationTag{},
	}
	col1 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN_1",
			PrimaryKey: true,
			ForeignKey: a.ForeignKey{},
			Tags:       []a.ClassificationTag{},
		},
	}
	col2 := a.Varchar{
		Length:       16777216,
		Collation:    "",
		DefaultValue: "",
		Nullable:     false,
		Unique:       false,
		ColumnFields: a.ColumnFields{
			Name:       "VARCHAR_COLUMN_2",
			PrimaryKey: true,
			ForeignKey: a.ForeignKey{},
			Tags:       []a.ClassificationTag{},
		},
	}
	tbl.Columns.Put(&col1)
	tbl.Columns.Put(&col2)

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
	assert.Equal(t, 2, len(dbTable.Columns))
	dbCol1, ok := lo.Find(dbTable.Columns, func(i e.Column) bool { return i.Name == col1.Name })
	assert.True(t, ok)
	assert.Equal(t, "TEXT", dbCol1.ColumnType.Type)
	assert.False(t, dbCol1.ColumnType.Nullable)
	assert.Equal(t, col1.Length, dbCol1.ColumnType.Length)
	assert.True(t, bool(dbCol1.PrimaryKey))
	assert.False(t, bool(dbCol1.UniqueKey))
	assert.Equal(t, "", dbCol1.Default)
	assert.Equal(t, "", dbCol1.Expression)
	assert.Equal(t, "", dbCol1.Check)
	assert.Equal(t, "", dbCol1.PolicyName)
}
