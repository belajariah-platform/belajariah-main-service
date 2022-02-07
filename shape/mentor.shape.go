package shape

import (
	"time"
)

type Mentors struct {
	ID            int
	Code          string
	Role_Code     string
	Email         string
	Password      string
	Old_Password  string
	New_Password  string
	Full_Name     string
	Phone         int64
	Verified_Code string
	Is_Verified   bool
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}

type MentorInfo struct {
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
	Minimum_Rate         int
	Allow_Contact_From   string
	Country_Number_Code  string
	Is_Active            bool
	Created_By           string
	Created_Date         time.Time
	Modified_By          string
	Modified_Date        time.Time
	Mentor_Schedule      []MentorSchedule
	Mentor_Experience    []MentorExperience
	Mentor_Class         []MentorClass
	Mentor_Package       []MentorPackage
}

type MentorSchedule struct {
	ID            int
	Code          string
	Mentor_Code   string
	Class_Code    string
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
	ID              int
	Code            string
	Mentor_Code     string
	Experience      string
	Experience_Type string
	Is_Active       bool
	Created_By      string
	Created_Date    time.Time
	Modified_By     string
	Modified_Date   time.Time
	Is_Deleted      bool
}

type MentorClass struct {
	ID                 int
	Code               string
	Mentor_Code        string
	Mentor_Name        string
	Class_Code         string
	Class_Name         string
	Class_Initial      string
	Minimum_Rate       int
	Allow_Contact_From string
	Is_Active          bool
	Created_By         string
	Created_Date       time.Time
	Modified_By        string
	Modified_Date      time.Time
	Is_Deleted         bool
}

type MentorPackage struct {
	ID                 int
	Code               string
	Class_Code         string
	Mentor_Code        string
	Type               string
	Price_Package      string
	Price_Discount     string
	Description        string
	Duration           int
	Duration_Frequence int
	Is_Active          bool
	Created_By         string
	Created_Date       time.Time
	Modified_By        string
	Modified_Date      time.Time
	Is_Deleted         bool
}
