package chars

import (
	"time"
)

//数据库中时间一般以时间戳或时间格式的形式存在

//AddTime 用法备忘
func AddTime(t time.Time, year, month, day int) time.Time {
	return t.AddDate(year, month, day)
}

//TimeNumPRC 中国纯数字时间
func TimeNumPRC(t time.Time) string {
	china, _ := time.LoadLocation("PRC")
	cstTime := t.In(china).Format("20060102150405")

	return cstTime
}

//TimeNumYMD 中国数字时间，只取年月日
func TimeNumYMD(t time.Time) string {
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

//RT 休息，非营业时间
type RT struct {
	restTimeFrom  time.Duration
	restTimeTo    time.Duration
	crossMidNight bool

	extraWaitTime time.Duration
}

//RestTime 初始化
func RestTime(hFrom, mFrom, hTo, mTo int, crossMidNight bool) *RT {
	rt := RT{}
	rt.SetRestTime(hFrom, mFrom, hTo, mTo, crossMidNight)
	return &rt
}

//SetRestTime 配置
func (rt *RT) SetRestTime(hFrom, mFrom, hTo, mTo int, crossMidNight bool) {
	rt.restTimeFrom = time.Duration(hFrom)*time.Hour + time.Duration(mFrom)*time.Minute
	rt.restTimeTo = time.Duration(hTo)*time.Hour + time.Duration(mTo)*time.Minute + time.Second*60
	rt.crossMidNight = crossMidNight
}

//SetExtWaitTime set extra wait time, only for rest time, not working in other time
func (rt *RT) SetExtWaitTime(extra time.Duration) {
	rt.extraWaitTime = extra
}

//IsRestTime 判断是否是非营业时间
func (rt *RT) IsRestTime(t time.Time) bool {
	return !rt.IsWorkingTime(t)
}

//IsWorkingTime 判断是否是工作时间
func (rt *RT) IsWorkingTime(t time.Time) bool {
	now := time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second

	if !rt.crossMidNight && now >= rt.restTimeFrom && now < rt.restTimeTo {
		return false
	}

	if rt.crossMidNight && (now >= rt.restTimeFrom || now < rt.restTimeTo) {
		return false
	}

	return true

}

//WaitTime 当前等待开业所需时间
func (rt *RT) WaitTime(t time.Time) time.Duration {
	now := time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second

	if !rt.crossMidNight && now >= rt.restTimeFrom && now < rt.restTimeTo {
		return rt.restTimeTo - now + rt.extraWaitTime
	}

	if rt.crossMidNight && now >= rt.restTimeFrom {
		return Time24Sub(t) + rt.restTimeTo + rt.extraWaitTime
	}

	if rt.crossMidNight && now < rt.restTimeTo {
		return rt.restTimeTo - now + rt.extraWaitTime
	}

	return 0

}
