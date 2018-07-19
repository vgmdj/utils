package chars

import (
	"github.com/vgmdj/utils/logger"
	"math"
	"strconv"
)

func TakeLeftInt(num int, n int) int {
	for num >= int(math.Pow10(n)) {
		num /= 10
	}

	return num
}

//TODO finish this function
func TekeRightInt(num int, n int) int {

	return 0
}

func ToInt(num interface{}) int {
	switch num.(type) {
	default:
		logger.Error("invalid type ", num)
		return 0

	case string:
		result, _ := strconv.Atoi(num.(string))
		return result

	case int:
		return num.(int)

	case int32:
		return int(num.(int32))

	case int64:
		return int(num.(int64))

	case bool:
		if num.(bool) {
			return 1
		}
		return 0

	case []byte:
		result := 0
		for k, v := range num.([]byte) {
			result += (int(v) - 48) * int(math.Pow10(len(num.([]byte))-k-1))
		}
		return result

	}

}
