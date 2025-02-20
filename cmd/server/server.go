package main

import (
	"agent-controller/internal/database"
	"agent-controller/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	db, err := database.InitDB("/app/data/data.db")

	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))

}
