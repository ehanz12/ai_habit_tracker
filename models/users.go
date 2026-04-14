package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Password  string `gorm:"type:text"`
	CreatedAt time.Time

	Habits []Habit `gorm:"foreignKey:UserID"`
}
