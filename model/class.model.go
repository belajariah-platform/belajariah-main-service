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

type ClassQuranRequest struct {
	Action string     `form:"action" json:"action" xml:"action"`
	Data   ClassQuran `form:"data" json:"data" xml:"data"`
	Query  Query      `form:"query" json:"query" xml:"query"`
}

type ClassQuran struct {
	ID                int            `form:"id" json:"id" xml:"id" db:"id"`
	Code              string         `form:"code" json:"code" xml:"code" db:"code"`
	ClassCategoryCode string         `form:"class_category_code" json:"class_category_code" xml:"class_category_code" db:"class_category_code"`
	ClassCategory     string         `form:"class_category" json:"class_category" xml:"class_category" db:"class_category"`
	ClassName         string         `form:"class_name" json:"class_name" xml:"class_name" db:"class_name"`
	ClassInitial      sql.NullString `form:"class_initial" json:"class_initial" xml:"class_initial" db:"class_initial"`
	ClassDescription  sql.NullString `form:"class_description" json:"class_description" xml:"class_description" db:"class_description"`
	ClassImage        sql.NullString `form:"class_image" json:"class_image" xml:"class_image" db:"class_image"`
	ClassVideo        sql.NullString `form:"class_video" json:"class_video" xml:"class_video" db:"class_video"`
	ClassDocument     sql.NullString `form:"class_document" json:"class_document" xml:"class_document" db:"class_document"`
	IsDirect          bool           `form:"is_direct" json:"is_direct" xml:"is_direct" db:"is_direct"`
	IsActive          bool           `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy         string         `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate       time.Time      `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy        sql.NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate      sql.NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted         bool           `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
