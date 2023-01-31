package number

import "strconv"

func IntToInt16(i int) int16 {
	return int16(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

func IntToInt32(i int) int32 {
	return int32(i)
}

func Int64ToFloat64(i int64) float64 {
	return float64(i)
}

func Float64ToString(f float64, pre int) string {
	return strconv.FormatFloat(f, 'g', pre, 64)
}
