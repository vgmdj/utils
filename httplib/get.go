package httplib

import (
	"fmt"
	"net/http"

	"github.com/vgmdj/utils/logger"
)

//Get get method
func (c *Client) Get(host string, query map[string]interface{}, respInfo interface{}, headers map[string]string) (err error) {
	reqURL := fmt.Sprintf("%s?%s", host, c.QueryEncode(query))

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	resp, err := c.httpCli.Do(req)
	if err != nil {
		logger.Error("发送请求错误")
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if specified, ok := headers[ResponseResultContentType]; ok {
		contentType = specified
	}

	logger.Info(req.Method, req.URL)

	return respParser(resp.Body, contentType, respInfo)

}
