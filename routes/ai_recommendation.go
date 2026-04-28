package routes

import (
	"github.com/ehanz12/ai_habit_tracker/handlers"
	"github.com/ehanz12/ai_habit_tracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAIRoutes(api fiber.Router) {
	ai := api.Group("/ai")
	ai.Get("/recommendation/:id/recap", middleware.ProtectedRoute,handlers.GetHabitRecapHandler)
}