package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/utils"
)

// auth middlware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
