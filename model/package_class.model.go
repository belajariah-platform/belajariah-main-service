package model

import (
	"database/sql"
	"time"
)

type PackageClass struct {
	ID             int
	Code           string
	ClassCode      string
	Type           string
	PricePackage   string
	PrinceDiscount sql.NullString
	IsActive       bool
	CreatedBy      string
	CreatedDate    time.Time
	ModifiedBy     sql.NullString
	ModifiedDate   sql.NullTime
	DeletedBy      sql.NullString
	DeletedDate    sql.NullTime
}
