package sms

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/vgmdj/utils/httplib"
	"github.com/vgmdj/utils/logger"
)

type (
	RLYun struct {
		RestURL     *url.URL
		Account     string
		Token       string
		AppId       string
		SoftVersion string

		DefaultTemplate Template
	}

	rlSMRequest struct {
		AppId      string   `json:"appId"`
		To         string   `json:"to"`
		TemplateId string   `json:"templateId"`
		Datas      []string `json:"datas"`
	}

	rlSMResponse struct {
		StatusCode string `json:"statusCode"`
		StatusMsg  string `json:"statusMsg"`
	}
)

const (
	APIVISION = "2013-12-26"
)

var (
	rlc    *RLYun
	rlSync sync.Once
)

//NewRlClient
//params are same as SetConfig(params map[string]interface{})
func NewRlClient(params map[string]interface{}) *RLYun {
	rlSync.Do(func() {
		rlc = &RLYun{}
	})

	if len(params) != 0 {
		rlc.SetConfig(params)
	}

	return rlc
}

//SetConfig "serverIP", "serverPort", "account", "token", "appId"
func (client *RLYun) SetConfig(params map[string]interface{}) {
	strParams := paramsToString(params)

	u := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%s", strParams["serverIP"], strParams["serverPort"]),
		Path:   fmt.Sprintf("/%s/Accounts/%s/SMS/TemplateSMS", APIVISION, strParams["account"]),
	}

	client.RestURL = u
	client.Account = strParams["account"]
	client.Token = strParams["token"]
	client.AppId = strParams["appId"]
}

func (client *RLYun) SetDefaultTemplate(template Template) {
	logger.Warning("don not need to set template")
}

func (client *RLYun) SendMsg(to string, args ...string) (err error) {

	sig, auth := client.sigParamater()

	values := url.Values{}
	values.Add("sig", sig)
	client.RestURL.RawQuery = values.Encode()

	headers := make(map[string]string)
	headers["Authorization"] = auth
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["Accept"] = "application/json"

	body := rlSMRequest{AppId: client.AppId, TemplateId: client.DefaultTemplate.TemplateId, To: to, Datas: args}
	resp := rlSMResponse{}

	err = httplib.NewClient(nil).PostJSON(client.RestURL.String(), body, &resp, headers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != "000000" {
		return errors.New(resp.StatusMsg)
	}

	return
}

func (client *RLYun) SendMsgWithTemplate(template Template, to string, args ...string) (err error) {
	sig, auth := client.sigParamater()

	values := url.Values{}
	values.Add("sig", sig)
	client.RestURL.RawQuery = values.Encode()

	headers := make(map[string]string)
	headers["Authorization"] = auth
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["Accept"] = "application/json"

	body := rlSMRequest{AppId: client.AppId, TemplateId: template.TemplateId, To: to, Datas: args}
	resp := rlSMResponse{}

	err = httplib.NewClient(nil).PostJSON(client.RestURL.String(), body, &resp, headers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != "000000" {
		return errors.New(resp.StatusMsg)
	}

	return
}

func (client *RLYun) sigParamater() (string, string) {
	date := time.Now()
	sig := getMd5String([]byte(fmt.Sprintf("%s%s%s", client.Account, client.Token, date.Format("20060102150405"))))
	auth := getBase64String([]byte(fmt.Sprintf("%s:%s", client.Account, date.Format("20060102150405"))))
	return sig, auth
}

func getMd5String(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func getBase64String(data []byte) string {
	h := base64.StdEncoding
	return h.EncodeToString(data)
}
