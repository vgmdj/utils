package excel

import (
	"log"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	file, err := os.OpenFile("test.xlsx", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err.Error())
		return
	}
	OpenReader(file)
}
