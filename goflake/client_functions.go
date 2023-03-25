package goflake

import (
	"context"
	"encoding/json"
	"strings"

	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	m "github.com/tsanton/goflake-client/goflake/models/mergeables"
	u "github.com/tsanton/goflake-client/goflake/utilities"
	"golang.org/x/exp/constraints"

	"github.com/snowflakedb/gosnowflake"
)

type executeScalarConstraint interface {
	constraints.Float | constraints.Integer | string | bool
}

func ExecuteScalar[T executeScalarConstraint](g *GoflakeClient, query string) (T, error) {
	var ret T
	err := g.db.Get(&ret, query)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func RegisterAsset(g *GoflakeClient, asset ai.ISnowflakeAsset, stack *u.Stack[ai.ISnowflakeAsset]) error {
	err := CreateAsset(g, asset)
	if err == nil {
		stack.Put(asset)
	}
	return err
}

func CreateAsset(g *GoflakeClient, asset ai.ISnowflakeAsset) error {
	query, numStatements := asset.GetCreateStatement()
	multiStatementContext, _ := gosnowflake.WithMultiStatement(context.Background(), numStatements)
	_, err := g.db.ExecContext(multiStatementContext, query)
	return err
}

func DeleteAssets(g *GoflakeClient, stack *u.Stack[ai.ISnowflakeAsset]) {
	for !stack.IsEmpty() {
		err := DeleteAsset(g, stack.Get())
		if err != nil {
			panic("unable to delete all assets in stack")
		}
	}
}

func DeleteAsset(g *GoflakeClient, asset ai.ISnowflakeAsset) error {
	query, numStatements := asset.GetDeleteStatement()
	multiStatementContext, _ := gosnowflake.WithMultiStatement(context.Background(), numStatements)
	_, err := g.db.ExecContext(multiStatementContext, query)
	return err
}

func DescribeOne[T e.ISnowflakeEntity](g *GoflakeClient, obj d.ISnowflakeDescribable) (T, error) {
	var ret T
	if obj.IsProcedure() {
		var procedureResponse string
		err := g.db.Get(&procedureResponse, obj.GetDescribeStatement())
		if err != nil {
			return ret, err
		}
		err = json.Unmarshal([]byte(procedureResponse), &ret)
		if err != nil {
			return ret, err
		}
	} else {
		err := g.db.Get(&ret, obj.GetDescribeStatement())
		if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
			return ret, err
		}
	}

	return ret, nil
}

func DescribeMany[T e.ISnowflakeEntity](g *GoflakeClient, obj d.ISnowflakeDescribable) ([]T, error) {
	var ret []T
	if obj.IsProcedure() {
		var procedureResponse string
		err := g.db.Get(&procedureResponse, obj.GetDescribeStatement())
		if err != nil {
			return ret, err
		}
		err = json.Unmarshal([]byte(procedureResponse), &ret)
		if err != nil {
			return ret, err
		}
	} else {
		err := g.db.Get(&ret, obj.GetDescribeStatement())
		if err != nil {
			return ret, err
		}
	}

	return ret, nil
}

func MergeInto(g *GoflakeClient, obj m.ISnowflakeMergeable) error {
	return nil
}

func GetMergable(g *GoflakeClient, obj m.ISnowflakeMergeable) (m.ISnowflakeMergeable, error) {
	return nil, nil
}

// TODO: change any for entity.Procedure
// TODO: also need assets.Procedure
func ExecuteProcedure(g *GoflakeClient, proc any) error {
	return nil
}
