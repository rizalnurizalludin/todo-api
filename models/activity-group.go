package models

import "time"

type ActivityGroup struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:nil" json:"deleted_at"`
}
