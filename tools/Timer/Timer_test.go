package Timer

import (
	"log"
	"testing"
)

func TestTimer(t *testing.T) {
	log.Println("GetNowStr", GetNowStr())
	log.Println("GetNowUnix", GetNowUnix())
	log.Println("GetPrevHourId", GetPrevHourId())
	log.Println("GetAgoHourId 10", GetAgoHourId(10))
	log.Println("GetNowHourId", GetNowHourId())
	log.Println("GetMilliSecond", GetMilliSecond())
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
}
