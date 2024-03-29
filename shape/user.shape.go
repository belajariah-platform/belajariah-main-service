package shape

import (
	"time"
)

type Users struct {
	ID                  int
	Code                string
	Role_Code           string
	Email               string
	Password            string
	Old_Password        string
	New_Password        string
	Full_Name           string
	Phone               int64
	Verified_Code       string
	Is_Verified         bool
	Country_Number_Code string
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
	Deleted_By          string
	Deleted_Date        time.Time
}

type UserInfo struct {
	ID                  int
	Code                string
	Role_Code           string
	Role                string
	Email               string
	Full_Name           string
	Phone               int
	Profession          string
	Gender              string
	Age                 int
	Birth               string
	Province            string
	City                string
	Address             string
	Image_Profile       string
	Country_Number_Code string
	Is_Verified         bool
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}

type UsersPost struct {
	ID                  int
	User_Code           string
	Role_Code           string
	Role                string
	Email               string
	Full_Name           string
	Phone               int64
	Profession          string
	Gender              string
	Age                 int
	Birth               time.Time
	Province            string
	City                string
	Address             string
	Image_Code          string
	Image_Filename      string
	Country_Number_Code string
	Is_Verified         bool
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}
