package chars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversion(t *testing.T) {
	ast := assert.New(t)

	c1 := NewConversion("32", 10)
	ast.Equal(c1.BaseValue(), "32")
	ast.Equal(c1.ZoomOut().ToString(0), "320")
	ast.Equal(c1.ZoomOut().ToInt(), 320)
	ast.Equal(c1.ZoomIn().ToString(1), "3.2")
	ast.Equal(c1.ZoomIn().ToString(2), "3.20")
	ast.Equal(c1.ZoomIn().ToInt(), 3)
	ast.Equal(c1.ZoomIn(320).ToFloat64(), 32.00)

	c2 := NewConversion("32", 100)
	ast.Equal(c2.ZoomOut().ToString(1), "3200.0")
	ast.Equal(c2.ZoomOut().ToInt(), 3200)
	ast.Equal(c2.ZoomIn().ToString(2), "0.32")
	ast.Equal(c2.ZoomIn().ToString(1), "0.3")
	ast.Equal(c2.ZoomIn().ToInt(), 0)
	ast.Equal(c2.ZoomIn().ToFloat64(), 0.32)

	ast.Equal(Sc().ZoomOut("1.130").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.134").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.135").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.136").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.137").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.138").ToInt(), 113)
	ast.Equal(Sc().ZoomOut(1.13).ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.1349").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.1350").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.1355").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.1356").ToInt(), 113)
	ast.Equal(Sc().ZoomOut("1.1399").ToInt(), 113)
	ast.Equal(Sc().ZoomIn(1139).ToString(), "11.39")
	ast.Equal(Sc().ZoomIn(1139).ToFloat64(), 11.39)
	ast.Equal(Sc().ZoomIn(1139).ToInt(), 11)

}
