package model

import (
	"database/sql"
	"time"
)

type UserClassRequest struct {
	Action string    `form:"action" json:"action" xml:"action"`
	Data   UserClass `form:"data" json:"data" xml:"data"`
	Query  Query     `form:"query" json:"query" xml:"query"`
}

type UserClass struct {
	ID                int
	Code              string
	UserCode          string
	ClassCode         string
	ClassName         string
	ClassInitial      sql.NullString
	ClassCategory     string
	ClassDescription  sql.NullString
	ClassImage        sql.NullString
	ClassRating       float64
	TotalUser         int
	PackageCode       string
	PackageType       string
	TypeCode          string
	Type              string
	StatusCode        string
	Status            string
	IsExpired         bool
	PromoCode         sql.NullString
	StartDate         sql.NullTime
	ExpiredDate       sql.NullTime
	TimeDuration      int
	Progress          sql.NullFloat64
	ProgressCount     sql.NullInt64
	ProgressIndex     sql.NullInt64
	ProgressSubindex  sql.NullInt64
	PreTestScores     sql.NullFloat64
	PostTestScores    sql.NullFloat64
	PostTestDate      sql.NullTime
	PreTestTotal      sql.NullInt64
	PostTestTotal     sql.NullInt64
	TotalConsultation sql.NullInt64
	TotalWebinar      sql.NullInt64
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	IsDeleted         bool
}

type TimeInterval struct {
	Interval1 int
	Interval2 int
}

type UserClassHistory struct {
	ID                int
	Code              string
	UserClassCode     string
	PackageCode       string
	PaymentMethodCode string
	PromoCode         string
	Price             int
	StartDate         sql.NullTime
	ExpiredDate       sql.NullTime
	IsActive          bool
	CreatedBy         string
	CreatedDate       time.Time
	ModifiedBy        sql.NullString
	ModifiedDate      sql.NullTime
	IsDeleted         bool
}

type UserClassQuranRequest struct {
	Action string     `form:"action" json:"action" xml:"action"`
	Data   ClassQuran `form:"data" json:"data" xml:"data"`
	Query  Query      `form:"query" json:"query" xml:"query"`
}

type UserClassQuran struct {
	ID           int        `form:"id" json:"id" xml:"id" db:"id"`
	Code         string     `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode    string     `form:"classs_code" json:"classs_code" xml:"classs_code" db:"classs_code"`
	user_code    string     `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	PackageCode  string     `form:"package_code" json:"package_code" xml:"package_code" db:"package_code"`
	PromoCode    NullString `form:"promo_code" json:"promo_code" xml:"promo_code" db:"promo_code"`
	IsActive     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
