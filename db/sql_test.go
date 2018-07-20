package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSql(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(checkOp("="), true)
	ast.Equal(checkOp("!"), false)

}

func TestSql2(t *testing.T) {
	ast := assert.New(t)

	sql := `select * from test`
	sql = AttachAnd(sql, "a", 1, EQ)
	sql = AttachAnd(sql, "b", 2, GE)
	sql = AttachAnd(sql, "c", 3, LE)
	sql = AttachOr(sql, "d", 4, LIKE)

	ast.Equal(sql, `select * from test where a = '1' and b >= '2' and c <= '3' or d like '4'`)

	ast.Equal(Count(sql),`select count(*) as count from test where a = '1' and b >= '2' and c <= '3' or d like '4'`)
	ast.Equal(Count(sql,"a"),`select count(a) as count from test where a = '1' and b >= '2' and c <= '3' or d like '4'`)

}
