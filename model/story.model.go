package model

import (
	"database/sql"
	"time"
)

type Story struct {
	ID           int
	Code         string
	CategoryCode string
	ImageCode    sql.NullString
	VideoCode    sql.NullString
	Title        string
	Content      string
	Source       sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
