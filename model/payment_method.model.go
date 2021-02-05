package model

import (
	"database/sql"
	"time"
)

type PaymentMethod struct {
	ID            int
	Code          string
	Type          string
	Value         string
	AccountName   sql.NullString
	AccountNumber sql.NullString
	MethodImage   sql.NullString
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	DeletedBy     sql.NullString
	DeletedDate   sql.NullTime
}
