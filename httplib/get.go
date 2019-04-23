package httplib

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/vgmdj/utils/chars"
)

//Get get method
func (c *Client) Get(uri string, query map[string]interface{}, respInfo interface{}, headers map[string]string) (err error) {
	if query == nil{
		return c.Raw(http.MethodGet,uri,nil,respInfo,nil)
	}

	return c.Raw(http.MethodGet, fmt.Sprintf("%s?%s", uri, c.QueryEncode(query)),
		nil, respInfo, headers)

}

//QueryEncode query url encode
func (c *Client) QueryEncode(query map[string]interface{}) string {
	values := url.Values{}

	for k, v := range query {
		values.Add(k, chars.ToString(v))
	}

	return values.Encode()

}
