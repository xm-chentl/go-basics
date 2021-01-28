package timeex

import (
	"time"

	"github.com/xm-chentl/go-basics/timeex/enum"
)

const missingDefaultTimeMsg = "请先设置timeex.DefaultTime"

// DefaultTime is 默认ITime
var DefaultTime ITime

// Now is 当前时间
func Now() time.Time {
	return DefaultTime.Now()
}

// NowUnix is 当前Unix
func NowUnix() int64 {
	if DefaultTime == nil {
		panic(missingDefaultTimeMsg)
	}

	return DefaultTime.Unix()
}

// NowUnixNano is 当前NanoUnix
func NowUnixNano() int64 {
	if DefaultTime == nil {
		panic(missingDefaultTimeMsg)
	}

	return DefaultTime.UnixNano()
}

// ParseDateUnix is 转(yyyy-MM-dd)时间戳
func ParseDateUnix(timeStr string) (int64, error) {
	timeObj, err := parseUnixByLayOut(enum.LayOutDate, timeStr)
	if err != nil {
		return 0, err
	}

	return timeObj, nil
}

// ParseUnix is 转(yyyy-MM-dd hh:mm:ss)时间戳
func ParseUnix(timeStr string) (int64, error) {
	timeObj, err := parseUnixByLayOut(enum.LayOutDateTime, timeStr)
	if err != nil {
		return 0, err
	}

	return timeObj, nil
}

// ParseDateString is 转(yyyy-MM-dd)字符串
func ParseDateString(unix int64) string {
	if unix == 0 {
		return ""
	}
	return time.Unix(unix, 0).Format(enum.LayOutDate)
}

// ParseString is 转(yyyy-MM-dd hh:mm:ss)字符串
func ParseString(unix int64) string {
	if unix == 0 {
		return ""
	}
	return time.Unix(unix, 0).Format(enum.LayOutDateTime)
}

func parseUnixByLayOut(layout, timeStr string) (int64, error) {
	timeObj, err := time.ParseInLocation(layout, timeStr, time.Local)
	if err != nil {
		return 0, err
	}

	return timeObj.Unix(), nil
}
