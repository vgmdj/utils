package chars

const (
	//DefaultMultiples 缺省缩放倍数
	DefaultMultiples = 100
)

type conversion struct {
	base      float64
	result    Result
	multiples int
}

func NewConversion(base interface{}, mutiples ...int) *conversion {
	c := new(conversion)

	c.multiples = DefaultMultiples
	if len(mutiples) != 0 {
		c.multiples = mutiples[0]
	}

	c.base = ToFloat64(base)

	return c
}

func (c *conversion) SetMultiples(multiples int) {
	c.multiples = multiples
}

func (c *conversion) BaseValue() float64 {
	return c.base
}

func (c *conversion) ZoomOut() Result {
	c.result = Result(c.base * float64(c.multiples))
	return c.result
}

func (c *conversion) ZoomIn() Result {
	c.result = Result(c.base / float64(c.multiples))
	return c.result
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
