package handlers

import "gorm.io/gorm"

type MemoryLogHandler struct {
	DB *gorm.DB
}
