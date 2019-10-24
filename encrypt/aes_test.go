package encrypt

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCBC(t *testing.T) {
	ast := assert.New(t)

	test, err := AesCBCEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("286f67b10fb151b27efe76a95b2ba9c9", hex.EncodeToString(test))

	plainText, err := AesCBCDecrypt(test, []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(plainText))
}

func TestAesECB(t *testing.T) {
	ast := assert.New(t)

	test, err := AesECBEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("46099ff6c9ba16007ecb17647a7c398a", hex.EncodeToString(test))

	text, err := hex.DecodeString("46099ff6c9ba16007ecb17647a7c398a")
	if err != nil {
		t.Error(err.Error())
		return
	}

	plaint, err := AesECBDecrypt(text, []byte("zhwwtoo786bbsome"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(plaint))
}

func TestAesGCM(t *testing.T) {
	ast := assert.New(t)
	test, err := AesGCMEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("0123456789ab"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("7d5abbb1b5dc629f979d842360fedd7f66bacd0b9ed75821d76b1d", hex.EncodeToString(test))

	result, err := AesGCMDecrypt(test, []byte("zhwwtoo786bbsome"), []byte("0123456789ab"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(result))

}

func TestAesCFB(t *testing.T) {
	ast := assert.New(t)
	test, err := AesCFBEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("bd72e4ae5c9d10fe29005521c9d22273", hex.EncodeToString(test))

	result, err := AesCFBDecrypt(test, []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(result))
}

func TestAesCTR(t *testing.T) {
	ast := assert.New(t)
	test, err := AesCTREncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("bdab2bc580bff9a9458217df82091df5", hex.EncodeToString(test))

	result, err := AesCTRDecrypt(test, []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(result))
}

func TestAesOFB(t *testing.T) {
	ast := assert.New(t)
	test, err := AesOFBEncrypt([]byte("vgmdj.utils"), []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("bd53152444a90e99e9e7646cfc19ceb0", hex.EncodeToString(test))

	result, err := AesOFBDecrypt(test, []byte("zhwwtoo786bbsome"), []byte("0123456789abcdef"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	ast.Equal("vgmdj.utils", string(result))
}
