package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// logger and recovery middlewares
	r.Use(gin.Logger(), gin.Recovery())

	// for no verbose debug logs in prod
	mode := os.Getenv("MODE")
	if mode == "prod" {
		gin.SetMode("release")
	}

	r.SetTrustedProxies(nil)

	r.GET("/health", handler.HealthCheck)

	return r
}
