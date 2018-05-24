package chars

import (
	"errors"
	"log"
)

//HideString 将给定的str，从offset位开始，一直替换limit位的repStr
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
