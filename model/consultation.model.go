package model

import (
	"database/sql"
	"time"
)

type Consultation struct {
	ID                int
	UserCode          int
	UserName          string
	ClassCode         string
	ClassName         string
	RecordingCode     sql.NullInt64
	RecordingPath     sql.NullString
	RecordingName     sql.NullString
	RecordingDuration sql.NullInt64
	StatusCode        string
	Status            string
	Description       sql.NullString
	TakenCode         sql.NullInt64
	TakenName         sql.NullString
	IsPlay            sql.NullFloat64
	IsRead            sql.NullFloat64
	IsActionTaken     sql.NullFloat64
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	DeletedBy         sql.NullString
	DeletedDate       sql.NullTime
}
