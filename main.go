package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"code/modules"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	token := os.Getenv("TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
	defer session.Close()

	session.AddHandlerOnce(readyHandler)
	session.AddHandler(modules.InteractionHandler)

	if err := session.Open(); err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	waitForInterrupt()

	fmt.Println("Shutting down gracefully...")
}

func loadEnv() error {
	return godotenv.Load(".env")
}

func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Bot is now ready.")
	if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "idle"}); err != nil {
		log.Printf("Error setting bot status: %v", err)
	}
	modules.RegisterSlashCommands(s)
}

func waitForInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}