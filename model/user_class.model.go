package model

import (
	"database/sql"
	"time"
)

type UserClass struct {
	ID                int
	Code              string
	UserCode          string
	ClassCode         string
	ClassName         string
	ClassInitial      sql.NullString
	ClassCategory     string
	ClassDescription  sql.NullString
	ClassImage        sql.NullString
	ClassRating       float64
	TotalUser         int
	PackageCode       string
	PackageType       string
	TypeCode          string
	Type              string
	StatusCode        string
	Status            string
	IsExpired         bool
	PromoCode         sql.NullString
	StartDate         time.Time
	ExpiredDate       time.Time
	TimeDuration      int
	Progress          sql.NullFloat64
	ProgressCount     sql.NullInt64
	ProgressIndex     sql.NullInt64
	ProgressSubindex  sql.NullInt64
	PreTestScores     sql.NullFloat64
	PostTestScores    sql.NullFloat64
	PostTestDate      sql.NullTime
	PreTestTotal      sql.NullInt64
	PostTestTotal     sql.NullInt64
	TotalConsultation sql.NullInt64
	TotalWebinar      sql.NullInt64
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	DeletedBy         sql.NullString
	DeletedDate       sql.NullTime
}

type TimeInterval struct {
	Interval1 int
	Interval2 int
}
