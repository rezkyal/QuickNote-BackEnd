package util

import "strings"

func Ellipsis(s string, l int) string {
	if len(s) < l {
		return s
	}
	final := []string{s[:l], "..."}
	return strings.Join(final, "")
}
