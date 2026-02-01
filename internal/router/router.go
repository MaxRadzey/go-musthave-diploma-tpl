package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/config"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/handler"
)

// SetupRouter создаёт Gin-роутер с группой /api/user (register, login).
func SetupRouter(h *handler.Handler, cfg *config.Config) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/user")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
	}
	return r
}
