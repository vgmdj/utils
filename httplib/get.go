package httplib

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//Get http method get
func Get(encode bool, baseURL string, respInfo interface{}, query ...map[string]string) (err error) {
	var (
		reqURL string
		resp   *http.Response
	)

	reqURL = JointURL(encode, baseURL, query[0])

	if resp, err = http.Get(reqURL); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	return respParser(resp.Body, contentType, respInfo)
}

//JointURL arrange the query and return the complete url
func JointURL(encode bool, basePath string, query map[string]string) string {
	queryInfo := "?"

	for k, v := range query {
		values := v
		if encode {
			values = url.QueryEscape(values)
		}

		queryInfo += fmt.Sprintf("%s=%s", k, values)
		queryInfo += "&"
	}

	queryInfo = strings.TrimRight(queryInfo, "&")

	return basePath + queryInfo

}
