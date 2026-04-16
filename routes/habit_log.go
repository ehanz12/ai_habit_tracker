package routes

import (
	"github.com/ehanz12/ai_habit_tracker/handlers"
	"github.com/ehanz12/ai_habit_tracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupHabitLogRoutes(api fiber.Router) {
	habitLog := api.Group("/check-habit")
	habitLog.Post("/:id", middleware.ProtectedRoute, handlers.CheckHabitLogHandler)
}