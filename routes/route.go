package routes

import (
	"github.com/go-interview/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Routes(app *fiber.App) {

	jwtToken := jwtware.New(jwtware.Config{
		SigningKey: []byte(config.App.JWTKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Missing or malformed JWT", "status": false})
			} else {

				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid or expired JWT", "status": false})
			}
		},
	})

	route := app.Group("/")

	bankRoutes(route, jwtToken)
	thirdPartyRoutes(route, jwtToken)
}
