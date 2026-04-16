package handlers

import (
	"github.com/ehanz12/ai_habit_tracker/service"
	"github.com/gofiber/fiber/v2"
)

func CheckHabitLogHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	habitID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid habit ID",
		})
	}
	// Call the service function to check habit completion
	err = service.CheckHabitCompletion(userID, uint(habitID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Habit log updated successfully",
	})
}