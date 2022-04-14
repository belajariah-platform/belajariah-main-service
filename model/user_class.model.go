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
	Action       string                 `form:"action" json:"action" xml:"action"`
	Data         UserClassQuran         `form:"data" json:"data" xml:"data"`
	DataSchedule UserClassQuranSchedule `form:"data_schedule" json:"data_schedule" xml:"data_schedule"`
	Query        Query                  `form:"query" json:"query" xml:"query"`
}

type UserClassQuran struct {
	ID                   int                    `form:"id" json:"id" xml:"id" db:"id"`
	Code                 string                 `form:"code" json:"code" xml:"code" db:"code"`
	ClassCode            string                 `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	ClassName            string                 `form:"class_name" json:"class_name" xml:"class_name" db:"class_name"`
	ClassInitial         NullString             `form:"class_initial" json:"class_initial" xml:"class_initial" db:"class_initial"`
	ClassCategory        string                 `form:"class_category" json:"class_category" xml:"class_category" db:"class_category"`
	ClassDescription     NullString             `form:"class_description" json:"class_description" xml:"class_description" db:"class_description"`
	ClassImage           NullString             `form:"class_image" json:"class_image" xml:"class_image" db:"class_image"`
	ColorPath            NullString             `form:"color_path" json:"color_path" xml:"color_path" db:"color_path"`
	UserCode             string                 `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	PackageCode          NullString             `form:"package_code" json:"package_code" xml:"package_code" db:"package_code"`
	PackageType          NullString             `form:"package_type" json:"package_type" xml:"package_type" db:"package_type"`
	IsActive             bool                   `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy            string                 `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate          time.Time              `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy           NullString             `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate         NullTime               `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted            bool                   `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	UserClassQuranDetail []UserClassQuranDetail `form:"user_class_quran_detail" json:"user_class_quran_detail" xml:"user_class_quran_detail" db:"user_class_quran_detail"`
}

type UserClassQuranDetail struct {
	ID                     int                      `form:"id" json:"id" xml:"id" db:"id"`
	Code                   string                   `form:"code" json:"code" xml:"code" db:"code"`
	UserClassCode          string                   `form:"user_class_code" json:"user_class_code" xml:"user_class_code" db:"user_class_code"`
	UserName               string                   `form:"user_name" json:"user_name" xml:"user_name" db:"user_name"`
	PackageCode            string                   `form:"package_code" json:"package_code" xml:"package_code" db:"package_code"`
	PackageType            string                   `form:"package_type" json:"package_type" xml:"package_type" db:"package_type"`
	MentorCode             string                   `form:"mentor_code" json:"mentor_code" xml:"mentor_code" db:"mentor_code"`
	MentorName             NullString               `form:"mentor_name" json:"mentor_name" xml:"mentor_name" db:"mentor_name"`
	MentorImage            NullString               `form:"mentor_image" json:"mentor_image" xml:"mentor_image" db:"mentor_image"`
	MentorCity             NullString               `form:"mentor_city" json:"mentor_city" xml:"mentor_city" db:"mentor_city"`
	MentorPhone            NullInt64                `form:"mentor_phone" json:"mentor_phone" xml:"mentor_phone" db:"phone"`
	IsCompleted            bool                     `form:"is_completed" json:"is_completed" xml:"is_completed" db:"is_completed"`
	IsActive               bool                     `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy              string                   `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate            time.Time                `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy             NullString               `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate           NullTime                 `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted              bool                     `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	UserClassQuranSchedule []UserClassQuranSchedule `form:"user_class_quran_schedule" json:"user_class_quran_schedule" xml:"user_class_quran_schedule" db:"user_class_quran_schedule"`
}

type UserClassQuranSchedule struct {
	ID                  int         `form:"id" json:"id" xml:"id" db:"id"`
	Code                string      `form:"code" json:"code" xml:"code" db:"code"`
	UserClassDetailCode string      `form:"user_class_detail_code" json:"user_class_detail_code" xml:"user_class_detail_code" db:"user_class_detail_code"`
	StartDate           NullTime    `form:"start_date" json:"start_date" xml:"start_date" db:"start_date"`
	FinishDate          NullTime    `form:"finish_date" json:"finish_date" xml:"finish_date" db:"finish_date"`
	UserMessage         NullString  `form:"user_message" json:"user_message" xml:"user_message" db:"user_message"`
	MentorMessage       NullString  `form:"mentor_message" json:"mentor_message" xml:"mentor_message" db:"mentor_message"`
	Sequence            int         `form:"sequence" json:"sequence" xml:"sequence" db:"sequence"`
	IsActive            bool        `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	CreatedBy           string      `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	CreatedDate         time.Time   `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	ModifiedBy          NullString  `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	ModifiedDate        NullTime    `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	IsDeleted           bool        `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	Material            NullString  `form:"material" json:"material" xml:"material" db:"material"`
	UserScore           NullFloat64 `form:"user_score" json:"user_score" xml:"user_score" db:"user_score"`
}
