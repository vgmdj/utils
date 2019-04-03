package httplib

import "testing"

func TestClient_Get(t *testing.T) {
	c := NewClient()
	err := c.Get("http://www.baidu.com",
		map[string]interface{}{
			"q": "你是谁",
			"a": "你大爷",
		}, nil, nil)

	if err != nil {
		t.Error(err.Error())
		return
	}

}
