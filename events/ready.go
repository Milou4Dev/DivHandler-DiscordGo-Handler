package events

import (
    "log"
    "github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
    log.Println("Bot is now ready.")
    if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "idle"}); err != nil {
        log.Printf("Error setting bot status: %v", err)
    }
}