package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHideString(t *testing.T) {
	ast := assert.New(t)

	str1, _ := Replace("1234567890", 1, 1, "*")
	ast.Equal(str1, "*234567890")

	str2, _ := Replace("1234567890", 2, 1, "*")
	ast.Equal(str2, "1*34567890")

	str3, _ := Replace("1234567890", 1, 3, "*")
	ast.Equal(str3, "***4567890")

	str4, _ := Replace("1234567890", 10, 1, "*")
	ast.Equal(str4, "123456789*")

}
