package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupInstanceRoutes(api *echo.Group, db *gorm.DB) {
	handler := &handlers.InstanceHandler{DB: db}

	instanceGroup := api.Group("/instances")
	instanceGroup.GET("", handler.GetInstances)
	instanceGroup.POST("", handler.CreateInstance)
}
