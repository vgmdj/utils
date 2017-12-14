package lifeApi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	Baseurl    = "http://apis.eolinker.com/common/oil/getOilPriceToday?"
	ProductKey = "9ktLPua9ea88ebd153a973dafedd2734073002e31d82e4f"
)

type ResOilPrice struct {
	StatusCode string `json:"statusCode"` //状态码
	Desc       string `json:"desc"`       //状态码说明
	Result     result `json:"result"`
}

type result struct {
	List    []list `json:"list"`
	RetCode int    `json:"ret_code"`
}
type list struct {
	Ct   string `json:"ct"`
	P0   string `json:"p0"`
	P89  string `json:"p89"`
	P90  string `json:"p90"`
	P92  string `json:"p92"`
	P93  string `json:"p93"`
	P95  string `json:"p95"`
	P97  string `json:"p97"`
	Prov string `json:"prov"`
}

func GetPriceToday(prov ...string) (*ResOilPrice, error) {
	values := make(url.Values)
	values.Add("productKey", ProductKey)
	if len(prov) != 0 {
		values.Add("province", prov[0])
	}
	postbytesReader := bytes.NewReader([]byte(values.Encode()))
	fmt.Println("GetOilPriceToday 链接:", Baseurl)
	req, err := http.NewRequest("POST", Baseurl, postbytesReader)
	if err != nil {
		log.Println("newRequest:=", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if resp, err := client.Do(req); err != nil {
		log.Println("Send request error is：", err)
		return nil, err
	} else {
		defer resp.Body.Close()
		if ss, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Println("err occured when ioutil to string, error: ", err)
			return nil, err
		} else {
			res := new(ResOilPrice)
			if err := json.Unmarshal(ss, res); err != nil {
				log.Println("err occured when unmarshal,error: ", err, ",res: ", string(ss))
				return nil, err
			}
			if res.StatusCode == "000000" {
				return res, nil
			} else {
				return nil, fmt.Errorf(res.Desc)
			}
		}

	}
}
