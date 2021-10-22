package shape

import (
	"time"
)

type Mentor struct {
	ID                   int
	Code                 string
	Mentor_Code          int
	Role_Code            string
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
	Birth                time.Time
	ImageProfile         string
	Description          string
	Account_Owner        string
	Account_Name         string
	Account_Number       string
	Learning_Method      string
	Learning_Method_Text string
	Rating               float64
	Is_Active            bool
	Created_By           string
	Created_Date         time.Time
	Modified_By          string
	Modified_Date        time.Time
	Mentor_Schedule      []MentorSchedule
	Mentor_Experience    []MentorExperience
}

type MentorSchedule struct {
	ID            int
	Code          string
	Mentor_Code   string
	Shift_Name    string
	Start_Date    time.Time
	End_Date      time.Time
	Time_Zone     string
	Sequence      int
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Is_Deleted    bool
}

type MentorExperience struct {
	ID            int
	Code          string
	Mentor_Code   string
	Experience    string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Is_Deleted    bool
}
