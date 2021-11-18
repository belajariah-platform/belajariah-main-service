package model

import "time"

type ScheduleRequest struct {
	Action string   `form:"action" json:"action" xml:"action"`
	Data   Schedule `form:"data" json:"data" xml:"data"`
	Query  Query    `form:"query" json:"query" xml:"query"`
}

type Schedule struct {
	ID                       int        `form:"id" json:"id" xml:"id" db:"id"`
	Code                     string     `form:"code" json:"code" xml:"code" db:"code"`
	User_Class_Code          string     `form:"user_class_code" json:"user_class_code" xml:"user_class_code" db:"user_class_code"`
	Class_Code               string     `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	Class_Name               string     `form:"class_name" json:"class_name" xml:"class_name" db:"class_name"`
	User_Code                string     `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	User_Name                string     `form:"user_name" json:"user_name" xml:"user_name" db:"user_name"`
	Mentor_Code              string     `form:"mentor_code" json:"mentor_code" xml:"mentor_code" db:"mentor_code"`
	Mentor_Name              string     `form:"mentor_name" json:"mentor_name" xml:"mentor_name" db:"mentor_name"`
	Description              NullString `form:"description" json:"description" xml:"description" db:"description"`
	Shift_Name               string     `form:"shift_name" json:"shift_name" xml:"shift_name" db:"shift_name"`
	Planning_Start_Time      NullTime   `form:"planning_start_time" json:"planning_start_time" xml:"planning_start_time" db:"planning_start_time"`
	Planning_End_Time        NullTime   `form:"planning_end_time" json:"planning_end_time" xml:"planning_end_time" db:"planning_end_time"`
	Actual_User_Start_Date   NullTime   `form:"actual_user_start_date" json:"actual_user_start_date" xml:"actual_user_start_date" db:"actual_user_start_date"`
	Actual_User_End_Date     NullTime   `form:"actual_user_end_date" json:"actual_user_end_date" xml:"actual_user_end_date" db:"actual_user_end_date"`
	Actual_Mentor_Start_Date NullTime   `form:"actual_mentor_start_date" json:"actual_mentor_start_date" xml:"actual_mentor_start_date" db:"actual_mentor_start_date"`
	Actual_Mentor_End_Date   NullTime   `form:"actual_mentor_end_date" json:"actual_mentor_end_date" xml:"actual_mentor_end_date" db:"actual_mentor_end_date"`
	Sequence                 int        `form:"sequence" json:"sequence" xml:"sequence" db:"sequence"`
	Is_Completed_User        bool       `form:"is_completed_user" json:"is_completed_user" xml:"is_completed_user" db:"is_completed_user"`
	Is_Completed_Mentor      bool       `form:"is_completed_mentor" json:"is_completed_mentor" xml:"is_completed_mentor" db:"is_completed_mentor"`
	Is_Active                bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By               string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date             time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By              string     `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date            NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted               bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
