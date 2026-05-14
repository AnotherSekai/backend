package handler

import (
	"github.com/AnotherSekai/backend/service"
	"github.com/gofiber/fiber/v3"
)

func GetGachas(c fiber.Ctx) error {
	region := c.Params("region")

	resp, err := service.GetLatestGachas(region, 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(resp)
}
