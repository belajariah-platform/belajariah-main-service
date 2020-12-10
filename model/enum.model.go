package model

import (
	"database/sql"
	"time"
)

type Enum struct {
	ID           int
	Code         string
	EnumType     string
	EnumValue    string
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}
