package shape

import (
	"time"
)

type Exercise struct {
	ID            int
	Code          string
	Subtitle_Code string
	Option        string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
