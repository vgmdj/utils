package dispatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Get(encode bool, baseURL string, respInfo interface{}, query ...Query) (err error) {
	var (
		getUrl string
		resp   *http.Response
		temp   []byte
	)

	if len(query) == 0 || len(query[0].QMap) == 0 {
		getUrl = baseURL
	} else {
		getUrl = JointURL(encode, baseURL, query[0])
	}

	if resp, err = http.Get(getUrl); err != nil {
		log.Println("发送请求错误")
		return
	}
	defer resp.Body.Close()

	if temp, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println("resp body read err ")
		return
	}

	if err = json.Unmarshal(temp, respInfo); err != nil {
		log.Println("resp status code ", resp.StatusCode)
		log.Println("err info ", string(temp))
		return
	}

	return
}

func JointURL(encode bool, basePath string, query Query) string {
	queryInfo := "?"

	if query.IsOrdered && len(query.QKeys) != 0 {
		for _, v := range query.QKeys {
			values := query.QMap[v]
			if encode {
				values = url.QueryEscape(values)
			}

			queryInfo += fmt.Sprintf("%s=%s", v, values)
			queryInfo += "&"
		}
	} else {
		for k, v := range query.QMap {
			values := v
			if encode {
				values = url.QueryEscape(values)
			}

			queryInfo += fmt.Sprintf("%s=%s", k, values)
			queryInfo += "&"
		}
	}

	queryInfo = strings.TrimRight(queryInfo, "&")

	return basePath + queryInfo

}
