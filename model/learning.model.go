package model

import (
	"database/sql"
	"time"
)

type Learning struct {
	ID               int
	Code             string
	ClassCode        string
	Title            string
	DocumentPath     sql.NullString
	DocumentName     sql.NullString
	Sequence         int
	IsExercise       bool
	IsDirectLearning bool
	IsActive         bool
	CreatedBy        string
	CreatedDate      time.Time
	ModifiedBy       sql.NullString
	ModifiedDate     sql.NullTime
	IsDeleted        bool
	Exercises        ExerciseReading
	SubTitles        []SubLearning
}

type SubLearning struct {
	ID            int
	Code          string
	TitleCode     string
	SubTitle      sql.NullString
	VideoDuration sql.NullFloat64
	Video         sql.NullString
	Document      sql.NullString
	Poster        sql.NullString
	Sequence      int
	IsExercise    bool
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	IsDeleted     bool
	Exercises     []Exercise
}

type LearningQuranRequest struct {
	Action string   `form:"action" json:"action" xml:"action"`
	Data   Learning `form:"data" json:"data" xml:"data"`
	Query  Query    `form:"query" json:"query" xml:"query"`
}

type LearningQuran struct {
	ID           int            `form:"id" json:"id" xml:"id" db:"id"`
	Code         string         `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode    string         `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	Title        string         `form:"title" json:"title" xml:"title" db:"title"`
	DocumentCode NullString     `form:"learning_document" json:"learning_document" xml:"learning_document" db:"learning_document"`
	Sequence     int            `form:"sequence" json:"sequence" xml:"sequence" db:"sequence"`
	IsActive     bool           `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy    string         `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate  time.Time      `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy   NullString     `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate NullTime       `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted    bool           `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	ClassBenefit []BenefitQuran `form:"class_benefit" json:"class_benefit" xml:"class_benefit" db:"class_benefit"`
}
