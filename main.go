package main

import (
	"os"
	"plms-backend/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// go install github.com/githubnemo/CompileDaemon
// CompileDaemon -command="./plms-backend"

// Migrate using this
// go run ./migrate/migrate.go

func main() {
    env.LoadEnv()
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(os.Getenv("PORT"))
}