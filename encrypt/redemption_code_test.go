package encrypt

import (
	"github.com/vgmdj/utils/chars"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedemptionCode_Produce(t *testing.T) {
	ast := assert.New(t)

	codeClt := NewClient("vgmdj@github.com", 8)

	codes, err := codeClt.Produce(2)
	if err != nil {
		t.Error(err.Error())
		return
	}

	checkStr := []string{}
	for _, v := range codes {
		checkStr = append(checkStr, v.FinalCode)
	}

	ast.Equal(codeClt.CheckCode(checkStr[:]...), nil)
	ast.Equal(false,chars.IsDuplicates(checkStr))

}
