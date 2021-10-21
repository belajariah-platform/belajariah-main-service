package model

import (
	"database/sql"
	"time"
)

type ClassTest struct {
	ID           int
	Code         string
	ClassCode    string
	TestTypeCode string
	Question     string
	OptionA      string
	OptionB      string
	OptionC      string
	OptionD      string
	Answer       int
	TestImage    sql.NullString
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
