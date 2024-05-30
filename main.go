package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Welcome to Tasks Applications!!!")
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})
	log.Fatal(app.Listen(":4000"))
}
