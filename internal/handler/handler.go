package handler

import "github.com/MaxRadzey/go-musthave-diploma-tpl/internal/service"

// Handler — HTTP-хендлеры.
type Handler struct {
	UserService *service.UserService
}

// New создаёт Handler с переданным UserService.
func New(userService *service.UserService) *Handler {
	return &Handler{UserService: userService}
}
