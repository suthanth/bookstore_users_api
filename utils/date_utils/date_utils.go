package date_utils

import "time"

const (
	dateFormat = "2006-01-02T15:04:05Z"
)

func GetDateNow() time.Time {
	return time.Now().UTC()
}

func GetDateNowString() string {
	return GetDateNow().Format(dateFormat)
}
