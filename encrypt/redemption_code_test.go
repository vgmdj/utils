package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedemptionCode_Produce(t *testing.T) {
	ast := assert.New(t)

	codeClt := NewClient("vgmdj@github.com", 8)

	codes, err := codeClt.Produce(1000000)
	if err != nil {
		t.Error(err.Error())
	}

	checkStr := []string{}
	for _, v := range codes {
		checkStr = append(checkStr, v.FinalCode)
	}

	ast.Equal(codeClt.CheckCode(checkStr[:]...), nil)

}
