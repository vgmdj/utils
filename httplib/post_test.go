package httplib

import (
	"testing"
)

func TestPostJSON(t *testing.T) {
	sn := ServerNow{}
	c := UniqueClient(nil)
	err := c.PostBytes("http://api.baidu.com/now", nil, &sn, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if sn.Code != -405 {
		t.Errorf("expected code -405 , but get %d", sn.Code)
		return
	}

}
