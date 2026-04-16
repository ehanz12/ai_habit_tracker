package response

type HabitResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	Description     string `json:"description"`
	TargetPerDay    int    `json:"target_per_day"`
	PreferredTime   string `json:"preferred_time"`
	TimeZone        string `json:"time_zone"`
	ReminderEnabled bool   `json:"reminder_enabled"`
	HabitStats	  *HabitStatResponse `json:"habit_stats,omitempty"`
}

type HabitStatResponse struct {
	Streak      int     `json:"streak"`
}
