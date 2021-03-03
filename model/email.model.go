package model

import (
	"time"
)

type EmailBody struct {
	Subject          string
	BodyTemp         string
	UserName         string
	UserEmail        string
	UserCode         int
	Count            int
	ExpiredDate      time.Time
	VerificationCode string

	//Payment
	InvoiceNumber     string
	PackageCode       string
	PaymentMethodCode string
	PaymentMethod     string
	AccountName       string
	AccountNumber     string
	ClassCode         string
	ClassName         string
	ClassPrice        int
	PromoCode         string
	PromoTitle        string
	PromoDiscount     string
	PromoPrice        int
	TotalTransfer     int
	TotalConsultation int
	TotalWebinar      int
}
