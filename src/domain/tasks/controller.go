package tasks

import (
	"github.com/balticoldschool/go-react-tutorial/src/utils/errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Controller interface {
	CreateTask(ctx *fiber.Ctx) error
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
		Status(http.StatusOK).
		JSON(*c.service.AddTask(&newTask))
}
