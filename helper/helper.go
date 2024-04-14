package helper

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, data interface{}, message string, status bool, code int) error {
	var response ApiResponse
	response.Message = message
	response.Status = status
	response.Data = data
	return c.Status(code).JSON(response)
}
