package mysql

import (
	"github.com/jinzhu/gorm"
)

type SqlExpr struct {
	*gorm.SqlExpr
}

func Expr(expression string, args ...interface{}) *SqlExpr {
	return &SqlExpr{SqlExpr: gorm.Expr(expression, args...)}
}

func transformSliceSqlExpr(args []interface{}) []interface{} {
	retArgs := make([]interface{}, 0, len(args))

	for _, arg := range args {
		if expr, ok := arg.(*SqlExpr); ok {
			retArgs = append(retArgs, expr.SqlExpr)
		} else {
			retArgs = append(retArgs, arg)
		}
	}

	return retArgs
}

func transformMapSqlExpr(args interface{}) interface{} {
	if argsMap, ok := args.(map[string]interface{}); ok {
		for key, arg := range argsMap {
			if expr, ok := arg.(*SqlExpr); ok {
				argsMap[key] = expr.SqlExpr
			}
		}

		return argsMap
	}

	return args
}
