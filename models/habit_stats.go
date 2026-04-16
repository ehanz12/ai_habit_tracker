package models

import "time"

type HabitStat struct {
	ID          uint `gorm:"primaryKey"`
	HabitID     uint `gorm:"uniqueIndex"`
	SuccessRate float64
	Streak      int
	LastUpdated time.Time

	Habit *Habit `gorm:"foreignKey:HabitID"`
}
