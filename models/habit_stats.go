package models

import "time"

type HabitStat struct {
	ID          uint `gorm:"primaryKey"`
	HabitID     uint `gorm:"uniqueIndex"`
	Streak      int
	LastUpdated time.Time `gorm:"autoUpdateTime"`

	Habit *Habit `gorm:"foreignKey:HabitID"`
}
