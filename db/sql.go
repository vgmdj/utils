package db

import (
	"fmt"
	"github.com/vgmdj/utils/chars"
)

const (
	EQ   = "="
	LT   = "<"
	LE   = "<="
	NE   = "!="
	GT   = ">"
	GE   = ">="
	LIKE = "like"
)

func AttachOr(sql string, query interface{}, data interface{}, op string) string {
	return attach(sql, query, data, op, "or")
}

func AttachAnd(sql string, query interface{}, data interface{}, op string) string {
	return attach(sql, query, data, op, "and")
}

func Attach(sql string, query interface{}, data interface{}, op string) string {
	return attach(sql, query, data, op, " ")
}

func attach(sql string, query interface{}, data interface{}, op string, relation string) string {
	if data == "" || !checkOp(op) {
		return sql
	}

	sql += fmt.Sprintf(" %s %v %v '%v' ", relation, query, op, data)
	return sql
}

func checkOp(op string) bool {
	ops := []string{EQ, LT, LE, NE, GT, GE}

	return chars.IsStringContain(op, ops)
}
