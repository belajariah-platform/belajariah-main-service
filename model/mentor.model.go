package model

import (
	"database/sql"
	"time"
)

type MentorRequest struct {
	Action string  `form:"action" json:"action" xml:"action"`
	Data   Mentors `form:"data" json:"data" xml:"data"`
	Query  Query   `form:"query" json:"query" xml:"query"`
}

type Mentors struct {
	ID           int
	Code         string
	RoleCode     string
	Email        string
	Password     string
	NewPassword  string
	OldPassword  string
	Full_Name    string
	Phone        NullInt64
	VerifiedCode sql.NullString
	IsVerified   bool
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime

	//
	Profession      string
	Gender          string
	Age             int
	Province        string
	City            string
	Address         string
	Birth           time.Time
	Description     string
	Account_Owner   string
	Account_Name    string
	Account_Number  string
	Learning_Method string
}

type MentorInfo struct {
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
	MinimumRate        sql.NullInt64
	Class_Code         string
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
	ClassCode    string
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

type MentorClass struct {
	ID           int
	Code         string
	MentorCode   string
	MentorName   sql.NullString
	ClassCode    string
	ClassName    string
	ClassInitial sql.NullString
	MinimumRate  sql.NullInt64
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}
