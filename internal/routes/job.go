package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupJobRoutes(api *echo.Group, db *gorm.DB) {
	handler := &handlers.JobHandler{DB: db}

	jobGroup := api.Group("/jobs")
	jobGroup.GET("", handler.GetJobs)
	jobGroup.POST("", handler.CreateJob)
}
