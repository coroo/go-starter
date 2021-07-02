package utils

import "time"

func DateNow(formatDate string) string {
	if formatDate == "" {
		return time.Now().Format("2006-01-02")
	}
	return time.Now().Format(formatDate)
}
