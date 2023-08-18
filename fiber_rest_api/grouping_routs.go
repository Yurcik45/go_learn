package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func get_handler(path string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println("Path handled:" + path)
		return c.Next()
	}
}

func main() {
	app := fiber.New()
  
	api := app.Group("/api", get_handler("/api"))  // /api
  
	v1 := api.Group("/v1", get_handler("/api/v1"))   // /api/v1
	v1.Get("/list", get_handler("/api/v1/list"))          // /api/v1/list
	v1.Get("/user", get_handler("/api/v1/user"))          // /api/v1/user
  
	v2 := api.Group("/v2", get_handler("api/v2"))   // /api/v2
	v2.Get("/list", get_handler("api/v2/list"))          // /api/v2/list
	v2.Get("/user", get_handler("api/v2/user"))          // /api/v2/user
  
	app.Listen(":3000")
  }