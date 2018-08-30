package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

//SaltPwd 加盐加密
func SaltPwd(pwd string, salt string) string {
	for i := 0; i < 10; i++ {
		if k, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32); err == nil {
			m5 := md5.New()
			m5.Write(k)
			return hex.EncodeToString(m5.Sum(nil))
		}
	}
	return salt + pwd
}

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}
