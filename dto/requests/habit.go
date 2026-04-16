package requests

type CreateHabitRequest struct {
	Name            string `json:"name" validate:"required"`
	Category        string `json:"category" validate:"required,oneof=low mid high"`
	Description     string `json:"description"`
	TargetPerDay    int    `json:"target_per_day" validate:"required,min=1"`
	PreferredTime   string `json:"preferred_time" validate:"required"`
	TimeZone        string `json:"time_zone" validate:"required"`
	ReminderEnabled *bool   `json:"reminder_enabled"`
}
