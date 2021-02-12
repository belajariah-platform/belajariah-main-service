package model

import (
	"database/sql"
)

type Instructor struct {
	ID            int
	Code          string
	RoleCode      string
	Role          string
	FullName      sql.NullString
	Biografi      sql.NullString
	Description   sql.NullString
	ImageCode     sql.NullString
	ImageFilename sql.NullString
	ImageFilepath sql.NullString
}
