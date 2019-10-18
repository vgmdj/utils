package httplib

import (
	"net/http"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	ba := &BasicAuth{
		UserName: "user",
		Pwd:      "pwd",
	}

	c := NewClient(&ClientConf{Auth: ba})

	req, err := c.NewRequest(http.MethodGet, "https://api.bilibili.com/x/server/now", nil, map[string]string{"test": "yes"})
	if err != nil {
		t.Error(err.Error())
		return
	}

	user, pwd, ok := req.BasicAuth()
	if !ok || user != ba.UserName || pwd != ba.Pwd {
		t.Errorf("expected basic auth user: %s, pwd: %s, but get %s, %s\n", ba.UserName, ba.Pwd, user, pwd)
		return
	}

	if len(req.Header) != 2 {
		t.Errorf("expected header length 2, but get %d\n", len(req.Header))
		t.Error("headers", req.Header)
		return
	}

	data, headers, err := c.DoWithData(req)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(data), headers)

}
