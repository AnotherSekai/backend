package middleware

import (
	"github.com/gofiber/fiber/v3"
)

var validRegions = map[string]bool{
	"jp": true,
	"en": true,
	"cn": true,
	"tc": true,
	"kr": true,
}

func ValidateRegion(c fiber.Ctx) error {
	region := c.Params("region")
	if !validRegions[region] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid region",
		})
	}
	return c.Next()
}
