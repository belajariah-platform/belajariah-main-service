package shape

import (
	"time"
)

type ClassUser struct {
	ID                int
	User_Code         int
	Class_Code        string
	Class_Name        string
	Class_Initial     string
	Class_Category    string
	Class_Description string
	Class_Image       string
	Class_Rating      float64
	Total_User        int
	Status_Code       string
	Status            string
	Is_Expired        bool
	Start_Date        time.Time
	Expired_Date      time.Time
	Time_Duration     int
	Progress          float64
	Pre_Test_Scores   float64
	Post_Test_Scores  float64
	Post_Test_Date    time.Time
	Is_Active         bool
	Created_By        string
	Created_Date      time.Time
	Modified_By       string
	Modified_Date     time.Time
	Deleted_By        string
	Deleted_Date      time.Time
}
