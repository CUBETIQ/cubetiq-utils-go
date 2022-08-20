package string

import (
	"strconv"
	"strings"
)

func GetPartOfLast(s, sep string) string {
	return strings.Split(s, sep)[len(strings.Split(s, sep))-1]
}

func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsNotEmpty(s string) bool {
	return s != ""
}
