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
	ClassInitial      string
	PackageCode       string
	PackageType       string
	PaymentMethodCode string
	PaymentMethod     string
	InvoiceNumber     string
	StatusPaymentCode string
	StatusPayment     string
	TotalTransfer     int
	SenderBank        sql.NullString
	SenderName        sql.NullString
	ImageCode         sql.NullInt64
	ImageProof        sql.NullString
	PaymentTypeCode   string
	PaymentType       string
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
}
