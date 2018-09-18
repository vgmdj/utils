package httplib

import "testing"

func TestPostJSON(t *testing.T) {
	var result string
	err := PostJSON("http://www.baidu.com",nil,&result,
		map[string]string{ResponseResultContentType:ContentTypeDefault})
	if err != nil{
		t.Error(err)
		return
	}

	t.Log(result)


}
