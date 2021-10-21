package model

import (
	"database/sql"
	"time"
)

type Promotion struct {
	ID            int
	Code          string
	ClassCode     string
	Title         string
	Description   sql.NullString
	PromoCode     string
	PromoTypeCode sql.NullString
	PromoType     sql.NullString
	Discount      sql.NullFloat64
	ImageBanner   sql.NullString
	ImageHeader   sql.NullString
	ExpiredDate   sql.NullTime
	QuotaUser     sql.NullInt64
	QuotaUsed     sql.NullInt64
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	IsDeleted     bool
}
