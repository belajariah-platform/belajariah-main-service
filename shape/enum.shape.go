package shape

import "time"

type Enum struct {
	ID            int
	Code          string
	Type          string
	Value         string
	Is_Active     bool
	Created_By    string
	Created_Date  time.Time
	Modified_By   string
	Modified_Date time.Time
	Deleted_By    string
	Deleted_Date  time.Time
}
