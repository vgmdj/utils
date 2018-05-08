package navigation

import (
	"strconv"
)

//Distance
//先判断是否在1000米以上，不足1000米，则显示米
//超过1000米，则显示千米，保留小数点后一位
func Distance(dis float64) string {
	var (
		km int
		m  int
	)
	km = int(dis / 1000)

	if km == 0 {
		m = int(dis)
		return strconv.Itoa(m) + "m"
	}

	return strconv.FormatFloat(dis/1000, 'f', 2, 32) + "km"

}
