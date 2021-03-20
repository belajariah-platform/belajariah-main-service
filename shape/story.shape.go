package shape

import (
	"time"
)

type Story struct {
	ID            int
	Code          string
	Category_Code string
	Header_Image  string
	Banner_Image  string
	Video_Code    string
	Title         string
	Content       string
	Source        string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
