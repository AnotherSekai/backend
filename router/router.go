package router

import (
	"github.com/AnotherSekai/backend/handler"
	"github.com/AnotherSekai/backend/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	region := app.Group("/api/:region", middleware.ValidateRegion)

	region.Get("/gachas", handler.GetGachas)
}
