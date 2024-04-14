package routes

import (
	"github.com/go-interview/handlers"
	"github.com/gofiber/fiber/v2"
)

func bankRoutes(route fiber.Router, token fiber.Handler) {

	route.Post("login", handlers.Login)
	route.Post("create-transaction", handlers.CreateTransaction)

}
