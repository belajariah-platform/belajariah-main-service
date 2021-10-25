package model

import (
	"time"
)

type EmailBody struct {
	Subject          string
	BodyTemp         string
	UserName         string
	UserEmail        string
	UserCode         string
	Count            int
	ExpiredDate      time.Time
	VerificationCode string
	CopyRight        string

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
	ClassImage        string
	PromoCode         string
	PromoTitle        string
	PromoDiscount     string
	PromoPrice        int
	PromoImage        string
	TotalTransfer     int
	TotalConsultation int
	TotalWebinar      int

	//Social Media
	Facebook   string
	WhatsApp   string
	Youtube    string
	Instagram  string
	GooglePLay string
}
