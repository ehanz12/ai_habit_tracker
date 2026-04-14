package handlers

import (
	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/dto/requests"
	"github.com/ehanz12/ai_habit_tracker/models"
	"github.com/ehanz12/ai_habit_tracker/service"
	"github.com/ehanz12/ai_habit_tracker/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req requests.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})
	}

	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "login successfully", "token": jwtToken})
}

func Register(c *fiber.Ctx) error {
	var req requests.RegisterRequest
	if err := c.BodyParser(&req);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "invalid payload"})
	}

	if err := service.Register(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message" : "register successfully"})
}

func Me(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "hello me"})
}
