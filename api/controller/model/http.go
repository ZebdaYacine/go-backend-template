package model

type LoginResponse struct {
	Token    string `json:"token"`
	UserData any    `json:"userdata"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
