package encrypt

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/vgmdj/utils/chars"
)

var (
	defaultIV = []byte("0123456789012345")
)

type CryptoCode struct {
	baseLen     int
	extendedLen int
	key         string
}

type RedemptionCode struct {
	BaseCode  string
	FinalCode string
}

func NewClient(key string, baseLen int, extendedLen ...int) *CryptoCode {
	extLen := baseLen
	if len(extendedLen) != 0 {
		extLen = extendedLen[0]
	}

	if err := checkFormat(key, baseLen, extLen); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &CryptoCode{
		baseLen:     baseLen,
		extendedLen: extLen,
		key:         key,
	}
}

func (rc *CryptoCode) Produce(nums ...int) (codes []RedemptionCode, err error) {
	if rc == nil {
		return nil, fmt.Errorf("invalid client")
	}

	count := 1
	if len(nums) != 0 {
		count = nums[0]
	}

	for i := 0; i < count; i++ {
		base := chars.RandomAlphanumeric(uint(rc.baseLen))
		code, err := rc.produce(base)
		if err != nil {
			return nil, err
		}

		codes = append(codes, code)
	}

	return
}

func (rc *CryptoCode) CheckCode(codes ...string) (err error) {
	if rc == nil {
		return fmt.Errorf("invalid client")
	}

	for _, v := range codes {
		if len(v) != rc.baseLen+rc.extendedLen {
			return fmt.Errorf("invalid code,code length should be %d", rc.baseLen+rc.extendedLen)
		}

		base := v[:rc.baseLen]
		code, err := rc.produce(base)
		if err != nil {
			return err
		}

		if code.FinalCode != v {
			return fmt.Errorf("invalid code %s", v)
		}

	}

	return
}

func (rc *CryptoCode) produce(base string) (code RedemptionCode, err error) {

	btsc, err := AesCBCEncrypt([]byte(base), []byte(rc.key), defaultIV)
	if err != nil {
		return code, err
	}

	encrpt := base64.StdEncoding.EncodeToString(btsc)[:rc.extendedLen]
	encrpt = strings.Replace(encrpt, "/", "a", -1)
	encrpt = strings.Replace(encrpt, "+", "A", -1)
	code = RedemptionCode{
		BaseCode:  base,
		FinalCode: base + encrpt,
	}

	return
}

func checkFormat(key string, baseLen int, extendLen int) (err error) {
	if len(key) != 16 {
		return fmt.Errorf("invalid  encrypt key")
	}

	maxLen := ((baseLen >> 4) + 1) << 4
	if extendLen <= 0 || extendLen > maxLen {
		return fmt.Errorf("extended length must bigger than 0")
	}

	return
}
