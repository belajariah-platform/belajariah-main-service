package shape

import (
	"time"
)

type PaymentMethod struct {
	ID             int
	Code           string
	Type           string
	Value          string
	Account_Name   string
	Account_Number string
	Method_Image   string
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Deleted_By     string
	Deleted_Date   time.Time
}
