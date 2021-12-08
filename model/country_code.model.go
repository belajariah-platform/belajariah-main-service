package model

import "time"

type CountryCodeRequest struct {
	Action string      `form:"action" json:"action" xml:"action"`
	Data   CountryCode `form:"data" json:"data" xml:"data"`
	Query  Query       `form:"query" json:"query" xml:"query"`
}

type CountryCode struct {
	ID            int        `form:"id" json:"id" xml:"id" db:"id"`
	Code          string     `form:"code" json:"code" xml:"code" db:"code"`
	Country       string     `form:"country" json:"country" xml:"country" db:"country"`
	Number_Code   string     `form:"number_code" json:"number_code" xml:"number_code" db:"number_code"`
	Flag          NullString `form:"flag" json:"flag" xml:"flag" db:"flag"`
	Is_Active     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
