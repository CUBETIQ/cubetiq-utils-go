package date

import "time"

func MillisToDateTime(m int64) time.Time {
	return time.UnixMilli(m)
}
