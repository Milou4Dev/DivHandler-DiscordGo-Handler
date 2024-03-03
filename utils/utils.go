package utils

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func RespondWithMessage(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: message, Flags: flags},
	})
	if err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}
