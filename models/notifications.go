package models

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint
	Title     string    `gorm:"size:100"`
	Body      string    `gorm:"type:text"`
	IsRead    bool      `gorm:"default:false"`
	CreatedAt time.Time

	User User
}