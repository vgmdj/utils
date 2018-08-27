package excel

import (
	"io"

	"github.com/Luxurioust/excelize"
	"github.com/vgmdj/utils/logger"
)

//OpenReader 读入excel
func OpenReader(file io.Reader) (rows [][]string, err error) {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	rows = xlsx.GetRows("Sheet1")
	return
}
