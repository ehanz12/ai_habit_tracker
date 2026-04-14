package models

import "time"

type HabitLog struct {
	ID        uint      `gorm:"primaryKey"`
	HabitID   uint
	Date      time.Time `gorm:"type:date"`
	Completed bool      `gorm:"default:false"`
	Note      string    `gorm:"type:text"`

	Habit Habit
}