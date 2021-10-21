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
	Image_Banner    string
	Image_Header    string
	Expired_Date    string
	Quota_User      int
	Quota_Used      int
	Is_Active       bool
	Created_By      string
	Created_Date    time.Time
	Modified_By     string
	Modified_Date   time.Time
	Is_Deleted      bool
}

type PromotionClaim struct {
	User_Code  int
	Class_Code string
	Promo_Code string
}
