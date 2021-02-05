package model

import (
	"database/sql"
	"time"
)

type Class struct {
	ID                    int
	Code                  string
	ClassCategoryCode     string
	ClassCategory         string
	ClassName             string
	ClassInitial          sql.NullString
	ClassDescription      sql.NullString
	ClassImage            sql.NullString
	ClassVideo            sql.NullString
	ClassRating           float64
	TotalReview           int
	InstructorName        string
	InstructorDescription sql.NullString
	InstructorBiografi    sql.NullString
	InstructorImage       sql.NullString
	IsActive              bool
	CreatedBy             string
	CreatedDate           time.Time
	ModifiedBy            string
	ModifiedDate          time.Time
	DeletedBy             string
	DeletedDate           time.Time
}
