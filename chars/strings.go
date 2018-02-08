package chars

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//将给定的str，从offset位开始，一直替换limit位的repStr
func HideString(str string, offset int, limit int, repStr string) (result string, err error) {
	var (
		temp     []byte
		strlen   = len(str)
		endPoint = offset + limit - 1
		repByte  byte
	)

	if strlen < 2 || strlen-1 < endPoint || len(repStr) != 1 {
		err = errors.New("字符长度不合法或函数使用错误")
		log.Println(str)
		return
	}

	repByte = repStr[0]
	for i := 0; i < strlen; i++ {
		if i < offset || i > endPoint {
			temp = append(temp, str[i])
			continue
		}
		temp = append(temp, repByte)
	}
	result = string(temp)

	return
}

func IsIntContain(char int, strArr []int) bool {
	if len(strArr) == 0 {
		return false
	}

	for _, v := range strArr {
		if char == v {
			return true
		}
	}

	return false

}

func IsStringContain(char string, strArr []string) bool {
	if len(strArr) == 0 {
		return false
	}

	for _, v := range strArr {
		if char == v {
			return true
		}
	}

	return false

}

func DBBytesToInt(b []byte) int {
	temp, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println(err.Error(), string(b))
		return 0
	}

	return temp
}

func DBBytesToBool(b []byte) bool {
	if string(b) == "1" {
		return true
	}

	return false
}

func ConvertArrayMapBytesToString(dbMaps []map[string][]byte) (maps []map[string]string) {
	maps = make([]map[string]string, 0)
	for _, dbmap := range dbMaps {
		strMap := make(map[string]string)
		for k, v := range dbmap {
			strMap[k] = string(v)
		}
		maps = append(maps, strMap)
	}

	return

}

func ConvertMapBytesToString(dbMaps map[string][]byte) (maps map[string]string) {
	maps = make(map[string]string)

	for k, v := range dbMaps {
		maps[k] = string(v)
	}

	return

}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func PrepareData(strQuery string, strRep ...string) (data string) {
	strRep = append(strRep, " ")

	query := strings.Split(strQuery, "?")
	lenQuery := len(query)
	lenReq := len(strRep)
	if lenQuery == 0 || lenReq == 0 {
		return strQuery
	}

	if lenQuery != lenReq {
		log.Println("数量对应错误", lenQuery, lenReq)
		return strQuery
	}

	for k, v := range query {
		data += v
		data += strRep[k]
	}

	return
}

func CheckInt(strs ...string) bool {
	for _, v := range strs {
		match, err := regexp.Match(`^[0-9]*$`, []byte(v))
		if err != nil || !match {
			log.Println(err, match, v)
			return false
		}
	}

	return true
}
