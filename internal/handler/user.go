package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/repository"
	"github.com/sumitst05/patiently/utils"
)

func Me(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please sign in"})
		return
	}

	jwtClaims := claims.(*utils.Claims)

	user, err := repository.GetUserById(jwtClaims.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}
