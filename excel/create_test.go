package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFile(t *testing.T) {
	exl := Excel{}
	exl.FileName = "./test.xlsx"

	exl.TitleKey = []string{"a", "b", "c", "d"}
	titles := make(map[string]string)
	titles["a"] = "公司报表A"
	titles["b"] = "公司报表B"
	titles["c"] = "公司报表C"
	titles["d"] = "公司报表D"
	exl.TitleName = titles

	content1 := make(map[string]string)
	content1["a"] = "A1"
	content1["b"] = "B1"
	content1["c"] = "C1"
	content1["d"] = "D1"
	exl.Content = append(exl.Content, content1)

	content2 := make(map[string]string)
	content2["a"] = "A2"
	content2["b"] = "B2"
	content2["c"] = "C2"
	content2["d"] = "D2"
	exl.Content = append(exl.Content, content2)

	exl.CreateFile()

}

func TestColumnTitle(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(columnTitle(1), "A")
	ast.Equal(columnTitle(2), "B")
	ast.Equal(columnTitle(3), "C")
	ast.Equal(columnTitle(26), "Z")
	ast.Equal(columnTitle(27), "AA")
	ast.Equal(columnTitle(701), "ZY")

}
