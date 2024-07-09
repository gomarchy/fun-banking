package api

import (
	"github.com/bytebury/fun-banking/internal/api/handler"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Start() {
	// Load Templates
	router.LoadHTMLGlob("public/templates/*.html")
	// Load Static Files
	router.Static("/static", "public/static")
	// Middleware
	// Setup Routes
	setupHomepageRoutes()
	// Run the application
	router.Run()
}

func setupHomepageRoutes() {
	handler := handler.NewHomePageHandler()

	router.
		GET("", handler.Index).
		POST("click", handler.Click)
}
