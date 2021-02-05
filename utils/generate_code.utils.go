package utils

import (
	"math/rand"
	"strings"
	"time"
)

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return StringWithCharset(length, charset)
}

func GenerateVerifyCode(val string) string {
	strNum := val[0:1] + val[2:3] +
		"BL-" + GenerateRandomCode(4)

	return strings.ToUpper(strNum)
}
