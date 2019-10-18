package httplib

import (
	"testing"
	"time"
)

func TestClientGet(t *testing.T) {
	sn := ServerNow{}
	c := UniqueClient(nil)
	err := c.Get("http://api.bilibili.com/x/server/now", nil, &sn, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	now := time.Now()
	if now.Unix()-sn.Data.Now > 30 {
		t.Errorf("client get test err , server time:%d, now :%d\n", sn.Data.Now, now.Unix())
		return
	}

}

func TestClientGetNull(t *testing.T) {
	c := UniqueClient(nil)
	err := c.Get("http://localhost:9090/test", nil, nil, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

}

func TestClientGetSearch(t *testing.T) {
	html := ""
	c := NewClient(nil)
	err := c.Get(
		"https://cn.bing.com/search",
		map[string]interface{}{
			"q":  "who are you",
			"ie": "utf-8",
		},
		&html,
		map[string]string{
			ResponseResultContentType: MIMEPlain,
		},
	)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if html == "" {
		t.Errorf("parse data err\n")
		return
	}

}
