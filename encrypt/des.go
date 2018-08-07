package encrypt

import (
	"crypto/des"
	"github.com/vgmdj/utils/logger"
)

func DesEncrypt(plaintext, key string) (cipherText []byte, err error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	blockSize := block.BlockSize()

	origData := pkcs5Padding([]byte(plaintext), blockSize)

	cipherText = make([]byte, len(origData))

	block.Encrypt(cipherText, origData)

	return
}

func DesTripleEncrypt(plaintext, key string) (cipherText []byte, err error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	blockSize := block.BlockSize()

	origData := pkcs5Padding([]byte(plaintext), blockSize)

	cipherText = make([]byte, len(origData))

	block.Encrypt(cipherText, origData)

	return
}
