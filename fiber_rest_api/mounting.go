package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	subapp := fiber.New()
	app := fiber.New()

	app.Mount("/api/some/mounted/prefix", subapp)
	
	subapp.Get("/service1", func(c *fiber.Ctx) error {
		return c.SendString("Service 1")
	})
	subapp.Get("/service2", func(c *fiber.Ctx) error {
		return c.SendString("Service 2")
	})

	app.Listen(":3000")
}