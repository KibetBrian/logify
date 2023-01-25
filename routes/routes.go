package routes

import (
	"github.com/gofiber/fiber/v2"
	"shoppy/handlers"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.Hello)
}
