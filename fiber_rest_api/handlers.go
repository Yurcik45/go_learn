package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	
	//Just to send a file
	app.Get("/file", func(c *fiber.Ctx) error {
		return c.SendFile("./fiber_rest_api/index.html")
	})

	//Processing duplicate uri-s 
	app.Get("/duplicate", func(c *fiber.Ctx) error {
		fmt.Println("1st route!")
		return c.Next()
	})
	  
	app.Get("/duplicate", func(c *fiber.Ctx) error {
		fmt.Println("2nd route!")
		return c.Next()
	})
	
	app.Get("/duplicate", func(c *fiber.Ctx) error {
		fmt.Println("3rd route!")
		return c.SendString("Hello, World!")
	}) 

	//Attaching multiple processors to one uri
	app.Get("/multiple/processors", 
		func(c *fiber.Ctx) error {
			fmt.Println("First processor of /multiple/processors")
			return c.Next()
		},
		func(c *fiber.Ctx) error {
			fmt.Println("Second processor of /multiple/processors")
			return c.SendString("Hello World")
		},
	)

	app.Listen(":3000")
}