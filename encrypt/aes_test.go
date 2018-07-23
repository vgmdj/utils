package encrypt

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleNewCBCEncrypter(t *testing.T) {
	ast := assert.New(t)

	test, _ := AesCBCEncrypt("vgmdj.utils", "zhwwtoo786bbsome", "1234567890123456")
	ast.Equal(base64.StdEncoding.EncodeToString(test), "SGntjIjo9/rEFhZ8FvwcZg==")
}
