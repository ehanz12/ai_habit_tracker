package handlers

import (
	"strconv"

	"github.com/ehanz12/ai_habit_tracker/service"
	"github.com/gofiber/fiber/v2"
)

func GetHabitRecapHandler(c *fiber.Ctx) error {
	habitID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid habit id",
		})
	}

	data, err := service.GetFullHabitRecap(uint(habitID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil mengambil recap",
		"data":    data,
	})
}