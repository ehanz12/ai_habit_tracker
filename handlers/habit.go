package handlers

import (
	"github.com/ehanz12/ai_habit_tracker/dto/requests"
	"github.com/ehanz12/ai_habit_tracker/service"
	"github.com/gofiber/fiber/v2"
)

func CreateHabitHandler(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(uint)
	var r requests.CreateHabitRequest
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "kesalahan payload"})
	}
	

	if err := service.CreateHabit(UserID, r); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message" : "berhasil membuat habit !"})
}

func UpdateHabitHandler(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(uint)
	HabitID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "ID habit tidak valid"})
	}
	var r requests.CreateHabitRequest
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "kesalahan payload"})
	}

	if err := service.UpdateHabit(UserID, uint(HabitID), r); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "berhasil memperbarui habit !"})
}

func DeleteHabitHandler(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(uint)
	HabitID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "params harus integer"})
	}

	if err := service.DeleteHabit(UserID, uint(HabitID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "berhasil menghapus habit !"})
}