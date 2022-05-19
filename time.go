package core


import (
	"time"
)


func NowUnixTime(mode ...string) int64 {
	now := time.Now()

	if len(mode) != 0 {
		switch mode[0] {
		case "nano":
			return now.UnixNano()
		case "milli":
			return now.UnixMilli()
		case "micro":
			return now.UnixMicro()
		case "second":
			return now.Unix()
		default:
			return now.Unix()
		}		
	} else {
		return now.Unix()
	}
}

func ParseUnixTime(nowUnix int64, mode ...string) time.Time {

	if len(mode) != 0 {
		switch mode[0] {
		case "milli":
			return time.UnixMilli(nowUnix)
		case "micro":
			return time.UnixMicro(nowUnix)
		case "second":
			return time.Unix(nowUnix, 0)
		default:
			return time.Unix(nowUnix, 0)
		}
	} else {
		return time.Unix(nowUnix, 0)
	}
}