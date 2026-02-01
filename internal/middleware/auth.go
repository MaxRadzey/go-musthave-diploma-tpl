package middleware

import (
	"net/http"

	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/auth"
	"github.com/gin-gonic/gin"
)

// KeyUserID — ключ в gin.Context, под которым сохраняется userID (int64) после успешной проверки куки.
const KeyUserID = "user_id"

// RequireAuth возвращает Gin middleware: проверяет куку, при валидной куке кладёт userID в контекст и вызывает Next.
// При отсутствии или невалидной куке отвечает 401 и прерывает цепочку.
func RequireAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(auth.CookieName)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userID, err := auth.ValidateCookie(cookie, secret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(KeyUserID, userID)
		c.Next()
	}
}
