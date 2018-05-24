package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIdCodeHelper(t *testing.T) {
	ast := assert.New(t)

	idCode, err := ParseIdCard("320321199007111234")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	idCode2, err := ParseIdCard("320321199005243903")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	idCode3, err := ParseIdCard("320321201705241239")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ast.Equal(idCode.CardNum, "320321199007111234")

	//随时间变化
	ast.Equal(idCode.GetAge(), 27)
	ast.Equal(idCode2.GetAge(), 28)
	ast.Equal(idCode3.GetAge(), 1)

	ast.Equal(idCode.GetBirthday().Year(), 1990)
	ast.Equal(idCode.GetBirthday().Month(), time.July)
	ast.Equal(idCode.GetBirthday().Day(), 11)
	ast.Equal(idCode.GetLastFour(), "1234")

}
