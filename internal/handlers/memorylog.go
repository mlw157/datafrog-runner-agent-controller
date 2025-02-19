package handlers

import (
	"agent-controller/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type MemoryLogHandler struct {
	DB *gorm.DB
}

func (h *MemoryLogHandler) GetMemoryLogs(c echo.Context) error {
	var memoryLogs []models.MemoryLog

	result := h.DB.Find(&memoryLogs)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch memory logs",
		})
	}

	return c.JSON(http.StatusOK, memoryLogs)
}

func (h *MemoryLogHandler) CreateMemoryLog(c echo.Context) error {
	memoryLog := new(models.MemoryLog)

	err := c.Bind(memoryLog)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid memory log data",
		})
	}

	instance := new(models.Instance)
	result := h.DB.Where("instance_id = ?", memoryLog.InstanceID).First(instance)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Instance not found",
		})
	}

	result = h.DB.Create(memoryLog)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create memory log",
		})
	}

	return c.JSON(http.StatusCreated, memoryLog)
}
