package routes

import (
	"github.com/gofiber/fiber/v2"
	"shoppy/handlers"
)

const rootUrl = "/api/v1/"

func Routes(app *fiber.App) {
	app.Get("/", handlers.Hello)

	//Auth
	app.Post(rootUrl+"auth/register", handlers.Register)
	app.Post(rootUrl+"auth/login", handlers.Login)
	app.Get(rootUrl+"auth/logout", handlers.Logout)

	//user
	app.Post(rootUrl+"user/update", handlers.UpdateInfo)

}
