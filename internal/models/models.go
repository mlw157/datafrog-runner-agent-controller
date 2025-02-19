package models

import (
	"gorm.io/gorm"
	"time"
)

type Instance struct {
	gorm.Model
	InstanceID       string `gorm:"uniqueIndex"`
	Type             string
	AvailabilityZone string
	PrivateIPAddress string
	LastSeenAt       time.Time
	MemoryMetrics    []MemoryLog `gorm:"foreignKey:InstanceID;references:InstanceID;constraint:OnDelete:CASCADE"`
	JobMetrics       []Job       `gorm:"foreignKey:InstanceID;references:InstanceID;constraint:OnDelete:CASCADE"`
}

type MemoryLog struct {
	gorm.Model
	InstanceID string `gorm:"index"`
	Total      int
	Used       int
	Free       int
	Timestamp  time.Time `gorm:"index"`
}

type Job struct {
	gorm.Model
	InstanceID string `gorm:"index"`
	Repository string
	Workflow   string
	StartTime  time.Time `gorm:"index"`
	EndTime    time.Time `gorm:"index"`
}
