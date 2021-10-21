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
	ClassDocument         sql.NullString
	ClassRating           float64
	TotalReview           int
	TotalVideo            int
	TotalVideoDuration    int
	InstructorName        sql.NullString
	InstructorDescription sql.NullString
	InstructorBiografi    sql.NullString
	InstructorImage       sql.NullString
	IsDirect              bool
	IsActive              bool
	CreatedBy             string
	CreatedDate           time.Time
	ModifiedBy            sql.NullString
	ModifiedDate          sql.NullTime
	IsDeleted             bool
}
