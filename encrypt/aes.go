package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//AesCBCEncrypt aes cbc 128 pkcs7padding mode
func AesCBCEncrypt(plaintext, key, iv string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := pkcs5Padding([]byte(plaintext), blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

func AesCBCDecrypt() {

}

func aesEncrypt(ext, key string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := pkcs5Padding([]byte(ext), blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:blockSize])

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
