package httplib

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/vgmdj/utils/chars"
)

var (
	hc   *Client
	once sync.Once
)

//Client httplib.Client
type Client struct {
	httpCli *http.Client
}

//NewClient return the only client
func NewClient() *Client {
	once.Do(func() {
		hc = &Client{
			httpCli: http.DefaultClient,
		}
	})

	return hc
}

//SetTimeout set cli.Timeout
func (c *Client) SetTimeout(duration time.Duration) {
	c.httpCli.Timeout = duration
}

//SetTimeout set Client.cli
func (c *Client) SetHttpClient(cli *http.Client) {
	c.httpCli = cli
}

//QueryEncode query url encode
func (c *Client) QueryEncode(query map[string]interface{}) string {
	values := url.Values{}

	for k, v := range query {
		values.Add(k, chars.ToString(v))
	}

	return values.Encode()

}
