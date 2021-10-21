package shape

import (
	"database/sql"
	"time"
)

type Story struct {
	ID                  int
	Code                string
	Story_Category_Code string
	Image_Banner_Story  string
	Image_Header_Story  string
	Video_Story         string
	Title               string
	Content             string
	Source              string
	Is_Active           bool
	Created_By          string
	Created_Date        time.Time
	Modified_By         string
	Modified_Date       time.Time
	ModifiedDate        sql.NullTime
	Is_Deleted          bool
}
