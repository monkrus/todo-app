package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/monkrus/todo-app.git/database"
	todo "github.com/monkrus/todo-app.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	fmt.Println("Database connection successful")
	database.DBConn.AutoMigrate(&todo.Todo{})
	fmt.Println("Database migrated")
}
func setupRoutes(app *fiber.App) {
	app.Get("/todos", todo.GetTodos)
	app.Get("/todos/:id", todo.GetTodoById)
	app.Post("/todos", todo.CreateTodo)
	app.Put("/todos/:id", todo.UpdateTodo)
	app.Delete("/todos/:id", todo.DeleteTodo)
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
