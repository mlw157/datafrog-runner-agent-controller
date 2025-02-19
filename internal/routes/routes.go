package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	frontendHandler := &handlers.FrontendHandler{}
	instanceHandler := &handlers.InstanceHandler{DB: db}
	jobHandler := &handlers.JobHandler{DB: db}
	memoryLogHandler := &handlers.MemoryLogHandler{DB: db}

	SetupFrontendRoutes(e, frontendHandler, instanceHandler, jobHandler, memoryLogHandler)

	api := e.Group("/api")
	SetupInstanceRoutes(api, instanceHandler)
	SetupJobRoutes(api, jobHandler)
	SetupMemoryLogRoutes(api, memoryLogHandler)
}
