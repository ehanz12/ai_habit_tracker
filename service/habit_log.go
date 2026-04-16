package service

import (
	"errors"
	"time"

	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/models"
	"gorm.io/gorm"
)

func CheckHabitCompletion(userid, id uint) error {
	today := time.Now().Truncate(24 * time.Hour)

	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("terjadi kesalahan sistem")
	}

	var habit models.Habit
	if err := tx.Where("id = ? AND user_id = ?", id, userid).First(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("habit tidak ditemukan")
	}

	var log models.HabitLog

	err := tx.Where("habit_id = ? AND DATE(date) = CURDATE()", id).First(&log).Error
	if err == nil {
		log.Completed = !log.Completed
		tx.Save(&log)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		newLog := models.HabitLog{
			HabitID:   id,
			Date:      today,
			Completed: true,
		}
		tx.Create(&newLog)
	} else {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("gagal commit")
	}

	streak, err := CalculateStreak(id)
	if err != nil {
		return err
	}

	var stat models.HabitStat
	database.DB.FirstOrCreate(&stat, models.HabitStat{
		HabitID: id,
	})

	// 🔥 update
	database.DB.Model(&stat).Updates(map[string]interface{}{
		"streak":       streak,
		"last_updated": time.Now(),
	})
	return nil
}
