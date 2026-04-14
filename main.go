package main

import (
	"github.com/ehanz12/ai_habit_tracker/config"
	"github.com/ehanz12/ai_habit_tracker/database"
	"github.com/ehanz12/ai_habit_tracker/models"
	"github.com/ehanz12/ai_habit_tracker/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//lOAD ENV
	config.LoadEnv()

	//connect database
	database.ConnectDB()
	database.DB.AutoMigrate(
    &models.User{},
    &models.Habit{},
    &models.HabitLog{},
    &models.HabitStat{},
    &models.AIRecommendation{},
    &models.Notification{},
)
	//register fiber
	app := fiber.New()
	//register route
	routes.SetupRoutes(app)
	
	//start server
	println("Server Start in port 3000", app.Listen(":3000"))
}