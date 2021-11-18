package model

import (
	"database/sql"
	"time"
)

type Package struct {
	ID                int
	Code              string
	ClassCode         string
	Type              string
	PricePackage      string
	PriceDiscount     sql.NullString
	Description       sql.NullString
	Duration          int
	DurationFrequence sql.NullInt64
	Webinar           sql.NullInt64
	Consultation      sql.NullInt64
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	IsDeleted         bool
}

type Benefit struct {
	ID           int
	Code         string
	ClassCode    string
	Description  string
	IconBenefit  sql.NullString
	Sequence     int
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}
