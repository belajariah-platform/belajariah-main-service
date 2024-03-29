package shape

import (
	"time"
)

type Consultation struct {
	ID                 int
	User_Code          int
	User_Name          string
	User_Image         string
	Class_Code         string
	Class_Name         string
	Recording_Code     int
	Recording_Path     string
	Recording_Name     string
	Recording_Duration int
	Status_Code        string
	Status             string
	Description        string
	Taken_Code         int
	Taken_Name         string
	Is_Play            bool
	Is_Read            bool
	Is_Action_Taken    bool
	Is_Active          bool
	Expired_Date       string
	Created_By         string
	Created_Date       time.Time
	Modified_By        string
	Modified_Date      time.Time
	Deleted_By         string
	Deleted_Date       time.Time
}

type ConsultationPost struct {
	ID                 int
	User_Code          int
	Action             string
	User_Name          string
	Class_Code         string
	Class_Name         string
	Recording_Code     int64
	Recording_Path     string
	Recording_Name     string
	Recording_Duration int64
	Status_Code        string
	Status             string
	Description        string
	Taken_Code         int64
	Taken_Name         string
	Is_Play            bool
	Is_Read            bool
	Is_Action_Taken    bool
	Is_Active          bool
	Expired_Date       time.Time
	Created_By         string
	Created_Date       time.Time
	Modified_By        string
	Modified_Date      time.Time
	Deleted_By         string
	Deleted_Date       time.Time
}
