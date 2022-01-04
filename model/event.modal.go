package model

import "time"

type EventRequest struct {
	Action string `form:"action" json:"action" xml:"action"`
	Data   Event  `form:"data" json:"data" xml:"data"`
	Query  Query  `form:"query" json:"query" xml:"query"`
}

type Event struct {
	ID              int                `form:"id" json:"id" xml:"id" db:"id"`
	Code            string             `form:"code" json:"code" xml:"code" db:"code"`
	Event_Name      string             `form:"event_name" json:"event_name" xml:"event_name" db:"event_name"`
	Event_Type      string             `form:"event_type" json:"event_type" xml:"event_type" db:"event_type"`
	Event_Type_Desc string             `form:"event_type_desc" json:"event_type_desc" xml:"event_type_desc" db:"event_type_desc"`
	Event_Image     NullString         `form:"event_image" json:"event_image" xml:"event_image" db:"event_image"`
	Is_Active       bool               `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By      string             `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date    time.Time          `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By     NullString         `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date   NullTime           `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted      bool               `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	User_Code       string             `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	EventFormDetail []EventMappingForm `form:"event_form_detail" json:"event_form_detail" xml:"event_form_detail"`
}

type EventMappingForm struct {
	ID              int        `form:"id" json:"id" xml:"id" db:"id"`
	Code            string     `form:"code" json:"code" xml:"code" db:"code"`
	Event_Code      string     `form:"event_code" json:"event_code" xml:"event_code" db:"event_code"`
	Event_Name      string     `form:"event_name" json:"event_name" xml:"event_name" db:"event_name"`
	Event_Type      string     `form:"event_type" json:"event_type" xml:"event_type" db:"event_type"`
	Event_Form_Code string     `form:"event_form_code" json:"event_form_code" xml:"event_form_code" db:"event_form_code"`
	Question        string     `form:"question" json:"question" xml:"question" db:"question"`
	Question_Type   string     `form:"question_type" json:"question_type" xml:"question_type" db:"question_type"`
	Choice          NullString `form:"choice" json:"choice" xml:"choice" db:"choice"`
	Answer          NullString `form:"answer" json:"answer" xml:"answer" db:"answer"`
	Is_Required     string     `form:"is_required" json:"is_required" xml:"is_required" db:"is_required"`
	Sort            string     `form:"sort" json:"sort" xml:"sort" db:"sort"`
	Is_Active       bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By      string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date    time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By     NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date   NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted      bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
	User_Code       string     `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
}
