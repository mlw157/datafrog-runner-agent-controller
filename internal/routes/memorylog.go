package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupMemoryLogRoutes(api *echo.Group, handler *handlers.MemoryLogHandler) {
	memoryLogGroup := api.Group("/memorylogs")

	memoryLogGroup.GET("", handler.GetMemoryLogs)
	memoryLogGroup.POST("", handler.CreateMemoryLog)
}
