package util

import (
	"math/rand"
	"strings"
	"time"
)

func Ellipsis(s string, l int) string {
	if len(s) < l {
		return s
	}
	final := []string{s[:l], "..."}
	return strings.Join(final, "")
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return stringWithCharset(length, charset)
}
