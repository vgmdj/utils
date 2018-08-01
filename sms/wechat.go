package sms

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
)

type WeChat struct {
	Ctl *core.Client
}

//NewWxClient
//params are same as SetConfig(params map[string]interface{})
func NewWxClient(params map[string]interface{}) *WeChat {
	wechat := &WeChat{}
	wechat.SetConfig(params)
	return wechat
}

//SetConfig 两种方式，任选其一
//- "appid", "secret"
//- "ctl"
func (wx *WeChat) SetConfig(params map[string]interface{}) {
	if ctl, ok := params["ctl"]; ok {
		if wx.Ctl, ok = ctl.(*core.Client); ok {
			return
		}
	}

	strParams := paramsToString(params)

	srv := core.NewDefaultAccessTokenServer(strParams["appid"], strParams["secret"], nil)
	wx.Ctl = core.NewClient(srv, nil)

}

type wxData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func (wx *WeChat) SendMsg(templateId string, to string, args ...string) (err error) {
	tm2 := template.TemplateMessage2{
		ToUser:     to,
		TemplateId: templateId,
		Data:       wx.setData(args[:]),
	}

	var msgID int64
	msgID, err = template.Send(wx.Ctl, tm2)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(msgID)

	return
}

func (wx *WeChat) setData(args []string) map[string]wxData {
	data := make(map[string]wxData)

	data["first"] = wxData{Value: args[0]}
	for i := 1; i < len(args)-1; i++ {
		key := fmt.Sprintf("keyword%d", i)
		data[key] = wxData{
			Value: args[i],
		}
	}
	data["remark"] = wxData{Value: args[len(args)-1]}

	logger.Info(data)

	return data

}
