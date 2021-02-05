package model

import (
	"database/sql"
	"time"
)

type RatingClass struct {
	ID           int
	ClassCode    string
	ClassName    string
	ClassInitial sql.NullString
	UserCode     int
	Rating       float64
	Comment      sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
