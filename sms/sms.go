package sms

import "github.com/vgmdj/utils/logger"

type SMSClient interface {
	SetConfig(map[string]interface{})
	SetDefaultTemplate(Template)
	SendMsg(string, ...string) error
	SendMsgWithTemplate(Template, string, ...string) error
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

//NewSMSClient
func (sf SMSFactory) NewSMSClient(params map[string]interface{}) SMSClient {
	switch sf.SMS {
	default:
		logger.Error("invalid sms type")
		return nil

	case SMS_YUNTONGXUN:
		return NewRlClient(params)

	case SMS_WECHAT:
		return NewWxClient(params)

	case SMS_ALIYUN:
		return nil

	case SMS_TENCENTCLOUD:
		return nil
	}
}

func paramsToString(params map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for k, v := range params {
		value, ok := v.(string)
		if ok {
			result[k] = value
		}
	}

	return result
}

type Template struct {
	TemplateId string
	Keys       []string
	Color      map[string]string
	URL        string
}
