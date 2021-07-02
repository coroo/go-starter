package dto

type Response struct {
	Message string `json:"message"`
}

type ErrResponse struct {
	ErrCode 	string `json:"errCode"`
	ErrMessage 	string `json:"errMessage"`
}
