package model

import (
	"time"
)

type CoachingProgramRequest struct {
	Action   string                `form:"action" json:"action" xml:"action"`
	Data_Mcp MasterCoachingProgram `form:"data_mcp" json:"data_mcp" xml:"data_mcp"`
	Data_TCP CoachingProgram       `form:"data" json:"data" xml:"data"`
	Query    Query                 `form:"query" json:"query" xml:"query"`
}

type MasterCoachingProgram struct {
	ID            int        `form:"id" json:"id" xml:"id" db:"id"`
	Code          string     `form:"code" json:"code" xml:"code" db:"code"`
	Title         string     `form:"title" json:"title" xml:"title" db:"title"`
	Description   NullString `form:"description" json:"description" xml:"description" db:"description"`
	Expired_Date  NullTime   `form:"expired_date" json:"expired_date" xml:"expired_date" db:"expired_date"`
	Quota_User    NullInt64  `form:"quota_user" json:"quota_user" xml:"quota_user" db:"quota_user"`
	Image_Header  NullString `form:"image_header" json:"image_header" xml:"image_header" db:"image_header"`
	Image_Banner  NullString `form:"image_banner" json:"image_banner" xml:"image_banner" db:"image_banner"`
	Is_Active     bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By    string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date  time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By   NullString `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted    bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}

type CoachingProgram struct {
	ID                  int        `form:"id" json:"id" xml:"id" db:"id"`
	Code                string     `form:"code" json:"code" xml:"code" db:"code"`
	User_Code           string     `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	CP_Code             string     `form:"cp_code" json:"cp_code" xml:"cp_code" db:"cp_code"`
	Program_Description NullString `form:"program_description" json:"program_description" xml:"program_description" db:"program_description"`
	Fullname            NullString `form:"fullname" json:"fullname" xml:"fullname" db:"fullname"`
	Gender              NullString `form:"gender" json:"gender" xml:"gender" db:"gender"`
	Email               NullString `form:"email" json:"email" xml:"email" db:"email"`
	WA_No               NullInt64  `form:"wa_no" json:"wa_no" xml:"wa_no" db:"wa_no"`
	Age                 NullInt64  `form:"age" json:"age" xml:"age" db:"age"`
	Address             NullString `form:"address" json:"address" xml:"address" db:"address"`
	Profession          NullString `form:"profession" json:"profession" xml:"profession" db:"profession"`
	Question_1          NullString `form:"question_1" json:"question_1" xml:"question_1" db:"question_1"`
	Question_2          NullString `form:"question_2" json:"question_2" xml:"question_2" db:"question_2"`
	Question_3          NullString `form:"question_3" json:"question_3" xml:"question_3" db:"question_3"`
	Question_4          NullString `form:"question_4" json:"question_4" xml:"question_4" db:"question_4"`
	Question_5          NullString `form:"question_5" json:"question_5" xml:"question_5" db:"question_5"`
	Question_6          NullString `form:"question_6" json:"question_6" xml:"question_6" db:"question_6"`
	Question_7          NullString `form:"question_7" json:"question_7" xml:"question_7" db:"question_7"`
	Question_8          NullString `form:"question_8" json:"question_8" xml:"question_8" db:"question_8"`
	Is_Confirmed        bool       `form:"is_confirmed" json:"is_confirmed" xml:"is_confirmed" db:"is_confirmed"`
	Is_Active           bool       `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By          string     `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date        time.Time  `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By         string     `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date       NullTime   `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted          bool       `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}
