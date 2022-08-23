package env

import (
	"os"
	"strings"
)

func GetMapper(placeholderName string) string {
	split := strings.SplitN(placeholderName, ":", 2)
	defValue := ""
	if len(split) == 2 {
		placeholderName = split[0]
		defValue = split[1]
	}

	val, ok := os.LookupEnv(placeholderName)
	if !ok {
		return defValue
	}

	return val
}
