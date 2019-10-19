package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/vgmdj/utils/logger"
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
	if length == 0 {
		return origData
	}

	unpadding := int(origData[length-1])
	if unpadding < 0 {
		logger.Warning(fmt.Sprintf("index out of range, want to use left %d, but length is %d",
			length-unpadding, length))
	}

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
func AesCBCEncrypt(plaintext, key, iv []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(plaintext, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

func AesCBCDecrypt(ciphertext, key, iv []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	code = make([]byte, len(ciphertext))
	mode.CryptBlocks(code, ciphertext)

	return PKCS5UnPadding(code), nil
}

//AesECBEncrypt aes cbc 128 pkcs7padding mode
func AesECBEncrypt(plaintext, key []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(plaintext, blockSize)

	blockMode := NewECBEncrypter(block)

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

//AesECBDecrypt aes cbc 128 pkcs7padding mode
func AesECBDecrypt(cipherText, key []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := NewECBDecrypter(block)

	code = make([]byte, len(cipherText))

	blockMode.CryptBlocks(code, cipherText)

	return PKCS5UnPadding(code), nil
}

func aesEncrypt(ext, key []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	origData := PKCS5Padding(ext, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	code = make([]byte, len(origData))

	blockMode.CryptBlocks(code, origData)

	return
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
