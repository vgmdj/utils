package dispatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Get(encode bool, baseURL string, respInfo interface{}, query Query) (err error) {
	var (
		url  string
		resp *http.Response
		temp []byte
	)

	if len(query.QMap) == 0 {
		url = baseURL
	} else {
		url = JointURL(encode, baseURL, query)
	}

	if resp, err = http.Get(url); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	if temp, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println("resp body read err ")
		return
	}

	if err = json.Unmarshal(temp, respInfo); err != nil {
		log.Println("resp status code ", resp.StatusCode)
		log.Println("err info ", string(temp))
		return
	}

	return
}

func JointURL(encode bool, basePath string, query Query) string {
	queryInfo := "?"

	if query.IsOrdered && len(query.QKeys) != 0 {
		for _, v := range query.QKeys {
			queryInfo += fmt.Sprintf("%s=%s", v, query.QMap[v])
			queryInfo += "&"
		}
	} else {
		for k, v := range query.QMap {
			queryInfo += fmt.Sprintf("%s=%s", k, v)
			queryInfo += "&"
		}
	}

	queryInfo = strings.TrimRight(queryInfo, "&")

	if !encode {
		return basePath + queryInfo
	}

	return basePath + url.QueryEscape(queryInfo)

}
