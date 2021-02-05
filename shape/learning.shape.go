package shape

import (
	"time"
)

type Learning struct {
	ID            int
	Code          string
	Class_Code    string
	Title         string
	Document      string
	Sequence      int
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}

type SubLearning struct {
	ID            int
	Code          string
	TitleCode     string
	SubTitle      string
	VideoDuration float64
	Video         string
	Document      string
	Sequence      int
	IsActive      bool
	CreatedBy     string
	CreatedDate   time.Time
	ModifiedBy    string
	ModifiedDate  time.Time
	DeletedBy     string
	DeletedDate   time.Time
}
