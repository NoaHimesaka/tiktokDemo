package utils

import "strconv"

//todo 失败的默认值为0
func Int64(str string, defaultValue func() int64) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue()
	}
	return res
}
