package app

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	app = fiber.New()
)

func CreateApp() {
	mapRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
