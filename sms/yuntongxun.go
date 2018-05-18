package sms

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/vgmdj/utils/httplib"
	"net/url"
	"time"
)

type (
	rongLianSM struct {
		RestURL     *url.URL
		Account     string
		Token       string
		AppId       string
		SoftVersion string
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
	rlConfig = []string{"serverIP", "serverPort", "account", "token", "appId"}
)

func newRlClient(params map[string]string) (SMSClient, error) {
	checkRlParams(params)

	u := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%s", params["serverIP"], params["serverPort"]),
		Path:   fmt.Sprintf("/%s/Accounts/%s/SMS/TemplateSMS", APIVISION, params["account"]),
	}

	return &rongLianSM{
		RestURL: u,
		Account: params["account"],
		Token:   params["token"],
		AppId:   params["appId"],
	}, nil
}

func (client *rongLianSM) SendSM(templateId string, to string, args ...string) (err error) {
	sig, auth := client.sigParamater()

	values := url.Values{}
	values.Add("sig", sig)
	client.RestURL.RawQuery = values.Encode()

	headers := make(map[string]string)
	headers["Authorization"] = auth
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["Accept"] = "application/json"

	body := rlSMRequest{AppId: client.AppId, TemplateId: templateId, To: to, Datas: args}
	resp := rlSMResponse{}

	err = httplib.PostJSON(client.RestURL.String(), body, &resp, headers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != "000000" {
		return errors.New(resp.StatusMsg)
	}

	return
}

func (client *rongLianSM) sigParamater() (string, string) {
	date := time.Now()
	sig := getMd5String([]byte(fmt.Sprintf("%s%s%s", client.Account, client.Token, date.Format("20060102150405"))))
	auth := getBase64String([]byte(fmt.Sprintf("%s:%s", client.Account, date.Format("20060102150405"))))
	return sig, auth
}

func checkRlParams(params map[string]string) error {
	for _, v := range rlConfig {
		if _, ok := params[v]; !ok {
			return fmt.Errorf("config %s no present", v)
		}
	}

	return nil

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
