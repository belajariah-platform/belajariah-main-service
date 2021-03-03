package utils

import "github.com/yudapc/go-rupiah"

func FormatRupiah(number int) string {
	amountFloat := float64(number)
	formatRupiah := rupiah.FormatRupiah(amountFloat)
	return formatRupiah
}
