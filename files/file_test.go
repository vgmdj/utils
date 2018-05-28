package files

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type Ts struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func TestFile(t *testing.T) {
	ast := assert.New(t)

	file := NewFile("file.go")
	ast.Equal(file.IsExist(), true)

	file.SetFileName("test.json")
	ts := Ts{"testJson", 1}
	err := file.WriteJsonTo(ts)
	if err != nil {
		t.Error(err)
		return
	}

	tc := Ts{}
	err = file.ParseFileTo(&tc)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Equal(tc == ts, true)

	content, err := file.ReadFile()
	if err != nil {
		t.Error(err)
		return
	}

	tbs, err := json.Marshal(tc)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Equal(content, tbs)

	years := []string{}
	for i := 1980; i <= 2018; i++ {
		years = append(years, strconv.Itoa(i)+",")
	}

	t.Error(years)

}
