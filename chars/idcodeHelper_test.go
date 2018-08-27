package chars_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vgmdj/utils/chars"
)

func TestIdCodeHelper(t *testing.T) {
	ast := assert.New(t)

	idCode, err := chars.ParseIDCard("320321199007111234")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	idCode2, err := chars.ParseIDCard("420621199005243903")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	idCode3, err := chars.ParseIDCard("320321201705241239")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ast.Equal(idCode.CardNum, "320321199007111234")

	//随时间变化
	ast.Equal(idCode.GetAge(), 28)
	ast.Equal(idCode2.GetAge(), 28)
	ast.Equal(idCode3.GetAge(), 1)

	ast.Equal(idCode.GetSex(), chars.Male)
	ast.Equal(idCode.GetSex().String(), "male")
	ast.Equal(idCode.GetSex().CNString(), "男")

	ast.Equal(idCode.GetBirthday().Year(), 1990)
	ast.Equal(idCode.GetBirthday().Month(), time.July)
	ast.Equal(idCode.GetBirthday().Day(), 11)
	ast.Equal(idCode.GetLastFour(), "1234")

	ast.Equal(idCode.GetPlaceOfBirth().Province, "江苏省")
	ast.Equal(idCode.GetPlaceOfBirth().City, "徐州市")
	ast.Equal(idCode.GetPlaceOfBirth().County, "丰县")
	ast.Equal(idCode.GetPlaceOfBirth().FullName(), "江苏省徐州市丰县")

	ast.Equal(idCode2.GetPlaceOfBirth().FullName(), "湖北省襄樊市襄阳县")

	ast.Equal(string(chars.GetIDCodeCheckSum("320321199308227214")), "4")
	ast.Equal(string(chars.GetIDCodeCheckSum("32032119930822721")), "4")
}
