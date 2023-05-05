package main

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "github.com/joho/godotenv"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env: %v", err)
    }

    token := os.Getenv("TOKEN")

    s, err := discordgo.New("Bot " + token)
    if err != nil {
        log.Fatalf("Error creating Discord session: %v", err)
    }

    err = s.Open()
    if err != nil {
        log.Fatalf("Error opening session: %v", err)
    }

    registerSlashCommands(s)

    fmt.Printf("Bot is now running. Press Ctrl+C to exit.\n")

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop

    err = s.Close()
    if err != nil {
        log.Printf("Error closing session: %v", err)
    }
}