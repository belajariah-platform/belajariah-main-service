package model

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID                int
	UserCode          int
	UserName          string
	ClassCode         string
	ClassName         string
	ClassImage        sql.NullString
	ClassInitial      string
	PackageCode       string
	PackageType       string
	PackageDiscount   sql.NullInt64
	TotalConsultation sql.NullInt64
	TotalWebinar      sql.NullInt64
	PromoCode         sql.NullString
	PromoTitle        sql.NullString
	PromoDiscount     sql.NullFloat64
	PaymentMethodCode string
	PaymentMethod     string
	AccountName       sql.NullString
	AccountNumber     sql.NullString
	InvoiceNumber     string
	StatusPaymentCode string
	StatusPayment     string
	TotalTransfer     int
	SenderBank        sql.NullString
	SenderName        sql.NullString
	ImageCode         sql.NullInt64
	ImageProof        sql.NullString
	ImageFilename     sql.NullString
	PaymentTypeCode   string
	PaymentType       string
	Remarks           sql.NullString
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
}
