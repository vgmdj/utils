package excel

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/vgmdj/utils/logger"
	"strconv"
)

type (
	Excel struct {
		FileName  string
		TitleKey  []string
		TitleName map[string]string
		Content   []map[string]string
	}
)

//CreateFile
func CreateFile(excel Excel) {
	if err := checkExcel(excel); err != nil {
		logger.Error(err)
		return
	}

	xlsx := excelize.NewFile()
	values := make(map[string]string)
	setValues(excel, values)

	setCellValue(xlsx, "sheet1", values)

	err := xlsx.SaveAs(excel.FileName)
	if err != nil {
		logger.Error(err)
		return
	}
}

func checkExcel(excel Excel) (err error) {
	if excel.FileName == "" {
		return fmt.Errorf("no file name")
	}

	if len(excel.TitleKey) != len(excel.TitleName) {
		return fmt.Errorf("can not match the title name ")
	}

	for _, v := range excel.TitleKey {
		if _, ok := excel.TitleName[v]; !ok {
			return fmt.Errorf("can not match the title name ")
		}
	}

	return
}

func setCellValue(xlsx *excelize.File, sheet string, values map[string]string) {
	for k, v := range values {
		xlsx.SetCellValue(sheet, k, v)
	}
}

func setValues(excel Excel, values map[string]string) {
	setTitle(excel.TitleKey, excel.TitleName, values)
	setContent(excel.TitleKey, excel.Content, values)
}

func setTitle(key []string, name map[string]string, values map[string]string) {
	var titles []string
	for _, v := range key {
		titles = append(titles, name[v])
	}

	setLineValues(1, titles, values)

}

func setContent(key []string, values []map[string]string, cells map[string]string) {
	for k := 0; k < len(values); k++ {
		contents := []string{}
		for _, v := range key {
			contents = append(contents, values[k][v])
		}

		setLineValues(k+2, contents, cells)
	}

}

func setLineValues(line int, values []string, cellValue map[string]string) {
	for k, v := range values {
		axis := columnTitle(k) + strconv.Itoa(line)
		cellValue[axis] = v
	}

	return
}

func columnTitle(n int) string {
	res := ""
	for n > 0 {
		res = string((n-1)%26+65) + res
		n = (n - 1) / 26
	}
	return res
}
