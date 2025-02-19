package handlers

import (
	"agent-controller/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type InstanceHandler struct {
	DB *gorm.DB
}

func (h *InstanceHandler) GetInstances(c echo.Context) error {
	var instances []models.Instance

	result := h.DB.Find(&instances)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch instances",
		})
	}

	return c.JSON(http.StatusOK, instances)
}

func (h *InstanceHandler) CreateInstance(c echo.Context) error {
	instance := new(models.Instance)

	err := c.Bind(instance)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid instance data",
		})
	}
	result := h.DB.Create(instance)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create instance",
		})

	}
	return c.JSON(http.StatusCreated, instance)
}
