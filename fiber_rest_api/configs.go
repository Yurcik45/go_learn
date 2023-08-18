package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Gogi",
		AppName: "Gogi App",
	})
	
	app.Get("/gogi", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Be more polite. Gogi prefers his name to start with an uppercase letter.")
	})

	app.Get("/Gogi", func(c *fiber.Ctx) error {
		return c.SendString("Gogi is pleased. He's having sex with Diana, slapping her ass with a bouquet of flowers.")
	})

	app.Listen(":3000")
}