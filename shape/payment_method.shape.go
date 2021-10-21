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
	Icon_Account   string
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Is_Deleted     bool
}
