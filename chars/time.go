package chars

import (
	"time"
)

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

type restTime struct {
	restTimeFrom  time.Duration
	restTimeTo    time.Duration
	crossMidNight bool
}

func RestTime(hFrom, mFrom, hTo, mTo int, crossMidNight bool) *restTime {
	rt := restTime{}
	rt.SetRestTime(hFrom, mFrom, hTo, mTo, crossMidNight)
	return &rt
}

func (rt *restTime) SetRestTime(hFrom, mFrom, hTo, mTo int, crossMidNight bool) {
	rt.restTimeFrom = time.Duration(hFrom)*time.Hour + time.Duration(mFrom)*time.Minute
	rt.restTimeTo = time.Duration(hTo)*time.Hour + time.Duration(mTo)*time.Minute
	rt.crossMidNight = crossMidNight
}

func (rt *restTime) IsRestTime(t time.Time) bool {
	return !rt.IsWorkingTime(t)
}

func (rt *restTime) IsWorkingTime(t time.Time) bool {
	now := time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute

	if !rt.crossMidNight && now >= rt.restTimeFrom && now <= rt.restTimeTo {
		return false
	}

	if rt.crossMidNight && (now >= rt.restTimeFrom || now <= rt.restTimeTo) {
		return false
	}

	return true

}
