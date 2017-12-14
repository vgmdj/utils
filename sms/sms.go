package sms

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	SMS_GATEWAY  = "http://dx.ipyy.net/smsJson.aspx?"
	SMS_ACCOUNT  = "AB00031"
	SMS_PASSWORD = "Mink2501"
	SMS_COMPANY  = "榕树科技"
)

type SmsResponse struct {
	Returnstatus  string `json:"returnstatus"`
	Message       string `json:"message"`
	Remainpoint   string `json:"remainpoint"`
	SuccessCounts string `json:"successCounts"`
}

func SendMessage(mobile, content string) (bool, error) {
	values := make(url.Values)
	values.Add("action", "send")
	values.Add("userid", "")
	values.Add("password", SMS_PASSWORD)
	values.Add("account", SMS_ACCOUNT)
	values.Add("content", fmt.Sprintf("%s【%s】", content, SMS_COMPANY))
	values.Add("sendTime", "")
	values.Add("extno", "")
	values.Add("mobile", mobile)
	finalUrl := SMS_GATEWAY + values.Encode()
	fmt.Println("链接:", finalUrl)
	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		log.Println("newRequest:=", err)
		return false, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if resp, err := client.Do(req); err != nil {
		log.Println("Send request err  is：", err)
		return false, err
	} else {
		defer resp.Body.Close()
		if ss, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Println("err occured when ioutil to string :", mobile, ",err is", err)
			return false, err
		} else {
			res := new(SmsResponse)
			if err := json.Unmarshal(ss, res); err != nil {
				log.Println("err occured when send message to C:", mobile, ",err is", err, ",res is", string(ss))
				return false, err
			}
			if res.Returnstatus == "Success" {
				return true, nil
			} else {
				return false, fmt.Errorf(res.Message)
			}
		}

	}
}
