package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetList(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Hello world",
	}
	return c.JSON(data)
}

func GetDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	data := map[string]string{
		"message": fmt.Sprintf("Get detail id %s", id),
	}
	return c.Status(http.StatusOK).JSON(data)
}


func QueryParam(c *fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")
	data := map[string]string{
		"username": username,
		"password": password,
	}
	return c.Status(http.StatusOK).JSON(data)
}
