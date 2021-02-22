package shape

import (
	"time"
)

type Notification struct {
	ID                int
	Code              string
	Code_Ref          string
	Table_Ref         string
	Notification_Type string
	User_Code         int
	Sequence          int
	Is_Read           bool
	Is_Action_Taken   bool
	Is_Active         bool
	Created_By        string
	Created_Date      time.Time
	Modified_By       string
	Modified_Date     time.Time
	Deleted_By        string
	Deleted_Date      time.Time
}
