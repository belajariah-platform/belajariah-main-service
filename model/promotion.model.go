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
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   string
	ModifiedDate time.Time
	DeletedBy    string
	DeletedDate  time.Time
}
