package shape

import (
	"time"
)

type Payment struct {
	ID                  int
	User_Code           int
	User_Name           string
	Class_Code          string
	Class_Name          string
	Class_Initial       string
	Package_Code        string
	Package_Type        string
	Payment_Method_Code string
	Payment_Method      string
	Invoice_Number      string
	Status_Payment_Code string
	Status_Payment      string
	Total_Transfer      int
	Sender_Bank         string
	Sender_Name         string
	Image_Proof         string
	Payment_Type_Code   string
	Payment_Type        string
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}

type PaymentPost struct {
	ID                  int
	User_Code           int
	User_Name           string
	Class_Code          string
	Class_Name          string
	Class_Initial       string
	Package_Code        string
	Package_Type        string
	Payment_Method_Code string
	Payment_Method      string
	Invoice_Number      string
	Status_Payment_Code string
	Status_Payment      string
	Total_Transfer      int
	Sender_Bank         string
	Sender_Name         string
	Image_Code          int64
	Image_Proof         string
	Action              string
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}
