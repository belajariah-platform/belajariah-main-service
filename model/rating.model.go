package model

import (
	"database/sql"
	"time"
)

type Rating struct {
	ID           int
	ClassCode    string
	ClassName    string
	ClassInitial sql.NullString
	UserCode     int
	UserName     string
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

type RatingPost struct {
	ID           int
	ClassCode    string
	MentorCode   int
	UserCode     int
	Rating       int
	Comment      sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
