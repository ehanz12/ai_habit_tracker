package service

import (
	"errors"
	"time"

	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/dto/requests"
	"github.com/ehanz12/ai_habit_tracker/models"
	"github.com/ehanz12/ai_habit_tracker/utils"
	"gorm.io/gorm"
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

	var reminderEnabled bool
	if r.ReminderEnabled != nil {
		reminderEnabled = *r.ReminderEnabled
	}

	habit := models.Habit{
		Name:            r.Name,
		Category:        r.Category,
		UserID:          UserID,
		Description:     r.Description,
		TargetPerDay:    r.TargetPerDay,
		PreferredTime:   timeNew,
		TimeZone:        r.TimeZone,
		ReminderEnabled: reminderEnabled,
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

func GetHabits(UserID uint, page int, limit int) ([]models.Habit, int64, int, error) {
	var habits []models.Habit
	var totalData int64

	offset := (page - 1) * limit
	db := database.DB

	// count
	if err := db.Model(&models.Habit{}).
		Where("user_id = ?", UserID).
		Count(&totalData).Error; err != nil {
		return nil, 0, 0, err
	}

	// data
	if err := db.Model(&models.Habit{}).
		Select("id", "user_id", "name", "category", "description", "target_per_day", "time_zone", "preferred_time", "reminder_enabled").
		Where("user_id = ?", UserID).
		Preload("HabitStats", func(db *gorm.DB) *gorm.DB {
			return db.Select("habit_id", "streak")
		}).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&habits).Error; err != nil {
		return nil, 0, 0, err
	}

	for i := range habits {
		if habits[i].HabitStats != nil {
			realStreak := GetRealStreak(
				habits[i].ID,
				habits[i].HabitStats.Streak,
			)

			habits[i].HabitStats.Streak = realStreak
		}
	}

	totalPages := int((totalData + int64(limit) - 1) / int64(limit))

	return habits, totalData, totalPages, nil
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
	if err := tx.Where("id = ? AND user_id = ?", HabitID, UserID).First(&habit).Error; err != nil {
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

func GetRealStreak(habitID uint, storedStreak int) int {
	var countToday int64
	var countYesterday int64

	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	db := database.DB

	db.Model(&models.HabitLog{}).
		Where("habit_id = ? AND DATE(date) = ? AND completed = ?", habitID, today, true).
		Count(&countToday)

	db.Model(&models.HabitLog{}).
		Where("habit_id = ? AND DATE(date) = ? AND completed = ?", habitID, yesterday, true).
		Count(&countYesterday)

	// ❗ skip kemarin dan belum hari ini → streak reset
	if countYesterday == 0 && countToday == 0 {
		return 0
	}

	// ❗ belum lakukan hari ini → tetap lanjut dari kemarin
	if countToday == 0 {
		return storedStreak
	}

	return storedStreak
}

func CalculateStreak(habitID uint) (int, error) {
	streak := 0

	now := time.Now()
	today := now.Format("2006-01-02")

	var countToday int64
	err := database.DB.Model(&models.HabitLog{}).
		Where("habit_id = ? AND DATE(date) = ? AND completed = ?", habitID, today, true).
		Count(&countToday).Error

	if err != nil {
		return 0, err
	}

	var startDate time.Time

	if countToday > 0 {
		startDate = now
	} else {
		startDate = now.AddDate(0, 0, -1)
	}

	limitDate := startDate.AddDate(0, 0, -365).Format("2006-01-02")

	var logs []models.HabitLog
	err = database.DB.Model(&models.HabitLog{}).
		Select("date").
		Where("habit_id = ? AND completed = ? AND date >= ?", habitID, true, limitDate).
		Find(&logs).Error

	if err != nil {
		return 0, err
	}

	logMap := make(map[string]bool)
	for _, log := range logs {
		logMap[log.Date.Format("2006-01-02")] = true
	}

	for i := 0; i < 365; i++ {
		checkDate := startDate.AddDate(0, 0, -i).Format("2006-01-02")

		if !logMap[checkDate] {
			break
		}

		streak++
	}

	return streak, nil
}

func GetHabitRecap(habitID uint) (int, int, int, error) {
	now := time.Now()

	var habit models.Habit
	if err := database.DB.Select("created_at").Where("id = ?", habitID).First(&habit).Error; err != nil {
		return 0, 0, 0, err
	}

	startDate := now.AddDate(0, 0, -6).Format("2006-01-02") + " 00:00:00"
	endDate := now.Format("2006-01-02") + " 23:59:59"

	var logs []models.HabitLog

	err := database.DB.
		Model(&models.HabitLog{}).
		Select("date, completed").
		Where("habit_id = ? AND date BETWEEN ? AND ?", habitID, startDate, endDate).
		Find(&logs).Error

	if err != nil {
		return 0, 0, 0, err
	}

	logMap := make(map[string]bool)

	for _, log := range logs {
		dateKey := log.Date.Format("2006-01-02")
		if log.Completed {
			logMap[dateKey] = true
		}
	}

	completed := 0
	missed := 0
	createdAtDate := habit.CreatedAt.Truncate(24 * time.Hour)

	for i := 0; i < 7; i++ {
		checkTime := now.AddDate(0, 0, -i)
		dateStr := checkTime.Format("2006-01-02")
		
		// Hanya hitung jika tanggal pengecekan >= tanggal pembuatan habit
		if checkTime.Truncate(24 * time.Hour).Before(createdAtDate) {
			continue // Abaikan hari-hari sebelum habit dibuat
		}

		if logMap[dateStr] {
			completed++
		} else {
			missed++
		}
	}

	streak, _ := CalculateStreak(habitID)

	return completed, missed, streak, nil
}

func GetFullHabitRecap(habitID uint) (map[string]interface{}, error) {
	completed, missed, streak, err := GetHabitRecap(habitID)
	if err != nil {
		return nil, errors.New("gagal mengambil data recap habit !")
	}

	var habit models.Habit
	if err := database.DB.Select("name").Where("id = ?", habitID).First(&habit).Error; err != nil {
		return nil, errors.New("habit tidak ditemukan")
	}

	insight := utils.GenerateAIInsight(habit.Name, completed, missed, streak)

	return map[string]interface{}{
		"completed_days": completed,
		"missed_days":    missed,
		"streak":         streak,
		"insight":        insight,
	}, nil
}
