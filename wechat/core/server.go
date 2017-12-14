package core

import (
	"gitee.com/sailinst/utils"
	"sort"
	"strings"
)

type Signature struct {
	Signature string `json:"signature"`
	TimeStamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

func NewSignal(signalstr, echostr, timestamp, nonce string) (signal *Signature) {
	return &Signature{
		Signature: signalstr,
		Echostr:   echostr,
		TimeStamp: timestamp,
		Nonce:     nonce,
	}
}
func (signature *Signature) CheckSign(token string) (bool, string) {
	if signature != nil {
		var array []string
		array = append(array, token, signature.TimeStamp, signature.Nonce)
		sort.Strings(array)
		sign := utils.Sha1(strings.Join(array, ""))
		return strings.EqualFold(sign, signature.Signature), sign
	}
	return false, ""
}
