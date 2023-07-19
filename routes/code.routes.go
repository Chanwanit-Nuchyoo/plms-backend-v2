package routes

import (
	"plms-backend/controller"

	"github.com/gofiber/fiber/v2"
)

func CodeRoutes(a *fiber.App) {
	route := a.Group("/api/code")

	route.Post("/execute", controller.ExecuteCode)
}
