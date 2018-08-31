package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //需要padding的数目
	//只要少于256就能放到一个byte中，默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	//最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

//AesCBCEncrypt aes cbc 128 pkcs7padding mode
func AesCBCEncrypt(plaintext, key, iv string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding([]byte(plaintext), blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

//AesCBCEncrypt aes cbc 128 pkcs7padding mode
func AesEBCEncrypt(plaintext, key string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding([]byte(plaintext), blockSize)

	//存储每次加密的数据
	code = make([]byte, len(origData))
	tmpData := make([]byte, blockSize)

	//分组分块加密
	for index := 0; index < len(origData); index += blockSize {
		block.Encrypt(tmpData, origData[index:index+blockSize])
		copy(code, tmpData)
	}

	return
}

//AesCBCEncrypt aes cbc 128 pkcs7padding mode
func AesEBCDecrypt(cipherText, key string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := []byte(cipherText)

	//存储每次加密的数据
	tmpData := make([]byte, blockSize)

	//分组分块解密
	for index := 0; index < len(origData); index += blockSize {
		block.Decrypt(tmpData, origData[index:index+blockSize])
		copy(origData, tmpData)
	}

	return PKCS5UnPadding(origData), nil
}

func aesEncrypt(ext, key string) (code []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding([]byte(ext), blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:blockSize])

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}
