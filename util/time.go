package util

import "time"

func CustomFormat(t time.Time) string {
	format := "Mon, 2006-Jan-02 03:04:05 PM"
	return t.Format(format)
}
