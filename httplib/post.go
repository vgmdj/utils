package httplib

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/vgmdj/utils/logger"
)

//PostJSON http method post, content type application/json
func (c *Client) PostJSON(postUrl string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
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

	return c.Raw(http.MethodPost, postUrl, values, respInfo, headers)
}

//PostXML http method post , content type application/xml
func (c *Client) PostXML(postUrl string, body interface{}, respInfo interface{}, headers map[string]string) (err error) {
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

	return c.Raw(http.MethodPost, postUrl, values, respInfo, headers)
}

//PostForm http method post , content type x-www-form-urlencoded
func (c *Client) PostForm(postUrl string, formValues map[string]string, respInfo interface{}, headers map[string]string) (err error) {
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

	return c.Raw(http.MethodPost, postUrl, []byte(values.Encode()), respInfo, headers)
}

//PostBytes http method post,
func (c *Client) PostBytes(postUrl string, bytes []byte, respInfo interface{}, headers map[string]string) (err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "text/plain;charset=UTF-8"
	}

	return c.Raw(http.MethodPost, postUrl, bytes, respInfo, headers)
}
