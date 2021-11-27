package shape

import (
	"time"
)

type UserClass struct {
	ID                 int       `form:"id" json:"id" xml:"id" db:"id"`
	Code               string    `form:"code" json:"code" xml:"code" db:"code"`
	User_Code          string    `form:"user_code" json:"user_code" xml:"user_code" db:"user_code"`
	Class_Code         string    `form:"class_code" json:"class_code" xml:"class_code" db:"class_code"`
	Class_Name         string    `form:"class_name" json:"class_name" xml:"class_name" db:"class_name"`
	Class_Initial      string    `form:"class_initial" json:"class_initial" xml:"class_initial" db:"class_initial"`
	Class_Category     string    `form:"class_category" json:"class_category" xml:"class_category" db:"class_category"`
	Class_Description  string    `form:"class_description" json:"class_description" xml:"class_description" db:"class_description"`
	Class_Image        string    `form:"class_image" json:"class_image" xml:"class_image" db:"class_image"`
	Class_Rating       float64   `form:"class_rating" json:"class_rating" xml:"class_rating" db:"class_rating"`
	Color_Path         string    `form:"color_path" json:"color_path" xml:"color_path" db:"color_path"`
	Total_User         int       `form:"total_user" json:"total_user" xml:"total_user" db:"total_user"`
	Type_Code          string    `form:"type_code" json:"type_code" xml:"type_code" db:"type_code"`
	Type               string    `form:"type" json:"type" xml:"type" db:"type"`
	Status_Code        string    `form:"status_code" json:"status_code" xml:"status_code" db:"status_code"`
	Status             string    `form:"status" json:"status" xml:"status" db:"status"`
	Package_Code       string    `form:"package_code" json:"package_code" xml:"package_code" db:"package_code"`
	Package_Type       string    `form:"package_type" json:"package_type" xml:"package_type" db:"package_type"`
	Promo_Code         string    `form:"promo_code" json:"promo_code" xml:"promo_code" db:"promo_code"`
	Is_Expired         bool      `form:"is_expired" json:"is_expired" xml:"is_expired" db:"is_expired"`
	Start_Date         string    `form:"start_date" json:"start_date" xml:"start_date" db:"start_date"`
	Expired_Date       string    `form:"expired_date" json:"expired_date" xml:"expired_date" db:"expired_date"`
	Time_Duration      int       `form:"time_duration" json:"time_duration" xml:"time_duration" db:"time_duration"`
	Progress           float64   `form:"progress" json:"progress" xml:"progress" db:"progress"`
	Progress_Count     int       `form:"progress_count" json:"progress_count" xml:"progress_count" db:"progress_count"`
	Progress_Index     int       `form:"progress_index" json:"progress_index" xml:"progress_index" db:"progress_index"`
	Progress_Subindex  int       `form:"progress_subindex" json:"progress_subindex" xml:"progress_subindex" db:"progress_subindex"`
	Pre_Test_Scores    float64   `form:"pre_test_scores" json:"pre_test_scores" xml:"pre_test_scores" db:"pre_test_scores"`
	Post_Test_Scores   float64   `form:"post_test_scores" json:"post_test_scores" xml:"post_test_scores" db:"post_test_scores"`
	Post_Test_Date     string    `form:"post_test_date" json:"post_test_date" xml:"post_test_date" db:"post_test_date"`
	Pre_Test_Total     int       `form:"pre_test_total" json:"pre_test_total" xml:"pre_test_total" db:"pre_test_total"`
	Post_Test_Total    int       `form:"post_test_total" json:"post_test_total" xml:"post_test_total" db:"post_test_total"`
	Total_Consultation int       `form:"rotal_consultation" json:"rotal_consultation" xml:"rotal_consultation" db:"rotal_consultation"`
	Total_Webinar      int       `form:"total_webinar" json:"total_webinar" xml:"total_webinar" db:"total_webinar"`
	Is_Active          bool      `form:"is_active" json:"is_active" xml:"is_active" db:"is_active"`
	Created_By         string    `form:"created_by" json:"created_by" xml:"created_by" db:"created_by"`
	Created_Date       time.Time `form:"created_date" json:"created_date" xml:"created_date" db:"created_date"`
	Modified_By        string    `form:"modified_by" json:"modified_by" xml:"modified_by" db:"modified_by"`
	Modified_Date      time.Time `form:"modified_date" json:"modified_date" xml:"modified_date" db:"modified_date"`
	Is_Deleted         bool      `form:"is_deleted" json:"is_deleted" xml:"is_deleted" db:"is_deleted"`
}

type UserClassPost struct {
	ID                int
	Code              string
	User_Code         string
	Class_Code        string
	Class_Name        string
	Class_Initial     string
	Class_Category    string
	Class_Description string
	Class_Image       string
	Class_Rating      float64
	Total_User        int
	Status_Code       string
	Status            string
	Is_Expired        bool
	Start_Date        string
	Expired_Date      string
	Time_Duration     int
	Progress          float64
	Progress_Count    int64
	Progress_Index    int64
	Progress_Subindex int64
	Pre_Test_Scores   float64
	Post_Test_Scores  float64
	Post_Test_Date    string
	Is_Active         bool
	Created_By        string
	Created_Date      time.Time
	Modified_By       string
	Modified_Date     time.Time
	Is_Deleted        bool
}
