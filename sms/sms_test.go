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

	err := rlClient.SendMsg("1", "189xxxxxxxx", "010000", "5")
	if err != nil {
		t.Error(err)
		return
	}

}

func TestWeChat(t *testing.T) {
	params := make(map[string]interface{})
	params["appid"] = "wx6ddf008341937de1"
	params["secret"] = "3aefe696e17fa29ca0e1ad14c8ec36ee"

	wxsms := SMSFactory{SMS_WECHAT}.NewSMSClient(params)
	err := wxsms.SendMsg("FD-7gHRWUEj-fXOBRQh07uW9f1uSAdZF0Y8D-_YkZZo", "omHgCwm_DFWghRVayaJ35AggBLG8",
		"first", "k1", "k2", "k3", "k4", "remark")
	if err != nil {
		t.Error(err)
		return
	}

}
