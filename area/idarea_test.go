package area

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDArea(t *testing.T) {
	ast := assert.New(t)

	bj := NewArea(110101)
	ast.Equal(bj.GetProvince(), "北京市")
	ast.Equal(bj.GetCity(), "北京市市辖区")
	ast.Equal(bj.GetFullName(), "北京市东城区")

	fx := NewArea(320321)
	ast.Equal(fx.GetProvince(), "江苏省")
	ast.Equal(fx.GetCity(), "江苏省徐州市")
	ast.Equal(fx.GetFullName(), "江苏省徐州市丰县")

}
