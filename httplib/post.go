package httplib

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/vgmdj/utils/logger"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//PostJSON http method post, content type application/json
func PostJSON(url string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
		temp    []byte
	)

	if temp, err = json.Marshal(body); err != nil {
		logger.Error("request: ", body)
		return
	}

	if request, err = http.NewRequest("POST",
		url, bytes.NewReader(temp)); err != nil {
		logger.Error("url ", url)
		return
	}

	if headers == nil {
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	if resp, err = client.Do(request); err != nil {
		logger.Error("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	return respParser(resp.Body, contentType, respInfo)
}

//PostXML http method post , content type application/xml
func PostXML(url string, body interface{}, respInfo interface{}) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
		temp    []byte
	)

	if temp, err = xml.Marshal(body); err != nil {
		logger.Error("request: ", body)
		return
	}

	if request, err = http.NewRequest("POST",
		url, bytes.NewReader(temp)); err != nil {
		logger.Error("url ", url)
		return
	}

	request.Header.Set("Content-Type", "application/xml;charset=UTF-8")

	if resp, err = client.Do(request); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	return respParser(resp.Body, contentType, respInfo)
}

//PostForm http method post , content type x-www-form-urlencoded
func PostForm(postUrl string, respInfo interface{}, formValues map[string]string, headers map[string]string) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
	)

	values := make(url.Values)
	for k, v := range formValues {
		values[k] = []string{v}
	}

	if request, err = http.NewRequest("POST",
		postUrl, strings.NewReader(values.Encode())); err != nil {
		log.Println("url ", postUrl)
		return
	}

	if headers == nil {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	if resp, err = client.Do(request); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	return respParser(resp.Body, contentType, respInfo)
}
