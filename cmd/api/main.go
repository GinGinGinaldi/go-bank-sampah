package main

import (
	"bank-sampah/internal/database"
	"bank-sampah/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Setup routes
	routes.RegisterRoutes(e, db)

	// Start server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
