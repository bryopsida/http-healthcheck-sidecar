package main

import (
	"log"

	"github.com/bryopsida/http-healthcheck-sidecar/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		if health.IsHealthy() {
			return c.Status(fiber.StatusOK).SendString("OK")
		} else {
			return c.Status(fiber.StatusServiceUnavailable).SendString("UNAVAILABLE")
		}
	})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/*": "/health"},
		StatusCode: 301,
	}))

	log.Fatal(app.Listen(":3000"))
}
