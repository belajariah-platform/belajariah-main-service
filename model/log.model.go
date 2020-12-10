package model

type Logs struct {
	Logs []LogsData `form:"logs" json:"logs" xml:"logs"`
}

type LogsData struct {
	Date string `form:"date" json:"date" xml:"date"`
	Log  string `form:"log" json:"log" xml:"log"`
}
