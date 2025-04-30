package middleware

import (
	"net/http"
	"notes-system/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer") || util.IsValidToken(strings.TrimPrefix(token, "Bearer")) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})

			return
		}
		c.Next()
	}
}
