package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wine-shop/internal/auth"
)

func AuthRequired(db interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		userIDFloat, _ := claims["user_id"].(float64)
		role, _ := claims["role"].(string)

		c.Set("userID", uint(userIDFloat))
		c.Set("role", role)

		c.Next()
	}
}
