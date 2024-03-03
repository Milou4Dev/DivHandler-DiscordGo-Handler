package main

import (
	"code/modules"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var session *discordgo.Session

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	token := os.Getenv("TOKEN")
	session = modules.StartDiscordClient(token)
	waitForInterrupt()
	if session != nil {
		fmt.Println("Disconnecting the bot...")
		err := session.Close()
		if err != nil {
			return
		}
	}
	fmt.Println("Shutting down finished.")
}

func waitForInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
