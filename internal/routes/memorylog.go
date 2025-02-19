package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupMemoryLogRoutes(api *echo.Group, db *gorm.DB) {
	handler := &handlers.MemoryLogHandler{DB: db}

	memoryLogGroup := api.Group("/memorylogs")
	memoryLogGroup.GET("", handler.GetMemoryLogs)
	memoryLogGroup.POST("", handler.CreateMemoryLog)
}
