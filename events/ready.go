package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Bot is now ready.")
	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "idle"})
	if err != nil {
		log.Printf("Error setting bot status: %v", err)
	}
}
