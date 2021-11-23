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

type RatingQuranRequest struct {
	Action string      `form:"action" json:"action" xml:"action"`
	Data   RatingQuran `form:"data" json:"data" xml:"data"`
	Query  Query       `form:"query" json:"query" xml:"query"`
}

type RatingQuran struct {
	ID           int        `form:"id" json:"id" xml:"id" db:"id"`
	Code         string     `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode    string     `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	ClassName    string     `form:"class_name" json:"class_name" xml:"class_name" db:"class_name"`
	ClassInitial NullString `form:"class_initial" json:"class_initial" xml:"class_initial" db:"class_initial"`
	UserCode     string     `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	UserName     string     `form:"user_name" json:"user_name" xml:"user_name" db:"user_name"`
	Rating       float64    `form:"rating" json:"rating" xml:"rating" db:"rating"`
	Comment      NullString `form:"comment" json:"comment" xml:"comment" db:"comment"`
	IsActive     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
