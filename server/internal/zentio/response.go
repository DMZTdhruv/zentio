package zentio

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data any) error {
	response := fiber.Map{
		"status":  "success",
		"message": message,
		"data":    nil,
		"error":   false,
	}

	if data != nil {
		response["data"] = data
	}
	return c.Status(statusCode).JSON(response)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"data":    nil,
		"error":   true,
	})
}
