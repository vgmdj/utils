package sms

import (
	"fmt"
)

type SMSClient interface {
	SendSM(templateId string, to string, data ...string) error
}

type SelectSystem uint16

const (
	SMS_YUNTONGXUN SelectSystem = iota + 1
	SMS_ALIYUN
	SMS_TENCENTCLOUD
	SMS_WECHAT
)

type SMSFactory struct {
	SMS SelectSystem
}

func (sf SMSFactory) NewSMSClient(params map[string]string) (SMSClient, error) {
	switch sf.SMS {
	default:
		return nil, fmt.Errorf("invalid sms type")

	case SMS_YUNTONGXUN:
		return newRlClient(params)

	case SMS_ALIYUN:
		return nil, fmt.Errorf("not support now")

	case SMS_TENCENTCLOUD:
		return nil, fmt.Errorf("not support now")

	case SMS_WECHAT:
		return nil, fmt.Errorf("not support now")

	}
}
