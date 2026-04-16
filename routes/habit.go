package routes

import (
	"github.com/ehanz12/ai_habit_tracker/handlers"
	"github.com/ehanz12/ai_habit_tracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupHabitRoutes(api fiber.Router) {
	habit := api.Group("/habits")

	habit.Post("/", middleware.ProtectedRoute, handlers.CreateHabitHandler)
	habit.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateHabitHandler)
}