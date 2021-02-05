package shape

import (
	"time"
)

type Story struct {
	ID            int
	Code          string
	Category_Code string
	Image_Code    string
	Video_Code    string
	Title         string
	Content       string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
