package model

import (
	"database/sql"
	"time"
)

type Enum struct {
	ID           int
	Code         string
	Type         string
	Value        string
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}
