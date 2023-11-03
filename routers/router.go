package routers

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func DefineRoutes(app *fiber.App) {
	app.Get("/api/v1", controllers.GetList)
	app.Get("/:id", controllers.GetDetail)
	app.Get("/api/user", controllers.QueryParam)
}
