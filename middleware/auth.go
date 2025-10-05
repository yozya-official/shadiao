package middleware

import (
	"net/http"
	"shadiao/conf"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		// 格式：Authorization: ApiKey xxx
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && strings.ToLower(parts[0]) == "apikey" && conf.IsValidAPIKey(parts[1])) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的验证信息"})
			c.Abort()
			return
		}

		c.Next()
	}
}
