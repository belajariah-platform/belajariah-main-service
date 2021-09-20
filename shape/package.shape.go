package shape

import (
	"time"
)

type Package struct {
	ID             int
	Code           string
	Class_Code     string
	Type           string
	Price_Package  string
	Price_Discount string
	Description    string
	Duration       int
	Webinar        int
	Consultation   int
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Deleted_By     string
	Deleted_Date   time.Time
}

type Benefit struct {
	ID            int
	Code          string
	Class_Code    string
	Value         string
	Benefit_Image string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
