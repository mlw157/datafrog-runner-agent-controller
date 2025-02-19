package database

import (
	"agent-controller/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Instance{}, &models.MemoryLog{}, &models.Job{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
