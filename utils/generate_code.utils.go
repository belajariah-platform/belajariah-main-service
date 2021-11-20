package utils

import (
	"belajariah-main-service/shape"
	"fmt"
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
func GenerateRandomCode2(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	return StringWithCharset(length, charset)
}

func GenerateVerifyCode(val string) string {
	strNum := val[0:1] + val[2:3] +
		"BL-" + GenerateRandomCode(4)

	return strings.ToUpper(strNum)
}

func GenerateInvoiceNumber(value shape.PaymentPost, types string) string {
	var strNum string
	date := time.Now()

	if types == "class general" {
		strNum = "INV/cg" + fmt.Sprintf(`%02d`, date.Second()) +
			GenerateRandomCode2(8) + fmt.Sprintf(`/%02d/%d`, date.Month(), date.Year())
	} else if types == "class quran" {
		strNum = "INV/cq" + fmt.Sprintf(`%02d`, date.Second()) +
			GenerateRandomCode2(8) + fmt.Sprintf(`/%02d/%d`, date.Month(), date.Year())
	}

	return strNum
}
