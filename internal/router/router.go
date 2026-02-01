package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/config"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/handler"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/middleware"
)

// SetupRouter создаёт Gin-роутер.
func SetupRouter(h *handler.Handler, cfg *config.Config) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/user")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		protected := api.Group("", middleware.RequireAuth(cfg.CookieSecret))
		{
			protected.POST("/orders", h.PostOrders)
			protected.GET("/orders", h.GetOrders)
		}
	}
	return r
}
