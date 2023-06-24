package main

import (
	"os"
	"plms-backend/env"
	"plms-backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// go install github.com/githubnemo/CompileDaemon
// CompileDaemon -command="./plms-backend -log-prefix=0"

// Migrate using this
// go run ./migrate/migrate.go

func main() {
	env.LoadEnv()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/api*", swagger.HandlerDefault)

	routes.UserRoutes(app)

	app.Listen(os.Getenv("PORT"))
}
