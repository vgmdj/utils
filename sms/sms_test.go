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

/*
用于测试
{{first.DATA}} 订单商品：{{keyword1.DATA}} 订单编号：{{keyword2.DATA}} 支付金额：{{keyword3.DATA}} 支付时间：{{keyword4.DATA}} {{remark.DATA}}
*/
func TestWeChat(t *testing.T) {
	params := make(map[string]interface{})
	params["appid"] = "wx6ddf008341937de1"
	params["secret"] = "3aefe696e17fa29ca0e1ad14c8ec36ee"

	wxsms := SMSFactory{SMS_WECHAT}.NewSMSClient(params)
	template := Template{
		TemplateId: "FD-7gHRWUEj-fXOBRQh07uW9f1uSAdZF0Y8D-_YkZZo",
		Keys:       []string{"first", "keyword1", "keyword2", "keyword3", "keyword4", "remark"},
		Color: map[string]string{
			"first":    "#0000CD",
			"keyword1": "#0000FF",
			"keyword2": "#00BFFF",
			"keyword3": "#4682B4",
			"keyword4": "#1E90FF",
			"remark":   "#191970",
		},
		URL: "http://www.google.com",
	}

	err := wxsms.SendMsgWithTemplate(template, "omHgCwm_DFWghRVayaJ35AggBLG8",
		"first'", "k1'", "k2'", "k3'", "k4'", "remark'")
	if err != nil {
		t.Error(err)
		return
	}

	wxsms.SetDefaultTemplate(template)
	err = wxsms.SendMsg("omHgCwm_DFWghRVayaJ35AggBLG8",
		"first", "k1", "k2", "k3", "k4", "remark")
	if err != nil {
		t.Error(err)
		return
	}

}

/*
充值结果通知
{{first.DATA}} 平台：{{keyword1.DATA}} 订单编号：{{keyword2.DATA}} 充值金额：{{keyword3.DATA}} 时间：{{keyword4.DATA}} {{remark.DATA}}
*/
func TestWeChat1(t *testing.T) {
	params := make(map[string]interface{})
	params["appid"] = "wx6ddf008341937de1"
	params["secret"] = "3aefe696e17fa29ca0e1ad14c8ec36ee"

	wxsms := SMSFactory{SMS_WECHAT}.NewSMSClient(params)
	template := Template{
		TemplateId: "1vvCKguqKO_gQ3m--fQxGeLK629x_W9RZez2ZhJPTLE",
		Keys:       []string{"first", "keyword1", "keyword2", "keyword3", "keyword4", "remark"},
		Color: map[string]string{
			"first":    "#0000CD",
			"keyword1": "#0000FF",
			"keyword2": "#00BFFF",
			"keyword3": "#4682B4",
			"keyword4": "#1E90FF",
			"remark":   "#191970",
		},
	}

	err := wxsms.SendMsgWithTemplate(template, "omHgCwm_DFWghRVayaJ35AggBLG8",
		"加油卡*123456充值成功", "85生活", "20181314151617", "1000元", "2014-09-22 08:10", "感谢您的使用。")
	if err != nil {
		t.Error(err)
		return
	}

}

/*
购买成功通知
{{first.DATA}} 商品名称：{{product.DATA}} 商品价格：{{price.DATA}} 购买时间：{{time.DATA}} {{remark.DATA}}
*/
func TestWeChat2(t *testing.T) {
	params := make(map[string]interface{})
	params["appid"] = "wx6ddf008341937de1"
	params["secret"] = "3aefe696e17fa29ca0e1ad14c8ec36ee"

	wxsms := SMSFactory{SMS_WECHAT}.NewSMSClient(params)
	template := Template{
		TemplateId: "RoFyLz_xtqtx4r6xNS1LZuPzPTdKEmIy9Oh0v2xAdTM",
		Keys:       []string{"first", "product", "price", "time", "remark"},
		Color: map[string]string{
			"first":  "#191970",
			"remark": "#191970",
		},
	}

	err := wxsms.SendMsgWithTemplate(template, "omHgCwm_DFWghRVayaJ35AggBLG8",
		"订单20181314151617创建成功", "加油卡充值", "1000元", "2014-09-22 08:10", "请等待充值结果。")
	if err != nil {
		t.Error(err)
		return
	}

}
