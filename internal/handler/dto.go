package handler

// RegisterRequest — тело запроса регистрации.
type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// LoginRequest — тело запроса логина.
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
