package commands

import (
	"github.com/bwmarrin/discordgo"
)

func PongCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "pong",
		Description: "???",
	}
}

func HandlePongCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
}
