package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"shoppy/routes"
	"shoppy/configs"
)

func main() {
	_, err := configs.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	routes.Routes(app)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
