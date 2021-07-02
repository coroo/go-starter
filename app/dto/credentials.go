package dto

type Credentials struct {
	UserProfilename string `form:"userProfilename"`
	Password string `form:"password"`
}

type Token struct {
	Token 	string `json:"token"`
}
