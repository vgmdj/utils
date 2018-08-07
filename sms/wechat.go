package sms

import (
	"github.com/vgmdj/utils/logger"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	wxTemp "gopkg.in/chanxuehong/wechat.v2/mp/message/template"
	"sync"
)

var (
	wxc    *WeChat
	wxSync sync.Once
)

type WeChat struct {
	Ctl             *core.Client
	DefaultTemplate Template
}

//NewWxClient
//params are same as SetConfig(params map[string]interface{})
func NewWxClient(params map[string]interface{}) *WeChat {
	wxSync.Do(func() {
		wxc = &WeChat{}
	})

	if len(params) != 0 {
		wxc.SetConfig(params)
	}

	return wxc
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

func (wx *WeChat) SetDefaultTemplate(template Template) {
	wx.DefaultTemplate = template
}

type wxData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func (wx *WeChat) SendMsg(to string, args ...string) (err error) {
	tm2 := wxTemp.TemplateMessage2{
		ToUser:     to,
		TemplateId: wx.DefaultTemplate.TemplateId,
		Data:       wx.setData(args[:], wx.DefaultTemplate),
	}

	var msgID int64
	msgID, err = wxTemp.Send(wx.Ctl, tm2)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(msgID)

	return
}

func (wx *WeChat) SendMsgWithTemplate(template Template, to string, args ...string) (err error) {
	tm2 := wxTemp.TemplateMessage2{
		ToUser:     to,
		TemplateId: template.TemplateId,
		Data:       wx.setData(args[:], template),
	}

	var msgID int64
	msgID, err = wxTemp.Send(wx.Ctl, tm2)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(msgID)

	return
}

func (wx *WeChat) setData(args []string, template Template) map[string]wxData {
	data := make(map[string]wxData)

	if len(args) < len(template.Keys) {
		logger.Error("invalid input params")
		return data
	}

	for k, v := range template.Keys {
		data[v] = wxData{
			Value: args[k],
			Color: template.Color[v],
		}
	}

	logger.Info(data)

	return data

}
