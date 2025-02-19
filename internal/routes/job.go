package routes

import (
	"agent-controller/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupJobRoutes(api *echo.Group, handler *handlers.JobHandler) {
	jobGroup := api.Group("/jobs")

	jobGroup.GET("", handler.GetJobs)
	jobGroup.POST("", handler.CreateJob)
}
