package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupInstanceRoutes(api *echo.Group, handler *handlers.InstanceHandler) {
	instanceGroup := api.Group("/instances")

	instanceGroup.GET("", handler.GetInstances)
	instanceGroup.POST("", handler.CreateInstance)
}
