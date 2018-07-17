package chars

import "time"

//数据库中时间一般以时间戳或无其他符号纯数字形式存在

//TimeIntPRC 中国纯数字时间
func TimeIntPRC(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102150405")

	return cstTime
}

//TimeIntYMD 中国数字时间，只取年月日
func TimeIntYMD(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102")

	return cstTime

}

//Time24Sub 计算下一个零点
func Time24Sub(t time.Time) time.Duration {
	next := t.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	return next.Sub(t)

}

//TimeID 以时间为基准的ID生成
func TimeID() string {
	return time.Now().Format("20060102") + RandomNumeric(8)
}
