package model

import (
	"database/sql"
)

type InstructorInfo struct {
	ID             int
	RoleCode       string
	Role           string
	Email          string
	FullName       sql.NullString
	Biografi       sql.NullString
	Description    sql.NullString
	Phone          sql.NullInt64
	Profession     sql.NullString
	Gender         sql.NullString
	Age            sql.NullInt64
	Province       sql.NullString
	City           sql.NullString
	Address        sql.NullString
	ImageCode      sql.NullString
	ImageFilename  sql.NullString
	ImageFilepath  sql.NullString
	Rating         float64
	TaskCompleted  int
	TaskInprogress int
}
