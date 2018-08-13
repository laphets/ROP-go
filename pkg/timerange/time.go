package timerange

import "time"

func GetStatus(start, end time.Time) string {
	return GetStatusSync(time.Now(), start, end)
}

func GetStatusSync(now, start, end time.Time) string {
	if now.Before(start) {
		return "before"
	} else if now.After(end) {
		return "end"
	} else {
		return "cur"
	}
}