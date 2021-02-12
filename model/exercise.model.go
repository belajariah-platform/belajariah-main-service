package model

import (
	"database/sql"
	"time"
)

type Exercise struct {
	ID           int
	Code         string
	SubtitleCode string
	Option       string
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
