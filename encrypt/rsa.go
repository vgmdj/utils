package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

//RsaSign rsa签名
func RsaSign(plainText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	pInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	p, ok := pInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("should be rsa key, not dsa or ecdsa key ")
	}

	h := sha1.New()
	h.Write([]byte(plainText))
	hash := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, p, crypto.SHA1, hash[:])
}

//RsaCheckSign rsa解签，plainText是原始数据, cipherText签名后数据
func RsaCheckSign(plainText, cipherText, publicKey []byte) error {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return errors.New("should be rsa key, not dsa or ecdsa key ")
	}

	hash := sha1.New()
	hash.Write(cipherText)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash.Sum(nil), plainText)

}

//RsaEncrypt rsa加密
func RsaEncrypt(plainText, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("should be rsa key, not dsa or ecdsa key ")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

//RsaDecrypt rsa解密
func RsaDecrypt(cipherText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	pInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	p, ok := pInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("should be rsa key, not dsa or ecdsa key ")
	}

	return rsa.DecryptPKCS1v15(rand.Reader, p, cipherText)
}
