package models

import "time"

type AIRecommendation struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	HabitID   uint
	Message   string `gorm:"type:text"`
	CreatedAt time.Time

	User  User  `gorm:"foreignKey:UserID"`
	Habit Habit `gorm:"foreignKey:HabitID"`
}
