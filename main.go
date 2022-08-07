package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/monkrus/todo-app.git/database"
	"github.com/monkrus/todo-app.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=gorm password=gorm dbname=goTodo port=54321 "
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}
	fmt.Println("Database comnected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
}
func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
