package routes

import (
	"plms-backend/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(a *fiber.App) {
	route := a.Group("/api/user")

	route.Post("/create", controller.CreateUser)
}
