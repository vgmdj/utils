package httplib

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/vgmdj/utils/logger"
	"net/http"
	"net/url"
)

const (
	ResponseResultContentType = "Result-Parse-Content-Type-vgmdj"
)

//PostJSON http method post, content type application/json
func PostJSON(postUrl string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json;charset=UTF-8"
	}

	var values []byte
	if values, err = json.Marshal(body); err != nil {
		logger.Error("request: ", body)
		return
	}

	return post(postUrl, values, respInfo, headers)
}

//PostXML http method post , content type application/xml
func PostXML(postUrl string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/xml;charset=UTF-8"
	}

	var values []byte
	if values, err = xml.Marshal(body); err != nil {
		logger.Error("request: ", body)
		return
	}

	return post(postUrl, values, respInfo, headers)
}

//PostForm http method post , content type x-www-form-urlencoded
func PostForm(postUrl string, respInfo interface{}, formValues map[string]string, headers map[string]string) (err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8"
	}

	values := make(url.Values)
	for k, v := range formValues {
		values[k] = []string{v}
	}

	return post(postUrl, []byte(values.Encode()), respInfo, headers)
}

//post http post request
func post(url string, body []byte, respInfo interface{}, headers map[string]string) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
	)

	logger.Info(url, string(body), headers)

	if request, err = http.NewRequest("POST",
		url, bytes.NewReader(body)); err != nil {
		logger.Error("url ", url)
		return
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	if resp, err = client.Do(request); err != nil {
		logger.Error("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if specified, ok := headers[ResponseResultContentType]; ok {
		contentType = specified
	}

	return respParser(resp.Body, contentType, respInfo)
}
