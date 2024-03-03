package main

import (
	"code/modules"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	token := os.Getenv("TOKEN")
	session, err := modules.StartDiscordClient(token)
	if err != nil {
		log.Fatalf("Failed to start Discord client: %v", err)
	}
	defer session.Close()
	modules.WaitForInterrupt()
}
