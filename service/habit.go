package service

import (
	"errors"

	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/dto/requests"
	"github.com/ehanz12/ai_habit_tracker/models"
	"github.com/ehanz12/ai_habit_tracker/utils"
)

func CreateHabit(UserID uint, r requests.CreateHabitRequest) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("terjadi kesalahan sistem !")
	}

	timeNew, err := utils.NormalizeAndValidateTime(r.PreferredTime)
	if err != nil {
		tx.Rollback()
		return errors.New("waktu yang diberikan tidak valid !")
	}


	habit := models.Habit{
		Name:            r.Name,
		Category:        r.Category,
		UserID:          UserID,
		Description:     r.Description,
		TargetPerDay:    r.TargetPerDay,
		PreferredTime:   timeNew,
		TimeZone:        r.TimeZone,
		ReminderEnabled: *r.ReminderEnabled,
	}

	if err := tx.Create(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("terjadi kesalahan gagal membuat habit !")
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("terjadi kesalahan sistem !")
	}
	return nil
}

func UpdateHabit(UserID uint, HabitID uint, r requests.CreateHabitRequest) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("terjadi kesalahan sistem !")
	}

	var habit models.Habit
	if err := tx.Where("id = ? AND user_id = ?", HabitID, UserID).First(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("habit tidak ditemukan !")
	}

	if r.PreferredTime != "" {
		timeNew, err := utils.NormalizeAndValidateTime(r.PreferredTime)
		if err != nil {
			tx.Rollback()
			return errors.New("waktu yang diberikan tidak valid !")
		}
		habit.PreferredTime = timeNew
	}

	if r.Name != "" {
		habit.Name = r.Name
	}

	if r.Category != "" {
		habit.Category = r.Category
	}

	if r.Description != "" {
		habit.Description = r.Description
	}

	if r.TargetPerDay != 0 {
		habit.TargetPerDay = r.TargetPerDay
	}

	if r.TimeZone != "" {
		habit.TimeZone = r.TimeZone
	}

	if r.ReminderEnabled != nil {
		habit.ReminderEnabled = *r.ReminderEnabled
	}

	if err := tx.Save(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("terjadi kesalahan gagal memperbarui habit !")
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("terjadi kesalahan sistem !")
	}
	return nil
}

func DeleteHabit(UserID, HabitID uint) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("terjadi kesalahan sistem !")
	}

	var habit models.Habit
	if err := tx.Where("id = ? AND user_id = ?", HabitID, UserID).Error; err != nil {
		tx.Rollback()
		return errors.New("habit tidak ditemukan !")
	}

	if err := tx.Delete(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("terjadi kesalahan gagal menghapus habit !")
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("terjadi kesalahan sistem !")
	}
	return nil
}
