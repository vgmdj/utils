package chars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeLeftInt(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(TakeLeftInt(1234567, 3), 123)
}

func TestToFloat64(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(ToFloat64("32.00"), float64(32))
	ast.Equal(ToFloat64(32.01), 32.01)
	ast.Equal(ToFloat64(32), float64(32))
	ast.Equal(ToFloat64("32.11"), 32.11)
	ast.Equal(ToFloat64("32.111"), 32.111)
}

func TestToInt(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(ToInt("32"), 32)
	ast.Equal(ToInt(32), 32)
	ast.Equal(ToInt(32.11), 32)

}

func TestToString(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(ToString("32"), "32")
	ast.Equal(ToString(32), "32")
	ast.Equal(ToString(32.01), "32.0100")
	ast.Equal(ToString(32.00), "32.0000")
	ast.Equal(ToString(32.10, 2), "32.10")
	ast.Equal(ToString(32.001, 4), "32.0010")
}
