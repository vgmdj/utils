package excel

import (
	"fmt"
	"strconv"

	"github.com/Luxurioust/excelize"
	"github.com/vgmdj/utils/logger"
)

//CreateFile 创建excel
func (excel *Excel) CreateFile() (err error) {
	if err = excel.checkExcel(); err != nil {
		logger.Error(err)
		return
	}

	xlsx := excelize.NewFile()
	values := make(map[string]string)
	excel.setValues(values)
	setCellValue(xlsx, "sheet1", values)

	err = xlsx.SaveAs(excel.FileName)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (excel *Excel) checkExcel() (err error) {
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

func (excel *Excel) setValues(values map[string]string) {
	setTitle(excel.TitleKey, excel.TitleName, values)
	setContent(excel.TitleKey, excel.Content, values)
}

func setCellValue(xlsx *excelize.File, sheet string, values map[string]string) {
	for k, v := range values {
		xlsx.SetCellValue(sheet, k, v)
	}
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
		axis := columnTitle(k+1) + strconv.Itoa(line)
		cellValue[axis] = v
	}

	return
}
