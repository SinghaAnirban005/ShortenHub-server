package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/SinghaAnirban005/ShortenHub-server/prisma"
	"github.com/SinghaAnirban005/ShortenHub-server/prisma/db"
	"github.com/SinghaAnirban005/ShortenHub-server/utils"
)

func ShortenUrl(c *fiber.Ctx) error {
	var request struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	shortCode := utils.GenerateShortURL(6)

	_, err := prisma.Client.URL.CreateOne(
		db.URL.ShortCode.Set(shortCode),
		db.URL.OriginalURL.Set(request.URL),
	).Exec(context.Background())

	if err != nil {
		log.Println("Database error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(fiber.Map{"short_url": shortCode})

}

func RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("short_code")

	url, err := prisma.Client.URL.FindUnique(
		db.URL.ShortCode.Equals(shortCode),
	).Exec(context.Background())

	if err != nil || url == nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	_, err = prisma.Client.URL.UpsertOne(
		db.URL.ShortCode.Equals(shortCode),
	).Create(
		db.URL.ShortCode.Set(shortCode),
		db.URL.OriginalURL.Set(url.OriginalURL),
	).Update(
		db.URL.Clicks.Set(url.Clicks + 1),
	).Exec(context.Background())

	if err != nil {
		log.Println("Error during upsert:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.Redirect(url.OriginalURL, 301)
}
