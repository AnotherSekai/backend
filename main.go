package main

import (
	"log"

	"github.com/AnotherSekai/backend/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Sekai!")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":9000"))
}
