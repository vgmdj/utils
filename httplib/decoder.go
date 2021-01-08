package httplib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
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
	ContentType               = "Content-Type"
	ResponseResultContentType = "Result-Parse-Content-Type-vgmdj"
	ResponseResultAllowNull   = "Result-Parse-Allow-Null"
)

//RespDecoder 用于处理请求的返回数据
type RespDecoder interface {
	Unmarshal(reader []byte, v interface{}) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

// parser 用于解析返回内容的解析器
type parser struct {
	decoder       RespDecoder
	respAllowNull bool
}

func newParser(config map[string]string) (*parser, error) {
	p := new(parser)

	contentType := config[ContentType]
	if ct, ok := config[ResponseResultContentType]; ok && ct != "" {
		contentType = ct
	}

	if contentType == "" {
		contentType = MIMEPlain
	}

	contentType = strings.Split(contentType, ";")[0]
	decoder, ok := decoders[contentType]
	if !ok {
		return nil, fmt.Errorf("Cannot decode request by content-type %s ", contentType)
	}
	p.decoder = decoder

	p.respAllowNull = true
	if config[ResponseResultAllowNull] == "false" {
		p.respAllowNull = false
	}

	return p, nil

}

//respParser 对返回的body的处理
func (p *parser) respParser(body []byte, v interface{}) (err error) {
	if len(body) == 0 && !p.respAllowNull {
		return fmt.Errorf("body not allow null , but get nothing")
	}

	if len(body) == 0 {
		return
	}

	if err = p.decoder.Unmarshal(body, v); err != nil {
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
