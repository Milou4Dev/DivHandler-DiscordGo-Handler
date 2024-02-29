package modules

import (
    "fmt"
    "log"
    "github.com/bwmarrin/discordgo"
    "mpb/events"
)

func StartDiscordClient(token string) *discordgo.Session {
    session, err := discordgo.New("Bot " + token)
    if err != nil {
        log.Fatalf("Error creating Discord session: %v", err)
    }
    session.AddHandlerOnce(events.ReadyHandler)
    session.AddHandler(InteractionHandler)
    if err := session.Open(); err != nil {
        log.Fatalf("Error opening connection: %v", err)
    }
    if session.State.User != nil {
        RegisterSlashCommands(session)
    } else {
        log.Println("Session user is nil, cannot register commands.")
    }
    fmt.Println("Bot is now running. Press Ctrl+C to exit.")
    return session
}