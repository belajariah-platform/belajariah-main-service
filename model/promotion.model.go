package model

import (
	"time"
)

type PromotionRequest struct {
	Action string    `form:"action" json:"action" xml:"action"`
	Data   Promotion `form:"data" json:"data" xml:"data"`
	Query  Query     `form:"query" json:"query" xml:"query"`
}
type Promotion struct {
	ID            int       `form:"id" json:"id" xml:"id" db:"id"`
	Code          string    `form:"code" json:"code" xml:"code" db:"code"`
	PromoLevel    string    `form:"promo_level" json:"promo_level" xml:"promo_level" db:"promo_level"`
	ClassCode     string    `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	PackageCode   string    `form:"package_code" json:"package_code" xml:"package_code" db:"package_code"`
	Title         string    `form:"title" json:"title" xml:"title" db:"title"`
	Description   string    `form:"description" json:"description" xml:"description" db:"description"`
	PromoCode     string    `form:"promo_code" json:"promo_code" xml:"promo_code" db:"promo_code"`
	PromoTypeCode string    `form:"promo_type_code" json:"promo_type_code" xml:"promo_type_code" db:"promo_type_code"`
	PromoType     string    `form:"promo_type" json:"promo_type" xml:"promo_type" db:"promo_type"`
	Discount      float64   `form:"discount" json:"discount" xml:"discount" db:"discount"`
	ImageBanner   string    `form:"image_banner" json:"image_banner" xml:"image_banner" db:"image_banner"`
	ImageHeader   string    `form:"image_header" json:"image_header" xml:"image_header" db:"image_header"`
	ExpiredDate   time.Time `form:"expired_date" json:"expired_date" xml:"expired_date" db:"expired_date"`
	QuotaUser     int       `form:"quota_user" json:"quota_user" xml:"quota_user" db:"quota_user"`
	QuotaUsed     int       `form:"quota_used" json:"quota_used" xml:"quota_used" db:"quota_used"`
	IsActive      bool      `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy     string    `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate   time.Time `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy    string    `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate  time.Time `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted     bool      `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
