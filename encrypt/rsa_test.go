package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsaEncrypt(t *testing.T) {
	ast := assert.New(t)

	var privateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALPLWPASm/x0ZLAl
4XHWM3K8p6EmJ67MoL7X2srkHwTbJ/VHoXdjFKw0DGT5D3pWTqHokmJFjFyvGb7G
BFZ8NwLZ405L0rD/CZFYylJfg8k+Hi0BW0CobjPgASnXCY+bmkUyWdnxdDDOHs3J
kH+okoUyy+jIyiFwzG1j4mq9mWZ7AgMBAAECgYEAnHkAsg7ACnoRluugxL2ykMx2
5tyZ9JrJ2s1o8OKPzF4e7GymrYxhVW0GzGmlesbaMDaED1qPyanqMgmLhOkdxbsS
R7feu2mCX1NttoNsnkPOnvO+ercJsa5gI0YcKhsNIpcJ9sAoM3C/AjtTTD2vZIhs
iyS2Cdu52aX/InYo5kECQQDmZVKjeniVAJqRyTkLnLP8H16v49SabvLw1RbwPTQ8
ND7bEYeAT4Vux3PwVIEsYADc/sEbBYhFMMroCvbad0DZAkEAx8ZuXwwQhTDVyd9I
nMMdTJDmVFnAcpbQgjjCdZ8YUc+jZeTsS2jXSFmCjopmN2s46bvnT0FK1Lte1dt4
4/ZNcwJAOMvpn1tltnW7pQzR/0bWJ+Uj1oB3vMp1IWGmkfrEkcLfa+naWYtA/Zo1
vp1WarYQAGrc9+hZO5VXr/Rj/l8/oQJBAL1gn9Q+LZL1HlUF82GXnLiuS4n+ou59
hR9NCxpRPM6hFPZMsqsxsZMGNztEe21hmUwJMlbxQCy1iksUiF8hZ30CQQCw/DWY
Vv703cVD3HxI7T85x49xrgXsrMkV28mMdtZlgHIcEMK7KCYBBue7S7geSqNkNf/d
e8/AOABBMttS7noe
-----END PRIVATE KEY-----

`)

	var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCzy1jwEpv8dGSwJeFx1jNyvKeh
JieuzKC+19rK5B8E2yf1R6F3YxSsNAxk+Q96Vk6h6JJiRYxcrxm+xgRWfDcC2eNO
S9Kw/wmRWMpSX4PJPh4tAVtAqG4z4AEp1wmPm5pFMlnZ8XQwzh7NyZB/qJKFMsvo
yMohcMxtY+JqvZlmewIDAQAB
-----END PUBLIC KEY-----
`)

	cipherText, err := RsaEncrypt([]byte("this is rsa encrypt"), publicKey)
	if err != nil {
		t.Error(err.Error())
		return
	}

	plainText, err := RsaDecrypt(cipherText, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ast.Equal("this is rsa encrypt", string(plainText))

}
