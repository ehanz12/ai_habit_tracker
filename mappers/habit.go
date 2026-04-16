package mappers

import (
	"github.com/ehanz12/ai_habit_tracker/dto/response"
	"github.com/ehanz12/ai_habit_tracker/models"
)

func ToHabitResponse(h models.Habit) response.HabitResponse {
	return response.HabitResponse{
		ID:              h.ID,
		Name:            h.Name,
		Category:        h.Category,
		Description:     h.Description,
		TargetPerDay:    h.TargetPerDay,
		PreferredTime:   h.PreferredTime,
		TimeZone:        h.TimeZone,
		ReminderEnabled: h.ReminderEnabled,
		HabitStats: func() *response.HabitStatResponse {
			if h.HabitStats != nil {
				return &response.HabitStatResponse{
					Streak:      h.HabitStats.Streak,
				}
			}
			return &response.HabitStatResponse{
				Streak:      0,
			}
		}(),
	}
}

func ListToHabitResponse(habits []models.Habit) []response.HabitResponse {
	res := make([]response.HabitResponse, 0, len(habits))
	for _, h := range habits {
		res = append(res, ToHabitResponse(h))
	}
	return res
}