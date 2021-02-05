package shape

import (
	"time"
)

type Payment struct {
	ID                  int
	User_Code           string
	User_Name           string
	Class_Code          string
	Class_Name          string
	Class_Initial       string
	Payment_Method_Code string
	Payment_Method      string
	Invoice_Number      string
	Status_Payment_Code string
	Status_Payment      string
	Total_Transfer      int
	Sender_Bank         string
	Destination_Account string
	Image_Proof         string
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
}
