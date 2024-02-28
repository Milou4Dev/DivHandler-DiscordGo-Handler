package commands

import (
	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"pong": handlePongCommand,
}

func PongCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "pong",
		Description: "Responds with a playful message.",
	}
}

func handlePongCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	respondWithMessage(s, i, "Wait, that's illegal.", discordgo.MessageFlagsEphemeral)
}

func respondWithMessage(s *discordgo.Session, i *discordgo.InteractionCreate, message string, flags discordgo.MessageFlags) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   flags,
		},
	})
}