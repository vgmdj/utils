package excel

import "testing"

func TestCreateFile(t *testing.T) {
	exl := Excel{}
	exl.FileName = "./test.xlsx"

	exl.TitleKey = []string{"a", "b", "c", "d"}
	titles := make(map[string]string)
	titles["a"] = "北京榕树科技有限公司报表"
	titles["b"] = "B"
	titles["c"] = "C"
	titles["d"] = "D"
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

	CreateFile(exl)
}
