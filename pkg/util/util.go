package util

import (
	"time"
)

// CheckDate check date
func CheckDate(seconds int64) (expired bool) {
	return !time.Unix(seconds, 0).After(time.Now())
}

// GetPosition get position
func GetPosition(seconds int64, length int) int {
	return int(seconds % int64(length))
}
