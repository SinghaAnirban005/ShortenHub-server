package prisma

import (
	"log"

	"github.com/SinghaAnirban005/ShortenHub-server/prisma/db"
)

var Client *db.PrismaClient

func InitPrisma() {
	Client = db.NewClient()

	if err := Client.Connect(); err != nil {
		log.Fatal("Failed to connect prisma", err)
	}
}

func ClosePrisma() {
	if err := Client.Disconnect(); err != nil {
		log.Println("Failed to close Prisma connection:", err)
	}
}
