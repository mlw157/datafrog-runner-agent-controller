package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	SetupFrontendRoutes(e)

	api := e.Group("/api")
	SetupInstanceRoutes(api, db)
	SetupJobRoutes(api, db)
	SetupMemoryLogRoutes(api, db)
}
