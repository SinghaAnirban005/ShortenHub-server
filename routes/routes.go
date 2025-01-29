package routes

import (
	"github.com/SinghaAnirban005/ShortenHub-server/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", handlers.ShortenUrl)
	app.Get("/:short_code", handlers.RedirectURL)
}
