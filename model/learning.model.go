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
