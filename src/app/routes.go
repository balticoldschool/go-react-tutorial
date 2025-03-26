package app

import (
	"github.com/balticoldschool/go-react-tutorial/src/domain/tasks"
	"github.com/gofiber/fiber/v2"
)

var (
	taskController = tasks.NewTaskController(tasks.NewTaskService())
)

func mapRoutes(app *fiber.App) {
	app.Get("/tasks", taskController.GetAllTasks)
	app.Get("/task/:id", taskController.GetTaskById)
	app.Post("/create", taskController.CreateTask)
}
