package area_test

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/vgmdj/gb2260/gbdata"
	"github.com/vgmdj/utils/area"
	"testing"
)

func TestIDArea(t *testing.T) {
	ast := assert.New(t)

	gb2260 := area.NewArea(area.GB2260)
	bj := gb2260.Get("110101")
	ast.Equal(bj.Province, "北京市")
	ast.Equal(bj.County, "东城区")
	ast.Equal(bj.FullName(), "北京市东城区")

	fx := gb2260.Get("320321")
	ast.Equal(fx.Province, "江苏省")
	ast.Equal(fx.City, "徐州市")
	ast.Equal(fx.County, "丰县")
	ast.Equal(fx.FullName(), "江苏省徐州市丰县")

}
