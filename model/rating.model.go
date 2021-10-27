package model

import (
	"database/sql"
	"time"
)

type Rating struct {
	ID           int
	Code         string
	ClassCode    string
	ClassName    string
	ClassInitial sql.NullString
	UserCode     string
	UserName     string
	Rating       float64
	Comment      sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}

type RatingPost struct {
	ID           int
	Code         string
	ClassCode    string
	MentorCode   string
	UserCode     string
	Rating       int
	Comment      sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}
