package encrypt

import (
	"crypto/des"
)

func DesEncrypt(plaintext, key string) (cipherText []byte, err error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding([]byte(plaintext), blockSize)

	cipherText = make([]byte, len(origData))

	block.Encrypt(cipherText, origData)

	return
}

func DesTripleEncrypt(plaintext, key string) (cipherText []byte, err error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding([]byte(plaintext), blockSize)

	cipherText = make([]byte, len(origData))

	block.Encrypt(cipherText, origData)

	return
}
