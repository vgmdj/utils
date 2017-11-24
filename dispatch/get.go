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

func Get(encode bool, baseURL string, respInfo interface{}, query map[string]string) (err error) {
	var (
		url  string
		resp *http.Response
		temp []byte
	)

	if len(query) == 0 {
		url = baseURL
	} else {
		url = jointURL(encode, baseURL, query)
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

func jointURL(encode bool, basePath string, query map[string]string) string {
	queryInfo := "?"

	for k, v := range query {
		queryInfo += fmt.Sprintf("%s=%s", k, v)
		queryInfo += "&"
	}

	queryInfo = strings.TrimRight(queryInfo, "&")

	if !encode {
		return basePath + queryInfo
	}

	return basePath + url.QueryEscape(queryInfo)

}
