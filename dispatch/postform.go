package dispatch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func PostForm(postUrl string, respInfo interface{}, formValues map[string]string, headers map[string]string) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
		temp    []byte
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

	if temp, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println("resp body read err ")
		return
	}

	if err = json.Unmarshal(temp, respInfo); err != nil {
		log.Println("resp status code ", resp.StatusCode)
		return
	}

	return
}
