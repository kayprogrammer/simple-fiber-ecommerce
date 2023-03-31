package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kayprogrammer/simple-fiber-ecommerce/database"
	"github.com/kayprogrammer/simple-fiber-ecommerce/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome api")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoints
	app.Get("/api", welcome)
	
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
