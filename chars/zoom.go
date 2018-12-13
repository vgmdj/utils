package chars

import (
	"fmt"
	"strings"
	"sync"
)

const (
	//DefaultMultiples 缺省缩放倍数
	DefaultMultiples = 100
)

var (
	single *Conversion
	once   sync.Once

	zero = "0"
)

func init() {
	for i := 0; i < 100; i++ {
		zero = fmt.Sprintf("%s%s", zero, "0")
	}
}

//Conversion 放大缩小转换
type Conversion struct {
	base      string
	result    Result
	multiples int
}

//NewConversion first param is base, second param is multiples
func NewConversion(params ...interface{}) *Conversion {
	c := new(Conversion)

	c.base = "0"
	if len(params) != 0 {
		c.base = zero + ToString(params[0]) + zero
	}

	c.multiples = parseMultiples(params[:]...)

	return c
}

//Sc single pattern return the same conversion
func Sc() *Conversion {
	once.Do(func() {
		single = NewConversion()
	})
	return single
}

//SetMultiples 设置倍数
func (c *Conversion) SetMultiples(multiples int) {
	c.multiples = multiples
}

//BaseValue 返回基数
func (c Conversion) BaseValue() string {
	return ToString(c.base, 0)
}

//ZoomOut 放大
func (c Conversion) ZoomOut(base ...interface{}) Result {
	if len(base) != 0 {
		c.base = zero + ToString(base[0]) + zero
	}

	runes := []rune(c.base)

	move := 0
	for i := len(zero); i < len(runes)-len(zero)+2 && move < precision(c.multiples); i++ {
		if runes[i] == '.' {
			runes[i+1], runes[i] = runes[i], runes[i+1]
			move++
		}
	}

	if move == 0 {
		runes[len(c.base)-len(zero)+precision(c.multiples)] = '.'
	}

	c.result = Result(strings.Trim(string(runes), "0"))

	return c.result
}

//ZoomIn 缩小
func (c Conversion) ZoomIn(base ...interface{}) Result {
	if len(base) != 0 {
		c.base = zero + ToString(base[0]) + zero
	}

	runes := []rune(c.base)

	move := 0
	for i := len(runes) - len(zero) - 1; i > len(zero)-2 && move < precision(c.multiples); i-- {
		if runes[i] == '.' {
			runes[i-1], runes[i] = runes[i], runes[i-1]
			move++
		}
	}

	if move != 0 {
		c.result = Result(strings.Trim(string(runes), "0"))
		return c.result
	}

	runes[len(runes)-len(zero)] = '.'

	for i := len(runes) - len(zero); move < precision(c.multiples); i, move = i-1, move+1 {
		runes[i-1], runes[i] = runes[i], runes[i-1]
	}

	c.result = Result(strings.Trim(string(runes), "0"))

	return c.result
}

func parseMultiples(params ...interface{}) int {
	if len(params) != 2 {
		return DefaultMultiples
	}

	m, ok := params[1].(int)
	if !ok || m%10 != 0 {
		return DefaultMultiples
	}

	return m
}

func precision(multiples int) int {
	var p int
	for multiples != 0 {
		multiples /= 10
		p++
	}

	return p - 1

}

//Result 转换结果
type Result string

//ToString 字符显示
func (r Result) ToString(prec ...int) string {
	return ToString(string(r), prec[:]...)
}

//ToInt 整数显示
func (r Result) ToInt() int {
	return ToInt(string(r))
}

//ToFloat64 浮点数显示
func (r Result) ToFloat64() float64 {
	return ToFloat64(string(r))
}
