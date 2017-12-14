package excel

import (
	"github.com/Luxurioust/excelize"
	"log"
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

var (
	alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

//CreateFile
func CreateFile(excel Excel) {
	xlsx := excelize.NewFile()
	values := make(map[string]string)
	setValues(excel, values)

	setCellValue(xlsx, "sheet1", values)

	err := xlsx.SaveAs(excel.FileName)
	if err != nil {
		log.Println(err)
	}

}

//CreateFile
func CreateFileReader(excel Excel) *excelize.File {
	xlsx := excelize.NewFile()
	values := make(map[string]string)
	setValues(excel, values)

	setCellValue(xlsx, "sheet1", values)

	return xlsx
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
		axis := alphabet[k] + strconv.Itoa(line)
		cellValue[axis] = v
	}

	return
}
