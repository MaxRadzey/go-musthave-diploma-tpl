package handler

import "github.com/MaxRadzey/go-musthave-diploma-tpl/internal/service"

// Handler — HTTP-хендлеры.
type Handler struct {
	UserService   *service.UserService
	CookieSecret  string
}

// New создаёт Handler с переданным UserService и секретом для куки.
func New(userService *service.UserService, cookieSecret string) *Handler {
	return &Handler{UserService: userService, CookieSecret: cookieSecret}
}
