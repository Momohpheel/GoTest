package routes

import (
	"github.com/go-interview/handlers"
	"github.com/gofiber/fiber/v2"
)

func bankRoutes(route fiber.Router) {

	route.Post("create-transaction", handlers.CreateTransaction)

}
