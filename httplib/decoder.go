package httplib

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/vgmdj/utils/logger"
)

//MIME
const (
	MIMEJSON              = "application/json"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEHTML              = "text/html"
)

//ResponseResultContentType 指定返回数据解析方式
const (
	ResponseResultContentType = "Result-Parse-Content-Type-vgmdj"
)

//RespDecoder 用于处理请求的返回数据
type RespDecoder interface {
	Unmarshal(reader []byte, v interface{}) error
}

var decoders = map[string]RespDecoder{
	MIMEJSON:  new(JsonDecoder),
	MIMEPlain: new(TextDecoder),
	MIMEXML:   new(XmlDecoder),
	MIMEXML2:  new(XmlDecoder),

	//MIMEPOSTForm:          new(PostFormDecoder),
	//MIMEMultipartPOSTForm: new(MultipartFormDecoder),
	//MIMEHTML:              new(HtmlDecoder),
}

//JsonDecoder json处理类
type JsonDecoder struct{}

//Unmarshal json解析处理
func (jd JsonDecoder) Unmarshal(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

//XmlDecoder xml处理类
type XmlDecoder struct{}

//Unmarshal xml解析处理
func (xd XmlDecoder) Unmarshal(body []byte, v interface{}) error {
	return xml.Unmarshal(body, v)
}

//TextDecoder text处理类
type TextDecoder struct{}

//Unmarshal text解析处理
func (td TextDecoder) Unmarshal(body []byte, v interface{}) error {
	p, ok := v.(*string)
	if !ok {
		return fmt.Errorf("input not a  string pointer")
	}

	*p = string(body)
	return nil
}

//TODO
//HtmlDecoder html处理类
type HtmlDecoder struct{}

//Unmarshal html解析处理
func (hd *HtmlDecoder) Unmarshal(body []byte, v interface{}) error {
	return nil
}

//TODO
//PostFormDecoder form-urlencoded 处理类
type PostFormDecoder struct{}

//Unmarshal post form 解析处理
func (pfd *PostFormDecoder) Unmarshal(body []byte, v interface{}) error {
	return nil
}

//TODO
//MultipartFormDecoder multipart form data 处理类
type MultipartFormDecoder struct{}

//Unmarshal multipart form data 解析处理
func (mfd *MultipartFormDecoder) Unmarshal(body []byte, v interface{}) error {
	return nil
}

//respParser 对返回的body的处理
func respParser(body []byte, contentTypes string, v interface{}) (err error) {
	if len(body) == 0 {
		logger.Info("no body data")
		return
	}

	contentType := strings.Split(contentTypes, ";")[0]

	decoder, ok := decoders[contentType]
	if !ok {
		logger.Error("unexpected content type ,you can use ResponseResultContentType in headers to specified the decode way")
		logger.Error("data : ", string(body))

		return fmt.Errorf("Cannot decode request by content-type %s ", contentType)
	}

	if err = decoder.Unmarshal(body, v); err != nil {
		logger.Error("err info ", string(body))
		return
	}

	return

}

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
