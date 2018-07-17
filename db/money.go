package db

import (
	"github.com/vgmdj/utils/logger"
	"strconv"
)

//使用数据库存价格等数据的时候，存在小数情况，一般放大转为整数存储

const (
	//Enlarge 放大倍数
	Enlarge = 100
)

//MoneyStrToEnlargeInt32  string转放大int32
func MoneyStrToEnlargeInt32(strm string) int {
	f32m, err := strconv.ParseFloat(strm, 32)
	if err != nil {
		logger.Error(err.Error())
	}

	int32m := f32m * Enlarge

	return int(int32m)
}

//MoneyStrToEnlargeInt64 string转放大后int64
func MoneyStrToEnlargeInt64(strm string) int64 {
	f32m, err := strconv.ParseFloat(strm, 32)
	if err != nil {
		logger.Error(err.Error())
	}

	ff32m := f32m * Enlarge

	return int64(ff32m)
}

//MoneyStrToInt32 string转int32
func MoneyStrToInt32(strm string) int {
	temp, err := strconv.Atoi(strm)
	if err != nil {
		logger.Error(err.Error())
		return -1
	}

	return temp

}

//MoneyEnlargeInt32ToStr 放大后int32转string
func MoneyEnlargeInt32ToStr(i64m int) string {
	ff32m := float64(i64m / Enlarge)

	return strconv.FormatFloat(ff32m, 'f', 2, 32)

}

//MoneyEnlargeStrToStr 放大后string转string
func MoneyEnlargeStrToStr(str32m string) string {
	temp, _ := strconv.Atoi(str32m)

	ff32m := float64(temp / Enlarge)

	return strconv.FormatFloat(ff32m, 'f', 2, 32)

}

//MoneyStrToInt64 string转int64
func MoneyStrToInt64(strm string) int64 {
	f32m, err := strconv.ParseInt(strm, 10, 32)
	if err != nil {
		logger.Error(err.Error())
	}

	return f32m
}

//MoneyEnlargeInt64ToStr 放大后int64转string
func MoneyEnlargeInt64ToStr(i64m int64) string {
	ff32m := float64(i64m / Enlarge)

	return strconv.FormatFloat(ff32m, 'f', 2, 32)

}
