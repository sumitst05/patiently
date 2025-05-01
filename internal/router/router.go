package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/handler"
	"github.com/sumitst05/patiently/internal/middleware"
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

	// base api group
	api := r.Group("/api")

	// health route: /api/health
	api.GET("/health", handler.HealthCheck)

	// auth routes: /api/auth/*
	auth := api.Group("/auth")
	auth.POST("/signup", handler.Singup)
	auth.POST("/signin", handler.Signin)
	auth.POST("/logout", handler.Logout)

	// protected routes: routes that require authentication
	api.Use(middleware.Auth())
	{
		// route to get current user
		api.GET("/me", handler.Me)

		// patient routes: /api/patient/*
		patient := api.Group("/patient")
		{
			// routes accessible by both receptionist and doctor
			patient.GET("/fetch", handler.GetAllPatients)
			patient.GET("/fetch/:id", handler.GetPatientById)
			patient.POST("/update/:id", handler.UpdatePatient)

			// routes accessible only by receptionist
			patient.GET("/fetch/:id/history", handler.GetPatientRegistrationHistory)
			patient.POST("/create", handler.CreatePatient)
			patient.DELETE("/delete/:id", handler.DeletePatient)
		}
	}

	return r
}
