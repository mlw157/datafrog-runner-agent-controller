package models

import (
	"gorm.io/gorm"
	"time"
)

type Instance struct {
	gorm.Model
	InstanceID       string      `json:"instance_id" gorm:"uniqueIndex"`
	Type             string      `json:"type"`
	AvailabilityZone string      `json:"availability_zone"`
	PrivateIPAddress string      `json:"private_ip_address"`
	LastSeenAt       time.Time   `json:"last_seen_at"`
	MemoryMetrics    []MemoryLog `gorm:"foreignKey:InstanceID;references:InstanceID;constraint:OnDelete:CASCADE"`
	JobMetrics       []Job       `gorm:"foreignKey:InstanceID;references:InstanceID;constraint:OnDelete:CASCADE"`
}

type Job struct {
	gorm.Model
	InstanceID   string    `json:"instance_id" gorm:"index"`
	Organization string    `json:"organization"`
	Repository   string    `json:"repository"`
	Workflow     string    `json:"workflow"`
	StartTime    time.Time `gorm:"index"`
	EndTime      time.Time `gorm:"index"`
}

type MemoryLog struct {
	gorm.Model
	InstanceID string    `json:"instance_id" gorm:"index"`
	Total      int       `json:"total"`
	Used       int       `json:"used"`
	Free       int       `json:"free"`
	Timestamp  time.Time `json:"timestamp" gorm:"index"`
}
