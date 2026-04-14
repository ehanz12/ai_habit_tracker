package models

import "time"

type Habit struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint
	Name          string    `gorm:"size:100"`
	Description   string    `gorm:"type:text"`
	TargetPerDay  int       `gorm:"default:1"`
	PreferredTime string    // bisa pakai TIME kalau mau custom type
	CreatedAt     time.Time

	User       User
	HabitLogs  []HabitLog
	HabitStats *HabitStat
}