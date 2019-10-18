package httplib

import (
	"net/http"
	"testing"
	"time"
)

type ServerNow struct {
	Code int64 `json:"code"`
	Data Data  `json:"data"`
}

type Data struct {
	Now int64 `json:"now"`
}

func TestClientRaw(t *testing.T) {
	sn := ServerNow{}

	c := NewClient(DefaultClientConf)
	err := c.Raw(http.MethodGet, "https://api.bilibili.com/x/server/now", nil, &sn, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	now := time.Now()
	if now.Unix()-sn.Data.Now > 30 {
		t.Errorf("client get test err , server time:%d, now :%d\n", sn.Data.Now, now.Unix())
		return
	}

	err = c.Raw(http.MethodPost, "https://api.bilibili.com/x/server/now", nil, &sn, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if sn.Code != -405 {
		t.Errorf("expected code -405, but got code %d", sn.Code)
		return
	}
}

func TestClient_DoWithData(t *testing.T) {
	c := UniqueClient(DefaultClientConf)
	request, _ := c.NewRequest(http.MethodGet, "https://api.bilibili.com/x/server/now", nil, nil)

	data, headers, _ := c.DoWithData(request)

	t.Log(string(data), headers)

}
