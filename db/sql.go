package db

import (
	"fmt"
	"github.com/vgmdj/utils/chars"
	"strconv"
)

type OP string

const (
	EQ   OP = "="
	LT   OP = "<"
	LE   OP = "<="
	NE   OP = "!="
	GT   OP = ">"
	GE   OP = ">="
	LIKE OP = "like"
)

func AttachOr(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, "or")
}

func AttachAnd(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, "and")
}

func Attach(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, " ")
}

func Limit(sql string, l, o string) string {
	var (
		limit, _  = strconv.Atoi(l)
		offset, _ = strconv.Atoi(o)
	)

	if limit == 0 {
		return sql
	}

	if offset == 0 {
		return fmt.Sprintf(" %s limit %d ", sql, limit)
	}

	return fmt.Sprintf(" %s limit %d, %d ", sql, offset, limit)

}

func attach(sql string, query interface{}, data interface{}, op OP, relation string) string {
	if data == "" || !checkOp(op) {
		return sql
	}

	sql += fmt.Sprintf(" %s %v %v '%v' ", relation, query, op, data)
	return sql
}

func checkOp(op OP) bool {
	ops := []interface{}{EQ, LT, LE, NE, GT, GE}

	return chars.IsContain(ops, op)
}
