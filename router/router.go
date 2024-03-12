package router

import (
	"github.com/gin-gonic/gin"

	"block-cipher/middlewares"
	"block-cipher/handlers"
)

func ConfigureRouter(r *gin.Engine) {
	// Apply CORS middleware
	r.Use(middlewares.CorsMiddleware())

	// Define endpoints for back-end services
	// General handler
	r.GET("/api", handlers.HelloHandler)
	apis := r.Group("/api")

	// CBC Mode Endpoints
	apis.POST("/cbc", handlers.HandleCBCRequest)
}