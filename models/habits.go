package models

import "time"

type Habit struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	Category        string `gorm:"size:50;default:'mid';enum:'low','mid','high'"`
	Name            string `gorm:"size:100"`
	Description     string `gorm:"type:text"`
	TargetPerDay    int    `gorm:"default:1"`
	PreferredTime   string `gorm:"type:time"`
	TimeZone        string `gorm:"size:50"`
	ReminderEnabled bool  `gorm:"default:true"`
	CreatedAt       time.Time

	User       User       `gorm:"foreignKey:UserID"`
	HabitLogs  []HabitLog `gorm:"foreignKey:HabitID"`
	HabitStats *HabitStat `gorm:"foreignKey:HabitID"`
}
