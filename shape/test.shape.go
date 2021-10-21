package shape

import "time"

type ClassTest struct {
	ID             int
	Code           string
	Class_Code     string
	Test_Type_Code string
	Question       string
	Option_A       string
	Option_B       string
	Option_C       string
	Option_D       string
	Answer         int
	Test_Image     string
	Is_Active      bool
	Created_By     string
	Created_Date   time.Time
	Modified_By    string
	Modified_Date  time.Time
	Is_Deleted     bool
}

type ClassTestPost struct {
	ID        int
	Test_Type string
	Answers   []ClassTest
}
