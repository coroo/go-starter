package dto

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Token struct {
	AccessToken 	string `json:"access_token"`
	RefreshToken 	string `json:"refresh_token"`
}
