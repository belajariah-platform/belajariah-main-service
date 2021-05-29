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
	Exercises     ExerciseReading
	Document_Path string
	Document_Name string
	Sequence      int
	Is_Exercise   bool
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
	Poster         string
	Sequence       int
	Is_Done        bool
	Is_Exercise    bool
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Deleted_By     string
	Deleted_Date   time.Time
}
