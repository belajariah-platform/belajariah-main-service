package model

import (
	"database/sql"
	"time"
)

type Learning struct {
	ID           int
	Code         string
	ClassCode    string
	Title        string
	SubTitles    []SubLearning
	Exercises    ExerciseReading
	DocumentPath sql.NullString
	DocumentName sql.NullString
	Sequence     sql.NullInt64
	IsExercise   sql.NullBool
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}

type SubLearning struct {
	ID            int
	Code          string
	TitleCode     string
	SubTitle      sql.NullString
	VideoDuration sql.NullFloat64
	Video         sql.NullString
	Document      sql.NullString
	Exercises     []Exercise
	Sequence      sql.NullInt64
	IsExercise    sql.NullBool
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	DeletedBy     sql.NullString
	DeletedDate   sql.NullTime
}
