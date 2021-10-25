package main

import (
	"fmt"
	"time"
)

const (
	SecondsPerMinute = 60
	SecondsPerHour = SecondsPerMinute * 60
	SecondsPerDay = SecondsPerHour * 24
)

/*
时间转换函数
结果: 10:49:30
*/
func FormatTime(second float32)string{
	seconds := int(second)
	min := seconds % SecondsPerHour /SecondsPerMinute
	hour := seconds % SecondsPerDay /SecondsPerHour
	days := seconds /SecondsPerDay
	if days > 0{
		return fmt.Sprintf(" %d days,  %d:%02d",days,hour,min)
	}
	return fmt.Sprintf("  %d:%02d",hour,min)
}

func timeNow() string {
        now := time.Now()
        return now.Format("15:04:05")
}

