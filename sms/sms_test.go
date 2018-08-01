package sms

import "testing"

func TestRLYun(t *testing.T) {
	params := make(map[string]interface{})
	params["serverIP"] = "app.cloopen.com"
	params["serverPort"] = "8883"
	params["account"] = "accountSid"
	params["token"] = "accountToken"
	params["appId"] = "appId"

	rlClient := SMSFactory{SMS_YUNTONGXUN}.NewSMSClient(params)

	err := rlClient.SendMsg("1", "189xxxxxxxx", []string{"010000", "5"})
	if err != nil {
		t.Error(err)
		return
	}

}

func TestWeChat(t *testing.T) {
	params := make(map[string]interface{})
	params[""] = ""

	wxsms := SMSFactory{SMS_WECHAT}.NewSMSClient(params)
	err := wxsms.SendMsg()
	if err != nil {
		t.Error(err)
		return
	}

}
