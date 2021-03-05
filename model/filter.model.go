package model

type Filter struct {
	Type  string
	Value string
	Field string
}

type Query struct {
	Skip    int    `form:"skip" json:"skip" xml:"skip"`
	Take    int    `form:"take" json:"take" xml:"take"`
	Order   string `form:"order" json:"order" xml:"order"`
	Search  string `form:"search" json:"search" xml:"search"`
	Filter  string `form:"filter" json:"filter" xml:"filter"`
	Filters []Filter
}
