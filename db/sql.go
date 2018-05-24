package db

import (
	"fmt"
	"github.com/vgmdj/utils/chars"
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
