package chars

import "sync"

const (
	//DefaultMultiples 缺省缩放倍数
	DefaultMultiples = 100
)

var (
	single *conversion
	once   sync.Once
)

type conversion struct {
	base      float64
	result    Result
	multiples int
}

//NewConversion first param is base, second param is multiples
func NewConversion(params ...interface{}) *conversion {
	c := new(conversion)

	c.base = 0
	if len(params) != 0 {
		c.base = ToFloat64(params[0])
	}

	c.multiples = parseMultiples(params[:]...)

	return c
}

//Sc single pattern return the same conversion
func Sc() *conversion {
	once.Do(func() {
		single = NewConversion()
	})
	return single
}

func (c *conversion) SetMultiples(multiples int) {
	c.multiples = multiples
}

func (c conversion) BaseValue() float64 {
	return c.base
}

func (c conversion) ZoomOut(base ...interface{}) Result {
	if len(base) != 0 {
		c.base = ToFloat64(base[0])
	}

	c.result = Result(c.base * float64(c.multiples))
	return c.result
}

func (c conversion) ZoomIn(base ...interface{}) Result {
	if len(base) != 0 {
		c.base = ToFloat64(base[0])
	}

	c.result = Result(c.base / float64(c.multiples))
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

type Result float64

func (r Result) ToString(multiples ...int) string {
	var m int = DefaultMultiples
	if len(multiples) != 0 {
		m = multiples[0]
	}

	return ToString(float64(r), precision(m))
}

func (r Result) ToInt() int {
	return ToInt(float64(r))
}

func (r Result) ToFloat64() float64 {
	return ToFloat64(float64(r))
}
