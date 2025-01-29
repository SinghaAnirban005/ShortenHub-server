package main

import (
	"log"
	"os"

	"github.com/SinghaAnirban005/ShortenHub-server/prisma"

	"github.com/SinghaAnirban005/ShortenHub-server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	prisma.InitPrisma()
	defer prisma.ClosePrisma()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
