package shape

import (
	"time"
)

type Promotion struct {
	ID            int
	Code          string
	Class_Code    string
	Title         string
	Description   string
	Promo_Code    string
	Discount      float64
	Banner_Image  string
	Header_Image  string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
