package httplib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/vgmdj/utils/logger"
	"io"
	"io/ioutil"
	"strings"
)

const (
	contentTypeAppJson   = "application/json"
	contentTypeTextPlain = "text/plain"
	contentTypeAppXml    = "application/xml"
	contentTypeTextXml   = "text/xml"
	contentTypeDefault   = "decoder/default"
)

//RespDecoder 用于处理请求的返回数据
type RespDecoder interface {
	Unmarshal(reader []byte, v interface{}) error
}

var decoders = map[string]RespDecoder{
	contentTypeAppJson:   new(JsonDecoder),
	contentTypeTextPlain: new(JsonDecoder),
	contentTypeAppXml:    new(XmlDecoder),
	contentTypeTextXml:   new(XmlDecoder),
	contentTypeDefault:   new(DefaultDecoder),
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

//DefaultDecoder 默认方式
type DefaultDecoder struct{}

//Unmarshal 默认解析处理方式
func (dd DefaultDecoder) Unmarshal(body []byte, v interface{}) error {
	v = string(body)
	return nil
}

//respParser 对返回的body的处理
func respParser(body io.Reader, contentTypes string, respInfo interface{}) (err error) {
	var temp []byte
	if temp, err = ioutil.ReadAll(body); err != nil {
		logger.Error("resp body read err ")
		return
	}

	contentType := strings.Split(contentTypes, ";")[0]

	decoder, ok := decoders[contentType]
	if !ok {
		logger.Error("unexpected content type ", contentType)
		logger.Error("data : ", string(temp))

		return fmt.Errorf("Cannot decode request for %s data ", temp)
	}

	if err = decoder.Unmarshal(temp, respInfo); err != nil {
		logger.Error("err info ", string(temp))
		return
	}

	return

}
