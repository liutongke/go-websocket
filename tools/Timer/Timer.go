package timer

import (
	"time"
)

func currentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
	//return time.Now()
}

// GetStrToUnix 字符串时间转为unix时间戳
// timeStr "2023-06-07 02:51:00"
// layout 为空则使用默认
func GetStrToUnix(timeStr, layout string) int64 {
	//timeStr := "2023-06-07 02:51:00" // 提供的字符串时间
	timezone := "Asia/Shanghai" // 设置时区，例如 "Asia/Shanghai"

	// 加载时区
	loc, _ := time.LoadLocation(timezone)

	// 解析字符串时间并转换为指定时区的时间
	if layout == "" {
		layout = "2006-01-02 15:04:05" // 根据提供的时间字符串格式进行指定
	}
	t, _ := time.ParseInLocation(layout, timeStr, loc)

	// 转换为 Unix 时间戳
	return t.Unix()
}

// GetUnixToStr 根据提供的unix时间戳转为字符串 layout如果为空则根据默认格式
func GetUnixToStr(unixTimestamp int64, layout string) string {
	//unixTimestamp := int64(1686944400) // 例：提供的 Unix 时间戳

	// 转换为时间
	t := time.Unix(unixTimestamp, 0)
	if layout == "" {
		// 格式化为字符串
		layout = "2006-01-02 15:04:05" // 根据需要指定时间字符串的格式
	}
	return t.Format(layout)
}

// GetOffsetZeroUnix 获取偏移的零点时间戳 -3表示3天前  0表示当天 3表示未来3天
func GetOffsetZeroUnix(offset int) int64 {
	// 获取当前时间
	now := currentTime()
	// 计算 N 天前的时间
	nDaysAgo := now.AddDate(0, 0, offset) // N 为要获取的天数
	// 获取 N 天前零点时间
	zeroTime := time.Date(nDaysAgo.Year(), nDaysAgo.Month(), nDaysAgo.Day(), 0, 0, 0, 0, nDaysAgo.Location())
	// 输出 N 天前零点时间戳
	return zeroTime.Unix()
}

// GetTodayZeroUnix 获取今天零点unix时间戳
func GetTodayZeroUnix() int64 {
	// 获取当前时间
	now := currentTime()
	// 获取当天零点时间
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 输出当天零点时间戳
	return zeroTime.Unix()
}

// GetOffsetDateId 获取偏移id offset -3表示3天前  0表示当天 3表示未来3天
func GetOffsetDateId(offset int) string {
	offsetTime := currentTime().AddDate(0, 0, offset)
	return offsetTime.Format("20060102")
}

// GetOffsetUnix 获取偏移的时间戳 -n表示过去时间 n表示未来时间 0表示当前时间
func GetOffsetUnix(offset int) int64 {
	// 计算 N 小时前的时间
	unixTm := currentTime().Add(time.Hour * time.Duration(offset))

	// 获取 Unix 时间戳
	return unixTm.Unix()
}

// GetDateId 获取今日日期id
func GetDateId() string {
	return currentTime().Format("20060102")
}

// GetNowStr 获取时间戳
func GetNowStr() string {
	formattedTime := currentTime().Format("2006-01-02 15:04:05")
	return formattedTime
}

// GetPrevHourId 获取上一个小时
func GetPrevHourId() string {
	previousHour := currentTime().Add(-time.Hour)
	return previousHour.Format("2006010215")
}

// GetAgoHourId 获取前x小时
func GetAgoHourId(n int) string {
	// 计算 n 小时前的时间
	hoursAgo := currentTime().Add(-time.Duration(n) * time.Hour)

	//fmt.Println(n, "小时前的时间是:", hoursAgo)
	return hoursAgo.Format("2006010215")
}

// GetOffsetHourId 获取指定的偏移小时id
func GetOffsetHourId(n int) string {
	// 计算 n 小时前的时间
	hoursAgo := currentTime().Add(time.Duration(n) * time.Hour)

	//fmt.Println(n, "小时前的时间是:", hoursAgo)
	return hoursAgo.Format("2006010215")
}

// GetNowHourId 获取当前小时
func GetNowHourId() string {
	//fmt.Println("当前时间：", currentTime)
	//fmt.Println("上一个小时的时间：", previousHour)
	return currentTime().Format("2006010215")
}

// GetNowUnix 获取10位unix时间戳
func GetNowUnix() int64 {
	return currentTime().Unix()
}

// GetNowMilliSecond 获取当前毫秒
func GetNowMilliSecond() int64 {
	milliseconds := currentTime().UnixNano() / int64(time.Millisecond)

	//fmt.Println("当前毫秒时间：", milliseconds)
	return milliseconds
}

// GetMicroseconds 获取当前微秒
func GetMicroseconds() int64 {
	microseconds := currentTime().UnixNano() / int64(time.Microsecond)

	//fmt.Println("当前微秒时间：", microseconds)
	return microseconds
}

// GetFormat 根据提供的布局格式化时间，例如：layout 2006-01-02 15:04:05
func GetFormat(layout string) string {
	return currentTime().Format(layout)
}

// FormatTime 根据提供的布局和时间格式化时间，例如：layout 2006-01-02 15:04:05
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// GetUnixFromTime 获取指定时间unix时间，例如：获取 2023 年 7 月 31 日 12 时 0 分 0 秒的 Unix 时间戳
// unixTimestamp := GetUnixTimestamp(2023, 7, 31, 12, 0, 0)
func GetUnixFromTime(year, month, day, hour, minute, second int) int64 {
	// 构造指定时间
	targetTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, currentTime().Location())

	// 获取对应的 Unix 时间戳
	unixTimestamp := targetTime.Unix()
	return unixTimestamp
}

// GetMillisecondsFromTime 获取指定时间的毫秒，例如：2050年1月1日0点0分0秒的毫秒数
// milliseconds := GetMillisecondsFromTime(2050, 1, 1, 0, 0, 0, 0)
func GetMillisecondsFromTime(year int, month time.Month, day, hour, minute, second, millisecond int) int64 {
	targetTime := time.Date(year, month, day, hour, minute, second, millisecond*int(time.Millisecond), currentTime().Location())
	return targetTime.UnixNano() / int64(time.Millisecond)
}

// GetMicrosecondsFromTime 获取指定微秒，例如：获取 2023 年 7 月 31 日 12 时 0 分 0 秒 500 微秒的 Unix 时间戳（微秒）
// microseconds := GetMicroseconds(2023, 7, 31, 12, 0, 0, 500)
func GetMicrosecondsFromTime(year, month, day, hour, minute, second, microsecond int) int64 {
	// 构造指定时间
	targetTime := time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1000, currentTime().Location())

	// 获取对应的 Unix 时间戳
	unixTimestamp := targetTime.UnixNano() / 1000 // 转换为微秒

	return unixTimestamp
}

// GetHourUnixTimestamp 获取当天指定 24 小时中某个小时的 Unix 时间戳 例如：unixTimestamp := GetHourUnixTimestamp(23)
func GetHourUnixTimestamp(hour int) int64 {
	now := currentTime()
	date := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, currentTime().Location())
	return date.Unix()
}

// GetNextWeekTime 获取下一个星期几的特定小时、分钟、秒的时间 例如：获取下一个星期二的15点30分30秒
//func GetNextWeekTime(weekday time.Weekday, hour, minute, second int) time.Time {
//	now := currentTime()
//	daysUntilWeekday := (int(weekday) - int(now.Weekday()) + 7) % 7
//	if daysUntilWeekday == 0 {
//		// 如果今天就是指定的星期几，则计算下一个星期几
//		daysUntilWeekday = 7
//	}
//	targetDate := now.AddDate(0, 0, daysUntilWeekday)
//	return time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), hour, minute, second, 0, currentTime().Location())
//}

// GetNowWeekTime 获取当前星期几的特定小时、分钟、秒的时间 例如：获取下一个星期二的15点30分30秒
//func GetNowWeekTime(weekday time.Weekday, hour, minute, second int) time.Time {
//	now := currentTime()
//	daysUntilWeekday := (int(weekday) - int(now.Weekday()) + 7) % 7
//	targetDate := now.AddDate(0, 0, daysUntilWeekday)
//	return time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), hour, minute, second, 0, currentTime().Location())
//}

//var cstZone = time.FixedZone("CST", 8*3600) // 东八
//
////格式化后的时间戳
//func FormatNow() string {
//	tm := nowTm()
//	return tm.Format("2006-01-02 15:04:05")
//}
//
////获取当前UNIX时间戳
//func FormatNowUnix() int64 {
//	return nowTm().Unix()
//}

//格式化字符串获取UNIX时间戳
//func FormatStrToUnix(toBeCharge string) int64 {
//	timeLayout := "2006-01-02"                                      //转化所需模板
//	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
//	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
//	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
//	return sr
//}

//时间戳转时间
//func UnixToStr(timeUnix int64, layout string) string {
//	timeStr := time.Unix(timeUnix, 0).Format(layout)
//	return timeStr
//}

//时间转时间戳
//func StrToUnix(timeStr, layout string) (int64, error) {
//	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
//	if err != nil {
//		return 0, err
//	}
//	tt, err := time.ParseInLocation(layout, timeStr, local)
//	if err != nil {
//		return 0, err
//	}
//	timeUnix := tt.Unix()
//	return timeUnix, nil
//}
