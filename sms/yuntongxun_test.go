package sms

import "testing"

func TestRLSendSM(t *testing.T) {
	params := make(map[string]string)
	params["serverIP"] = "app.cloopen.com"
	params["serverPort"] = "8883"
	params["account"] = "accountSid"
	params["token"] = "accountToken"
	params["appId"] = "appId"

	rlClient, err := SMSFactory{SMS_YUNTONGXUN}.NewSMSClient(params)
	if err != nil {
		t.Error(err)
		return
	}

	err = rlClient.SendSM("1", "189xxxxxxxx", "010000", "5")
	if err != nil {
		t.Error(err)
		return
	}

}
