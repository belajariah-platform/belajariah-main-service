package model

import (
	"database/sql"
	"time"
)

type Quran struct {
	ID             int
	Code           string
	SuratCode      string
	SuratName      string
	SuratText      string
	SuratTranslate string
	CountAyat      int
	AyatNumber     int
	AyatText       string
	AyatTranslate  string
	JuzNumber      int
	PageNumber     int
	IsActive       bool
	CreatedBy      string
	CreatedDate    time.Time
	ModifiedBy     sql.NullString
	ModifiedDate   sql.NullTime
	DeletedBy      sql.NullString
	DeletedDate    sql.NullTime
}
