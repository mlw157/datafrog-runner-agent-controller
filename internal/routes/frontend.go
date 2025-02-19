package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupFrontendRoutes(e *echo.Echo) {
	e.GET("/", handlers.ServeDashboard)

}
