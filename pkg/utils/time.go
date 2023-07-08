package utils

import "time"

func ConvertMinutesToMS(minutes int) int {
	return minutes * 60 * 1000
}

func TimeNowInTimestamp() int {
	return int(time.Now().UnixNano() / int64(time.Millisecond))
}
