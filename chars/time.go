package chars

import "time"

//数据库中时间一般存在纯数字形式

//TimeIntPRC 中国纯数字时间
func TimeIntPRC(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102150405")

	return cstTime
}

//TimeIntYMD 中国数字时间，只取年月日
func TimeIntYMD(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102150405")

	return cstTime[:8]

}
