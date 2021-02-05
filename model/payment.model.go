package model

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID                 int
	UserCode           string
	UserName           string
	ClassCode          string
	ClassName          string
	ClassInitial       string
	PaymentMethodCode  string
	PaymentMethod      string
	InvoiceNumber      string
	StatusPaymentCode  string
	StatusPayment      string
	TotalTransfer      int
	SenderBank         sql.NullString
	DestinationAccount sql.NullString
	ImageProof         sql.NullString
	IsActive           bool
	CreatedBy          string
	CreatedDate        time.Time
	ModifiedBy         sql.NullString
	ModifiedDate       sql.NullTime
}
