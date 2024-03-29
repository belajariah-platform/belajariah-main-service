package shape

import (
	"time"
)

type PaymentnRequest struct {
	Action string      `form:"action" json:"action" xml:"action"`
	Data   PaymentPost `form:"data" json:"data" xml:"data"`
}

type Payment struct {
	ID                   int
	Code                 string
	User_Code            string
	User_Name            string
	Class_Code           string
	Class_Name           string
	Class_Initial        string
	Package_Code         string
	Package_Type         string
	Payment_Method_Code  string
	Payment_Method       string
	Payment_Method_Type  string
	Payment_Method_Image string
	Account_Name         string
	Account_Number       string
	Invoice_Number       string
	Status_Payment_Code  string
	Status_Payment       string
	Total_Transfer       int
	Sender_Bank          string
	Sender_Name          string
	Image_Proof          string
	Payment_Type_Code    string
	Payment_Type         string
	Expired_Date         string
	Schedule_Code_1      string
	Schedule_Code_2      string
	Payment_Reference    string
	Is_Active            bool
	Created_By           string
	Created_Date         time.Time
	Modified_By          string
	Modified_Date        time.Time
}

type PaymentPost struct {
	ID                  int
	Code                string
	User_Code           string
	User_Name           string
	Class_Code          string
	Class_Name          string
	Class_Initial       string
	Promo_Code          string
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
	Remarks             string
	Schedule_Code_1     string
	Schedule_Code_2     string
	Is_Direct           bool
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}
