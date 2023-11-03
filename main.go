package main

import (
	"log"
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.GetList)
	app.Get("/:id", controllers.GetDetail)
	app.Get("/api/user", controllers.QueryParam)

	log.Fatal(app.Listen(":7000"))
}
