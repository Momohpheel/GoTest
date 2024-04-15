package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	route := app.Group("/")
	troute := app.Group("/v1/")

	bankRoutes(route)
	thirdPartyRoutes(troute)
}
