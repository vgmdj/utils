package chars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIntegerOrAlphabet(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(IsIntegerOrAlphabet("adsaASDF1234"), true)
	ast.Equal(IsIntegerOrAlphabet("adsa1234"), true)
	ast.Equal(IsIntegerOrAlphabet("adsaasdgbfg"), true)
	ast.Equal(IsIntegerOrAlphabet("adsaASDF1234*&^%"), false)
	ast.Equal(IsIntegerOrAlphabet("adsaASDF*%"), false)
	ast.Equal(IsIntegerOrAlphabet("adsaA#$%^SDF"), false)

}
