package excel

import (
	"github.com/Luxurioust/excelize"
	"io"
	"log"
)

func OpenReader(file io.Reader) (rows [][]string, err error) {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rows = xlsx.GetRows("Sheet1")
	return
}
