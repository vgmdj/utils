package db

import (
	"fmt"
	"strings"

	"github.com/vgmdj/utils/chars"
)

//OP 操作类型
type OP string

const (
	//EQ 等于
	EQ OP = "="
	//LT 小于
	LT OP = "<"
	//LE 小于等于
	LE OP = "<="
	//NE 不等于
	NE OP = "!="
	//GT 大于
	GT OP = ">"
	//GE 大于等于
	GE OP = ">="
	//LIKE like
	LIKE OP = "like"
)

//Count 获取个数
func Count(sql string, param ...string) string {
	p := "*"
	if len(param) != 0 {
		p = param[0]
	}

	sql = strings.Replace(sql, "\t", " ", -1)
	sql = strings.Replace(sql, "\n", " ", -1)
	sql = strings.Replace(sql, "\r", " ", -1)

	start := strings.Index(sql, " from ")

	return fmt.Sprintf("select count(%s) as count %s", p, sql[start+1:])
}

//AttachOr 附加or
func AttachOr(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, "or")
}

//AttachAnd 附加and
func AttachAnd(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, "and")
}

//Attach 附加
func Attach(sql string, query interface{}, data interface{}, op OP) string {
	return attach(sql, query, data, op, " ")
}

//Limit limit
func Limit(sql string, pageCount, pageIndex interface{}) string {
	var (
		count = chars.ToInt(pageCount)
		index = chars.ToInt(pageIndex)
	)

	limit, offset := LimitQuery(count, index)

	return fmt.Sprintf(" %s limit %d, %d ", sql, offset, limit)

}

//LimitQuery 页码和显示数，转换成limit
func LimitQuery(pageCount, pageIndex int) (limit int, offset int) {
	limit = pageCount
	offset = limit * (pageIndex - 1)

	if limit == 0 {
		return 0, 0
	}

	if offset <= 0 {
		return limit, 0
	}

	return limit, offset

}

func attach(sql string, query interface{}, data interface{}, op OP, relation string) string {
	if data == "" || !checkOp(op) {
		return sql
	}

	sql = strings.Replace(sql, "\t", " ", -1)
	sql = strings.Replace(sql, "\n", " ", -1)
	sql = strings.Replace(sql, "\r", " ", -1)

	if strings.Contains(strings.ToLower(sql), " where ") || sql == "" {
		sql = fmt.Sprintf("%s %s %v %v '%v'", sql, relation, query, op, data)
		return sql
	}

	sql = fmt.Sprintf("%s where %v %v '%v'", sql, query, op, data)
	return sql
}

func checkOp(op OP) bool {
	ops := []interface{}{EQ, LT, LE, NE, GT, GE, LIKE}

	return chars.IsContain(ops, op)
}
