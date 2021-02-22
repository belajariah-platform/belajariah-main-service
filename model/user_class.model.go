package model

import (
	"database/sql"
	"time"
)

type UserClass struct {
	ID                  int
	UserCode            int
	ClassCode           string
	ClassName           string
	ClassInitial        sql.NullString
	ClassCategory       string
	ClassDescription    sql.NullString
	ClassImage          sql.NullString
	ClassRating         float64
	TotalUser           int
	PackageCode         string
	TypeCode            string
	Type                string
	StatusCode          string
	Status              string
	IsExpired           bool
	PromoCode           sql.NullString
	StartDate           time.Time
	ExpiredDate         time.Time
	TimeDuration        int
	Progress            sql.NullFloat64
	ProgressIndex       sql.NullInt64
	ProgressCurIndex    sql.NullInt64
	ProgressCurSubindex sql.NullInt64
	PreTestScores       sql.NullFloat64
	PostTestScores      sql.NullFloat64
	PostTestDate        sql.NullTime
	IsActive            bool
	CreatedBy           string
	CreatedDate         time.Time
	ModifiedBy          sql.NullString
	ModifiedDate        sql.NullTime
	DeletedBy           sql.NullString
	DeletedDate         sql.NullTime
}
