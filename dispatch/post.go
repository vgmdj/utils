package dispatch

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func PostJSON(url string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
		temp    []byte
	)

	if temp, err = json.Marshal(body); err != nil {
		log.Println("request: ", body)
		return
	}

	if request, err = http.NewRequest("POST",
		url, bytes.NewReader(temp)); err != nil {
		log.Println("url ", url)
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

func PostXML(url string, body interface{}, respInfo interface{}) (err error) {
	var (
		client  http.Client
		request *http.Request
		resp    *http.Response
		temp    []byte
	)

	if temp, err = xml.Marshal(body); err != nil {
		log.Println("request: ", body)
		return
	}

	if request, err = http.NewRequest("POST",
		url, bytes.NewReader(temp)); err != nil {
		log.Println("url ", url)
		return
	}

	request.Header.Set("Content-Type", "application/xml;charset=UTF-8")

	if resp, err = client.Do(request); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	if temp, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println("resp body read err ")
		return
	}

	if err = xml.Unmarshal(temp, respInfo); err != nil {
		log.Println("resp status code ", resp.StatusCode)
		return
	}

	return
}
