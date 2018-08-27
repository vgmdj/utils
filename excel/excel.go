package excel

import (
	"fmt"
	"strings"
)

type (
	//Excel excel
	Excel struct {
		FileName  string
		TitleKey  []string
		TitleName map[string]string
		Content   []map[string]string
	}
)

//SetFileName set the excel file name , and if filename is null then use 'default.xlsx' as default
func (excel *Excel) SetFileName(fileName string) {
	if fileName == "" {
		excel.FileName = "default.xlsx"
		return
	}

	if strings.HasSuffix(fileName, ".xlsx") || strings.HasSuffix(fileName, ".xls") {
		excel.FileName = fileName
		return
	}

	excel.FileName = fmt.Sprintf("%s.xlsx", fileName)
}

func columnTitle(n int) string {
	res := ""
	for n > 0 {
		res = string((n-1)%26+65) + res
		n = (n - 1) / 26
	}
	return res
}

func titleToNumber(s string) int {
	sum := 0
	for _, ch := range s {
		current := int(ch-'A') + 1
		sum = sum*26 + current
	}
	return sum

}
