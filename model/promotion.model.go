package model

import (
	"database/sql"
	"time"
)

type Promotion struct {
	ID           int
	Code         string
	ClassCode    string
	Title        string
	Description  sql.NullString
	PromoCode    string
	Discount     sql.NullFloat64
	BannerImage  sql.NullString
	HeaderImage  sql.NullString
	ExpiredDate  sql.NullTime
	QuotaUser    sql.NullInt64
	QuotaUsed    sql.NullInt64
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
