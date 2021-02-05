package model

import (
	"database/sql"
	"time"
)

type Users struct {
	ID           int
	RoleCode     string
	Email        string
	Password     string
	FullName     sql.NullString
	Phone        sql.NullInt64
	VerifiedCode sql.NullString
	IsVerified   bool
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
}

type UserInfo struct {
	ID            int
	RoleCode      string
	Role          string
	Email         string
	FullName      sql.NullString
	Phone         sql.NullInt64
	Profession    sql.NullString
	Gender        sql.NullString
	Age           sql.NullInt64
	Province      sql.NullString
	City          sql.NullString
	Address       sql.NullString
	ImageCode     sql.NullString
	ImageFilename sql.NullString
	ImageFilepath sql.NullString
	IsVerified    bool
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
}
