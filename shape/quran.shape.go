package shape

import "time"

type Quran struct {
	ID              int
	Code            string
	Surat_Code      string
	Surat_Name      string
	Surat_Text      string
	Surat_Translate string
	Count_Ayat      int
	Ayat_Number     int
	Ayat_Text       string
	Ayat_Translate  string
	Juz_Number      int
	Page_Number     int
	Is_Active       bool
	Created_By      string
	Created_Date    time.Time
	Modified_By     string
	Modified_Date   time.Time
	Deleted_By      string
	Deleted_Date    time.Time
}
