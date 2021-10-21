package shape

import (
	"time"
)

type Mentor struct {
	ID                   int
	Role_Code            string
	Class_Code           string
	Mentor_Code          int
	Role                 string
	Email                string
	Full_Name            string
	Phone                int
	Profession           string
	Gender               string
	Age                  int
	Province             string
	City                 string
	Address              string
	Description          string
	Image_Code           string
	Image_Filename       string
	Image_Filepath       string
	Rating               float64
	Learning_Method      string
	Learning_Method_Text string
	Task_Completed       int
	Task_Inprogress      int
	Is_Active            bool
	Created_By           string
	Created_Date         time.Time
	Modified_By          string
	Modified_Date        time.Time
	Schedule             []MentorSchedule
}

type MentorSchedule struct {
	ID            int
	Mentor_Code   int
	Shift_Name    string
	Start_At      time.Time
	End_At        time.Time
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Time_Zone     string
}
