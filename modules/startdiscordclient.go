package modules

import (
    "code/events"
    "github.com/bwmarrin/discordgo"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func StartDiscordClient(token string) (*discordgo.Session, error) {
    session, err := discordgo.New("Bot " + token)
    if err != nil {
        return nil, err
    }
    session.AddHandlerOnce(events.ReadyHandler)
    session.AddHandler(InteractionHandler)
    if err := session.Open(); err != nil {
        return nil, err
    }
    RegisterSlashCommands(session)
    log.Println("Bot is now running. Press Ctrl+C to exit.")
    return session, nil
}

func WaitForInterrupt() {
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop
}