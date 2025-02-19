package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupFrontendRoutes(e *echo.Echo, handler *handlers.FrontendHandler, instanceHandler *handlers.InstanceHandler, jobHandler *handlers.JobHandler, memoryLogHandler *handlers.MemoryLogHandler) {

	e.GET("/", handler.ServeDashboard)
}
