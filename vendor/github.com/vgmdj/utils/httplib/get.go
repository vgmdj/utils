package httplib

import (
	"fmt"
	"github.com/vgmdj/utils/logger"
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

	reqURL = baseURL
	if len(query) != 0 {
		reqURL = JointURL(encode, baseURL, query[0])
	}

	logger.Info(reqURL)

	if resp, err = http.Get(reqURL); err != nil {
		logger.Error("发送请求错误")
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
