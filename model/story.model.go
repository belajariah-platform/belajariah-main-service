package model

import (
	"database/sql"
	"time"
)

type Story struct {
	ID                int
	Code              string
	StoryCategoryCode string
	ImageBannerStory  sql.NullString
	ImageHeaderStory  sql.NullString
	VideoStory        sql.NullString
	Title             string
	Content           string
	Source            sql.NullString
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	IsDeleted         bool
}
