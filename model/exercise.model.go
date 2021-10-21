package model

import (
	"database/sql"
	"time"
)

type Exercise struct {
	ID            int
	Code          string
	SubtitleCode  string
	ImageExercise sql.NullString
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	IsDeleted     bool
}

type ExerciseReading struct {
	ID           int
	Code         string
	TitleCode    string
	SuratCode    int
	AyatStart    int
	AyatEnd      int
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}

type UserExerciseReading struct {
	ID            int
	UserCode      int
	ClassCode     string
	RecordingCode int
	Duration      int
	ExpiredDate   time.Time
	TitleCode     sql.NullString
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	IsDeleted     bool
}
