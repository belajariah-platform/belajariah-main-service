package shape

import (
	"time"
)

type Promotion struct {
	ID              int
	Code            string
	Class_Code      string
	Title           string
	Description     string
	Promo_Code      string
	Promo_Type_Code string
	Promo_Type      string
	Discount        float64
	Banner_Image    string
	Header_Image    string
	Expired_Date    string
	Quota_User      int
	Quota_Used      int
	Is_Active       bool
	Created_By      string
	Created_Date    time.Time
	Modified_By     string
	Modified_Date   time.Time
	Deleted_By      string
	Deleted_Date    time.Time
}

type PromotionClaim struct {
	User_Code  int
	Class_Code string
	Promo_Code string
}
