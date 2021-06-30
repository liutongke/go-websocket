package Timer

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func nowTm() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

//func UnixFormatStr(unixTm int) string {
//	return unixTm.Format("2006-01-02 15:04:05")
//}

//将字符串转换成UNIX时间戳
func StrFormatUnix(toBeCharge string) int {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	//ts, _ := time.ParseInLocation("2006-01-02 15:04:05", strings.Replace(toBeCharge, "T", " ", -1), time.Local)
	ts, _ := time.ParseInLocation("2006/01/02", toBeCharge, time.Local)
	tms := cast.ToInt(ts.In(cstSh).Unix())
	if tms <= 0 {
		return 0
	}
	return tms
}

//获取当前的时间戳
func NowUnix() int {
	nowUnix := nowTm().Unix() //获取当前时间戳
	return cast.ToInt(nowUnix)
}

//获取偏移的时间戳
func OffsetUinx(offset int) int {
	return NowUnix() + offset
}

//获取当前时间戳字符串
func NowStr() string {
	return nowTm().Format("2006-01-02 15:04:05")
}

func DateId() string {
	return nowTm().Format("20060102")
}

//获取偏移id
func OffsetDateId(offset int) {
	oldTime := nowTm().AddDate(0, 0, offset) //若要获取3天前的时间，则应将-2改为-3
	fmt.Println(oldTime.Format("20060102"))
}

//获取今日零点
func DateZeroUnix() int {
	timeStr := nowTm().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return cast.ToInt(t.Unix())
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
