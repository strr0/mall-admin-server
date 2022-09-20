package util

import (
	"errors"
	"time"
)

var (
	yyyyMMdd            = "20060102"
	yyyy_MM_dd          = "2006-01-02"
	yyyyMMddHHmmss      = "20060102150405"
	yyyy_MM_dd_HH_mm_ss = "2006-01-02 15:04:05"
)

func ParseTimeWithErr(timeStr string) (time.Time, error) {
	var res time.Time
	var err error
	switch len(timeStr) {
	case 10:
		res, err = time.Parse(yyyy_MM_dd, timeStr)
	case 19:
		res, err = time.Parse(yyyy_MM_dd_HH_mm_ss, timeStr)
	default:
		err = errors.New("time error")
	}
	return res, err
}

func DateToStr(current time.Time) string {
	return current.Format(yyyyMMdd)
}

func TimeToStr(current time.Time) string {
	return current.Format(yyyyMMddHHmmss)
}
