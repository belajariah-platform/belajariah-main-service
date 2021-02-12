package shape

import (
	"time"
)

type Learning struct {
	ID            int
	Code          string
	Class_Code    string
	Title         string
	SubTitles     []SubLearning
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
	ID             int
	Code           string
	Title_Code     string
	Sub_Title      string
	Video_Duration float64
	Video          string
	Document       string
	Exercise_Image string
	Exercises      []Exercise
	Sequence       int
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Deleted_By     string
	Deleted_Date   time.Time
}
