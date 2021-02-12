package shape

import (
	"time"
)

type Mentor struct {
	ID              int
	Role_Code       string
	Role            string
	Email           string
	Full_Name       string
	Phone           int
	Profession      string
	Gender          string
	Age             int
	Province        string
	City            string
	Address         string
	Image_Code      string
	Image_Filename  string
	Image_Filepath  string
	Rating          float64
	Task_Completed  int
	Task_Inprogress int
	Is_Active       bool
	Created_By      string
	Created_Date    time.Time
	Modified_By     string
	Modified_Date   time.Time
}
