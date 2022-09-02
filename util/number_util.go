package util

import "strconv"

// 字符串转数字
func ParseInt(numStr string, num int) int {
	res, err := ParseIntWithErr(numStr)
	if err != nil {
		return num
	}
	return res
}

func ParseIntWithErr(numStr string) (int, error) {
	res, err := strconv.ParseInt(numStr, 10, 0)
	return int(res), err
}

func ParseInt64(numStr string, num int64) int64 {
	res, err := ParseInt64WithErr(numStr)
	if err != nil {
		return num
	}
	return res
}

func ParseInt64WithErr(numStr string) (int64, error) {
	return strconv.ParseInt(numStr, 10, 64)
}

func ParseInt32(numStr string, num int32) int32 {
	res, err := ParseInt32WithErr(numStr)
	if err != nil {
		return num
	}
	return res
}

func ParseInt32WithErr(numStr string) (int32, error) {
	res, err := strconv.ParseInt(numStr, 10, 32)
	return int32(res), err
}

// 字符串数组
func ParseInt64Array(numsStr []string) []int64 {
	var res []int64
	for _, numStr := range numsStr {
		num, err := ParseInt64WithErr(numStr)
		if err == nil {
			res = append(res, num)
		}
	}
	return res
}
