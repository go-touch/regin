package utils

import (
	time2 "time"
)

const TimeFormat = "2006-01-02 15:04:05"

// 时间处理器
type Timer struct {
	time2.Time
}

// 获取 Timer
func Time(timeStr ...string) Timer {
	var timer *Timer
	if timeStr != nil {
		if t, err := time2.Parse(TimeFormat, timeStr[0]); err == nil {
			timer = &Timer{
				t,
			}
		}
	}
	if timer == nil {
		timer = &Timer{
			time2.Now(),
		}
	}
	return *timer
}

// 格式化输出
func (t Timer) FormatStr() string {
	return t.Format(TimeFormat)
}
