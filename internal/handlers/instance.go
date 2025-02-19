package handlers

import "gorm.io/gorm"

type InstanceHandler struct {
	DB *gorm.DB
}
