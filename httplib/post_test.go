package httplib

import (
	"testing"
)


/*
{
api: "mtop.common.getTimestamp",
v: "*",
ret: [
"SUCCESS::接口调用成功"
],
data: {
t: "1571388909183"
}
}

*/


func TestPostJSON(t *testing.T) {
	c := UniqueClient(nil)
	result := make(map[string]interface{})
	err := c.PostBytes("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp", nil, &result, nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(result)

}
