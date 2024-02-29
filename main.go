package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "github.com/joho/godotenv"
    "mpb/modules"
    "github.com/bwmarrin/discordgo"
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
        session.Close()
    }
    fmt.Println("Shutting down finished.")
}

func waitForInterrupt() {
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop
}