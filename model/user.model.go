package model

import (
	"database/sql"
	"time"
)

type Users struct {
	ID                int
	Code              string
	RoleCode          string
	Email             string
	Password          string
	NewPassword       string
	OldPassword       string
	Full_Name         sql.NullString
	Phone             sql.NullInt64
	VerifiedCode      sql.NullString
	CountryNumberCode sql.NullString
	IsVerified        bool
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
}

type UserInfo struct {
	ID                int
	Code              string
	RoleCode          string
	Role              string
	Email             string
	FullName          sql.NullString
	Phone             sql.NullInt64
	Profession        sql.NullString
	Gender            sql.NullString
	Age               sql.NullInt64
	Birth             sql.NullTime
	Province          sql.NullString
	City              sql.NullString
	Address           sql.NullString
	ImageProfile      sql.NullString
	CountryNumberCode sql.NullString
	IsVerified        bool
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
}

type UserHeader struct {
	ID        int
	Code      string
	Role_Code string
	Role      string
	Email     string
	Full_Name string
	Phone     int
}
