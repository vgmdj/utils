package encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleNewCBCEncrypter(t *testing.T) {
	ast := assert.New(t)

	test, _ := AesCBCEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("1234567890123456"))
	ast.Equal(base64.StdEncoding.EncodeToString(test), "SGntjIjo9/rEFhZ8FvwcZg==")
}

func TestNewECBEncrypter(t *testing.T) {
	ast := assert.New(t)

	test, _ := AesECBEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"))
	ast.Equal(hex.EncodeToString(test), "92401b0fa38ce448d23af9ce5a16aa19")

	text, _ := hex.DecodeString("92401b0fa38ce448d23af9ce5a16aa19")
	plaint, _ := AesECBDecrypt(text, []byte("zhwwtoo786bbsome"))
	ast.Equal(string(plaint), "vgmdj.utils")
}

func TestAesCBCDecrypt(t *testing.T) {
	ast := assert.New(t)

	origin, _ := base64.StdEncoding.DecodeString("SGntjIjo9/rEFhZ8FvwcZg==")
	test, _ := AesCBCDecrypt(origin, []byte("zhwwtoo786bbsome"), []byte("1234567890123456"))
	ast.Equal(string(test), "vgmdj.utils")
}
