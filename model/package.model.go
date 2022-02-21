package model

import (
	"database/sql"
	"time"
)

type Package struct {
	ID                int
	Code              string
	ClassCode         string
	Type              string
	PricePackage      string
	PriceDiscount     sql.NullString
	Description       sql.NullString
	Duration          int
	DurationFrequence sql.NullInt64
	Webinar           sql.NullInt64
	Consultation      sql.NullInt64
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	IsDeleted         bool
}

type Benefit struct {
	ID           int
	Code         string
	ClassCode    string
	Description  string
	IconBenefit  sql.NullString
	Sequence     int
	IsActive     bool
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   sql.NullString
	ModifiedDate sql.NullTime
	IsDeleted    bool
}

type PackageQuranRequest struct {
	Action string       `form:"action" json:"action" xml:"action"`
	Data   PackageQuran `form:"data" json:"data" xml:"data"`
	Query  Query        `form:"query" json:"query" xml:"query"`
}

type PackageQuran struct {
	ID                int        `form:"id" json:"id" xml:"id" db:"id"`
	Code              string     `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode         string     `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	MentorCode        string     `form:"mentor_code" json:"mentor_code" xml:"mentor_code" db:"mentor_code"`
	Type              string     `form:"type" json:"type" xml:"type" db:"type"`
	PricePackage      string     `form:"price_package" json:"price_package" xml:"price_package" db:"price_package"`
	PriceDiscount     NullString `form:"price_discount" json:"price_discount" xml:"price_discount" db:"price_discount"`
	Description       NullString `form:"description" json:"description" xml:"description" db:"description"`
	Duration          int        `form:"duration" json:"duration" xml:"duration" db:"duration"`
	DurationFrequence NullInt64  `form:"duration_frequence" json:"duration_frequence" xml:"duration_frequence" db:"duration_frequence"`
	IsActive          bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy         string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate       time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy        NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate      NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted         bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	TotalMembers      NullInt64  `form:"total_members" json:"total_members" xml:"total_members" db:"total_members"`
	TotalHours        NullString `form:"total_hours" json:"total_hours" xml:"total_hours" db:"total_hours"`
	AgeRange          NullString `form:"age_range" json:"age_range" xml:"age_range" db:"age_range"`
}

type BenefitQuranRequest struct {
	Action string       `form:"action" json:"action" xml:"action"`
	Data   BenefitQuran `form:"data" json:"data" xml:"data"`
	Query  Query        `form:"query" json:"query" xml:"query"`
}

type BenefitQuran struct {
	ID           int        `form:"id" json:"id" xml:"id" db:"id"`
	Code         string     `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode    string     `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	IconBenefit  NullString `form:"icon_benefit" json:"icon_benefit" xml:"icon_benefit" db:"icon_benefit"`
	Sequence     int        `form:"sequence" json:"sequence" xml:"sequence" db:"sequence"`
	Description  NullString `form:"description" json:"description" xml:"description" db:"description"`
	IsActive     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}

type TermConditionQuranRequest struct {
	Action string             `form:"action" json:"action" xml:"action"`
	Data   TermConditionQuran `form:"data" json:"data" xml:"data"`
	Query  Query              `form:"query" json:"query" xml:"query"`
}

type TermConditionQuran struct {
	ID           int        `form:"id" json:"id" xml:"id" db:"id"`
	Code         string     `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode    string     `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	IconTerm     NullString `form:"icon_term" json:"icon_term" xml:"icon_term" db:"icon_term"`
	Sequence     int        `form:"sequence" json:"sequence" xml:"sequence" db:"sequence"`
	Description  NullString `form:"description" json:"description" xml:"description" db:"description"`
	IsActive     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
