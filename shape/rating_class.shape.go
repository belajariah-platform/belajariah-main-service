package shape

import "time"

type RatingClass struct {
	ID            int
	Class_Code    string
	Class_Name    string
	Class_Initial string
	User_Code     int
	Rating        float64
	Comment       string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
