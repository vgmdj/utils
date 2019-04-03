package httplib

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/vgmdj/utils/chars"

	"github.com/vgmdj/utils/logger"
)

var (
	hc   *Client
	once sync.Once
)

type Client struct {
	timeout time.Time
	cli     *http.Client
}

func NewClient() *Client {
	once.Do(func() {
		hc = &Client{
			cli: http.DefaultClient,
		}
	})

	return hc
}

func (c *Client) Get(host string, query map[string]interface{}, respInfo interface{}, headers map[string]string) (err error) {
	reqURL := fmt.Sprintf("%s?%s", host, c.QueryEncode(query))

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	resp, err := c.cli.Do(req)
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

func (c *Client) QueryEncode(query map[string]interface{}) string {
	values := url.Values{}

	for k, v := range query {
		values.Add(k, chars.ToString(v))
	}

	return values.Encode()

}
