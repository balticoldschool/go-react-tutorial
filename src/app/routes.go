package app

import (
	"github.com/balticoldschool/go-react-tutorial/src/domain/tasks"
	"github.com/gofiber/fiber/v2"
)

var (
	taskController = tasks.NewTaskController(tasks.NewTaskService())
)

func mapRoutes(app *fiber.App) {
	app.Post("/", taskController.CreateTask)
}
