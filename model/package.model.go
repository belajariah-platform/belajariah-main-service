package model

import (
	"database/sql"
	"time"
)

type Package struct {
	ID            int
	Code          string
	ClassCode     string
	Type          string
	PricePackage  string
	PriceDiscount sql.NullString
	Duration      int
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	DeletedBy     sql.NullString
	DeletedDate   sql.NullTime
}
