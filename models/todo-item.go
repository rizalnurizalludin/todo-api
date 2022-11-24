package models

import (
	"time"
)

type TODO struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	ActivityGroupID uint      `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       time.Time `gorm:"default:nil" json:"deleted_at"`
}
