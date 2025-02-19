package handlers

import (
	"gorm.io/gorm"
)

type JobHandler struct {
	DB *gorm.DB
}
