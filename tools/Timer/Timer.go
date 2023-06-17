package Timer

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

// GetOffsetZeroUnix 获取偏移的时间戳 -3表示3天前  0表示当天 3表示未来3天
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

// GetNowUnix 获取10位unix时间戳
func GetNowUnix() int64 {
	return currentTime().Unix()
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

// GetNowHourId 获取当前小时
func GetNowHourId() string {
	//fmt.Println("当前时间：", currentTime)
	//fmt.Println("上一个小时的时间：", previousHour)
	return currentTime().Format("2006010215")
}

// GetMilliSecond 获取当前毫秒
func GetMilliSecond() int64 {
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
