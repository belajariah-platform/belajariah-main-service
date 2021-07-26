package shape

import (
	"time"
)

type Class struct {
	ID                     int
	Code                   string
	Class_Category_Code    string
	Class_Category         string
	Class_Name             string
	Class_Initial          string
	Class_Description      string
	Class_Image            string
	Class_Video            string
	Class_Rating           float64
	Total_Review           int
	Total_Video            int
	Total_Video_Duration   float64
	Instructor_Name        string
	Instructor_Description string
	Instructor_Biografi    string
	Instructor_Image       string
	Is_Direct              bool
	Is_Active              bool
	Created_By             string
	Created_Date           time.Time
	Modified_By            string
	Modified_Date          time.Time
	Deleted_By             string
	Deleted_Date           time.Time
	Price_Start            string
	Price_Start_Discount   string
	Price_End              string
	Price_End_Discount     string
}
