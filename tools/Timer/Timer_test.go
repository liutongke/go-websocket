package timer

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	log.Println("GetNowStr", GetNowStr())
	log.Println("GetNowUnix", GetNowUnix())
	log.Println("GetPrevHourId", GetPrevHourId())
	log.Println("GetAgoHourId 10", GetAgoHourId(10))
	log.Println("GetNowHourId", GetNowHourId())
	log.Println("GetMilliSecond", GetNowMilliSecond())
	log.Println("GetMicroseconds", GetMicroseconds())
	log.Println("GetFormat", GetFormat("2006-01-02 15:04:05"))
	log.Println("GetDateId", GetDateId())
	log.Println("GetOffsetUnix -5", GetOffsetUnix(-5))
	log.Println("GetOffsetUnix 5", GetOffsetUnix(5))
	log.Println("GetOffsetUnix 0", GetOffsetUnix(0))
	log.Println("GetOffsetDateId -3", GetOffsetDateId(-3))
	log.Println("GetOffsetDateId 0", GetOffsetDateId(0))
	log.Println("GetOffsetDateId 3", GetOffsetDateId(3))
	log.Println("GetTodayZeroUnix", GetTodayZeroUnix())
	log.Println("GetOffsetZeroUnix -10", GetOffsetZeroUnix(-10))
	log.Println("GetOffsetZeroUnix 0", GetOffsetZeroUnix(0))
	log.Println("GetOffsetZeroUnix 10", GetOffsetZeroUnix(10))
	log.Println("GetUnixToStr", GetUnixToStr(int64(1686945563), "")) //2023-06-17 03:59:23
	log.Println("GetStrToUnix", GetStrToUnix("2023-06-17 03:59:23", "2006-01-02 15:04:05"))
	log.Println("GetMilliseconds", GetMillisecondsFromTime(2050, 1, 1, 0, 0, 0, 0))
	log.Println("GetUnixTimestamp", GetUnixFromTime(2023, 7, 31, 12, 0, 0))
	log.Println("GetMicrosecondsFromTime", GetMicrosecondsFromTime(2023, 7, 31, 12, 0, 0, 500))
	log.Println("GetOffsetHourId 10", GetOffsetHourId(10))
	log.Println("GetOffsetHourId -10", GetOffsetHourId(-10))
	log.Println("GetOffsetHourId 0", GetOffsetHourId(0))
	log.Println("FormatTime", FormatTime(time.Now(), "2006-01-02 15:04:05"))
	log.Println("GetHourUnixTimestamp 9", GetHourUnixTimestamp(9))
	//log.Println("GetSpecificTime 获取下一个星期日的15点30分30秒", GetNextWeekTime(time.Sunday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期一的15点30分30秒", GetNextWeekTime(time.Monday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期二的15点30分30秒", GetNextWeekTime(time.Tuesday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期三的15点30分30秒", GetNextWeekTime(time.Wednesday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期四的15点30分30秒", GetNextWeekTime(time.Thursday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期五的15点30分30秒", GetNextWeekTime(time.Friday, 15, 30, 30))
	//log.Println("GetSpecificTime 获取下一个星期六的15点30分30秒", GetNextWeekTime(time.Saturday, 15, 30, 30))

	//log.Println("GetNowWeekTime 获取当前星期二的15点30分30秒", GetNowWeekTime(time.Monday, 15, 30, 30))
}
