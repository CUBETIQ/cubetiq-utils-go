package string

import (
	"crypto/rand"
	"log"
	"math/big"
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

func ToBoolean(s string) (bool, error) {
	result, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsNotEmpty(s string) bool {
	return s != ""
}

func HasPrefix(existString, newString string) bool {
	return strings.HasPrefix(existString, newString)
}

func HasSuffix(existString, newString string) bool {
	return strings.HasSuffix(existString, newString)
}

func TrimWhiteSpace(s string) string {
	return strings.Trim(s, " ")
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
