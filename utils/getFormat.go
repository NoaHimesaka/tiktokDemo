package utils

import "strings"

func GetFormat(str string) string {
	arr := strings.Split(str, ".")
	len := len(arr)
	return arr[len-1]
}
