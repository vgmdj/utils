package httplib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/vgmdj/utils/logger"
)

const (
	ContentTypeAppJson   = "application/json"
	ContentTypeTextPlain = "text/plain"
	ContentTypeAppXml    = "application/xml"
	ContentTypeTextXml   = "text/xml"
	ContentTypeDefault   = "decoder/default"
)

const (
	ResponseResultContentType = "Result-Parse-Content-Type-vgmdj"
)

//RespDecoder 用于处理请求的返回数据
type RespDecoder interface {
	Unmarshal(reader []byte, v interface{}) error
}

var decoders = map[string]RespDecoder{
	ContentTypeAppJson:   new(JsonDecoder),
	ContentTypeTextPlain: new(TextDecoder),
	ContentTypeAppXml:    new(XmlDecoder),
	ContentTypeTextXml:   new(XmlDecoder),
	ContentTypeDefault:   new(DefaultDecoder),
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

//DefaultDecoder 默认方式
type DefaultDecoder struct{}

//Unmarshal 默认解析处理方式
func (dd DefaultDecoder) Unmarshal(body []byte, v interface{}) error {
	return TextDecoder{}.Unmarshal(body, v)
}

//respParser 对返回的body的处理
func respParser(body io.Reader, contentTypes string, respInfo interface{}) (err error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		logger.Error("resp body read err ")
		return err
	}

	if data == nil {
		logger.Info("no body data")
		return
	}

	contentType := strings.Split(contentTypes, ";")[0]

	decoder, ok := decoders[contentType]
	if !ok {
		logger.Error("unexpected content type ", contentType)
		logger.Error("data : ", string(data))

		return fmt.Errorf("Cannot decode request for %s data ", data)
	}

	if err = decoder.Unmarshal(data, respInfo); err != nil {
		logger.Error("err info ", string(data))
		return
	}

	return

}
