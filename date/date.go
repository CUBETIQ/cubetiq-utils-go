package date

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MillisToDateTime(m int64) time.Time {
	return time.UnixMilli(m)
}

func PrimitiveDateTimeToMillis(p primitive.DateTime) int64 {
	return p.Time().UnixMilli()
}
