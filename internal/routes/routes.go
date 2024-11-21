package routes

import (
	"bank-sampah/internal/controllers"
	"bank-sampah/internal/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	e = echo.New()
	e.POST("/login", controllers.Login(db))
	e.POST("/register", controllers.Register(db))
	e.GET("/users", controllers.GetAllUsers(db))
	e.GET("/statistics", controllers.GetStatistics(db))
	e.GET("/user", controllers.TokenMiddleware(controllers.GetUserData))

	// Menggunakan middleware admin untuk endpoint tertentu
	adminGroup := e.Group("/admin")
	adminGroup.Use(middlewares.AdminMiddleware) // Terapkan middleware admin
	adminGroup.GET("/dashboard", controllers.TokenMiddleware(controllers.GetAdminDashboard)) // Contoh endpoint admin

	e.Logger.Fatal(e.Start(":8080"))
}