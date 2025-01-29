package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateShortURL(length int) string {
	short := make([]byte, length)
	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}

	return string(short)
}
