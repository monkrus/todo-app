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
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 "
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect")
	}
	fmt.Println("Database connected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")

}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.CreateTodoById)
	app.Post("/todos", models.CreateTodo)
	app.Put("/todos/:id", models.UpdateTodo)
}
func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
