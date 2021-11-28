package utils

import (
	"time"
)

//获取相差时间
func GetDaysDiffer(start_time, end_time string) int64 {
    var days int64
    t1, err := time.ParseInLocation("2006-01-02", start_time, time.Local)
    t2, err := time.ParseInLocation("2006-01-02", end_time, time.Local)
    if err == nil && t1.Before(t2) {
        diff := t2.Unix() - t1.Unix()
        days = diff / (3600 * 24)
        return days
    } else {
        return days
    }
}