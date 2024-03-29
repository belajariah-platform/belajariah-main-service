package shape

import (
	"time"
)

type Exercise struct {
	ID             int
	Code           string
	Subtitle_Code  string
	Image_Exercise string
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Is_Deleted     bool
}

type ExerciseReading struct {
	ID            int
	Code          string
	Title_Code    string
	Surat_Code    int
	Ayat_Start    int
	Ayat_End      int
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Is_Deleted    bool
}

type UserExerciseReading struct {
	ID             int
	User_Code      int
	Class_Code     string
	Recording_Code int
	Duration       int
	Expired_Date   time.Time
	Title_Code     string
	Max_Recording  int
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Is_Deleted     bool
}
