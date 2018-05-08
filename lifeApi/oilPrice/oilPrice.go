package oilPrice

import (
	"github.com/vgmdj/utils/dispatch"
	"log"
)

const (
	Baseurl    = "http://apis.eolinker.com/common/oil/getOilPriceToday"
	ProductKey = ""
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
	res := new(ResOilPrice)
	values := make(map[string]string)
	values["productKey"] = ProductKey
	err := dispatch.PostForm(Baseurl, res, values, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil

}
