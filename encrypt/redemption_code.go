package encrypt

import (
	"fmt"

	"github.com/vgmdj/utils/chars"
)

type RedemptionCode struct {
	baseCode    string
	baseLen     int
	encryptCode string
	encryptLen  int
	key         string
	extended    string
}

func NewClient(baseCode, key string, encryptLen ...int) *RedemptionCode {
	encLen := len(baseCode)
	if len(encryptLen) != 0 {
		encLen = encryptLen[0]
	}

	return &RedemptionCode{
		baseCode:   baseCode,
		baseLen:    len(baseCode),
		encryptLen: encLen,
		key:        key,
		extended:   "*",
	}
}

func (rc *RedemptionCode) Produce(count int) (codes []string, err error) {
	if err = rc.checkFormat(); err != nil {
		return
	}

	return
}

func (rc *RedemptionCode) CheckCode() (err error) {
	if err = rc.checkFormat(); err != nil {
		return
	}

	return
}

func (rc *RedemptionCode) SetExtendedStr(str string) {
	rc.extended = str
}

func (rc *RedemptionCode) checkFormat() (err error) {
	if len(rc.baseCode) == 0 || len(rc.key) == 0 {
		return fmt.Errorf("invalid base code or encrypt key")
	}

	if rc.encryptLen < rc.baseLen {
		return fmt.Errorf("encrypt code length must longer than base code length")
	}

	if !chars.IsIntegerOrAlphabet(rc.baseCode) {
		return fmt.Errorf("base code can only contain 0-9 or a-z or A-Z")
	}

	return
}

func extendedCode(code string, length int) (result string) {
	if length <= len(code) {
		return code
	}

	for i := 0; i < length-len(code); i++ {

	}

}
