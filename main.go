package main

import (
	"github.com/go-interview/database"
	"github.com/go-interview/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//initialize database here
	database.Start()
	//migrate database here
	database.Migrate()

	routes.Routes(app)

	port := "4500"

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
