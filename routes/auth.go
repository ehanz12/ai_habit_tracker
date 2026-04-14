package routes

import (
	"github.com/ehanz12/ai_habit_tracker/handlers"
	"github.com/ehanz12/ai_habit_tracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)
	auth.Get("/me", middleware.ProtectedRoute, handlers.Me)
}