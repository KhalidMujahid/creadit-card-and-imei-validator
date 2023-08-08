package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khalidMujahid/src/utils"
)

// Get request
func RenderWelcomePage(c *fiber.Ctx) error {
	return c.Status(200).Render("index", fiber.Map{
		"Message": "Credit card validtion system",
	})
}

// Get request
func RenderCCPage(c *fiber.Ctx) error {
	return c.Status(200).Render("validator", fiber.Map{})
}

// body
type bodyData struct {
	Value string `json:"value"`
}

// Post request
func Handler(c *fiber.Ctx) error {
	var body bodyData

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := utils.CreditCardValidator(body.Value)

	if result {
		return c.Status(200).JSON(fiber.Map{
			"value": "Credit card is valid",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"value": "Credit card is not valid",
		})
	}
}
