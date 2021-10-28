package model

import (
	"time"
)

type Email struct {
	ID            int       `form:"id" json:"id" xml:"id" db:"id"`
	Code          string    `form:"code" json:"code" xml:"code" db:"code"`
	Type          string    `form:"type" json:"type" xml:"type" db:"type"`
	Body          string    `form:"body" json:"body" xml:"body" db:"body"`
	Subject       string    `form:"subject" json:"subject" xml:"subject" db:"subject"`
	Is_Active     bool      `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By    string    `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date  time.Time `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By   string    `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date NullTime  `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted    bool      `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}

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
