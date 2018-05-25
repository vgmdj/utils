package area

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDArea(t *testing.T) {
	ast := assert.New(t)

	bj := NewArea(GB2260)
	bji := bj.Get("110010")
	ast.Equal(bji.Province, "北京市")
	ast.Equal(bji.City, "北京市市辖区")
	ast.Equal(bji.FullName(), "北京市东城区")

	fx := NewArea(POST)
	fxi := fx.Get("320321")
	ast.Equal(fxi.Province, "江苏省")
	ast.Equal(fxi.City, "江苏省徐州市")
	ast.Equal(fxi.FullName(), "江苏省徐州市丰县")

}
