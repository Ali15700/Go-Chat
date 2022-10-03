package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1481390",
		Key:     "6a9eafb896efa21e4b0f",
		Secret:  "4c35fe0b35f3afbfbe58",
		Cluster: "ap2",
		Secure:  true,
	}
	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)
		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
