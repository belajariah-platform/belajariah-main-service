package model

import (
	"database/sql"
	"time"
)

type ClassUser struct {
	ID               int
	UserCode         int
	ClassCode        string
	ClassName        string
	ClassInitial     sql.NullString
	ClassCategory    string
	ClassDescription sql.NullString
	ClassImage       sql.NullString
	ClassRating      float64
	TotalUser        int
	StatusCode       string
	Status           string
	IsExpired        bool
	StartDate        time.Time
	ExpiredDate      time.Time
	TimeDuration     int
	Progress         sql.NullFloat64
	PreTestScores    sql.NullFloat64
	PostTestScores   sql.NullFloat64
	PostTestDate     sql.NullTime
	IsActive         bool
	CreatedBy        string
	CreatedDate      time.Time
	ModifiedBy       string
	ModifiedDate     time.Time
	DeletedBy        string
	DeletedDate      time.Time
}
