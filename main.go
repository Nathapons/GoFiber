package main

import (
	"log"
	"main/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.DefineRoutes(app)

	log.Fatal(app.Listen(":7000"))
}
