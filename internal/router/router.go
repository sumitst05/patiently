package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/handler"
	"github.com/sumitst05/patiently/middleware"
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

	// health route
	r.GET("/health", handler.HealthCheck)

	// auth routes
	r.POST("/signup", handler.Singup)
	r.POST("/signin", handler.Signin)
	r.POST("/logout", handler.Logout)

	protected := r.Group("/api")
	protected.Use(middleware.Auth())
	{
		protected.GET("/me", handler.Me) // route to get current user
	}

	return r
}
