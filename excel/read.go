package excel

import (
	"io"

	"github.com/Luxurioust/excelize"
)

//OpenReader 读入excel
func OpenReader(file io.Reader) (rows [][]string, err error) {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return
	}

	rows = xlsx.GetRows("Sheet1")
	return
}
