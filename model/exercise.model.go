package model

import (
	"database/sql"
	"time"
)

type Exercise struct {
	ID            int
	Code          string
	SubtitleCode  string
	ImageCode     int
	ExerciseImage string
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	DeletedBy     sql.NullString
	DeletedDate   sql.NullTime
}

type ExerciseReading struct {
	ID           int
	Code         string
	TitleCode    string
	SuratCode    string
	AyatStart    int
	AyatEnd      int
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	DeletedBy    sql.NullString
	DeletedDate  sql.NullTime
}

type UserExerciseReading struct {
	ID            int
	UserCode      int
	ClassCode     string
	RecordingCode int
	Duration      int
	ExpiredDate   string
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    sql.NullString
	ModifiedDate  sql.NullTime
	DeletedBy     sql.NullString
	DeletedDate   sql.NullTime
}
