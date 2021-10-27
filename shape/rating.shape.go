package shape

import "time"

type Rating struct {
	ID            int
	Code          string
	Class_Code    string
	Class_Name    string
	Class_Initial string
	User_Code     string
	User_Name     string
	Rating        float64
	Comment       string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Is_Deleted    bool
}

type RatingPost struct {
	ID            int
	Code          string
	Class_Code    string
	Mentor_Code   string
	User_Code     string
	Rating        int
	Comment       string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Is_Deleted    bool
}
