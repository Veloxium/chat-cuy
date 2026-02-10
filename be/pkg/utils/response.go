package utils

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

func SendError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(ErrorResponse{
		Error:   "error",
		Message: message,
	})
}

func SendSuccess(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(SuccessResponse{
		Message: message,
		Data:    data,
	})
}
