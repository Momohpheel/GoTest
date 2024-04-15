package routes

import (
	"github.com/go-interview/handlers"
	"github.com/gofiber/fiber/v2"
)

func thirdPartyRoutes(route fiber.Router) {

	route.Post("payments", handlers.ProcessTransaction)
	route.Get("payments/:ref", handlers.GetTransaction)

}
