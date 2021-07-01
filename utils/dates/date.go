package dates

import "time"

const dateLayout = "2006-01-02 15:04:05"

func GetNowString() string {
	return GetNow().Format(dateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
