package tasks

import (
	"github.com/balticoldschool/go-react-tutorial/src/utils/errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Controller interface {
	CreateTask(ctx *fiber.Ctx) error
	GetAllTasks(ctx *fiber.Ctx) error
	GetTaskById(ctx *fiber.Ctx) error
	UpdateTaskCompletion(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
}

type controller struct {
	service Service
}

func NewTaskController(service Service) Controller {
	return &controller{
		service,
	}
}

func (c controller) CreateTask(ctx *fiber.Ctx) error {
	newTask := Task{}

	if err := ctx.BodyParser(&newTask); err != nil {
		restErr := errors.NewBadRequestErr("invalid request body")
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	return ctx.
		Status(http.StatusCreated).
		JSON(*c.service.AddTask(&newTask))
}

func (c controller) GetAllTasks(ctx *fiber.Ctx) error {
	tasks := c.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(tasks)
}

func (c controller) GetTaskById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		restErr := errors.NewBadRequestErr("invalid task id")
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	response, restErr := c.service.GetById(uint(id))
	if restErr != nil {
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (c controller) UpdateTaskCompletion(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		restErr := errors.NewBadRequestErr("invalid id param")
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	task, responseErr := c.service.UpdateTaskCompletion(uint(id))
	if responseErr != nil {
		return ctx.Status(responseErr.Status).JSON(responseErr)
	}

	return ctx.Status(http.StatusOK).JSON(task)
}

func (c controller) DeleteTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		restErr := errors.NewBadRequestErr("invalid param")
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	task, restErr := c.service.RemoveTask(uint(id))
	if restErr != nil {
		return ctx.Status(restErr.Status).JSON(restErr)
	}

	return ctx.Status(http.StatusOK).JSON(task)
}
