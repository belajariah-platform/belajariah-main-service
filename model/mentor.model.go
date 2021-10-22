package model

import (
	"database/sql"
	"time"
)

type Mentor struct {
	ID                 int
	Code               string
	MentorCode         int
	RoleCode           string
	Role               string
	Email              string
	FullName           sql.NullString
	Phone              sql.NullInt64
	Profession         sql.NullString
	Gender             sql.NullString
	Age                sql.NullInt64
	Province           sql.NullString
	City               sql.NullString
	Address            sql.NullString
	Birth              sql.NullTime
	ImageProfile       sql.NullString
	Description        sql.NullString
	AccountOwner       sql.NullString
	AccountName        sql.NullString
	AccountNumber      sql.NullString
	LearningMethod     sql.NullString
	LearningMethodText sql.NullString
	Rating             float64
	IsActive           bool
	CreatedBy          string
	CreatedDate        time.Time
	ModifiedBy         sql.NullString
	ModifiedDate       sql.NullTime
}

type MentorSchedule struct {
	ID           int
	Code         string
	MentorCode   string
	ShiftName    string
	StartDate    time.Time
	EndDate      time.Time
	TimeZone     string
	Sequence     int
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}

type MentorExperience struct {
	ID           int
	Code         string
	MentorCode   string
	Experience   string
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}
