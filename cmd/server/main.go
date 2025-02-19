package server

import (
	"agent-controller/internal/database"
	"agent-controller/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		log.Fatal("Provide DB_PATH environment variable")
	}

	db, err := database.InitDB(dbPath)

	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))

}
