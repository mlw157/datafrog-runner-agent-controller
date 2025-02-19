package handlers

import (
	"agent-controller/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type JobHandler struct {
	DB *gorm.DB
}

func (h *JobHandler) GetJobs(c echo.Context) error {
	var jobs []models.Job

	result := h.DB.Find(&jobs)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch jobs",
		})
	}

	return c.JSON(http.StatusOK, jobs)
}

func (h *JobHandler) CreateJob(c echo.Context) error {
	job := new(models.Job)

	err := c.Bind(job)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid job data",
		})
	}

	instance := new(models.Instance)
	result := h.DB.Where("instance_id = ?", job.InstanceID).First(instance)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Instance not found",
		})
	}

	result = h.DB.Create(job)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create job",
		})
	}

	return c.JSON(http.StatusCreated, job)
}
