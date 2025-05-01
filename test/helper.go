package test

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/router"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return router.SetupRouter()
}
