package service

import (
	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/dto/requests"
	"github.com/ehanz12/ai_habit_tracker/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(req requests.RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hash),
	}

	// simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}