package model

type Login struct {
	Email    string `form:"Email" json:"Email" xml:"Email"`
	Password string `form:"Password" json:"Password" xml:"Password"`
	JwtToken string `form:"JwtToken" json:"JwtToken" xml:"JwtToken"`
	Action   string `form:"Action" json:"Action" xml:"Action"`
	IsMobile bool   `form:"IsMobile" json:"IsMobile" xml:"IsMobile"`
}

type LoginResult struct {
	IsAuthorized bool
	Token        string
	ErrorMessage string
	HTTPStatus   int
}
