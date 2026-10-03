package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	liatrio_app := fiber.New()
	liatrio_app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Max Dieterle",
			"timestamp": time.Now().Unix(),
		})
	})

	liatrio_app.Listen(":3000")
}
